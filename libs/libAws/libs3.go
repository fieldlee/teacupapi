package libAws

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
	"net/http"
	"teacupapi/config"
)

var (
	svcS3 *s3.S3
	accessKeyID = config.GetS3Conf().AccessKeyID
	secretAccessKey = config.GetS3Conf().SecretAccessKey
	s3Region = config.GetS3Conf().S3Region
	s3Bucket = config.GetS3Conf().S3Bucket
)

func Init() error  {
	creds := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, "")
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3Region),
		Credentials: creds,
	})
	if err != nil {
		return err
	}
	svcS3 = s3.New(s)
	return nil
}

func Upload(file *multipart.FileHeader) (string,error) {
	f, err := file.Open()
	if err != nil {
		return "",err
	}
	var size int64 = file.Size
	buffer := make([]byte, size)
	f.Read(buffer)

	outObj, err := svcS3.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(s3Bucket),
		Key:                aws.String(fmt.Sprintf("teacup-upload/%s",file.Filename)),
		ACL:                aws.String("private"),
		Body:               bytes.NewReader(buffer),
		ContentLength:      aws.Int64(size),
		ContentType:        aws.String(http.DetectContentType(buffer)),
		ContentDisposition: aws.String("attachment"),
	})
	if err != nil {
		return "",err
	}
	return *outObj.ETag,nil
}


func Read(path string) ([]byte,error) {
	var buffer bytes.Buffer
	// Using the key, get the object from the bucket
	obj, err := svcS3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil,err
	}
	defer obj.Body.Close()
	scanner := bufio.NewScanner(obj.Body)
	for scanner.Scan() {
		buffer.Write(scanner.Bytes())
	}
	return buffer.Bytes(),nil
}