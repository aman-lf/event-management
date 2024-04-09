package service

import (
	"context"
	"strings"

	"github.com/aman-lf/event-management/database"
	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/model"
	"github.com/aman-lf/event-management/utils"
)

var participantSortableCol = []string{"user_id", "event_id", "role"}

func CreateParticipant(ctx context.Context, input graphModel.NewParticipant) (*model.Participant, error) {
	participant := model.Participant{
		UserID:  uint(input.UserID),
		EventID: uint(input.EventID),
		Role:    input.Role,
	}
	result := database.DB.Create(&participant)
	if result.Error != nil {
		return nil, result.Error
	}

	return &participant, nil
}

func GetParticipants(ctx context.Context, options *graphModel.ParticipantQueryOptions) ([]*model.Participant, error) {
	var participants []*model.Participant

	query := database.DB.Model(&model.Participant{})

	if options != nil {
		filter := options.Filter
		if filter != nil {
			if filter.ID != nil && *filter.ID != "" {
				query = query.Where("id = ?", *filter.ID)
			}
			if filter.UserID != nil && *filter.UserID != "" {
				query = query.Where("user_id = ?", *filter.UserID)
			}
			if filter.EventID != nil && *filter.EventID != "" {
				query = query.Where("event_id = ?", *filter.EventID)
			}
			if filter.Role != nil && *filter.Role != "" {
				query = query.Where("LOWER(role) ILIKE ?", "%"+strings.ToLower(*filter.Role)+"%")
			}
		}

		// Apply sorting
		if options.SortBy != nil && *options.SortBy != "" {
			sortColumn := "id" // Default sort column
			if utils.ContainsString(participantSortableCol, *options.SortBy) {
				sortColumn = *options.SortBy
			}
			order := "ASC" // Default sort order
			if strings.ToUpper(*options.SortOrder) == "DESC" {
				order = "DESC"
			}
			query = query.Order(sortColumn + " " + order)
		}

		// Apply limit and offset
		query = query.Limit(*options.Limit).Offset(*options.Offset)
	}

	// Execute the query and fetch participants
	result := query.Find(&participants)

	if result.Error != nil {
		return nil, result.Error
	}

	return participants, nil
}
