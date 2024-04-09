package controller

import (
	"context"
	"strconv"

	graphModel "github.com/aman-lf/event-management/graph/model"
	"github.com/aman-lf/event-management/service"
)

func GetActivityHandler(ctx context.Context, filter *graphModel.ActivityFilter, pagination *graphModel.Pagination) ([]*graphModel.Activity, error) {
	activities, err := service.GetActivities(ctx, filter, pagination)
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
