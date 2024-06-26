package controller

import (
	"context"
	"strconv"

	"github.com/aman-lf/event-management/data"
	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/middleware"
	"github.com/aman-lf/event-management/service"
)

func GetEventsHandler(ctx context.Context, filter *graphModel.EventFilter, pagination *graphModel.Pagination) ([]*graphModel.Event, error) {
	if filter != nil && filter.ID != nil {
		eventId, _ := strconv.Atoi(*filter.ID)
		err := hasEventAccess(ctx, eventId)
		if err != nil {
			return nil, err
		}
	}

	userId := middleware.GetCurrentUserIDFromContext(ctx)
	events, err := service.GetEventsByUserId(ctx, *userId, filter, pagination)
	if err != nil {
		return nil, err
	}

	var returnEvent []*graphModel.Event
	for _, event := range events {
		eventID := strconv.FormatUint(uint64(event.ID), 10)
		returnEvent = append(returnEvent, &graphModel.Event{
			ID:          eventID,
			Name:        event.Name,
			StartDate:   event.StartDate.Format("2006-01-02"),
			EndDate:     event.EndDate.Format("2006-01-02"),
			Location:    event.Location,
			Type:        new(string),
			Description: new(string),
		})
	}

	return returnEvent, nil
}

func GetEventByIdHandler(ctx context.Context, id int) (*graphModel.Event, error) {
	event, err := service.GetEventById(ctx, id)
	if err != nil {
		return nil, err
	}

	eventID := strconv.FormatUint(uint64(event.ID), 10)
	returnEvent := graphModel.Event{
		ID:          eventID,
		Name:        event.Name,
		StartDate:   event.StartDate.Format("2006-01-02"),
		EndDate:     event.EndDate.Format("2006-01-02"),
		Location:    event.Location,
		Type:        event.Type,
		Description: event.Description,
	}

	return &returnEvent, nil
}

func CreateEventHandler(ctx context.Context, input graphModel.NewEvent) (*graphModel.Event, error) {
	event, err := service.CreateEvent(ctx, input)
	if err != nil {
		return nil, err
	}

	userId := middleware.GetCurrentUserIDFromContext(ctx)
	newParticipant := graphModel.NewParticipant{
		UserID:  *userId,
		EventID: int(event.ID),
		Role:    data.ADMIN,
	}
	service.CreateParticipant(ctx, newParticipant)

	eventID := strconv.FormatUint(uint64(event.ID), 10)
	return &graphModel.Event{
		ID:          eventID,
		Name:        event.Name,
		StartDate:   event.StartDate.Format("2006-01-02"),
		EndDate:     event.EndDate.Format("2006-01-02"),
		Location:    event.Location,
		Type:        event.Type,
		Description: event.Description,
	}, nil
}

func UpdateEventHandler(ctx context.Context, idStr string, input graphModel.UpdateEvent) (*graphModel.Event, error) {
	id, _ := strconv.Atoi(idStr)
	err := hasAdminAccess(ctx, id)
	if err != nil {
		return nil, err
	}

	event, err := service.UpdateEvent(ctx, id, input)
	if err != nil {
		return nil, err
	}

	eventID := strconv.FormatUint(uint64(event.ID), 10)
	return &graphModel.Event{
		ID:          eventID,
		Name:        event.Name,
		StartDate:   event.StartDate.Format("2006-01-02"),
		EndDate:     event.EndDate.Format("2006-01-02"),
		Location:    event.Location,
		Type:        event.Type,
		Description: event.Description,
	}, nil
}

func DeleteEventHandler(ctx context.Context, idStr string) (bool, error) {
	id, _ := strconv.Atoi(idStr)
	err := hasAdminAccess(ctx, id)
	if err != nil {
		return false, err
	}

	return service.DeleteEvent(ctx, id)
}
