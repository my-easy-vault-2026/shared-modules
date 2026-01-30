package utils

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func TestAWSVerify(t *testing.T) {
	message := "This is a secret"
	const keyID = "arn:aws:kms:ap-southeast-1:411650974882:key/3cd3a589-3449-435c-a350-68067b5ed466"

	// 設定 AWS 凭证
	creds := credentials.NewStaticCredentials("AKIAV7WCPMSRFBT6GOOZ", "VIt0UkQXDnnP1mlO0y9vlcGiTHaW40YcQHAhmHU8", "")

	// 建立 AWS Session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: creds,
	})

	if err != nil {
		fmt.Printf("error: %s", err)
	}

	// 創建 KMS 客戶端
	svc := kms.New(sess)

	// 簽名訊息
	signMethod := &kms.SignInput{
		KeyId:            aws.String(keyID),
		Message:          []byte(message),
		MessageType:      aws.String("RAW"),
		SigningAlgorithm: aws.String("RSASSA_PKCS1_V1_5_SHA_256"),
	}

	signature, err := svc.Sign(signMethod)
	if err != nil {
		fmt.Printf("Sign error: %s", err)
		return
	}

	fmt.Printf("Signature: %x\n", signature.Signature)

	// 驗證簽名
	verifyMethod := &kms.VerifyInput{
		KeyId:            aws.String(keyID),
		Message:          []byte(message),
		Signature:        signature.Signature,
		MessageType:      aws.String("RAW"),
		SigningAlgorithm: aws.String("RSASSA_PKCS1_V1_5_SHA_256"),
	}

	result, err := svc.Verify(verifyMethod)
	if err != nil {
		fmt.Printf("Verify error: %s", err)
		return
	}

	fmt.Printf("Signature valid: %v\n", *result.SignatureValid)
}
