package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/scrypt"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var salt = []byte{0xa4, 0x98, 0xc2, 0xf8, 0xed, 0x51, 0xb3, 0x76}

func encodeBase32(b []byte) string {
	str := base32.StdEncoding.EncodeToString(b)
	return strings.Split(str, "=")[0]
}

func decodeBase32(s string) []byte {
	data, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("decode base32 error: %s", err)
	}

	return data
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Fatalf("decode base64 err: %s", err)
	}
	return data
}

func encodeHex(b []byte) string {
	return hex.EncodeToString(b)
}

func decodeHex(s string) []byte {
	decodeStr, err := hex.DecodeString(s)

	if err != nil {
		log.Fatalf("Decode hex string error: %s", err)
	}
	return decodeStr
}

func getEncryptionKey() []byte {
	secretKey := os.Getenv("SECRET_KEY")
	dk, err := scrypt.Key([]byte(secretKey), salt, 32768, 8, 1, 32)
	if err != nil {
		log.Fatalf("Get encryption key error: %s", err)
	}

	return dk
}

func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher(getEncryptionKey())
	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encodeHex(cipherText), nil
}

func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher(getEncryptionKey())
	if err != nil {
		return "", err
	}

	cipherText := decodeHex(text)
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil

}
