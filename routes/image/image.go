package image

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dsummers91/go-bmdc/app"
	"github.com/dsummers91/go-bmdc/database"
	"github.com/mongodb/mongo-go-driver/bson"
)

func uploadToS3(file io.Reader, filename string, username string) {
	bucket := os.Getenv("S3_BUCKET")

	//select Region to use.
	conf := aws.Config{Region: aws.String("us-west-2")}
	sess := session.New(&conf)
	svc := s3manager.NewUploader(sess)

	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("/public/assets/profile/" + username + filepath.Ext(filename)),
		Body:   file,
	})

	if err != nil {
		fmt.Println("error", err)
	}
	collection, context, cancel := database.Collection("members")
	defer cancel()
	_, err = collection.UpdateOne(context, bson.M{"oauth": username}, bson.M{"$set": bson.M{"image": result.Location}})

}

func FileUpload(r *http.Request, user string) (string, error) {
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

	uploadToS3(file, handler.Filename, user)
	return file_name, errors

}

func PostImageHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")
	oauthProfile := session.Values["profile"]
	oauthObject := oauthProfile.(map[string]interface{})
	user := oauthObject["sub"].(string)

	imageName, err := FileUpload(r, user)
	//here we call the function we made to get the image and save it
	if err != nil {
		panic("no image found")
		//checking whether any error occurred retrieving image
	}
	json.NewEncoder(w).Encode(bson.M{"status": 1, "name": imageName})
}
