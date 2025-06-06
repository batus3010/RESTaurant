package uploadprovider

import (
	"RESTaurant_v2/common"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

type awsS3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewAWSS3Provider(bucketName, region, apiKey, secret, domain string) *awsS3Provider {
	provider := &awsS3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			""),             //Token can be ignored
	})
	if err != nil {
		log.Fatalln(err)
	}
	provider.session = s3Session

	return provider
}

func (provider *awsS3Provider) SaveFileUploaded(ctx context.Context, data []byte, destination string) (*common.Image, error) {
	// TODO: uncomment for s3 configuration later
	//fileBytes := bytes.NewReader(data)
	//fileType := http.DetectContentType(data)
	//
	//_, err := s3.New(provider.session).PutObject(&s3.PutObjectInput{
	//	Bucket:      aws.String(provider.bucketName),
	//	Key:         aws.String(destination), // file path stored in s3
	//	ACL:         aws.String("private"),
	//	ContentType: aws.String(fileType),
	//	Body:        fileBytes,
	//})
	//
	//if err != nil {
	//	return nil, err
	//}

	img := &common.Image{
		Url:       fmt.Sprintf("%s/%s", provider.domain, destination),
		CloudName: "s3",
	}
	return img, nil
}
