package s3

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader struct {
	manager s3manager.Uploader
	src     string
	bucket  string
}

func NewUploader(s *session.Session, src string, bucket string) *Uploader {

	return &Uploader{
		manager: *s3manager.NewUploader(s),
		src:     src,
		bucket:  bucket,
	}
}

func (u Uploader) UploadInDir() error {

	// ディレクトリか判定
	finfo, _ := os.Stat(u.src)
	if finfo.IsDir() == false {
		return errors.New("not Dir")
	}

	files, err := ioutil.ReadDir(u.src)
	if err != nil {
		return err
	}

	for _, file := range files {

		body, _ := os.Open(u.src + "/" + file.Name())
		u.manager.Upload(&s3manager.UploadInput{
			Bucket: aws.String(u.bucket),
			Key:    aws.String(file.Name()),
			Body:   body,
		})
	}

	if err != nil {
		fmt.Println("error")
	}

	return nil
}
