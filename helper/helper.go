package helper

import (
	"encoding/base64"
	"math/rand"
	"os"
	"regexp"
)

// RandomStringBase64 function for random string and base64 encoded
func RandomStringBase64(length int) string {
	rb := make([]byte, length)
	_, err := rand.Read(rb)

	if err != nil {
		return ""
	}
	rs := base64.URLEncoding.EncodeToString(rb)

	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		return ""
	}

	return reg.ReplaceAllString(rs, "")
}

// GetEnv read `ENV` variable from os system
func GetEnv() (env string) {
	env = os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	return env
}