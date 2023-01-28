package argon

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/devNica/mochileros/configurations"
	"github.com/devNica/mochileros/exceptions"
	"golang.org/x/crypto/argon2"
)

func GeneratePassworHash(password string, agc *configurations.Argon2Config) string {
	saltBytes := generateRandomBytes(agc.SaltLength)
	_, err := rand.Read(saltBytes)
	exceptions.PanicLogging(err)

	argon2Hash := argon2.IDKey([]byte(password), saltBytes, agc.Iterations, agc.Memory, agc.Parallelism, agc.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(saltBytes)
	b64Argon2Hash := base64.RawStdEncoding.EncodeToString(argon2Hash)

	hash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, agc.Memory,
		agc.Iterations, agc.Parallelism, b64Salt, b64Argon2Hash)

	return hash
}

func generateRandomBytes(n uint32) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	exceptions.PanicLogging(err)

	return b
}
