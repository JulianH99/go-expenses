package model

import (
	"errors"
	"fmt"
	"time"
)

type MovementType string

const (
	Expense MovementType = "Expense"
	Income  MovementType = "Income"
)

func ToMovementType(value string) MovementType {
	switch value {
	case "0":
		return Expense
	case "1":
		return Income
	}

	return ""
}

func (m *MovementType) Scan(value interface{}) error {

	if value == nil {
		*m = Expense
		return nil
	}

	if value.(int64) == 1 {
		*m = Income
		return nil
	}

	if value.(int64) == 0 {
		*m = Expense
		return nil
	}

	return errors.New(fmt.Sprintf("Value is invalid type of movement %d", m))

}

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
