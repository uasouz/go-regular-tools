package utils

import (
	"regexp"

	"github.com/codahale/blake2"
	"golang.org/x/crypto/bcrypt"
)

//RemoveNonNumeric - Retira todos os caracters não numéricos da string inserida
func RemoveNonNumeric(text string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return "", err
	}
	processedString := reg.ReplaceAllString(text, "")
	return processedString, nil
}

func HashToBCrypt(pass string) ([]byte, error) {
	password := []byte(pass)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil

}

func HashToBlake2b(pass string) ([]byte, error) {
	// password := []byte(pass)

	h := blake2.NewBlake2B()
	h.Write([]byte(pass))
	hashedPassword := h.Sum(nil)
	return hashedPassword, nil

}

func HashToBlake2bShort(pass string) ([]byte, error) {
	// password := []byte(pass)

	h := blake2.New(&blake2.Config{Size: 32})
	h.Write([]byte(pass))
	hashedPassword := h.Sum(nil)
	return hashedPassword, nil

}
