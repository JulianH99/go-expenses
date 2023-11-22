package main

import "time"

type MovementType string

const (
	Expense MovementType = "Expense"
	Income  MovementType = "Income"
)

func (m MovementType) Int() int {
	switch m {
	case Expense:
		return 0
	case Income:
		return 1
	}

	return -1
}

type InitialValue struct {
	value int
}

type Movement struct {
	Value        int
	Date         time.Time
	MovementType MovementType
}
