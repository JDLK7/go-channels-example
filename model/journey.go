package model

type Journey struct {
	Id					int 		`json:id`
	Time				int			`json:journey_time`
	Destination	string	`json:destination`
}