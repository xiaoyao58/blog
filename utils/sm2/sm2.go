package sm2

import (
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"io/ioutil"
	"log"
)

type KeyPair struct{
	privateKey string
	publicKey string
}

func generateKeyPair() (KeyPair,error){

	priv,_:=sm2.GenerateKey()
	privateKeyByte,err1:=sm2.WritePrivateKeytoMem(priv,[]byte("hello"))
	if err1 !=nil{
		return KeyPair{},err1
	}
	privateKey := string(privateKeyByte)

	publicKeyByte,err2:=sm2.WritePublicKeytoMem(&priv.PublicKey,[]byte("hello"))
	if err2 !=nil{
		return KeyPair{},nil
	}
	publicKey := string(publicKeyByte)

	keyPair := KeyPair{
		privateKey: privateKey,
		publicKey: publicKey,
	}
	return keyPair,nil
}

func encrypt(sourceText,publicKey string) (string,error){
	public_key,err1:=sm2.ReadPublicKeyFromMem([]byte(publicKey),[]byte("hello"))
	if err1 !=nil{
		return "",err1
	}
	cipherByte,err2:=sm2.Encrypt(public_key,[]byte(sourceText))
	if err2 !=nil{
		return "",err2
	}
	return base64.StdEncoding.EncodeToString(cipherByte),nil
}

func decrypt(cipher,privateKey string) ([]byte,error){
	cipherByte,err1:=base64.StdEncoding.DecodeString(cipher)
	if err1!=nil{
		return []byte{},err1
	}
	private_key,err2:=sm2.ReadPrivateKeyFromMem([]byte(privateKey),[]byte("hello"))
	if err2!=nil{
		return []byte{},err2
	}
	plainByte,err3:=sm2.Decrypt(private_key,cipherByte)
	if err3!=nil{
		return []byte{},err3
	}
	return plainByte,nil
}

func Test(){
	keyPair,err:=generateKeyPair()
	if err!=nil{
		log.Fatal(err.Error())
	}


	publicKey:=keyPair.publicKey
	privateKey:=keyPair.privateKey

	pubErr:=ioutil.WriteFile("./public.txt",[]byte(publicKey),0644)
	priErr:=ioutil.WriteFile("./private.txt",[]byte(privateKey),0644)
	if pubErr!=nil || priErr!=nil{
		log.Fatal("文件写入错误")
	}

	cipher,err1:=encrypt(privateKey,publicKey)
	if err1!=nil{
		log.Fatal(err1.Error())
	}
	fmt.Println(publicKey)

	plainByte,err2:=decrypt(cipher,privateKey)
	if err2!=nil{
		log.Fatal(err2.Error())
	}

	plain:=string(plainByte)
	fmt.Println(plain)
}

