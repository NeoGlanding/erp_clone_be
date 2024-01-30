package helpers

// package helpers

// import (
// 	// "encoding/base64"
// 	// "fmt"
// 	// "log"
// 	"encoding/base64"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/automa8e_clone/config"
// 	"github.com/golang-jwt/jwt"
// 	// "github.com/goccy/go-json"
// )

// func GenerateGoogleJWT() (string, error) {
// 	// secretKey := config.FirebaseConfig.PrivateKey

// 	// newData := strings.Replace(secretKey,"\\n", "\n", -1);

// 	var header map[string]interface{} = map[string]interface{}{}
// 	var claims jwt.MapClaims = jwt.MapClaims{}

// 	header["alg"] = "RS256"
// 	header["typ"] = "JWT"
// 	header["kid"] = config.FirebaseConfig.ProjectKeyId

// 	claims["exp"] = time.Now().Add(time.Minute * time.Duration(30)).Unix()
// 	claims["iss"] = "firebase-adminsdk-fzpuv@file-uploading-e17ad.iam.gserviceaccount.com"
// 	claims["scope"] = "https://www.googleapis.com/auth/devstorage.full_control"
// 	claims["aud"] = "https://oauth2.googleapis.com/token"
// 	claims["iat"] = time.Now().Unix()

// 	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	// token.Header["kid"] = config.FirebaseConfig.ProjectKeyId
// 	// token.Header["alg"] = "RS256"

// 	// signedToken, err := token.SignedString([]byte(secretKey))

// 	// if (err != nil) {
// 	// 	log.Fatalf("while signing token %v \n", err)
// 	// }

// 	// fmt.Println("signed token -> ", signedToken)

// 	jsonHeader, err := json.Marshal(header)

// 	if (err != nil) {
// 		log.Fatal(err)
// 	}

// 	jsonClaims, err := json.Marshal(claims)

// 	if (err != nil) {
// 		log.Fatal(err)
// 	}

// 	headerEncoded := base64.URLEncoding.EncodeToString(jsonHeader)
// 	claimsEncoded := base64.URLEncoding.EncodeToString(jsonClaims)
// 	signature := fmt.Sprintf("%s.%s", headerEncoded, claimsEncoded)
// 	signatureEncoded := base64.URLEncoding.EncodeToString([]byte(signature))

// 	token := fmt.Sprintf("%s.%s.%s", headerEncoded, claimsEncoded, signatureEncoded);

// 	fmt.Println(token)

// 	return "tokenString", nil
// }