package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func encryptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		plaintext := []byte("my secret message")
		key := []byte("a very very very very secret key")

		ciphertext, err := encrypt(plaintext, key)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Ciphertext: %x\n", ciphertext)

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(encryptMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, world!",
		})
	})

	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
