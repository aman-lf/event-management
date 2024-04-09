package controller

import (
	"context"
	"strconv"

	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/middleware"
	"github.com/aman-lf/event-management/service"
)

func GetActivityHandler(ctx context.Context, filter *graphModel.ActivityFilter, pagination *graphModel.Pagination) ([]*graphModel.Activity, error) {
	userId := middleware.GetCurrentUserIDFromContext(ctx)
	activities, err := service.GetActivitiesByUserId(ctx, *userId, filter, pagination)
	if err != nil {
		return nil, err
	}

	var returnActivities []*graphModel.Activity
	for _, activity := range activities {
		activityID := strconv.FormatUint(uint64(activity.ID), 10)
		returnActivities = append(returnActivities, &graphModel.Activity{
			ID:          activityID,
			Name:        activity.Name,
			StartTime:   activity.StartTime.Format("2006-01-02T15:04:05"),
			EndTime:     activity.EndTime.Format("2006-01-02T15:04:05"),
			Description: &activity.Description,
			EventID:     int(activity.EventID),
		})
	}
	return returnActivities, nil
}

func CreateActivityHandler(ctx context.Context, input graphModel.NewActivity) (*graphModel.Activity, error) {
	err := hasAdminOrContributerAccess(ctx, input.EventID)
	if err != nil {
		return nil, err
	}

	activity, err := service.CreateActivity(ctx, input)
	if err != nil {
		return nil, err
	}

	activityID := strconv.FormatUint(uint64(activity.ID), 10)
	return &graphModel.Activity{
		ID:          activityID,
		Name:        input.Name,
		StartTime:   input.StartTime,
		EndTime:     input.EndTime,
		Description: input.Description,
		EventID:     int(activity.EventID),
	}, nil
}

func UpdateActivityHandler(ctx context.Context, idStr string, input graphModel.UpdateActivity) (*graphModel.Activity, error) {
	id, _ := strconv.Atoi(idStr)
	err := hasAdminOrContributerAccess(ctx, id)
	if err != nil {
		return nil, err
	}

	activity, err := service.UpdateActivity(ctx, id, input)
	if err != nil {
		return nil, err
	}

	activityID := strconv.FormatUint(uint64(activity.ID), 10)
	return &graphModel.Activity{
		ID:          activityID,
		Name:        activity.Name,
		StartTime:   activity.StartTime.Format("2006-01-02T15:04:05"),
		EndTime:     activity.EndTime.Format("2006-01-02T15:04:05"),
		Description: &activity.Description,
		EventID:     int(activity.EventID),
	}, nil
}

func DeleteActivityHandler(ctx context.Context, idStr string) (bool, error) {
	id, _ := strconv.Atoi(idStr)
	err := hasAdminOrContributerAccess(ctx, id)
	if err != nil {
		return false, err
	}

	return service.DeleteActivity(ctx, id)
}
