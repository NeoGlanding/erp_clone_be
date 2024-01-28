package initializers

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"github.com/automa8e_clone/config"
	"github.com/automa8e_clone/libs"
	"google.golang.org/api/option"
)



func FirebaseInit() {
	opt := option.WithCredentialsFile("authentication/firebase.key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	fmt.Println("Success initializing firebase app")
	libs.Firebase = app

	storage, err := app.Storage(context.Background())

	if err != nil {
		log.Fatalf("error accessing storage %v \n", err)
	}

	libs.FirebaseStorage = storage

	bucket, err := storage.Bucket(config.FirebaseConfig.BucketURL)

	if err != nil {
		log.Fatalf("error while initializing bucket %v \n",err)
	}

	libs.FirebaseStorageBucket = bucket
}