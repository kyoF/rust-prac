package usecase

import (
	"context"
	"dddWithJWT/pkg/domain/model"
	"dddWithJWT/pkg/domain/repository"
	"dddWithJWT/pkg/myerror"
	"dddWithJWT/pkg/utils"
	"errors"
	"time"
)

type UseCase interface {
	Signup(c context.Context, username, email, password string) (*model.User, error)
	Login(c context.Context, email, password string) (string, *model.User, error)
}

type useCase struct {
	repository repository.Repository
	timeout    time.Duration
}

func NewUseCase(userRepo repository.Repository) UseCase {
	return &useCase{
		repository: userRepo,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (uc *useCase) Signup(c context.Context, username, email, password string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()

	existUser, err := uc.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}
	if existUser.ID != 0 {
		return nil, &myerror.BadRequestError{Err: errors.New("user already exists")}
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, &myerror.InternalServerError{Err: err}
	}

	u := &model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	user, err := uc.repository.CreateUser(ctx, u)
	if err != nil {
		return nil, &myerror.InternalServerError{Err: err}
	}

	return user, nil
}

func (uc *useCase) Login(c context.Context, email, password string) (string, *model.User, error) {
	ctx, cancel := context.WithTimeout(c, uc.timeout)
	defer cancel()

	user, err := uc.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", nil, &myerror.InternalServerError{Err: err}
	}
	if user.ID == 0 {
		return "", nil, &myerror.BadRequestError{Err: errors.New("user is not exist")}
	}

	err = utils.CheckPassword(user.Password, password)
	if err != nil {
		return "", nil, &myerror.BadRequestError{Err: errors.New("password is incorrect")}
	}

	signedString, err := utils.GenerateSignedString(user.ID, user.Username)
	if err != nil {
		return "", nil, &myerror.InternalServerError{Err: err}
	}

	return signedString, user, nil
}
