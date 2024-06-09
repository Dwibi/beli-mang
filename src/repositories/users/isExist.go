package userrepository

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type IsExistParams struct {
	Username string
	Email    string
	Role     string
}

func (i sUserRepository) IsExist(p *IsExistParams) (bool, error) {
	tempQuery := "SELECT 1 FROM users WHERE 1=1"

	conditions := []string{}
	params := []interface{}{}

	if p.Username != "" {
		conditions = append(conditions, "username = $"+strconv.Itoa(len(conditions)+1))
		params = append(params, p.Username)
	}
	if p.Email != "" {
		conditions = append(conditions, "email = $"+strconv.Itoa(len(conditions)+1))
		params = append(params, p.Email)

	}
	if p.Role != "" {
		conditions = append(conditions, "role = $"+strconv.Itoa(len(conditions)+1))
		params = append(params, p.Role)
	}

	if len(conditions) > 0 {
		tempQuery += " AND "
	}

	tempQuery += strings.Join(conditions, " AND ")

	query := fmt.Sprintf("SELECT EXISTS (%v);", tempQuery)

	// log.Println("query : " + query)

	var exists bool
	err := i.DB.QueryRow(query, params...).Scan(&exists)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return exists, nil
}
