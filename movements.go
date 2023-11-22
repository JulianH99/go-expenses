package main

type Store interface {
	initialize()
	setInitialValue(initialValue InitialValue) error
	getMovements() ([]Movement, error)
	createMovement(movement Movement) error
}

type MysqlStore struct {
	host     string
	port     int
	user     string
	password string
}

func (mysql MysqlStore) initalize() {

}

func (m MysqlStore) getMovements() ([]Movement, error) {
	return nil, nil

}
func (m MysqlStore) createMovement(movement Movement) error {
	return nil
}
