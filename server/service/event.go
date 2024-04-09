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

var eventSortableCol = []string{"name", "start_date", "end_date", "location", "type"}

func CreateEvent(ctx context.Context, input graphModel.NewEvent) (*model.Event, error) {
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return nil, err
	}
	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return nil, err
	}
	event := model.Event{
		Name:        input.Name,
		StartDate:   &startDate,
		EndDate:     &endDate,
		Location:    input.Location,
		Type:        input.Type,
		Description: input.Description,
	}

	result := database.DB.Create(&event)
	if result.Error != nil {
		return nil, result.Error
	}

	return &event, err
}

func GetEvents(ctx context.Context, filter *graphModel.EventFilter, pagination *graphModel.Pagination) ([]*model.Event, error) {
	var events []*model.Event

	query := database.DB.Model(&model.Event{})

	if filter != nil {
		if filter != nil {
			if filter.ID != nil && *filter.ID != "" {
				query = query.Where("id = ?", *filter.ID)
			}
			if filter.Name != nil && *filter.Name != "" {
				query = query.Where("LOWER(name) ILIKE ?", "%"+strings.ToLower(*filter.Name)+"%")
			}
			if filter.StartDate != nil && *filter.StartDate != "" {
				query = query.Where("start_date = ?", *filter.StartDate)
			}
			if filter.EndDate != nil && *filter.EndDate != "" {
				query = query.Where("end_date = ?", *filter.EndDate)
			}
			if filter.Location != nil && *filter.Location != "" {
				query = query.Where("LOWER(location) ILIKE ?", "%"+strings.ToLower(*filter.Location)+"%")
			}
			if filter.Type != nil && *filter.Type != "" {
				query = query.Where("LOWER(type) ILIKE ?", "%"+strings.ToLower(*filter.Type)+"%")
			}
		}

		// Apply pagination and sort
		query = utils.ApplyPagination(query, pagination, "id", eventSortableCol)
	}

	// Execute the query and fetch events
	result := query.Find(&events)

	if result.Error != nil {
		return nil, result.Error
	}

	return events, result.Error
}

func GetEventById(ctx context.Context, id int) (*model.Event, error) {
	var event *model.Event

	result := database.DB.First(&event, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}
