package image

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func uploadToS3() {
	if len(os.Args) != 3 {
		fmt.Printf("usage: %s <bucket> <filename>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	bucket := os.Args[1]
	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file", filename, err)
		os.Exit(1)
	}
	defer file.Close()

	//select Region to use.
	conf := aws.Config{Region: aws.String("us-west-2")}
	sess := session.New(&conf)
	svc := s3manager.NewUploader(sess)

	fmt.Println("Uploading file to S3...")
	_, err = svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filepath.Base(filename)),
		Body:   file,
	})
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
}

func FileUpload(r *http.Request) (string, error) {
	//this function returns the filename(to save in database) of the saved file or an error if it occurs
	r.ParseMultipartForm(32 << 20)

	//ParseMultipartForm parses a request body as multipart/form-data

	var file_name string
	var errors error

	file, handler, err := r.FormFile("file") //retrieve the file from form data
	defer file.Close()                       //close the file when we finish

	if err != nil {
		errors = err

	}
	//this is path which  we want to store the file
	f, err := os.OpenFile("./images/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {

		errors = err
	}
	file_name = handler.Filename
	defer f.Close()
	io.Copy(f, file)
	//here we save our file to our path

	return file_name, errors

}

func PostImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName, err := FileUpload(r)
	//here we call the function we made to get the image and save it
	if err != nil {
		panic("no image found")
		//checking whether any error occurred retrieving image
	}
	fmt.Println(imageName)
}
