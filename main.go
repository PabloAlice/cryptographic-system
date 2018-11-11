package main

import (
	"bytes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PabloAlice/cryptographic-system/src"

	present "github.com/PabloAlice/cryptographic-system/src/presentCipher"
	"github.com/gin-gonic/gin"
)

func parseForm(c *gin.Context, decryption bool) (*present.Cipher, []byte, []byte, string, []byte, string, []byte) {
	file, _ := c.FormFile("file")
	method := c.PostForm("method")
	rawKey := c.PostForm("key")
	rawIv := c.PostForm("iv")
	key := []byte(rawKey)
	block, err := present.New(key)
	if err != nil {
		panic(err)
	}
	blockSize64 := int64(block.BlockSize())
	src := make([]byte, file.Size+blockSize64-file.Size%blockSize64)
	cipherData := make([]byte, len(src)+block.BlockSize())
	cipherDataBuffer := new(bytes.Buffer)
	cipherDataBuffer.Write(([]byte(rawIv))[:block.BlockSize()])
	iv := cipherData[:block.BlockSize()]
	cipherDataBuffer.Read(iv)
	encryptedData := cipherData[block.BlockSize():]
	encryptedFileName := fmt.Sprintf("ENC%v", file.Filename)
	if decryption {
		encryptedFileName = fmt.Sprintf("DEC%v", file.Filename)
	}
	srcFile, _ := file.Open()
	defer srcFile.Close()
	srcFile.Read(src)

	return block, encryptedData, src, encryptedFileName, cipherData, method, iv
}

func saveFile(data []byte, fileName string) {
	ioutil.WriteFile(fmt.Sprintf("./build/static/%v", fileName), data, 0644)
	ioutil.WriteFile(fmt.Sprintf("/Users/pabloalice/Desktop/%v", fileName), data, 0644)
}

func encrypt(block *present.Cipher, iv []byte, method string, dest []byte, src []byte) {
	switch method {
	case "CBC":
		println("CBC")
		blockMode := cipher.NewCBCEncrypter(block, iv)
		blockMode.CryptBlocks(dest, src)
	case "CFB":
		println("CFB")
		blockMode := cipher.NewCFBEncrypter(block, iv)
		blockMode.XORKeyStream(dest, src)
	case "OFB":
		println("OFB")
		blockMode := cipher.NewOFB(block, iv)
		blockMode.XORKeyStream(dest, src)
	default:
		println(fmt.Sprintf("default: %v", method))
		blockMode := ecb.NewECBEncrypter(block)
		blockMode.CryptBlocks(dest, src)
	}
}

func decrypt(block *present.Cipher, iv []byte, method string, dest []byte, src []byte) {
	switch method {
	case "CBC":
		println("CBC")
		blockMode := cipher.NewCBCDecrypter(block, iv)
		blockMode.CryptBlocks(dest, src)
	case "CFB":
		println("CFB")
		blockMode := cipher.NewCFBDecrypter(block, iv)
		blockMode.XORKeyStream(dest, src)
	case "OFB":
		println("OFB")
		blockMode := cipher.NewOFB(block, iv)
		blockMode.XORKeyStream(dest, src)
	default:
		println(fmt.Sprintf("default: %v", method))
		blockMode := ecb.NewECBDecrypter(block)
		blockMode.CryptBlocks(dest, src)
	}
}

func main() {
	r := gin.Default()
	r.POST("/api/encryption", func(c *gin.Context) {
		block, encryptedData, src, encryptedFileName, cipherData, method, iv := parseForm(c, false)
		// end data parsed
		encrypt(block, iv, method, encryptedData, src)
		saveFile(cipherData[block.BlockSize():], encryptedFileName)
		c.JSON(http.StatusOK, gin.H{
			"fileName": encryptedFileName,
			"method":   method,
		})
	})
	r.POST("/api/decryption", func(c *gin.Context) {
		block, encryptedData, src, encryptedFileName, cipherData, method, iv := parseForm(c, true)
		decrypt(block, iv, method, encryptedData, src)
		saveFile(cipherData[block.BlockSize():], encryptedFileName)
		c.JSON(http.StatusOK, gin.H{
			"fileName": encryptedFileName,
			"method":   method,
		})
	})

	r.Static("/static", "./build/static")

	r.GET("/", func(c *gin.Context) {
		c.File("./build/index.html")
	})
	log.Println("Listening...")
	r.Run() // listen and serve on 0.0.0.0:8080

}

var _ cipher.Block = (*present.Cipher)(nil)
