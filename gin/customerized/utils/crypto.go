package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

var (
	pubKey []byte
	prvKey []byte

	aesKey string
)

/*
	TODO: 定義一個資料結構把多種加密方法包裝起來
*/
func InitPubPrvKeyAndAesSert(pubKeyPath, prvKeyPath string, aesString string) error {
	var err error
	pubKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		return err
	}

	prvKey, err = ioutil.ReadFile(prvKeyPath)
	if err != nil {
		return err
	}

	aesKey = aesString

	return nil
}

func RsaEncrypt(origData []byte) (string, error) {
	// 用 pem.Decode 把公鑰解密，之後用 block 把金鑰緩存起來
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return "", errors.New("public key error")
	}

	// x509.ParsePKIPublicKey 解析公鑰
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", errors.New(fmt.Sprintf("error happend in parse key: %v", err))
	}

	pub := pubInterface.(*rsa.PublicKey)

	// TODO: 將密鑰封裝起來，只要呼叫 rsa.EncryptPKCS1v15 就好
	// 將要加密的的字串做 PKCS1v15 加密
	encryptBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	if err != nil {
		return "", errors.New(fmt.Sprintf("rsa encrypt err : %v", err))
	}

	return string(encryptBytes), nil
}

func RsaDecrypt(encryMsg string) ([]byte, error) {
	block, _ := pem.Decode(prvKey)
	if block == nil {
		return nil, errors.New("private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("s1 private key err: %v", err))
	}

	// TODO: 將密鑰封裝起來，只要呼叫 rsa.DecryptPKCS1v15 就好
	// 將要解密的的字串做 PKCS1v15 解密
	decryptBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, []byte(encryMsg))

	return decryptBytes, nil
}

/*
	TODO: 用裝飾把，把AES加解密的行為包裝起來
*/
// Decrypt : (解密)將[]byte轉換為string
func Decrypt(ciphertext []byte, key string) (string, error) {
	// 如果解密的時候沒有帶aesKey，使用預設的aesKey
	if key == "" {
		key = aesKey
	}

	// c, err := aes.NewCipher([]byte(key)) // passphrasewhichneedstobe32bytes!
	c, err := aes.NewCipher([]byte(GetMD5Hash(key))) // 需要先將加密的字串做MD5然後轉換成[]byte才能放入Cipher, 不然會報錯除非剛好32bytes "e.g. passphrasewhichneedstobe32bytes!"
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("Invalid length")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// Encrypt : (加密)將string轉換為[]byte
func Encrypt(text string, key string) ([]byte, error) {
	// 如果解密的時候沒有帶aesKey，使用預設的aesKey
	if key == "" {
		key = aesKey
	}

	ciphertext := []byte(text)
	// c, err := aes.NewCipher([]byte(GetMD5Hash(key))) // passphrasewhichneedstobe32bytes!
	c, err := aes.NewCipher([]byte(GetMD5Hash(key))) // 需要先將加密的字串做MD5然後轉換成[]byte才能放入Cipher, 不然會報錯除非剛好32bytes "e.g. passphrasewhichneedstobe32bytes!"
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	btxt := gcm.Seal(nonce, nonce, ciphertext, nil)

	return btxt, nil
}

// GetMD5Hash : 獲取MD5值
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
