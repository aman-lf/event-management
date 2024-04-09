package controller

import (
	"context"
	"strconv"

	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/service"
)

func GetParticipantHandler(ctx context.Context, filter *graphModel.ParticipantFilter, pagination *graphModel.Pagination) ([]*graphModel.Participant, error) {
	participants, err := service.GetParticipants(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}

	var returnParticipant []*graphModel.Participant
	for _, participant := range participants {
		participantID := strconv.FormatUint(uint64(participant.ID), 10)
		returnParticipant = append(returnParticipant, &graphModel.Participant{
			ID:      participantID,
			UserId:  int(participant.UserID),
			EventID: int(participant.EventID),
			Role:    participant.Role,
		})
	}
	return returnParticipant, nil
}

func CreateParticipantHandler(ctx context.Context, input graphModel.NewParticipant) (*graphModel.Participant, error) {
	participant, err := service.CreateParticipant(ctx, input)
	if err != nil {
		return nil, err
	}

	participantId := strconv.FormatUint(uint64(participant.ID), 10)
	return &graphModel.Participant{
		ID:      participantId,
		UserId:  int(participant.UserID),
		EventID: int(participant.EventID),
		Role:    input.Role,
	}, nil
}

func UpdateParticipantHandler(ctx context.Context, idStr string, input graphModel.UpdateParticipant) (*graphModel.Participant, error) {
	id, _ := strconv.Atoi(idStr)
	participant, err := service.UpdateParticipant(ctx, id, input)
	if err != nil {
		return nil, err
	}

	participantID := strconv.FormatUint(uint64(participant.ID), 10)
	return &graphModel.Participant{
		ID:      participantID,
		UserId:  int(participant.UserID),
		EventID: int(participant.EventID),
		Role:    participant.Role,
	}, nil
}

func DeleteParticipantHandler(ctx context.Context, idStr string) (bool, error) {
	id, _ := strconv.Atoi(idStr)
	return service.DeleteParticipant(ctx, id)
}
