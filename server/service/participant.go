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

func GetParticipantsByEventId(ctx context.Context, eventId int, filter *graphModel.ParticipantFilter, pagination *graphModel.Pagination) ([]*model.Participant, error) {
	var participants []*model.Participant

	query := database.DB.Model(&model.Participant{})

	if filter != nil {
		if filter.ID != nil && *filter.ID != "" {
			query = query.Where("id = ?", *filter.ID)
		}
		if filter.UserID != nil && *filter.UserID != "" {
			query = query.Where("user_id = ?", *filter.UserID)
		}
		if filter.Role != nil && *filter.Role != "" {
			query = query.Where("LOWER(role) ILIKE ?", "%"+strings.ToLower(*filter.Role)+"%")
		}

		// Apply pagination and sort
		query = utils.ApplyPagination(query, pagination, "id", participantSortableCol)
	}

	// Execute the query and fetch participants
	result := query.Where("event_id = ?", eventId).Find(&participants)

	if result.Error != nil {
		return nil, result.Error
	}

	return participants, nil
}

func GetParticipantByUserIDAndEventId(ctx context.Context, userId, eventId int) (*model.Participant, error) {
	var participant *model.Participant
	result := database.DB.Where("event_id = ? AND user_id = ?", eventId, userId).First(&participant)
	if result.Error != nil {
		return nil, result.Error
	}
	return participant, nil
}

func UpdateParticipant(ctx context.Context, id int, input graphModel.UpdateParticipant) (*model.Participant, error) {
	var participant model.Participant
	result := database.DB.First(&participant, id)
	if result.Error != nil {
		return nil, result.Error
	}

	// Update participant fields
	if input.UserID != nil {
		participant.UserID = uint(*input.UserID)
	}
	if input.EventID != nil {
		participant.EventID = uint(*input.EventID)
	}
	if input.Role != nil {
		participant.Role = *input.Role
	}

	// Save updated participant
	result = database.DB.Save(&participant)
	if result.Error != nil {
		return nil, result.Error
	}

	return &participant, nil
}

func DeleteParticipant(ctx context.Context, id int) (bool, error) {
	result := database.DB.Delete(&model.Participant{}, id)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
