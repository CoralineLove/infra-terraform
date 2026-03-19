package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sts"
)

// AWSConfig represents AWS configuration
type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

// AWSClient represents an AWS client
type AWSClient struct {
	*sts.Sts
	*S3
}

// NewAWSClient returns a new AWS client instance
func NewAWSClient(config *AWSConfig) (*AWSClient, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
	}, config)
	if err!= nil {
		return nil, err
	}
	return &AWSClient{
		Sts:     sts.New(sess),
		S3:     s3.New(sess),
	}, nil
}

// GetSTSClient returns the STS client instance
func (c *AWSClient) GetSTSClient() *sts.Sts {
	return c.Sts
}

// GetS3Client returns the S3 client instance
func (c *AWSClient) GetS3Client() *s3.S3 {
	return c.S3
}

// GenerateRandomString generates a random string of specified length
func GenerateRandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent 64 unique letters (a-z, A-Z)
		letterIdxMask = 1<<letterIdxBits - 1 // All letter/number bits set to 1, all other bits 0
	)
	var result []byte
	if length > 0 {
		for i := 0; i < length; {
			if i%2 == 0 {
				result = append(result, letterBytes[rand.Int63()%int64(letterIdxMask)&letterIdxMask])
			} else {
				result = append(result, letterBytes[rand.Int63()%int64(letterIdxMask)>>letterIdxBits])
			}
			i++
		}
	}
	return string(result)
}

// GetEnvironmentVariable returns the value of the specified environment variable
func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

// GetAbsolutePath returns the absolute path of the specified file
func GetAbsolutePath(file string) (string, error) {
	return filepath.Abs(file)
}

// LogError logs an error message
func LogError(err error) {
	log.Printf("Error: %v\n", err)
}

// LogInfo logs an info message
func LogInfo(msg string) {
	log.Printf("Info: %v\n", msg)
}

// LogDebug logs a debug message
func LogDebug(msg string) {
	log.Printf("Debug: %v\n", msg)
}

// IsNil checks if the specified value is nil
func IsNil(value interface{}) bool {
	if value == nil {
		return true
	}
	return false
}

// GetRandomUUID returns a random UUID
func GetRandomUUID() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", GetRandomBytes(4), GetRandomBytes(2), GetRandomBytes(2), GetRandomBytes(2), GetRandomBytes(6))
}

// GetRandomBytes returns a random byte slice of specified length
func GetRandomBytes(length int) []byte {
	b := make([]byte, length)
	if _, err := rand.Read(b); err!= nil {
		panic(err)
	}
	return b
}