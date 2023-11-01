package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ankeshnirala/sqlscan"
	"github.com/gorilla/mux"

	"github.com/product-mgmt/common-service/constants/messages"
	"github.com/product-mgmt/common-service/constants/procedures"
	"github.com/product-mgmt/common-service/constants/tables"
	"github.com/product-mgmt/common-service/types"
	"github.com/product-mgmt/common-service/utils/commfunc"
)

func (s *Storage) Authorization(roles ...string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID := r.Context().Value(types.CTXKey{Key: "userID"})

			// create a context to timeout db operation once work end
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// checking user is registered or not, if not then throw USERNOTREGISTERED error
			rows, err := s.mysqlStore.GetRecordByArgs(ctx, procedures.SP_GETRECORD, tables.USERS, "id", userID)
			if err != nil {
				s.logger.Error(fmt.Errorf("GetRecordByArgs.Error: %s", messages.INVALIDPERMISSION))
				commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.INVALIDPERMISSION})
				return
			}

			// maping db user user details to user struct
			var user types.User
			if err := sqlscan.Row(&user, rows); err != nil {
				s.logger.Error(fmt.Errorf("sqlscan.Error: %s", messages.INVALIDPERMISSION))
				commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.INVALIDPERMISSION})
				return
			}

			joinRoles := strings.Join(roles, ",")
			isRoleIncluded := strings.Contains(joinRoles, user.Role)

			if !isRoleIncluded || user.Role == "" {
				s.logger.Error(fmt.Errorf("isRoleIncluded.Error: %s", messages.INVALIDPERMISSION))
				commfunc.WriteJSON(w, http.StatusForbidden, commfunc.ApiError{Error: messages.INVALIDPERMISSION})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
