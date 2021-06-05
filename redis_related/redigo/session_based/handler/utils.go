package handler

import uuid "github.com/satori/go.uuid"

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json: "password"`
	Username string `json: "username"`
}

func ReturnUUidNewV4() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
