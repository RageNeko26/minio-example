package main

import (
	"context"
	"log"

	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	ENDPOINT   = "127.0.0.1:9000"
	ACCESS_KEY = "HE0BGWT4CTYQ90AQ9A8D"
	SECRET     = "vk7J+Bat6I6NoQi17LNiEe4zOUyv7N+6d7xYJSgD"
)

func main() {
	client, err := minio.New(ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(ACCESS_KEY, SECRET, ""),
		Secure: false,
	})

	if err != nil {
		log.Fatalf("Error initial connection: %v \n", err)
	}

	//CreateBucket(client, "test-bucket")
	//ListBucket(client)
	UploadItem(client, "bijou-lagi")
}

func ListBucket(client *minio.Client) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	buckets, _ := client.ListBuckets(ctx)

	for _, bucket := range buckets {
		log.Println("Bucket Name:", bucket.Name)
	}
}

func CreateBucket(client *minio.Client, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	err := client.MakeBucket(ctx, name, minio.MakeBucketOptions{
		Region: "us-east-1",
	})

	if err != nil {
		log.Fatalf("Failed to make bucket because error: %v \n", err)
	}

	log.Println("Bucket has been created!")
}

func UploadItem(client *minio.Client, name string) {
	file, err := os.Open("bijou.jpeg")

	if err != nil {
		log.Fatal("Failed to open file because error:", err)
	}

	stat, err := file.Stat()

	if err != nil {
		log.Fatal("Failed to get information of file:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	upload, err := client.PutObject(
		ctx,
		"test-bucket",
		name,
		file,
		stat.Size(),
		minio.PutObjectOptions{
			ContentType: "image/jpeg",
		},
	)

	if err != nil {
		log.Fatal("Failed to upload file into Storage Object:", err)
	}

	signedURL := "http://" + ENDPOINT + "/test-bucket/" + name

	log.Println("File has been uploaded:", upload.Bucket)
	log.Println("Signed URL: ", signedURL)

}
