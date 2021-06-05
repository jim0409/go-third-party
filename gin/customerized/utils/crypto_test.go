package utils

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestMD5(t *testing.T) {
	verified := "202cb962ac59075b964b07152d234b70"
	v := GetMD5Hash("123")
	assert.Equal(t, verified, v)

}

func TestAesCryptor(t *testing.T) {
	cs := "This is over 32 bytes gooddddddddddddddddddddddddddgoodddddddddddddddddddddddddd"
	etext, err := Encrypt("test", cs)
	assert.NoError(t, err)

	dtext, err := Decrypt(etext, cs)
	assert.NoError(t, err)
	assert.Equal(t, "test", dtext)

	// 考慮如果沒帶加密Key的情況，使用設定檔(預設的)AesKey
	etext1, err := Encrypt("test1", "")
	assert.NoError(t, err)

	dtext1, err := Decrypt(etext1, "")
	assert.NoError(t, err)
	assert.Equal(t, "test1", dtext1)
}

func TestRsaCryptor(t *testing.T) {
	err := InitPubPrvKeyAndAesSert("../conf/rsa_pub.key", "../conf/rsa_prv.key", "tester")
	assert.NoError(t, err)

	testString := "hello jim"
	encryptString, err := RsaEncrypt([]byte(testString))
	assert.NoError(t, err)

	decryptBytes, err := RsaDecrypt(encryptString)
	assert.NoError(t, err)
	assert.Equal(t, testString, string(decryptBytes))

}
