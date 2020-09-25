package sm4

import (
	"bytes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm4"
	"log"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
func generateKey() (key,iv string){
	return RandString(16),RandString(16)
}


func Test(){
	randKey,randIV := generateKey()
	// 128比特密钥
	key := []byte(randKey)
	// 128比特iv
	iv := []byte(randIV)
	data := []byte("一拳一个嘤嘤怪")
	ciphertxt,err1 := sm4Encrypt(key,iv, data)
	if err1 != nil{
		log.Fatal(err1)
	}
	plaintxt,err2 := sm4Decrypt(key,iv,ciphertxt)
	if err2 !=nil{
		log.Fatal(err2)
	}

	fmt.Printf("加密结果: %v\n", ciphertxt)

	fmt.Printf("解密结果: %v\n", string(plaintxt))
}

func sm4Encrypt(key, iv, plainText []byte) (string, error) {
	block, err := sm4.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData := pkcs5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryted := make([]byte, len(origData))
	blockMode.CryptBlocks(cryted, origData)
	cipherText:=base64.StdEncoding.EncodeToString(cryted)
	return cipherText, nil
}

func sm4Decrypt(key, iv []byte, cipherStr string ) ([]byte, error) {
	cipherText,err1:=base64.StdEncoding.DecodeString(cipherStr)
	if err1!=nil{
		return nil,err1
	}
	block, err := sm4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(origData, cipherText)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}
// pkcs5填充
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	if(length==0){
		return nil
	}
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}