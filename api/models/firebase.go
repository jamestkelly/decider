package models

/*
FirebaseConfig
Interface to represent Firebase configuration details structure.
*/
type FirebaseConfig struct {
	ProjectID     string
	APIKey        string
	AuthDomain    string
	DatabaseURL   string
	StorageBucket string
	MessagingID   string
	AppID         string
}
