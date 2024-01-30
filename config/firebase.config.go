package config

type firebaseConfigType struct {
	ProjectID		string
	ProjectKeyId	string
	BucketURL		string
	PrivateKey		string
}

var FirebaseConfig firebaseConfigType = firebaseConfigType{
	ProjectID: "",
	BucketURL: "",
	PrivateKey: "",
}