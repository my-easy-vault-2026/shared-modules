package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func getConfig() aws.Config {
	access_key := "AKIAV7WCPMSRJLIUZGWS"
	secret_key := "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"
	//end_point := "https://euinside-dev.s3.ap-east-1.amazonaws.com" //endpoint设置，不要动
	region := "ap-east-1"

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: access_key, SecretAccessKey: secret_key,
			},
		}),
		config.WithRegion(region),
	)

	if err != nil {
		fmt.Printf("init config error: %v", err)
		return cfg
	}

	return cfg
}

func TestS3ListBucket(t *testing.T) {

	ctx := context.Background()
	sdkConfig := getConfig()

	// 初始化 S3 客戶端
	s3Client := s3.NewFromConfig(sdkConfig)
	count := 10
	fmt.Printf("Let's list up to %v buckets for your account.\n", count)

	// 列出 S3 Buckets
	result, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		log.Fatalf("List bucket error: %v", err)
	}

	// 處理沒有 Bucket 的情況
	if len(result.Buckets) == 0 {
		fmt.Println("You don't have any buckets!")
		return
	}

	// 限制輸出 Bucket 數量
	if count > len(result.Buckets) {
		count = len(result.Buckets)
	}

	// 列出 Buckets
	fmt.Println("Your Buckets:")
	for _, bucket := range result.Buckets[:count] {
		fmt.Printf("\t%v\n", aws.ToString(bucket.Name))
	}
}

// CreateBucket creates a bucket with the specified name in the specified Region.
func TestCreateBucket(t *testing.T) {

	ctx := context.Background()
	sdkConfig := getConfig()

	bucketName := "my-valid-bucket-20241217"
	s3Client := s3.NewFromConfig(sdkConfig)

	// 檢查 Bucket 是否已存在
	_, err := s3Client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err == nil {
		fmt.Println("Bucket already exists.")
		return
	}

	result, err := s3Client.CreateBucket(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(sdkConfig.Region),
		},
	})

	if err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	} else {
		fmt.Printf("Bucket created: %s\n", aws.ToString(result.Location))
	}

	// 等待 Bucket 存在
	waiter := s3.NewBucketExistsWaiter(s3Client)
	err = waiter.Wait(ctx, &s3.HeadBucketInput{Bucket: aws.String(bucketName)}, time.Minute)
	if err != nil {
		log.Printf("Failed to wait for bucket %s to exist: %v\n", bucketName, err)
	} else {
		fmt.Printf("Bucket %s is ready to use.\n", bucketName)
	}

}

func TestUploadFile(t *testing.T) {

	ctx := context.Background()
	sdkConfig := getConfig()

	s3Client := s3.NewFromConfig(sdkConfig)
	bucketName := "euinside-dev"

	file, err := os.Open("/Users/henk/Downloads/resetPincode.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	output, err := s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("202412/resetPincode.png"),
		Body:   file, // 这里也可以使用其他 io.Reader 实例实现对数据流的上传
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("ETag: %s\n", aws.ToString(output.ETag))
}

func TestUploadFileBytes(t *testing.T) {

	ctx := context.Background()
	sdkConfig := getConfig()

	s3Client := s3.NewFromConfig(sdkConfig)
	bucketName := "euinside-dev"

	output, err := s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("202412/test.txt"),
		Body:   strings.NewReader("Hello This is Test!"),
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("ETag: %s\n", aws.ToString(output.ETag))
}

func TestCreateDownloadURL(t *testing.T) {

	sdkConfig := getConfig()

	bucketName := "euinside-dev"

	s3Client := s3.NewFromConfig(sdkConfig)
	s3PresignClient := s3.NewPresignClient(s3Client)
	req, err := s3PresignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("resetPincode.png"),
	}, s3.WithPresignExpires(5*time.Minute))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(req.URL)
}

// func TestAWSS3ListBucket(t *testing.T) {
// 	access_key := "AKIAV7WCPMSRJLIUZGWS"
// 	secret_key := "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"
// 	end_point := "https://euinside-dev.s3.ap-east-1.amazonaws.com/" //endpoint设置，不要动

// 	sess, err := session.NewSession(&aws.Config{
// 		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
// 		Endpoint:    aws.String(end_point),
// 		Region:      aws.String("ap-east-1"),
// 		// DisableSSL:       aws.Bool(true),
// 		// S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
// 	})

// 	svc := s3.New(sess)
// 	result, err := svc.ListBuckets(nil)
// 	if err != nil {
// 		fmt.Printf("Unable to list buckets, %v", err)
// 	}

// 	fmt.Println("Buckets:")

// 	for _, b := range result.Buckets {
// 		fmt.Printf("* %s created on %s\n",
// 			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
// 	}

// 	for _, b := range result.Buckets {
// 		fmt.Printf("%s\n", aws.StringValue(b.Name))
// 	}
// }

