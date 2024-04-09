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

func GetParticipants(ctx context.Context, filter *graphModel.ParticipantFilter, pagination *graphModel.Pagination) ([]*model.Participant, error) {
	var participants []*model.Participant

	query := database.DB.Model(&model.Participant{})

	if filter != nil {
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

		// Apply pagination and sort
		query = utils.ApplyPagination(query, pagination, "id", participantSortableCol)
	}

	// Execute the query and fetch participants
	result := query.Find(&participants)

	if result.Error != nil {
		return nil, result.Error
	}

	return participants, nil
}
