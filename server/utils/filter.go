package utils

import (
	"strings"

	graphModel "github.com/aman-lf/event-management/graph/model"
	"gorm.io/gorm"
)

func ApplyPagination(query *gorm.DB, pagination *graphModel.Pagination, defaultSort string, sortableCol []string) *gorm.DB {
	if pagination.SortBy != nil && *pagination.SortBy != "" {
		sortColumn := defaultSort
		if ContainsString(sortableCol, *pagination.SortBy) {
			sortColumn = *pagination.SortBy
		}
		order := "ASC" // Default sort order
		if strings.ToUpper(*pagination.SortOrder) == "DESC" {
			order = "DESC"
		}
		query = query.Order(sortColumn + " " + order)
	}

	query = query.Limit(*pagination.Limit).Offset(*pagination.Offset)

	return query
}
