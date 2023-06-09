package models

type ContactInfo struct {
	Email       string `bson:"email"`
	PhoneNumber string `bson:"phone_number"`
}

type User struct {
	ID          string      `bson:"_id,omitempty"`
	FirstName   string      `bson:"name"`
	LastName    string      `bson:"email"`
	ContactInfo ContactInfo `bson:"contact_info"`
}
