package middlewares

import (
	"net/http"
	"strconv"

	"github.com/automa8e_clone/types"
	"github.com/gin-gonic/gin"
)

func generatePageError(msg string) types.ApiError {
	return types.ApiError{
		Field: "page",
		Message: msg,
	}
}

func generatePageSizeError(msg string) types.ApiError {
	return types.ApiError{
		Field: "page_size",
		Message: msg,
	}
}


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
		slice := append(errors, generatePageError("Page should be number"))
		errors = slice
	}

	if errPageSize != nil {
		slice := append(errors, generatePageSizeError("Page Size should be number"))
		errors = slice
	}

	if intPage <= 0 {
		slice := append(errors, generatePageError("Page should be more than 0"))
		errors = slice
	}

	if intPageSize <= 0 {
		slice := append(errors, generatePageSizeError("Page Size should be more than 0"))
		errors = slice
	}

	if (len(errors) > 0) {
		c.JSON(400,gin.H{"status": http.StatusBadRequest, "errors": errors})
		c.Abort()
	}


	offset := (intPage - 1) * intPageSize

	c.Set("pagination", types.PaginationQuery{Page: intPage, PageSize: intPageSize, Offset: offset})
	
}