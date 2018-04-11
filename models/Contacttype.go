package models

import (
	"github.com/satori/go.uuid"
)

type Contacttype struct {
	ID				uuid.UUID
	Description		string
	Name			string
}

