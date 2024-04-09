package controller

import (
	"context"
	"errors"

	"github.com/aman-lf/event-management/data"
	"github.com/aman-lf/event-management/database"
	"github.com/aman-lf/event-management/middleware"
	"github.com/aman-lf/event-management/model"
	"github.com/aman-lf/event-management/service"
)

var ErrUnauthorized = errors.New("Unauthorized")

func getCurrentUser(ctx context.Context, eventId int) (*model.Participant, error) {
	userId := middleware.GetCurrentUserIDFromContext(ctx)
	currentUser, err := service.GetParticipantByUserIDAndEventId(ctx, *userId, eventId)
	if err != nil {
		return nil, ErrUnauthorized
	}

	return currentUser, nil
}

func canManageParticipant(ctx context.Context, eventId int, requestedRole string) error {
	currentUser, err := getCurrentUser(ctx, eventId)
	if err != nil {
		return ErrUnauthorized
	}

	if currentUser.IsAdmin(eventId) {
		return nil
	} else if currentUser.IsContributor(eventId) && requestedRole == data.ATTENDEE {
		return nil
	}

	return ErrUnauthorized
}

func canDeleteParticipant(ctx context.Context, requestedUserId string) error {
	var participant *model.Participant
	result := database.DB.First(&participant, requestedUserId)
	if result.Error != nil {
		return ErrUnauthorized
	}

	eventId := int(participant.EventID)
	currentUser, err := getCurrentUser(ctx, eventId)
	if err != nil {
		return ErrUnauthorized
	}

	if currentUser.IsAdmin(eventId) {
		return nil
	} else if currentUser.IsContributor(eventId) && participant.Role == data.ATTENDEE {
		return nil
	}

	return ErrUnauthorized
}

func hasAdminAccess(ctx context.Context, eventId int) error {
	currentUser, err := getCurrentUser(ctx, eventId)
	if err != nil {
		return ErrUnauthorized
	}

	if !currentUser.IsAdmin(eventId) {
		return ErrUnauthorized
	}

	return nil
}

func hasAdminOrContributerAccess(ctx context.Context, eventId int) error {
	currentUser, err := getCurrentUser(ctx, eventId)
	if err != nil {
		return ErrUnauthorized
	}

	if !(currentUser.IsAdmin(eventId) || currentUser.IsContributor(eventId)) {
		return ErrUnauthorized
	}

	return nil
}

func hasEventAccess(ctx context.Context, eventId int) error {
	_, err := getCurrentUser(ctx, eventId)
	if err != nil {
		return ErrUnauthorized
	}

	return nil
}
