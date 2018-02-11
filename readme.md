# jwt
```
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

...
// first, login to get token
r.POST("/login", func(c *gin.Context) {
    token, err := jt.CreateToken(foreverClaims)
    if err == nil {
        c.JSON(200, gin.H{
            "token": token,
        })
    } else {
        c.JSON(401, gin.H{"status": "unauthorized"})
    }
})

// add token to request header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMC..."
auth := r.Group("/", JWTAuth())
{
    auth.GET("/data", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "data": "secret",
        })
    })
}
```