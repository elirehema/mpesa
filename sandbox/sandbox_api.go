package sandbox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

var MpesaClient = initHttpClient()

func GetSessionBearerAccessToken(apiKey string) string {
	godotenv.Load()

	publicKey := strings.TrimSpace(os.Getenv("API_PUBLIC_KEY"))
	encodedPublicKey, _ := base64.StdEncoding.DecodeString(publicKey)
	publicKeySpec, _ := x509.ParsePKIXPublicKey(encodedPublicKey)
	pk := publicKeySpec.(*rsa.PublicKey)
	cipher, _ := rsa.EncryptPKCS1v15(rand.Reader, pk, []byte(apiKey))
	encryptedApiKey := base64.StdEncoding.EncodeToString(cipher)
	return encryptedApiKey
}
func initHttpClient() *resty.Client {
	client := resty.New()
	godotenv.Load()

	BASE_URL := os.Getenv("BASE_URL")
	COUNTRY := os.Getenv("COUNTRY")
	baseUrl := fmt.Sprintf("%s/%s", BASE_URL, COUNTRY)

	client.SetBaseURL(baseUrl)
	client.SetHeader("Accept", "*/*")
	client.SetHeader("Origin", "*")
	client.SetHeader("Host", "openapi.m-pesa.com")
	client.SetHeader("Content-Type", "application/json")
	client.SetTimeout(1 * time.Minute)
	client.SetDebug(true)

	client.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		if r.URL == "/getSession/" {
			apiKey := strings.TrimSpace(os.Getenv("API_KEY"))
			r.SetAuthToken(GetSessionBearerAccessToken(apiKey))
		} else {
			godotenv.Load("session.env")
			apiKey := strings.TrimSpace(os.Getenv("API_SESSION_KEY"))
			r.SetAuthToken(GetSessionBearerAccessToken(apiKey))
		}
		return nil // if its success otherwise return error
	})
	client.OnError(func(req *resty.Request, err error) {
		if v, ok := err.(*resty.ResponseError); ok {
		} else {
			log.Println("Logging to custom fileeeeeeee: ", v.Error())
		}

		// Log the error, increment a metric, etc...
	})
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {

		return nil // if its success otherwise return error
	})

	client.SetDebug(true)

	return client
}
