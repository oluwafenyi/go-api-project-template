package db

import (
	"fmt"
)

// User ...
type User struct {
	tableName struct{} `pg:"users"`
	UUID      string   `pg:"id,pk,default:gen_random_uuid()"`
	Username  string   `pg:"username"`
	Email     string   `pg:"email"`
}

func (u User) String() string {
	return fmt.Sprintf("<User %s>", u.UUID)
}

// Insert ...
func (u *User) Insert() error {
	_, err := DB.Model(u).Insert()
	return err
}
