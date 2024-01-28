package config

type firebaseConfigType struct {
	ProjectID	string
	BucketURL	string
}

var FirebaseConfig firebaseConfigType = firebaseConfigType{
	ProjectID: "",
	BucketURL: "",
}