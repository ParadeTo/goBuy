package jwt

import (
	"testing"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"net/http"
	"encoding/json"
	"strings"
)
type CustomClaimsTest struct {
	*CustomClaims
	siginKey string
	wanted string
}
type ExpiredClaimsTest struct {
	CustomClaims
	siginKey string
}
var claims = []CustomClaimsTest{
	{
		&CustomClaims{
			1,
			"awh521",
			"1044176017@qq.com",
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
				Issuer: "test",
			},
		},
		"test",
		"",
	},
}
var expiredClaims = []ExpiredClaimsTest{
	{
		CustomClaims{
			1,
			"awh521",
			"1044176017@qq.com",
			jwt.StandardClaims{
				ExpiresAt: 1500,
				Issuer: "test",
			},
		},
		"test",
	},
}
var jt *JWT = &JWT{
	[]byte("test"),
}
var foreverClaims CustomClaims = CustomClaims{
	1000,
	"default",
	"default@qq.com",
	jwt.StandardClaims{
		ExpiresAt: 0,
		Issuer: "default",
	},
}
func TestCreateForeverTokens(t *testing.T) {
	token, err := jt.CreateToken(foreverClaims)

	assert.NoError(t, err)
	claims, err := jt.ParseToken(token)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), claims.StandardClaims.ExpiresAt)
}
func TestJWTCreateToken(t *testing.T) {
	for _, c := range claims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		token, err := j.CreateToken(*c.CustomClaims)
		assert.NoError(t, err)
		assert.IsType(t, "string", token)
	}
}
func TestJWTParseToken(t *testing.T) {
	for _, c := range claims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		var err error
		c.wanted, err = j.CreateToken(*c.CustomClaims)
		result, err := j.ParseToken(c.wanted)
		assert.NoError(t, err)
		assert.Equal(t, c.CustomClaims.ID, result.ID)
		assert.Equal(t, c.CustomClaims.Email, result.Email)
		assert.Equal(t, c.CustomClaims.Name, result.Name)
		assert.Equal(t, c.CustomClaims.StandardClaims.ExpiresAt, result.StandardClaims.ExpiresAt)
		assert.Equal(t, c.CustomClaims.StandardClaims.Issuer, result.StandardClaims.Issuer)
	}
}
func TestRefreshToken(t *testing.T) {
	for _, c := range expiredClaims {
		j := &JWT{SigningKey: []byte(c.siginKey)}
		token, err := j.CreateToken(c.CustomClaims)
		assert.NoError(t, err)
		claims, err := j.ParseToken(token)
		assert.EqualError(t, err, TokenExpired.Error())
		assert.Nil(t, claims)
		token, err = j.RefreshToken(token)
		assert.NoError(t, err)
		assert.IsType(t, "string", token)
	}
}


type RepToken struct {
	Token string
}

type Data struct {
	Data string
}

type CustomClaim struct {
	Id int
	Name string
}


func TestJwtAuth (t *testing.T) {
	token, err := jt.CreateToken(foreverClaims)

	r := gin.New()

	r.POST("/login", func(c *gin.Context) {

		if err == nil {
			c.JSON(200, gin.H{
				"token": token,
			})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	})

	auth := r.Group("/", JWTAuth())
	{
		auth.GET("/data", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": "secret",
			})
		})
	}

	w0 := httptest.NewRecorder()
	req0, _ := http.NewRequest("GET", "/data", nil)
	r.ServeHTTP(w0, req0)
	assert.Equal(t, 401, w0.Code)


	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/login", nil)
	r.ServeHTTP(w1, req1)
	assert.Equal(t, 200, w1.Code)
	repToken := &RepToken{}
	json.Unmarshal([]byte(w1.Body.String()), repToken)
	assert.Equal(t, token, repToken.Token)
	parts := strings.Split(repToken.Token, ".")
	customClaimsByte, err := jwt.DecodeSegment(parts[1])
	assert.NoError(t, err)
	customClaim := CustomClaim{}
	json.Unmarshal(customClaimsByte, &customClaim)
	assert.Equal(t, customClaim.Id, foreverClaims.ID)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/data", nil)
	req2.Header = map[string][]string{
		"Authorization": {repToken.Token},
	}
	r.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)
	data := &Data{}
	json.Unmarshal([]byte(w2.Body.String()), data)
	assert.Equal(t, "secret", data.Data)
}
