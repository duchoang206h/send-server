package util

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/duchoang206h/send-server/config"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

type ShortenUrlAPIResponse struct {
	Short string `json:"short"`
}

func ShortenUrl(url string) (string, error) {
	shortenAPIUrl := config.Config("SHORTEN_API_URL")
	resp, err := http.Get(shortenAPIUrl + fmt.Sprintf("?url=%s", url))
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var shortRes ShortenUrlAPIResponse
	err = json.Unmarshal(body, &shortRes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", config.Config("SHORTEN_HOST"), shortRes.Short), nil
}
