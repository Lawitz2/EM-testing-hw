package hw6

// Gomock для генерации моков интерфейса DB
//go:generate mockgen -destination mock_db.go -package hw5 . DB

import "fmt"

// Представим, что тут описано взаимодействие с настоящей БД
type DB interface {
	AddToDb(u User) error
	DeleteFromDb(userName string) error
}

type RealDB struct {
	DBname string
}

func NewRealDB() *RealDB {
	return &RealDB{
		DBname: "Real database",
	}
}

type User struct {
	Name string
	Age  int
}

func (r *RealDB) AddToDb(u User) error {
	if u.Age < 18 {
		return fmt.Errorf("User %s is underage, can't add to %s", u.Name, r.DBname)
	}
	fmt.Printf("User %s added to %s\n", u.Name, r.DBname)
	return nil
}

func (r *RealDB) DeleteFromDb(userName string) error {
	fmt.Printf("User %s deleted from %s\n", userName, r.DBname)
	return nil
}
