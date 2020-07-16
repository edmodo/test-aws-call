package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	f "path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	m "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/goamz/goamz/cloudfront"
)

const (
	// For S3 Download
	REGION     = "me-south-1"
	ACCESS_KEY = ""
	SECRET_KEY = ""
	BUCKET     = ""
	FILENAME   = "testfile1" // File in Bucket

	// For CloudFront
	CloudFrontBaseURL        = ""
	CloudFrontPemKeyFilePath = ""
	CloudFrontKeyPairId      = ""
	CloudFrontPath           = ""
	QueryString              = ""
)

func main() {
	fmt.Println("\n== Test S3 Download ==\n")
	testS3Download()

	fmt.Println("\n== Test CloudFront ==\n")
	//testCloudFront()
}

func testS3Download() {
	fmt.Printf("Try to download [%s] in bucket [%s] of region [%s]\n", FILENAME, BUCKET, REGION)
	s3 := NewS3Handler(ACCESS_KEY, SECRET_KEY, BUCKET)
	if s3 == nil {
		fmt.Println("Failed to create S3 Handler")
	}
	localfile, err := s3.DownloadFile("./tmp/", "testfile1")
	if err != nil {
		fmt.Printf("Failed to download file. err: %v\n", err)
	} else {
		fmt.Println("Downloaded to", localfile)
	}
}

func testCloudFront() {
	// https://godoc.org/github.com/goamz/goamz/cloudfront
	cloudFrontClient, err := NewCloudFront(CloudFrontBaseURL, CloudFrontPemKeyFilePath, CloudFrontKeyPairId)
	if err != nil {
		fmt.Printf("Failed to create cloudFrontClient. err: %v\n", err)
	}

	expires := time.Unix(100, 0)
	signedUrl, err := cloudFrontClient.CannedSignedURL(CloudFrontPath, QueryString, expires)

	if err == nil {
		fmt.Printf(signedUrl)
	} else {
		fmt.Printf("Failed to get signedURL. err: %v\n", err)
	}
}

type S3Handler struct {
	*m.Downloader
	bucket string
}

func NewS3Handler(awsKey string, awsSecret string, bucket string) *S3Handler {
	creds := credentials.NewStaticCredentials(awsKey, awsSecret, "")

	config := &aws.Config{
		Region:      aws.String(REGION),
		Credentials: creds,
	}
	sess := session.Must(session.NewSession(config))

	return &S3Handler{
		m.NewDownloader(sess),
		bucket,
	}
}

func (h *S3Handler) DownloadFile(downloadDir string, filepath string) (string, error) {
	filename := f.Base(filepath)
	localFile := downloadDir + filename
	fd, err := os.Create(localFile)
	if err != nil {
		return "", err
	}

	defer fd.Close()

	params := &s3.GetObjectInput{
		Bucket: aws.String(h.bucket),
		Key:    aws.String(filepath),
	}

	_, err = h.Downloader.Download(fd, params)

	if err != nil {
		return "", err
	}
	fmt.Printf("Downloaded file from S3: %s\n", localFile)
	return localFile, nil
}

type CloudFrontClientInterface interface {
	CannedSignedURL(path, querystrings string, expires time.Time) (string, error)
}

func NewCloudFront(baseUrl, keyFilePath, keyPairId string) (*cloudfront.CloudFront, error) {
	pemPrivateKey, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(pemPrivateKey))
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return cloudfront.New(baseUrl, key, keyPairId), nil
}

// Helper functions for encoding an s3 url when creating a presigned url
// There are 2 functions since the rules for encoding the url are different
// depending on whether you're encoding the filename or a query param.

// Helper method for encoding the filename part of a presigned url for s3
func EncodeStringForS3(toEncode string) string {
	// S3 uses query encoding for the filename in the path, except:
	// spaces must be encoded as +
	// exclamation marks stay as !
	// commas marks stay as ,

	encodedString := url.QueryEscape(toEncode)
	encodedString = strings.Replace(encodedString, "%20", "+", -1)
	encodedString = strings.Replace(encodedString, "%21", "!", -1)
	encodedString = strings.Replace(encodedString, "%2C", ",", -1)
	return encodedString
}

// Helper method for encoding the query values of a presigned url for s3
func EncodeQueryStringForS3(toEncode string) string {
	// S3 uses query encoding for the query string, except:
	// spaces must be encoded as %20

	encodedString := url.QueryEscape(toEncode)
	return strings.Replace(encodedString, "+", "%20", -1)
}
