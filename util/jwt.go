package util

import (
    "time"

    jwt "github.com/dgrijalva/jwt-go"

    "gintest/config"
)

var jwtSecret = []byte(config.JwtSecret)

type Claims struct {
    Username string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
    nowTime := time.Now()
    expireTime := nowTime.Add(3 * time.Hour)

    claims := Claims{
        username,
        password,
        jwt.StandardClaims {
            ExpiresAt : expireTime.Unix(),
            Issuer : "gintest",
        },
    }

    tokenClaims := jwt.NewWithClaims(jwt.GetSigningMethod(config.SigningMethod), claims)
    token, err := tokenClaims.SignedString(jwtSecret)

    return token, err
}

func ParseToken(token string) (*Claims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}