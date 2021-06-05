package utils

import uuid "github.com/satori/go.uuid"

func ReturnUUidNewV4() string {
	uuid, _ := uuid.NewV4()
	return uuid.String()
}
