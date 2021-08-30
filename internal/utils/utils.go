package utils

import (
	"gopkg.in/guregu/null.v3"
)

func CreateNullString(str string) null.String {
	return null.NewString(str, true)
}