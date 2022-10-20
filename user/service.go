package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(email string) (bool, error)
	GetUserbyId(id int) (User, error)
	ServiceChangeName(id int, input ChangeNameInput) (User, error)
	ChangeEmailService(id int, input ChangeEmailInput) (User, error)
	GetAllUsers() ([]User, error)
	ChangePassword(id int, input ChangePassword) (bool, error)
	Delete(input DeleteInput) (bool, error)
	ChangeDetailService(input ChangeDetailInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Active = 1
	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(Password)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil

}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, nil
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) ChangeDetailService(input ChangeDetailInput) (User, error) {
	user, err := s.repository.FindById(input.ID)
	if err != nil {
		return user, err
	}

	if input.Password != "" {
		Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return user, err
		}
		user.Password = string(Password)
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Active = input.Active

	changedDetail, err := s.repository.Update(user)
	if err != nil {
		return changedDetail, err
	}
	return changedDetail, nil

}

// mapping struct input ke struct user
// simpan struct user ke responsitory

func (s *service) IsEmailAvailable(email string) (bool, error) {

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserbyId(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.AllUser()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *service) ServiceChangeName(id int, input ChangeNameInput) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	updatedName, err := s.repository.Update(user)
	if err != nil {
		return updatedName, err
	}

	return updatedName, nil

}

func (s *service) ChangeEmailService(id int, input ChangeEmailInput) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Email = input.Email
	changedEmail, err := s.repository.Update(user)
	if err != nil {
		return changedEmail, err
	}
	return changedEmail, nil
}

func (s *service) ChangePassword(id int, input ChangePassword) (bool, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return false, err
	}
	Password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	user.Password = string(Password)
	_, err = s.repository.Update(user)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (s *service) Delete(input DeleteInput) (bool, error) {
	user, err := s.repository.FindById(input.ID)
	if err != nil {
		return false, err
	}

	_, err = s.repository.Delete(input.ID, user)
	if err != nil {
		return false, err
	}
	return true, nil
}
