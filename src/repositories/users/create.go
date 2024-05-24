package userrepository

type ParamsCreateUser struct {
	Username string
	Email    string
	Password string
	Role     string
}

func (i sUserRepository) Create(p *ParamsCreateUser) (int, error) {
	var userId int
	err := i.DB.QueryRow("INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id;", p.Username, p.Email, p.Password, p.Role).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
