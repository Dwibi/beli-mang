package userrepository

import (
	"database/sql"
	"log"
)

func (i sUserRepository) FindUserRole(userId int) (string, error) {
	query := "SELECT role FROM users WHERE id = $1"

	var role string
	err := i.DB.QueryRow(query, userId).Scan(&role)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return role, nil
}
