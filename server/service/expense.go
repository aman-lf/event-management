package service

import (
	"context"
	"strings"

	"github.com/aman-lf/event-management/database"
	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/model"
	"github.com/aman-lf/event-management/utils"
)

var expenseSortableCol = []string{"item_name", "cost", "type", "event_id"}

func CreateExpense(ctx context.Context, input graphModel.NewExpense) (*model.Expense, error) {
	expense := model.Expense{
		ItemName:    input.ItemName,
		Cost:        input.Cost,
		Description: input.Description,
		Type:        input.Type,
		EventID:     uint(input.EventID),
	}
	result := database.DB.Create(&expense)
	if result.Error != nil {
		return nil, result.Error
	}

	return &expense, nil
}

func GetExpenses(ctx context.Context, filter *graphModel.ExpenseFilter, pagination *graphModel.Pagination) ([]*model.Expense, error) {
	var expenses []*model.Expense

	query := database.DB.Model(&model.Expense{})

	if filter != nil {
		if filter.ID != nil && *filter.ID != "" {
			query = query.Where("id = ?", *filter.ID)
		}
		if filter.ItemName != nil && *filter.ItemName != "" {
			query = query.Where("LOWER(item_name) ILIKE ?", "%"+strings.ToLower(*filter.ItemName)+"%")
		}
		if filter.Type != nil && *filter.Type != "" {
			query = query.Where("LOWER(type) ILIKE ?", "%"+strings.ToLower(*filter.Type)+"%")
		}
		if filter.EventID != nil && *filter.EventID != "" {
			query = query.Where("event_id = ?", *filter.EventID)
		}

		// Apply pagination and sort
		query = utils.ApplyPagination(query, pagination, "id", expenseSortableCol)
	}

	// Execute the query and fetch expenses
	result := query.Find(&expenses)

	if result.Error != nil {
		return nil, result.Error
	}

	return expenses, nil
}
