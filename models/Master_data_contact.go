package models

import (
	"github.com/satori/go.uuid"
)

type Master_data_contact struct {
	ID				uuid.UUID
	UserID			uuid.UUID
	ContactType		uuid.UUID
	Description		string
	Value			string
	
}
