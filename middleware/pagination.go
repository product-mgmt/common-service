package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/product-mgmt/common-service/constants/messages"
	"github.com/product-mgmt/common-service/types"
)

func (s *Storage) Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		searchColumn := (r.URL.Query().Get("searchcolmn"))
		searchTerm := (r.URL.Query().Get("searchterm"))
		sortColumn := (r.URL.Query().Get("sortcolumn"))
		sortBy := (r.URL.Query().Get("sortby"))

		if sortColumn == "" {
			sortColumn = "id"
		}

		if sortBy == "" {
			sortBy = "ASC"
		}

		if searchColumn == "" {
			searchColumn = "name"
		}

		recordPerPage, err := strconv.Atoi(r.URL.Query().Get("recordPerPage"))
		if err != nil || recordPerPage < 0 {
			recordPerPage = 10
		}

		page, err1 := strconv.Atoi(r.URL.Query().Get("page"))
		if err1 != nil || page < 1 {
			page = 1
		}

		offset := (page - 1) * recordPerPage

		resp := types.Paginate{
			SearchColumn: searchColumn,
			SearchTerm:   searchTerm,
			SortColumn:   searchColumn,
			SortOrder:    sortBy,
			Limit:        recordPerPage,
			Offset:       offset,
		}

		// Set user data in the request context
		ctx := context.WithValue(r.Context(), types.CTXKey{Key: messages.PAGINATE}, resp)

		// Update the request with the new context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
