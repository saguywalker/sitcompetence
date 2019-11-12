package db

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3 struct
type S3 struct {
	s3     *s3.S3
	bucket string
	path   string
}

// NewS3 return S3 struct with contains S3 instance and its configuration
func NewS3(config *Config) (*S3, error) {
	creds := credentials.NewStaticCredentials(config.AwsAccessKey, config.AwsSecretKey, "")
	if _, err := creds.Get(); err != nil {
		return nil, err
	}

	cfg := aws.NewConfig().WithRegion(config.AwsRegion).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	s3Struct := S3{
		s3:     svc,
		bucket: config.AwsBucket,
		path:   config.AwsPath,
	}

	return &s3Struct, nil
}

// UploadProfilePicture push a new profile image to S3
func (s *S3) UploadProfilePicture(file multipart.File, header *multipart.FileHeader, studentID string) (string, error) {
	size := header.Size
	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := s.path + studentID + filepath.Ext(header.Filename)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}

	resp, err := s.s3.PutObject(params)
	if err != nil {
		return "", err
	}

	return awsutil.StringValue(resp), nil
}
