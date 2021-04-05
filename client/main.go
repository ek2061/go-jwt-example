package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Key很重要，最好從環境變數讀，不應直接在程式宣告
var SigningKey = []byte("secret")

// Get / 印出token
func homePage(w http.ResponseWriter, r *http.Request) {
	// 製造token
	validToken, err := GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	// 直接印在 http://localhost:9002
	fmt.Fprintf(w, validToken)
}

// 製造token
func GenerateJWT() (string, error) {
	// token算法
	token := jwt.New(jwt.SigningMethodHS256)

	// token資訊
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Yuchi"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	// 製造token
	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		fmt.Errorf("發生錯誤: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

// 起動server在 http://localhost:9002
func handleRequests() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":9002", nil))
}

func main() {
	fmt.Println("這是client端")

	// tokenString, err := GenerateJWT()
	// if err != nil {
	// 	fmt.Println("Error generating token string")
	// }

	// fmt.Println(tokenString)
	handleRequests()
}
