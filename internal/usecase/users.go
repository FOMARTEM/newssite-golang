package usecase

import (
	"github.com/FOMARTEM/newssite-golang/internal/entities"
)

// CreateUser
func (u *Usecase) CreateUser(user entities.User) (*entities.User, error) {
	if user, err := u.p.SelectUserByEmail(user.Email); err != nil {
		return nil, err
	} else if user != nil {
		return nil, entities.ErrUserEmailConflict
	}

	createdUser, err := u.p.InsertUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil

}

// CheckPasswordUser
func (u *Usecase) CheckPasswordUser(user entities.User) (*bool, error) {
	login := false

	password, err := u.p.SelectUserPasswordByEmail(user.Email)
	if err != nil {
		return &login, err
	} else if password == nil {
		return &login, entities.ErrUserNotFound
	}

	if user.Password == *password {
		login = true
	} else {
		err = entities.ErrUserLoginConflict
	}

	return &login, err
}

// SelectUserByID
func (u *Usecase) SelectUserByID(id int) (*entities.User, error) {
	user, err := u.p.SelectUserById(id)
	if err != nil {
		return nil, err
	} else if user == nil {
		return user, entities.ErrUserNotFound
	}

	return user, nil
}

// SelectUserByEmail
func (u *Usecase) SelectUserByEmail(email string) (*entities.User, error) {
	user, err := u.p.SelectUserByEmail(email)
	if err != nil {
		return nil, err
	} else if user == nil {
		return user, entities.ErrUserNotFound
	}

	return user, nil
}

// UpdateAdminRules
func (u *Usecase) UpdateAdminRules(email string, admin int) (*bool, error) {
	update := false

	err := u.p.UpdateUserAdminRulesByEmail(email, admin)

	if err != nil {
		return &update, err
	}

	update = true

	return &update, err
}

// UpdateUser
func (u *Usecase) UpdateUser(user entities.User) (*entities.User, error) {
	newUser, err := u.p.UpdateUserById(user)
	if err != nil {
		return nil, err
	} else if newUser == nil {
		return newUser, entities.ErrUserNotFound
	}

	return newUser, nil
}

// DeleteUserById
func (u *Usecase) DeleteUserById(id int) error {
	err := u.p.DeleteUserById(id)

	if err != nil {
		return err
	}

	return nil
}

// DeleteUserByEmail
func (u *Usecase) DeleteUserByEmail(email string) error {
	err := u.p.DeleteUserByEmail(email)

	if err != nil {
		return err
	}

	return nil
}
