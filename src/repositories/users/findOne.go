package userrepository

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	"github.com/Dwibi/beli-mang/src/entities"
)

type FindOneParams struct {
	Username string
	Role     string
}

func (i sUserRepository) FindOne(p *FindOneParams) (*entities.User, error) {
	query := "SELECT id, password FROM users WHERE 1=1"

	conditions := []string{}
	params := []interface{}{}

	if p.Username != "" {
		conditions = append(conditions, "username = $"+strconv.Itoa(len(conditions)+1))
		params = append(params, p.Username)
	}
	if p.Role != "" {
		conditions = append(conditions, "role = $"+strconv.Itoa(len(conditions)+1))
		params = append(params, p.Role)
	}

	if len(conditions) > 0 {
		query += " AND "
	}

	query += strings.Join(conditions, " AND ")

	// log.Println("query : " + query)

	var user entities.User
	err := i.DB.QueryRow(query, params...).Scan(&user.Id, &user.Password)

	if err != nil {
		log.Println(err)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
