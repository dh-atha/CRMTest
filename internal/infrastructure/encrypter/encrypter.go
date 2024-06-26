package encrypter

import (
	"crypto/md5"
	"encoding/hex"
)

// func EncryptPassword(password string, cost int) (string, error) {
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(hash), nil
// }

// func VerifyPassword(password, hashedPassword string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func calculateMD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(password, hashedPassword string) bool {
	return calculateMD5Hash(password) == hashedPassword
}
