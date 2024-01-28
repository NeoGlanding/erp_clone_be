package libs

import (
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	fbsStore "firebase.google.com/go/v4/storage"
)

var Firebase *firebase.App
var FirebaseStorage *fbsStore.Client
var FirebaseStorageBucket *storage.BucketHandle