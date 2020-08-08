
	package auth
	
	import (
		"automatepi/config"
		"automatepi/utils/console"
		"fmt"
		"net/http"
		//"reflect"
		"strconv"
	
		//"strconv"
		"strings"
		"time"
	
		"github.com/dgrijalva/jwt-go"
	)
	
	
	func CreateToken (user_id uint32) (string, error) {
		claims := jwt.MapClaims{}
		claims["authorized"] = true
		claims["user_id"] = user_id
		claims["exp"] = time.Now().Add(time.Hour * 100).Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString(config.SECRETKEY)
	}
	
	func TokenValid(r *http.Request ) error {
		tokenString := ExtractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return config.SECRETKEY, nil
		})
	
		if err != nil {
			return err
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			console.Pretty(claims)
		}
		return nil
	}
	
	func ExtractToken(r *http.Request) string {
		keys := r.URL.Query()
		token := keys.Get("token")
		if token != "" {
			return token
		}
		bearerToken := r.Header.Get("Authorization")
	
		if len(strings.Split(bearerToken, " ")) == 2 {
			return strings.Split(bearerToken, " ")[1]
		}
		return ""
	
	}
	
	func ExtractTokenID(r *http.Request) (uint32, error) {
		tokenString := ExtractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return config.SECRETKEY, nil
		})
	
		if err != nil {
			return 0, err
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
			if err != nil {
				return 0, err
			}
				
				return uint32(uid), nil
		}
		return 0, nil
	}