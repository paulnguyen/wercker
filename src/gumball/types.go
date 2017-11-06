/*
	Gumball API in Go (Version 1)
	Basic Version with Wercker
*/

package gumball

type gumballMachine struct {
	Id            int
	CountGumballs int
	ModelNumber   string
	SerialNumber  string
}

var machine gumballMachine = gumballMachine{
	Id:            1,
	CountGumballs: 989,
	ModelNumber:   "M102988",
	SerialNumber:  "1234998871109",
}

type order struct {
	Id          string
	OrderStatus string
}

var orders map[string]order
