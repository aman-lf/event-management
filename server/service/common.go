package service

import (
	"context"

	"github.com/aman-lf/event-management/database"
)

func DeleteItem(ctx context.Context, entity interface{}, id int) (bool, error) {
	result := database.DB.Delete(&entity, id)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
