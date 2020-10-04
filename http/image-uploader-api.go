package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	// http -f POST :8080/files imageFile@~/environment/GoProject/http/elpis_logo_small.jpg
	http.HandleFunc("/files", PostFile)
	// http  DELETE :8080/files/delete fileName=elpis_logo_small.jpg
	http.HandleFunc("/files/delete", DeleteFile)

	http.ListenAndServe(":8080", nil)
}

type ImageFile struct {
	Id   string
	Name string
	Url  string
}

type InputJsonSchema struct {
	FileName string `json:"fileName"`
}

const AWS_S3_BUCKET string = "image-uploader-golang"

func PostFile(w http.ResponseWriter, r *http.Request) {

	// Formからファイル情報を取得
	file, fileHeader, err := r.FormFile("imageFile")
	if err != nil {
		log.Fatal(err)
	}

	// S3へのアップローダーを取得
	uploader := getUploader()
	// ファイルのアップロード
	fileUpload(uploader, file, fileHeader)

}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

	// request bodyの読み取り
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("io error")
		return
	}

	// jsonのdecode
	jsonBytes := ([]byte)(body)
	data := new(InputJsonSchema)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Fatal("JSON Unmarshal error: ", err)
		return
	}
	log.Printf("JSONデータ: %s", data)

	// 指定したファイルを削除
	fileDelete(data.FileName)
}

func getUploader() *s3manager.Uploader {

	// sessionの作成
	// オプションの「SharedConfigState」に「SharedConfigEnable」を設定することで、~/.aws/config内を参照してくれる
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return s3manager.NewUploader(sess)
}

func fileUpload(uploader *s3manager.Uploader, file multipart.File, fileHeader *multipart.FileHeader) {
	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(fileHeader.Filename),
		Body:   file,
	})
	if err != nil {
		log.Fatal("Unable to uploaded. reason: ", err)
	}

	log.Println("Sucessfully uploaded!")
	log.Printf("file path: %s", res.Location)
	log.Printf("file name: %s", fileHeader.Filename)
}

func fileDelete(fileName string) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// オブジェクト削除
	svc := s3.New(sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(fileName),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(AWS_S3_BUCKET),
		Key:    aws.String(fileName),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Object is deleted.")
}
