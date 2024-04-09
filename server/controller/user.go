package controller

import (
	"context"
	"strconv"

	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/service"
)

func GetUsersHandler(ctx context.Context, filter *graphModel.UserFilter, pagination *graphModel.Pagination) ([]*graphModel.User, error) {
	users, err := service.GetUsers(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	var returnUser []*graphModel.User
	for _, user := range users {
		userID := strconv.FormatUint(uint64(user.ID), 10)
		returnUser = append(returnUser, &graphModel.User{
			ID:      userID,
			Name:    user.Name,
			Email:   user.Email,
			PhoneNo: *user.PhoneNo,
		})
	}

	return returnUser, nil
}

func GetUserByIdHandler(ctx context.Context, id int) (*graphModel.User, error) {
	user, err := service.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	userID := strconv.FormatUint(uint64(user.ID), 10)
	returnUser := graphModel.User{
		ID:      userID,
		Name:    user.Name,
		Email:   user.Email,
		PhoneNo: *user.PhoneNo,
	}

	return &returnUser, nil
}

func CreateUserHandler(ctx context.Context, input graphModel.NewUser) (*graphModel.User, error) {
	user, err := service.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	userID := strconv.FormatUint(uint64(user.ID), 10)
	return &graphModel.User{
		ID:      userID,
		Name:    user.Name,
		Email:   user.Email,
		PhoneNo: *user.PhoneNo,
	}, nil
}
