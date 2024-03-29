package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func decrypt(data []byte, passphrase string) ([]byte, error) {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func validate(checkstring string) (string, error) {
	for key, p := range validator {
		check, err := regexWith(checkstring, key, p.message, p.check)
		if err != nil {
			return "", err
		}
		if check != "" {
			return check, nil
		}
	}
	return "", nil
}

func regexWith(checkstring string, pat string, errmsg string, check bool) (string, error) {
	rs, err := regexp.MatchString(pat, checkstring)
	fmt.Println("In:", checkstring)
	fmt.Println("Pat:", pat)
	fmt.Println("Msg:", errmsg)
	fmt.Println("CheckParam:", check)
	fmt.Println("CHeckResult:", rs)

	if err != nil {
		return "", err
	}
	if check == rs {
		return "", nil
	}
	return errmsg, nil
}