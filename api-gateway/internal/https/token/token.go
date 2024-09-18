package token

import (
	"errors"
	"fmt"

	"api-gateway/internal/pkg/load"

	"github.com/form3tech-oss/jwt-go"
)

type Tokens struct {
	Token     string
	ExpiresIn string
}


// func GenereteJWTToken(req *pb.LoginUserResponse) *Tokens {
	
// 	refreshToken := jwt.New(jwt.SigningMethodHS256)

// 	rftclaims := refreshToken.Claims.(jwt.MapClaims)
// 	rftclaims["user_id"] = req.UserId
// 	rftclaims["user_email"] = req.Email
// 	rftclaims["iat"] = time.Now().Unix()
// 	rftclaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
// 	refresh, err := refreshToken.SignedString([]byte(tokenKey))
// 	if err != nil {
// 		log.Fatal("error while genereting refresh token : ", err)
// 	}

// 	expirationUnix := rftclaims["exp"].(int64)
// 	expirationTime := time.Unix(expirationUnix, 0).Format("2006-01-02 15:04:05")

// 	return &Tokens{
	// 		Token:     refresh,
	// 		ExpiresIn: expirationTime,
	// 	}
	// }
	
	func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
		cfgRun, err := load.Load("./config/config.yaml")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		tokenKey := cfgRun.TokenKey
		fmt.Println("tokenKey")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte(tokenKey), nil
		})
		fmt.Println(tokenKey)
		if err != nil {
			return nil, err
		}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
