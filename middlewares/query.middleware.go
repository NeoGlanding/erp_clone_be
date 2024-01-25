package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TypeQueryMiddleware struct {
	Search				string
	SearchExist			bool
	SortBy				string
	SortByExist			bool
	SortDirection		string
	SortDirectionExist	bool
}

func QueryMiddleware(c *gin.Context) {

	var query TypeQueryMiddleware = TypeQueryMiddleware{
		SearchExist: false,
	}

	search, exist := c.GetQuery("search")

	if exist {
		query.Search = fmt.Sprintf("%%%s%%", search)
		query.SearchExist = true
	}

	sortBy, exist := c.GetQuery("sort_by")

	if exist {
		query.SortBy = sortBy;
		query.SortByExist = true
	}

	sortDirection, exist := c.GetQuery("sort_direction")

	if exist {
		query.SortDirection = sortDirection
		query.SortDirectionExist = true
	} else {
		query.SortDirection = "asc"
		query.SortDirectionExist = false;
	}

	c.Set("query", query)
	
	c.Next()
}