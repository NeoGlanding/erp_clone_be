package middlewares

import (
	"net/http"
	"strconv"

	"github.com/automa8e_clone/types"
	"github.com/gin-gonic/gin"
)

func PaginationMiddleware(c *gin.Context) {
	page, pageExist := c.GetQuery("page")
	pageSize, pageSizeExist := c.GetQuery("page_size")
	
	if !pageExist {
		page = "1"
	}
	
	if !pageSizeExist {
		pageSize = "10"
	}

	var errors []types.ApiError = []types.ApiError{}

	intPage, err := strconv.Atoi(page);
	intPageSize, errPageSize := strconv.Atoi(pageSize)

	if err != nil {
		slice := append(errors, types.ApiError{Message: "Page should be number", Field: "page"})
		errors = slice
	}

	if errPageSize != nil {
		slice := append(errors, types.ApiError{Message: "Page Size should be number", Field: "page_size"})
		errors = slice
	}

	if (len(errors) > 0) {
		c.JSON(400,gin.H{"status": http.StatusBadRequest, "errors": errors})
		c.Abort()
	}

	offset := (intPage - 1) * intPageSize

	c.Set("pagination", types.PaginationQuery{Page: intPage, PageSize: intPageSize, Offset: offset})
	
}