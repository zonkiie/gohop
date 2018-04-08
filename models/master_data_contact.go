package models

import (
	"github.com/satori/go.uuid"
)

type master_data_contact struct {
	ID				uuid
	UserID			uuid
	ContactType		uuid
	Description		string
	Value			string
	
}
