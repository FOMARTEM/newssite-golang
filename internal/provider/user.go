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
		`INSERT INTO public.users (name, email, password, admin) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.Name, user.Email, user.Password, false,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &entities.User{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Admin:    user.Admin,
	}, nil
}

// поиск пользователя по id
func (p *Provider) SelectUserById(id int) (*entities.User, error) {
	var user entities.User

	err := p.conn.QueryRow(
		`SELECT id, name, email, password, admin FROM public.users WHERE id = $1`,
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// поиск пользователя по email
func (p *Provider) SelectUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := p.conn.QueryRow(
		`SELECT id, name, email, password, admin FROM public.users WHERE email = $1`,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Admin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// получение password по email
func (p *Provider) SelectUserPasswordByEmail(email string) (*string, error) {
	var password string

	err := p.conn.QueryRow(
		`SELECT password FROM public.users WHERE email = $1`,
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
	var updateduser entities.User

	err := p.conn.QueryRow(
		`UPDATE public.users SET name=$1, email=$2, password=$3 WHERE id = $4 RETURNING name, email, password`,
		user.Name, user.Email, user.Password, user.ID,
	).Scan(&updateduser.Name, &updateduser.Email, &updateduser.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	updateduser.ID = user.ID
	updateduser.Admin = user.Admin

	return &updateduser, nil
}

// обновление статуса admin по id
func (p *Provider) UpdateUserAdminRulesById(id int, admin bool) error {
	_, err := p.conn.Query(
		`UPDATE public.users SET admin = $1 WHERE id = $2`,
		admin, id,
	)

	if err != nil {
		return err
	}

	return nil
}

// обновление статуса admin по email
func (p *Provider) UpdateUserAdminRulesByEmail(email string, admin bool) error {
	_, err := p.conn.Query(
		`UPDATE public.users SET admin = $1 WHERE email = $2`,
		admin, email,
	)

	if err != nil {
		return err
	}

	return nil
}

// проверка статуса admin по id
func (p *Provider) CheckUserIsAdminById(id int) (*bool, error) {
	var admin bool

	err := p.conn.QueryRow(
		`SELECT admin FROM public.users WHERE id = $1`,
		id,
	).Scan(&admin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &admin, nil

}

// проверка статуса admin по email
func (p *Provider) CheckUserIsAdminByEmail(email string) (*bool, error) {
	var admin bool

	err := p.conn.QueryRow(
		`SELECT admin FROM public.users WHERE email = $1`,
		email,
	).Scan(&admin)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &admin, nil
}

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