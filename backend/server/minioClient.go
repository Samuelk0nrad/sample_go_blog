package server

import (
	"log"

	"simple_blog_api/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func CreateMinioClient(env *config.Env) *minio.Client {
	endpint := env.MinioEndpoint
	accessKey := env.MinioAccessKeyID
	secretKey := env.MinioSecretAccessKey
	useSSL := env.MinioUseSSL

	minioClient, err := minio.New(endpint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal("Error creating minio client, ", err)
	}
	return minioClient
}
