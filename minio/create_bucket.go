package minio

import (
	"context"
	"fmt"
	min "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Client *min.Client
}

var (
	endpoint        string
	accessKeyID     string
	secretAccessKey string
)

func GetMinio() (*Minio, error) {
	var err error
	minioCl := &Minio{}

	minioCl.Client, err = min.New(endpoint, &min.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		return minioCl, err
	}

	return minioCl, nil
}

func (m *Minio) CreateBucketWithPolicy(bucketName string) error {
	// Проверяем, существует ли уже bucket
	exists, err := m.Client.BucketExists(context.Background(), bucketName)
	if err != nil {
		return err
	}

	// Если bucket уже существует, возвращаем ошибку
	if exists {
		fmt.Println("Bucket '%s' already exists", bucketName)
		return fmt.Errorf("Bucket '%s' already exists", bucketName)
	}

	// Создаем новый bucket
	err = m.Client.MakeBucket(
		context.Background(), bucketName,
		min.MakeBucketOptions{})
	if err != nil {
		return err
	}

	// Устанавливаем политику доступа для нового bucket (public-read)
	//policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::` + bucketName + `/*"]}]}`
	policy := `{
   "Version":"2012-10-17",
   "Statement":[
      {
         "Effect":"Allow",
         "Principal":{
            "AWS":["*"]
         },
         "Action":["s3:GetObject", "s3:PutObject"],
         "Resource":["arn:aws:s3:::` + bucketName + `/*"]
      }
   ]
}`

	err = m.Client.SetBucketPolicy(context.Background(), bucketName, policy)
	if err != nil {
		return err
	}

	return nil
}
