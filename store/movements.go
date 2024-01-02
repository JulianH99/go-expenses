package store

import (
	"database/sql"
	"fmt"
	"github.com/julianh99/goexpenses/model"
	_ "github.com/lib/pq"
	"log"
)

type Store interface {
	Initialize()
	GetMovements() ([]*model.Movement, error)
	CreateMovement(movement model.Movement) error
}

type PostgresOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

type PostgresStore struct {
	db *sql.DB
}

func CreatePostgresStore(options PostgresOptions) PostgresStore {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", options.User, options.DbName, options.Password, options.Host))
	if err != nil {
		log.Fatal(err)
	}

	return PostgresStore{db: db}

}

func (s PostgresStore) Initialize() {
	query := `
    create table if not exists movements (
      id SERIAL,
      type int,
      date timestamp,
      value int
    )
  `

	_, err := s.db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

}

func (m PostgresStore) GetMovements() ([]*model.Movement, error) {
	rows, err := m.db.Query("select type, date, value from movements")

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	movements := []*model.Movement{}

	for rows.Next() {
		movement, err := rowToMovement(rows)
		if err != nil {
			return nil, err
		}
		movements = append(movements, movement)
	}

	return movements, nil

}
func (m PostgresStore) CreateMovement(movement model.Movement) error {

	_, err := m.db.Query("insert into movements(type, date, value) values ($1, $2, $3)", movement.MovementType.Int(), movement.Date.Format("2006-01-02"), movement.Value)

	if err != nil {
		return err
	}

	return nil
}

func rowToMovement(rows *sql.Rows) (*model.Movement, error) {

	movement := new(model.Movement)

	err := rows.Scan(
		&movement.MovementType,
		&movement.Date,
		&movement.Value,
	)

	return movement, err

}
