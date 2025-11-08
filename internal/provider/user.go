package provider

import (
	"database/sql"
	"errors"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
)

// Функции с таблицей user
// создание пользователя
func (p *Provider) InsertUser(user entities.User) (*entities.User, error) {
	var id int

	err := p.conn.QueryRow(
		`CALL create_user($1, $2, $3, n_id := NULL)`,
		user.Name, user.Email, user.Password,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// поиск пользователя по email
func (p *Provider) SelectUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := p.conn.QueryRow(
		`SELECT * FROM get_user(p_email => $1)`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.AdminRole)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// поиск пользователя по id
func (p *Provider) SelectUserById(id int) (*entities.User, error) {
	var user entities.User

	err := p.conn.QueryRow(
		`SELECT * FROM get_user(p_id => $1)`,
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.AdminRole)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// Получение прав пользователя по email
func (p *Provider) SelectUserRulesByEmail(email string) (*int, error) {
	var rules int

	err := p.conn.QueryRow(
		`SELECT * FROM get_admin(p_email => $1)`,
		email,
	).Scan(&rules)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &rules, nil
}

// Получение прав пользователя по id
func (p *Provider) SelectUserRulesById(id int) (*int, error) {
	var rules int

	err := p.conn.QueryRow(
		`SELECT * FROM get_admin(p_id => $1)`,
		id,
	).Scan(&rules)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &rules, nil
}

// получение password по email
func (p *Provider) SelectUserPasswordByEmail(email string) (*string, error) {
	var password string

	err := p.conn.QueryRow(
		`SELECT * FROM get_password(p_email => $1);`,
		email,
	).Scan(&password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &password, nil
}

// редактирование пользователя
// все данные
func (p *Provider) UpdateUserById(user entities.User) (*entities.User, error) {
	_, err := p.conn.Exec(
		`CALL update_user($1, $2, $3)`,
		user.ID, user.Name, user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// обновление статуса admin по email
func (p *Provider) UpdateUserAdminRulesByEmail(email string, adminRole int) error {
	_, err := p.conn.Exec(
		`CALL update_user($1, $2)`,
		email, adminRole,
	)

	if err != nil {
		return err
	}

	return nil
}

// Функции удаления пользователя
// Удаление по id
func (p *Provider) DeleteUserById(id int) error {
	_, err := p.conn.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.ErrUserNotFound
		}

		return err
	}

	return nil
}

// Удаление по email
func (p *Provider) DeleteUserByEmail(email string) error {
	_, err := p.conn.Exec("DELETE FROM users WHERE email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.ErrUserNotFound
		}

		return err
	}

	return nil
}
