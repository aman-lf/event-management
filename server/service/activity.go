package service

import (
	"context"
	"strings"
	"time"

	"github.com/aman-lf/event-management/database"
	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/model"
	"github.com/aman-lf/event-management/utils"
)

var activitySortableCol = []string{"name", "start_time", "end_time", "event_id"}

func CreateActivity(ctx context.Context, input graphModel.NewActivity) (*model.Activity, error) {
	startTime, err := time.Parse("2006-01-02T15:04:05", input.StartTime)
	if err != nil {
		return nil, err
	}
	endTime, err := time.Parse("2006-01-02T15:04:05", input.EndTime)
	if err != nil {
		return nil, err
	}
	activity := model.Activity{
		Name:        input.Name,
		StartTime:   &startTime,
		EndTime:     &endTime,
		Description: *input.Description,
		EventID:     uint(input.EventID),
	}
	result := database.DB.Create(&activity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &activity, nil
}

func GetActivitiesByUserId(ctx context.Context, userId int, filter *graphModel.ActivityFilter, pagination *graphModel.Pagination) ([]*model.Activity, error) {
	var activities []*model.Activity

	query := database.DB.Model(&model.Activity{}).
		Joins("JOIN participants ON participants.event_id = activity.event_id").
		Where("participants.user_id = ?", userId)

	if filter != nil {
		if filter.ID != nil && *filter.ID != "" {
			query = query.Where("id = ?", *filter.ID)
		}
		if filter.Name != nil && *filter.Name != "" {
			query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(*filter.Name)+"%")
		}
		if filter.StartTime != nil && *filter.StartTime != "" {
			query = query.Where("start_time = ?", *filter.StartTime)
		}
		if filter.EndTime != nil && *filter.EndTime != "" {
			query = query.Where("end_time = ?", *filter.EndTime)
		}
		if filter.EventID != nil && *filter.EventID != "" {
			query = query.Where("event_id = ?", *filter.EventID)
		}

		// Apply pagination and sort
		query = utils.ApplyPagination(query, pagination, "id", activitySortableCol)
	}

	// Execute the query and fetch activities
	result := query.Find(&activities)

	if result.Error != nil {
		return nil, result.Error
	}

	return activities, nil
}

func UpdateActivity(ctx context.Context, id int, input graphModel.UpdateActivity) (*model.Activity, error) {
	var activity model.Activity
	result := database.DB.First(&activity, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Update activity fields
	if input.Name != nil && *input.Name != "" {
		activity.Name = *input.Name
	}
	if input.StartTime != nil && *input.StartTime != "" {
		startTime, err := time.Parse("2006-01-02T15:04:05", *input.StartTime)
		if err != nil {
			return nil, err
		}
		activity.StartTime = &startTime
	}
	if input.EndTime != nil && *input.EndTime != "" {
		endTime, err := time.Parse("2006-01-02T15:04:05", *input.EndTime)
		if err != nil {
			return nil, err
		}
		activity.EndTime = &endTime
	}
	if input.Description != nil {
		activity.Description = *input.Description
	}
	if input.EventID != nil {
		activity.EventID = uint(*input.EventID)
	}

	// Save updated activity
	result = database.DB.Save(&activity)
	if result.Error != nil {
		return nil, result.Error
	}

	return &activity, nil
}

func DeleteActivity(ctx context.Context, id int) (bool, error) {
	result := database.DB.Delete(&model.Activity{}, id)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
