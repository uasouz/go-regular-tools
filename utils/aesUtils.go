package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

type Decodable struct {
	Iv    string `json:"iv"`
	Value string `json:"value"`
}

func Encrypt(key []byte, message string) (encmess string, err error) {
	plainText := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	//IV needs to be unique, but doesn't have to be secure.
	//It's common to put it at the beginning of the ciphertext.
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}

func Decrypt(key []byte, securemess string) (decodedmess string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}
	var decodable = Decodable{}
	err = json.Unmarshal(cipherText, &decodable)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// if len(cipherText) < aes.BlockSize {
	// 	err = errors.New("Ciphertext block size is too short!")
	// 	return
	// }

	fmt.Println(decodable.Iv, decodable.Value)
	iv, err := base64.StdEncoding.DecodeString(decodable.Iv)
	if err != nil {
		fmt.Println(err)
		return
	} //cipherText[:aes.BlockSize]
	cipherText, err = base64.StdEncoding.DecodeString(decodable.Value) //[aes.BlockSize:]
	if err != nil {
		fmt.Println(err)
		return
	}

	stream := cipher.NewCBCDecrypter(block, iv)
	fmt.Println("|||||||||STREAM|||||||||", string(cipherText))
	stream.CryptBlocks(cipherText, cipherText)

	decodedmess = string(cipherText)
	return
}
