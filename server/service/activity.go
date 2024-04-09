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

func GetActivities(ctx context.Context, filter *graphModel.ActivityFilter, pagination *graphModel.Pagination) ([]*model.Activity, error) {
	var activities []*model.Activity

	query := database.DB.Model(&model.Activity{})

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