// func TestAWSS3CreateBucket(t *testing.T) {
// 	access_key := "AKIAV7WCPMSRJLIUZGWS"
// 	secret_key := "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"
// 	end_point := "https://euinside-dev.s3.ap-east-1.amazonaws.com/" //endpoint设置，不要动

// 	sess, err := session.NewSession(&aws.Config{
// 		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
// 		Endpoint:    aws.String(end_point),
// 		Region:      aws.String("ap-east-1"),
// 		// DisableSSL:       aws.Bool(true),
// 		// S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
// 	})

// 	bucketName := "TEST-HENK"
// 	svc := s3.New(sess)
// 	params := &s3.CreateBucketInput{
// 		Bucket: aws.String(bucketName),
// 	}

// 	output, err := svc.CreateBucket(params)

// 	if err != nil {
// 		fmt.Printf("Unable to create bucket %s, %v", bucketName, err)
// 	}

// 	fmt.Printf("result: %d", output)

// 	// Wait until bucket is created before finishing
// 	fmt.Printf("Waiting for bucket %s to be created...\n", bucketName)

// 	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
// 		Bucket: aws.String(bucketName),
// 	})

// 	if err != nil {
// 		fmt.Printf("Error occurred while waiting for bucket to be created, %v", bucketName)
// 	}

// 	fmt.Printf("Bucket &q successfully created : %s", bucketName)
// }

func TestDownloadImageFromS3ToBytes(t *testing.T) {
	ctx := context.Background()
	bucketName := "euinside-dev"
	filePath := "/kyc2_face_photo1/2025/5/4d7a56504d16c39ddd9bfcbc72a200de" // 使用之前上传的图片文件
	Config = &Conf{
		S3: S3Config{
			AccessKey:     "AKIAV7WCPMSRJLIUZGWS",
			SecretKey:     "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",
			Region:        "ap-east-1",
			BucketName:    bucketName,
			BaseDirectory: "files",
		},
	}

	fmt.Printf("开始下载图片文件: %s/%s\n", bucketName, filePath)

	data, err := DownloadFileFromS3ToBytes(ctx, filePath)
	if err != nil {
		log.Fatalf("下载图片失败: %v", err)
	}

	fmt.Printf("成功下载图片，大小: %d bytes\n", len(data))

	// 检查是否是 PNG 文件（PNG 文件头是 89 50 4E 47）
	if len(data) >= 4 && data[0] == 0x89 && data[1] == 0x50 && data[2] == 0x4E && data[3] == 0x47 {
		fmt.Println("✅ 确认是 PNG 格式的图片文件！")
	} else {
		fmt.Println("❌ 文件格式验证失败")
	}
}

func TestDownloadNonExistentFile(t *testing.T) {
	ctx := context.Background()
	filePath := "non-existent-file.txt"
	bucketName := "euinside-dev"
	Config = &Conf{
		S3: S3Config{
			AccessKey:     "AKIAV7WCPMSRJLIUZGWS",
			SecretKey:     "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",
			Region:        "ap-east-1",
			BucketName:    bucketName,
			BaseDirectory: "files",
		},
	}

	fmt.Printf("测试下载不存在的文件: %s/%s\n", bucketName, filePath)

	_, err := DownloadFileFromS3ToBytes(ctx, filePath)
	if err != nil {
		fmt.Printf("✅ 预期的错误: %v\n", err)
	} else {
		fmt.Println("❌ 应该返回错误，但没有返回")
	}
}

// 比较两种方法的性能（可选）
func TestCompareDownloadMethods(t *testing.T) {
	ctx := context.Background()
	filePath := "202412/test.txt"
	bucketName := "euinside-dev"
	Config = &Conf{
		S3: S3Config{
			AccessKey:     "AKIAV7WCPMSRJLIUZGWS",
			SecretKey:     "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",
			Region:        "ap-east-1",
			BucketName:    bucketName,
			BaseDirectory: "files",
		},
	}

	fmt.Println("=== 比较两种下载方法 ===")

	// 方法1: []byte
	fmt.Println("方法1: DownloadFileFromS3ToBytes")
	data, err := DownloadFileFromS3ToBytes(ctx, filePath)
	if err != nil {
		log.Printf("方法1失败: %v", err)
	} else {
		fmt.Printf("方法1成功，数据大小: %d bytes，内容: %s\n", len(data), string(data))
	}

	// 方法2: *bytes.Buffer
	fmt.Println("\n方法2: DownloadFileFromS3ToBuffer")
	buffer, err := DownloadFileFromS3ToBuffer(ctx, filePath)
	if err != nil {
		log.Printf("方法2失败: %v", err)
	} else {
		fmt.Printf("方法2成功，数据大小: %d bytes，内容: %s\n", buffer.Len(), buffer.String())
	}

	// 比较结果
	if err == nil && len(data) == buffer.Len() && string(data) == buffer.String() {
		fmt.Println("✅ 两种方法结果一致！")
	}
}
