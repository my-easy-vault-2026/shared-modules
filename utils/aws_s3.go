package utils

import (
	"bytes"
	"context"
	"io"
	"shared-modules/logger"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func getAwsConfig() (aws.Config, error) {
	access_key := Config.S3.AccessKey
	secret_key := Config.S3.SecretKey
	region := Config.S3.Region

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
		logger.Warnf("init config error: %v", err)
		return cfg, err
	}

	return cfg, err
}

// filePath = path + fileName  ex: report/20240101.xls
func UploadFileBytesToS3(ctx context.Context, filePath string, buffer *bytes.Buffer) error {

	sdkConfig, err := getAwsConfig()
	if err != nil {
		return err
	}

	s3Client := s3.NewFromConfig(sdkConfig)
	bucketName := Config.S3.BucketName

	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
		Body:   buffer,
	})

	return err
}

func CreateS3DownloadURL(ctx context.Context, filePath string) (string, error) {

	sdkConfig, err := getAwsConfig()
	if err != nil {
		return "", err
	}

	bucketName := Config.S3.BucketName

	s3Client := s3.NewFromConfig(sdkConfig)
	s3PresignClient := s3.NewPresignClient(s3Client)
	req, err := s3PresignClient.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
	}, s3.WithPresignExpires(5*time.Minute))
	if err != nil {
		logger.Warnf("get Download URL FAIL : %v", err)
		return "", err
	}
	return req.URL, nil
}

func DownloadFileFromS3ToBytes(ctx context.Context, filePath string) ([]byte, error) {
	sdkConfig, err := getAwsConfig()
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(sdkConfig)
	bucketName := Config.S3.BucketName

	result, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
	})
	if err != nil {
		logger.Warnf("download file from S3 failed: %v", err)
		return nil, err
	}
	defer result.Body.Close()

	// 读取所有数据到 []byte
	data, err := io.ReadAll(result.Body)
	if err != nil {
		logger.Warnf("read S3 object body failed: %v", err)
		return nil, err
	}

	return data, nil
}

// 下载S3对象并返回 *bytes.Buffer
func DownloadFileFromS3ToBuffer(ctx context.Context, filePath string) (*bytes.Buffer, error) {
	sdkConfig, err := getAwsConfig()
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(sdkConfig)
	bucketName := Config.S3.BucketName

	result, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
	})
	if err != nil {
		logger.Warnf("download file from S3 failed: %v", err)
		return nil, err
	}
	defer result.Body.Close()

	// 创建 buffer 并复制数据
	var buffer bytes.Buffer
	_, err = io.Copy(&buffer, result.Body)
	if err != nil {
		logger.Warnf("copy S3 object to buffer failed: %v", err)
		return nil, err
	}

	return &buffer, nil
}

func DownloadFileFromS3Reader(ctx context.Context, filePath string, fn func(reader io.Reader) error) error {
	sdkConfig, err := getAwsConfig()
	if err != nil {
		return err
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	bucketName := Config.S3.BucketName
	result, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filePath),
	})
	if err != nil {
		logger.Warnf("download file from S3 failed: %v", err)
		return err
	}
	defer result.Body.Close()
	return fn(result.Body)
}

func UploadFileStreamToS3(ctx context.Context, key string, writeData func(io.Writer) error, length int64, expires *time.Time) error {
	sdkConfig, err := getAwsConfig()
	if err != nil {
		return err
	}

	s3Client := s3.NewFromConfig(sdkConfig)
	bucketName := Config.S3.BucketName

	// 建立 io.Pipe 作為 stream 來源
	reader, writer := io.Pipe()

	// 在另一個 goroutine 寫入資料
	go func() {
		defer writer.Close()
		if err := writeData(writer); err != nil {
			writer.CloseWithError(err)
		}
	}()

	// 上傳 stream 至 S3
	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(key),
		Body:          reader, // stream 來源
		ContentLength: aws.Int64(length),
		Expires:       expires,
	})

	return err
}

func UploadFileToS3(ctx context.Context, key string, reader io.Reader, expires *time.Time) error {
	sdkConfig, err := getAwsConfig()
	if err != nil {
		return err
	}

	s3Client := s3.NewFromConfig(sdkConfig)
	bucketName := Config.S3.BucketName

	// 上傳 stream 至 S3
	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:  aws.String(bucketName),
		Key:     aws.String(key),
		Body:    reader, // stream 來源
		Expires: expires,
	})

	return err
}

func CreateS3DownloadURLWithExpires(ctx context.Context, key string, expires time.Duration) (string, error) {

	sdkConfig, err := getAwsConfig()
	if err != nil {
		return "", err
	}

	bucketName := Config.S3.BucketName

	s3Client := s3.NewFromConfig(sdkConfig)
	s3PresignClient := s3.NewPresignClient(s3Client)
	req, err := s3PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expires))
	if err != nil {
		logger.Warnf("get Download URL FAIL : %v", err)
		return "", err
	}
	return req.URL, nil
}

// BuildImageURL 組合圖片完整 URL
func BuildImageURL(iconPath, iconName string) string {
	baseDomain := Config.System.S3ImgDomin
	return strings.TrimRight(baseDomain, "/") + "/" + strings.Trim(iconPath, "/") + "/" + iconName
}
