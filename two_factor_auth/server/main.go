package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()

	// func 在此處其實可以被拉出，作為另一層controller
	router.GET("/", func(c *gin.Context) {
		// 取決是否需要使用預設query參數
		// initStr := c.Query("init") // "abcdefghigklmnop"
		initStr := c.DefaultQuery("init", "abcdefghigklmnop")
		str, err := GetCode(initStr)
		if err != nil {
			panic(err)
		}

		// 回傳response
		c.JSON(200, gin.H{
			"message": str,
		})
	})

	return router
}

func main() {
	router := setUpRouter()
	router.Run()
}

func GetCode(key string) (string, error) {
	inputNoSpacesUpper := strings.ToUpper(key)
	encodeKey, err := base32.StdEncoding.DecodeString(inputNoSpacesUpper)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	epochSeconds := time.Now().Unix()
	pwd := oneTimePassword(encodeKey, toBytes(epochSeconds/30))
	log.Println(pwd)
	return fmt.Sprintf("%06d", pwd), nil
}

func oneTimePassword(key []byte, value []byte) uint32 {
	// sign the value using HMAC-SHA1
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(value)
	hash := hmacSha1.Sum(nil)

	// We're going to use a subset of the generated hash.
	// Using the last nibble (half-byte) to choose the index to start from.
	// This number is always appropriate as it's maximum decimal 15, the hash will
	// have the maximum index 19 (20 bytes of SHA1) and we need 4 bytes.
	offset := hash[len(hash)-1] & 0x0F

	// get a 32-bit (4-byte) chunk from the hash starting at offset
	hashParts := hash[offset : offset+4]

	// ignore the most significant bit as per RFC 4226
	hashParts[0] = hashParts[0] & 0x7F

	number := toUint32(hashParts)

	// size to 6 digits
	// one million is the first number with 7 digits so the remainder
	// of the division will always return < 7 digits
	pwd := number % 1000000

	return pwd
}

func toBytes(value int64) []byte {
	var result []byte
	mask := int64(0xFF)
	shifts := [8]uint16{56, 48, 40, 32, 24, 16, 8, 0}
	for _, shift := range shifts {
		result = append(result, byte((value>>shift)&mask))
	}
	return result
}

func toUint32(bytes []byte) uint32 {
	return (uint32(bytes[0]) << 24) + (uint32(bytes[1]) << 16) +
		(uint32(bytes[2]) << 8) + uint32(bytes[3])
}
