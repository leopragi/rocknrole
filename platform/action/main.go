package action

import (
	"github.com/google/uuid"
)

func generate() string {
	return uuid.New().String()
}

var MakeAction = BuildMakeAction(generate)
