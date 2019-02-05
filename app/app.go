package app

import (
	"encoding/gob"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/sessions"
)

var (
	Store *sessions.FilesystemStore
	svc   *s3.S3
)

func Init() error {
	Store = sessions.NewFilesystemStore("", []byte("ssomething-very-secret"))
	Store.Options = &sessions.Options{
		MaxAge: 86400,
	}
	gob.Register(map[string]interface{}{})

	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_ACCESS_SECRET"), os.Getenv("AWS_TOKEN"))

	_, err := creds.Get()
	if err != nil {
		// handle error
	}
	cfg := aws.NewConfig().WithRegion("us-west-2").WithCredentials(creds)
	svc = s3.New(session.New(), cfg)

	return nil
}
