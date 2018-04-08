/*
 * user's master data (german: Stammdaten) like address etc.
 */ 

package models

import (
	"github.com/satori/go.uuid"
)

type master_data struct {
	ID				uuid
	FirstName		string
	LastName		string
	NameAddition	string
	Street			string
	Streetno		string
	Postcode		string
	City			string
	Countrycode		string
	Contact			contacttype[]
}
