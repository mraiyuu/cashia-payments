package handlers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mraiyuu/cashia-payments/internal/utils"
	"github.com/joho/godotenv"
)

func init() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        slog.Error("Error loading .env file")
    }
}

func AuthenticateMerchant(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("APP_SECRET")
	if secret == "" {
		utils.WriteJSONError(w, "missing app secret", http.StatusInternalServerError)
		return
	}

	keyID := os.Getenv("APP_KEY_ID")
	if keyID == "" {
		utils.WriteJSONError(w, "missing app key id", http.StatusInternalServerError)
		return
	}
	
	host := "localhost:8000"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := "nonce"
	signatureRaw := host + "POST" + timestamp + nonce + keyID

	data := map[string]interface{}{
		"age":   16,
		"email": "user@gmail.com",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	headers := map[string]string{
		"X-Cashia-Key-ID":    keyID,
		"X-Cashia-Timestamp": timestamp,
		"X-Cashia-Signature": computeHmac256(signatureRaw, secret),
		"X-Cashia-Nonce":     nonce,
		"X-Cashia-Hash":      computeHmac256(string(jsonData), secret),
	}

	url := "http://" + host + "/api/v1/merchant-info"
	// url := "https://www.google.com"
	response, err := makeHttpRequest(url, "POST", headers, jsonData)
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println(string(response))
}

func computeHmac256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func makeHttpRequest(url, method string, headers map[string]string, jsonData []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	slog.Error(string(body))
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return body, nil
}
