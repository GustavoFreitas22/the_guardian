package cript

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/GustavoFreitas22/the_guardian/model"
)

func EncriptDataInfo(data model.Data) (model.Data, error) {
	cipherBlock, err := aes.NewCipher([]byte(data.UserKey))
	if err != nil {
		log.Println(err)
		return data, nil
	}

	salt, err := GenerateSalt()
	if err != nil {
		log.Println(err)
		return model.Data{}, err
	}

	data.Salt = salt

	cipherEncrypt := cipher.NewCFBEncrypter(cipherBlock, salt)
	plainValue := []byte(data.Value)
	keyCrypt := make([]byte, len(plainValue))

	cipherEncrypt.XORKeyStream(keyCrypt, []byte(data.Value))

	data.Value = encodeData(keyCrypt)

	return data, err
}

func encodeData(encryp []byte) string {
	return base64.StdEncoding.EncodeToString(encryp)
}

func decodeData(value string) []byte {
	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Panic(err)
	}
	return data
}

func DecryptDataInfo(data model.Data) (model.Data, error) {

	cipherBlock, err := aes.NewCipher([]byte(data.UserKey))
	if err != nil {
		log.Println(err)
		return data, err
	}

	cipherValue := decodeData(data.Value)

	cipherDecrypt := cipher.NewCFBDecrypter(cipherBlock, data.Salt)
	valueDecrypt := make([]byte, len(cipherValue))

	cipherDecrypt.XORKeyStream(valueDecrypt, cipherValue)

	data.Value = string(valueDecrypt)
	log.Println(data.Value)

	return data, err
}

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
