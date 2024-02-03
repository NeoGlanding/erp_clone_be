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

func filename() string {
	env := config.AppConfig.ENV

	if env == "local" {
		return "authentication/firebase.key.json"
	} else if env == "dev" {
		return "authentication/firebase.key.dev.json" 
	} else if env == "uat" {
		return "authentication/firebase.key.uat.json"
	} else if env == "prd" {
		return "authentication/firebase.key.prd.json"
	}
	return ""
}

func FirebaseInit() {
	opt := option.WithCredentialsFile(filename())
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