package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader struct {
	manager s3manager.Uploader
	src     string
	dest    string
}

func NewUploader(s *session.Session, src string) *Uploader {

	return &Uploader{
		manager: *s3manager.NewUploader(s),
		src:     src,
	}
}

func (u Uploader) Upload() error {

	// TODO アップリード処理実装
	result, err := u.manager.Upload(&s3manager.UploadInput{
		Bucket: "kari",
		Key:    "kari",
		Body:   "klari",
	})

	if err != nil {
		fmt.Println("error")
	}

	return nil
}
