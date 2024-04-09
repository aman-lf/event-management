package controller

import (
	"context"
	"strconv"

	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/service"
)

func GetExpenseByEventIdHandler(ctx context.Context, eventIdStr string, filter *graphModel.ExpenseFilter, pagination *graphModel.Pagination) ([]*graphModel.Expense, error) {
	eventId, _ := strconv.Atoi(eventIdStr)
	err := hasAdminOrContributerAccess(ctx, eventId)
	if err != nil {
		return nil, err
	}

	expenses, err := service.GetExpensesByEventId(ctx, eventId, filter, pagination)
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
	err := hasAdminAccess(ctx, input.EventID)
	if err != nil {
		return nil, err
	}

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
	err := hasAdminAccess(ctx, id)
	if err != nil {
		return nil, err
	}

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
	err := hasAdminAccess(ctx, id)
	if err != nil {
		return false, err
	}

	return service.DeleteExpense(ctx, id)
}

func GetExpenseReportHandler(ctx context.Context, idStr string) (*graphModel.ExpenseReport, error) {
	id, _ := strconv.Atoi(idStr)
	err := hasAdminOrContributerAccess(ctx, id)
	if err != nil {
		return nil, err
	}
	return service.GetExpenseReport(ctx, id)
}
