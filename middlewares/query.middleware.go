package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TypeQueryMiddleware struct {
	Search		string
	SearchExist	bool
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

	c.Set("query", query)
	
	c.Next()
}