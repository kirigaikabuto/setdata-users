package setdata_users

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

var marketplaceAppRepoQueries = []string{
	`CREATE TABLE IF NOT EXISTS users(
		id TEXT,
		username TEXT,
		email TEXT,
		login_type TEXT,
		first_name TEXT,
		last_name TEXT,
		PRIMARY KEY(id)
	);`,
}

type usersStore struct {
	db *sql.DB
}

func NewPostgresUsersStore(cfg PostgresConfig) (UsersStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range marketplaceAppRepoQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &usersStore{db: db}
	return store, nil
}

func (u *usersStore) Create(user *User) (*User, error) {
	result, err := u.db.Exec("INSERT INTO users (id, username, email, login_type, first_name, last_name) "+
		"VALUES ($1, $2, $3, $4, $5, $6)",
		user.Id, user.Username, user.Email, user.LoginType.ToString(), user.FirstName, user.LastName,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, errors.New("error in creating of element")
	}
	return user, nil
}

func (u *usersStore) Update(user *UserUpdate) (*User, error) {
	return nil, nil

}

func (u *usersStore) Delete(id string) error {
	return nil
}

func (u *usersStore) Get(id string) (*User, error) {
	return nil, nil
}
