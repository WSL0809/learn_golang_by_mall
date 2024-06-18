package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"log"
)

var Encrypt *Encryption

type Encryption struct {
	key string
}

func init() {
	Encrypt = NewEncryption()
}
func NewEncryption() *Encryption {
	return &Encryption{}
}
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)
	srcByte = append(srcByte, ret...)
	return srcByte
}
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)
	bLock, err := aes.NewCipher([]byte(k.key))
	if err != nil {
		log.Fatal("Error creating AES cipher: ", err)
	}
	NewSrcByte := PadPwd(srcByte, bLock.BlockSize()) // 填充
	dst := make([]byte, len(NewSrcByte))
	bLock.Encrypt(dst, NewSrcByte)

	pwd := base64.StdEncoding.EncodeToString(dst)
	return pwd
}
func UnPadPwd(dst []byte) ([]byte, error) {
    if len(dst) <= 0 {
        return dst, errors.New("长度有误")
    }

    unpadNum := int(dst[len(dst)-1]) // 获取填充值
    strErr := "error"
    op := []byte(strErr)

    if len(dst) < unpadNum {
        return op, errors.New("填充数值错误")
    }

    str := dst[:(len(dst) - unpadNum)] // 去除填充
    return str, nil
}
func (k *Encryption) AesDecoding(pwd string) string {
    pwdByte := []byte(pwd)
    pwdByte, err := base64.StdEncoding.DecodeString(pwd)
    if err != nil {
        // 处理 Base64 解码错误
		log.Fatal("Error decoding base64: ", err)
    }

    block, errBlock := aes.NewCipher([]byte(k.key))
    if errBlock != nil {
        // 处理 AES 创建错误
		log.Fatal("Error creating AES cipher: ", errBlock)
    }

    dst := make([]byte, len(pwdByte))
    block.Decrypt(dst, pwdByte)

    dst, err = UnPadPwd(dst) // 去除填充
    if err != nil {
		// 处理去除填充错误
		log.Fatal("Error creating AES cipher: ", err)
    }

    return string(dst)
}

func (k *Encryption) SetKey (key string) {
	k.key = key
}