package models

import (
	"github.com/satori/go.uuid"
)

type user struct {
	ID 				uuid
	Username 		string
	Password 		string
	Ctime 			time
	LastLogin		time
	MasterData		masterdata
}

