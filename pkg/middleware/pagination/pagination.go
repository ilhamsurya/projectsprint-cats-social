package pagination

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
	SortBy     string `json:"sort_by"`
	Search     string `json:"search"`
	TotalItems int    `json:"total_items"`
	TotalPages int    `json:"total_pages"`
}

type PaginationResponse struct {
	Pagination Pagination  `json:"pagination,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func ReturnPaginationResult(pagination Pagination, data interface{}) PaginationResponse {
	res := PaginationResponse{
		Pagination: pagination,
		Data:       data,
	}
	return res
}

func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	sort := "DESC"
	limit := 10
	page := 1
	sort_by := "updated_at"
	search := ""
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			if strings.ToLower(queryValue) == "desc" || strings.ToLower(queryValue) == "asc" {
				sort = queryValue
			}
		case "sort_by":
			sort_by = queryValue
		case "search":
			search = queryValue
		}
	}
	return Pagination{
		Limit:  limit,
		Page:   page,
		Sort:   sort,
		SortBy: sort_by,
		Search: search,
	}

}
