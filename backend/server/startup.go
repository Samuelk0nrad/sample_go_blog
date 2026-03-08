package server

import (
	"log"

	"simple_blog_api/config"

	"github.com/minio/minio-go/v7"
)

type App struct {
	Env *config.Env

	MinioClient *minio.Client
}

func Start() {
	env := config.NewEnv(".env", true)

	minioClient := CreateMinioClient(env)

	app := &App{
		Env:         env,
		MinioClient: minioClient,
	}

	r := NewRouter(app)

	log.Println("server running on :8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
