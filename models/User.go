package models

import (
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID 				uuid.UUID
	Username 		string
	Password 		string
	Ctime 			time.Time
	LastLogin		time.Time
	MasterData		Master_data
}

