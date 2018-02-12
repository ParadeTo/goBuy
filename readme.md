# JWT

## Create Token And Register Routes
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

auth := r.Group("/", JWTAuth())
{
    auth.GET("/data", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "data": "secret",
        })
    })
}
```

## Decode Token In Fronted-End To Get User Info
```
token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMCwibmFtZSI6ImRlZmF1bHQiLCJlbWFpbCI6ImRlZmF1bHRAcXEuY29tIiwiaXNzIjoiZGVmYXVsdCJ9.x06cm6t8AZfmW3WHPh31rYVJlfmt3LSWxN2COnH0CJg"

decodebase64(token.split('.')[1])

{"id":1000,"name":"default","email":"default@qq.com","iss":"default"}
```

## JWT Authorization
Just add token to request header:

```
"Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTAwMC..."
```
