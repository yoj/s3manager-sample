package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/yoj/s3-uploader/s3"
)

const (
	region = endpoints.ApNortheast1RegionID
	dir    = ""
	backet = ""
)

func main() {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region)}),
	)

	// S3 Managerの新規作成
	uploader := s3.NewUploader(sess, dir, backet)

	// S3へのアップロード処理
	uploader.UploadInDir()

}
