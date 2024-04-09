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

func GetExpensesByEventId(ctx context.Context, eventId int, filter *graphModel.ExpenseFilter, pagination *graphModel.Pagination) ([]*model.Expense, error) {
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

		// Apply pagination and sort
		query = utils.ApplyPagination(query, pagination, "id", expenseSortableCol)
	}

	// Execute the query and fetch expenses
	result := query.Where("event_id = ?", eventId).Find(&expenses)

	if result.Error != nil {
		return nil, result.Error
	}

	return expenses, nil
}

func UpdateExpense(ctx context.Context, id int, input graphModel.UpdateExpense) (*model.Expense, error) {
	var expense model.Expense
	result := database.DB.First(&expense, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Update expense fields
	if input.ItemName != nil && *input.ItemName != "" {
		expense.ItemName = *input.ItemName
	}
	if input.Cost != nil {
		expense.Cost = *input.Cost
	}
	if input.Description != nil {
		expense.Description = input.Description
	}
	if input.Type != nil {
		expense.Type = *input.Type
	}
	if input.EventID != nil {
		expense.EventID = uint(*input.EventID)
	}

	// Save updated expense
	result = database.DB.Save(&expense)
	if result.Error != nil {
		return nil, result.Error
	}

	return &expense, nil
}

func DeleteExpense(ctx context.Context, id int) (bool, error) {
	result := database.DB.Delete(&model.Expense{}, id)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}

func GetExpenseReport(ctx context.Context, eventId int) (*graphModel.ExpenseReport, error) {
	var totalExpenses int
	expensesByCategory := []*graphModel.ExpenseCategory{}

	// Retrieve total expenses for the event
	result := database.DB.Model(&model.Expense{}).Where("event_id = ?", eventId).Select("COALESCE(SUM(cost), 0)").Scan(&totalExpenses)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retrieve expenses breakdown by category
	rows, err := database.DB.Model(&model.Expense{}).Where("event_id = ?", eventId).Select("type, SUM(cost) as total_cost").Group("type").Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		var cost int
		if err := rows.Scan(&category, &cost); err != nil {
			return nil, err
		}
		expensesByCategory = append(expensesByCategory, &graphModel.ExpenseCategory{
			Category: category,
			Cost:     cost,
		})
	}

	return &graphModel.ExpenseReport{
		TotalExpenses:      totalExpenses,
		ExpensesByCategory: expensesByCategory,
	}, nil
}
