package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte("secret")

// 製造帶有重要資訊的首頁
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "超級機密文件")
}

// 検査是否帶有token
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 有token
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				// 解析token
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("解析token失敗")
				}
				return SigningKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			// token解析為有效就執行傳入的homePage()
			if token.Valid {
				endpoint(w, r)
			}
			// 沒token
		} else {
			fmt.Fprint(w, "未被授權token")
		}
	})
}

// 起動server在 http://localhost:9000
func handleRequests() {
	http.Handle("/", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("這是server端")
	handleRequests()
}
