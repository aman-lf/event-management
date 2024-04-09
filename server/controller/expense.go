package controller

import (
	"context"
	"strconv"

	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/service"
)

func GetExpenseHandler(ctx context.Context, filter *graphModel.ExpenseFilter, pagination *graphModel.Pagination) ([]*graphModel.Expense, error) {
	expenses, err := service.GetExpenses(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	var returnExpenses []*graphModel.Expense
	for _, expense := range expenses {
		expenseID := strconv.FormatUint(uint64(expense.ID), 10)
		returnExpenses = append(returnExpenses, &graphModel.Expense{
			ID:          expenseID,
			ItemName:    expense.ItemName,
			Cost:        expense.Cost,
			Description: expense.Description,
			Type:        expense.Type,
			EventID:     int(expense.EventID),
		})
	}
	return returnExpenses, nil
}

func CreateExpenseHandler(ctx context.Context, input graphModel.NewExpense) (*graphModel.Expense, error) {
	expense, err := service.CreateExpense(ctx, input)
	if err != nil {
		return nil, err
	}

	expenseID := strconv.FormatUint(uint64(expense.ID), 10)
	return &graphModel.Expense{
		ID:          expenseID,
		ItemName:    input.ItemName,
		Cost:        input.Cost,
		Description: input.Description,
		Type:        input.Type,
		EventID:     int(expense.EventID),
	}, nil
}

func UpdateExpenseHandler(ctx context.Context, idStr string, input graphModel.UpdateExpense) (*graphModel.Expense, error) {
	id, _ := strconv.Atoi(idStr)
	expense, err := service.UpdateExpense(ctx, id, input)
	if err != nil {
		return nil, err
	}

	expenseID := strconv.FormatUint(uint64(expense.ID), 10)
	return &graphModel.Expense{
		ID:          expenseID,
		ItemName:    expense.ItemName,
		Cost:        expense.Cost,
		Description: expense.Description,
		Type:        expense.Type,
		EventID:     int(expense.EventID),
	}, nil
}

func DeleteExpenseHandler(ctx context.Context, idStr string) (bool, error) {
	id, _ := strconv.Atoi(idStr)
	return service.DeleteExpense(ctx, id)
}
