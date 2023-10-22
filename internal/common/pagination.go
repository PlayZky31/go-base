package common

import "math"

type QueryPagination struct {
	Page          int    `json:"page"`
	ItemsPerPage  int    `json:"items_per_page"`
	TotalItems    int64  `json:"total_items"`
	SearchKeyword string `json:"search_keyword"`
}

type Pagination struct {
	NextPage  *int `json:"next_page"`
	PrevPage  *int `json:"prev_page"`
	TotalPage int  `json:"total_page"`
	*QueryPagination
}

func GetPaginationFromQuery(query *QueryPagination) Pagination {
	var (
		prevPage *int = nil
		nextPage *int = nil
	)

	totalPage := int(math.Ceil(float64(query.TotalItems) / float64(query.ItemsPerPage)))

	if query.Page < totalPage {
		val := query.Page + 1
		nextPage = &val
	}

	if query.Page >= 2 {
		val := query.Page - 1
		prevPage = &val
	}

	return Pagination{
		NextPage:        nextPage,
		PrevPage:        prevPage,
		TotalPage:       totalPage,
		QueryPagination: query,
	}
}
