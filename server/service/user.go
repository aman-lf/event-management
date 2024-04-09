package service

import (
	"context"
	"strings"

	"github.com/aman-lf/event-management/database"
	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/model"
	"github.com/aman-lf/event-management/utils"
)

var userSortableCol = []string{"name", "email"}

func CreateUser(ctx context.Context, input graphModel.NewUser) (*model.User, error) {
	user := model.User{
		Name:    input.Name,
		Email:   input.Email,
		PhoneNo: input.PhoneNo,
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, result.Error
}

func GetUsers(ctx context.Context, options *graphModel.UserQueryOptions) ([]*model.User, error) {
	var users []*model.User

	query := database.DB.Model(&model.User{})

	if options != nil {
		filter := options.Filter
		if filter != nil {
			if filter.ID != nil && *filter.ID != "" {
				query = query.Where("id = ?", *filter.ID)
			}
			if filter.Name != nil && *filter.Name != "" {
				query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(*filter.Name)+"%")
			}
			if filter.Email != nil && *filter.Email != "" {
				query = query.Where("LOWER(email) ILIKE ?", "%"+strings.ToLower(*filter.Email)+"%")
			}
			if filter.PhoneNo != nil && *filter.PhoneNo != "" {
				query = query.Where("phone_no ILIKE ?", "%"+strings.ToLower(*filter.PhoneNo)+"%")
			}
		}

		// Apply sorting
		sortColumn := "id" // Default sort column
		if options.SortBy != nil && *options.SortBy != "" {
			if utils.ContainsString(userSortableCol, *options.SortBy) {
				sortColumn = *options.SortBy
			}
		}
		order := "ASC" // Default sort order
		if options.SortOrder != nil && strings.ToUpper(*options.SortOrder) == "DESC" {
			order = "DESC"
		}
		query = query.Order(sortColumn + " " + order)

		// Apply limit and offset
		query = query.Limit(*options.Limit).Offset(*options.Offset)
	}

	// Execute the query and fetch users
	result := query.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetUserById(ctx context.Context, id int) (*model.User, error) {
	var user *model.User

	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}