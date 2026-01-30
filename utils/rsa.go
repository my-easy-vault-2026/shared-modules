package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 生成RSA 密钥
func GenRsaKey() ([]byte, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, nil, err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)

	if err != nil {
		return nil, nil, err
	}

	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prikey := pem.EncodeToMemory(block)
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey := pem.EncodeToMemory(block)
	return prikey, pubkey, nil
}

// Rsa Sign with Sha256
func RsaSignWithSha256(data []byte, key []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)

	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		priKey8, err2 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err2 != nil {
			return nil, err2
		}
		priKey = priKey8.(*rsa.PrivateKey)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// Rsa verify with Sha256
func RsaVerifySignWithSha256(data []byte, sign []byte, key []byte) (bool, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return false, errors.New("decode public key error")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], sign)
	if err != nil {
		if err.Error() == rsa.ErrVerification.Error() {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// 公钥加密
func RsaEncrypt(data, keyBytes []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("decode public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) ([]byte, error) {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("decode private key error")
	}
	//解析PKCS1格式的私钥
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		priKey8, err2 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err2 != nil {
			return nil, err2
		}
		priKey = priKey8.(*rsa.PrivateKey)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priKey, ciphertext)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 私钥解密
func RsaDecryptOAEP(ciphertext, keyBytes []byte) ([]byte, error) {

	//https://gist.github.com/jemygraw/4da51d58b349bfddd0c5
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, errors.New("decode private key error")
	}
	//解析PKCS1格式的私钥
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		priKey8, err2 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err2 != nil {
			return nil, err2
		}
		priKey = priKey8.(*rsa.PrivateKey)
	}

	// 解密
	data, err := priKey.Decrypt(nil, ciphertext, &rsa.OAEPOptions{
		Hash:    crypto.SHA1,
		MGFHash: crypto.SHA1,
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func RsaPublicKeyByteToPem(pubkey []byte) ([]byte, error) {
	pub, err := x509.ParsePKIXPublicKey(pubkey)
	if err != nil {
		return nil, err
	}

	derPkix, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	keyPem := pem.EncodeToMemory(block)
	return keyPem, nil
}
