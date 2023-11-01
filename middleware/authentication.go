package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ankeshnirala/sqlscan"
	"github.com/golang-jwt/jwt/v4"

	"github.com/product-mgmt/common-service/constants/messages"
	"github.com/product-mgmt/common-service/constants/procedures"
	"github.com/product-mgmt/common-service/constants/tables"
	"github.com/product-mgmt/common-service/types"
	"github.com/product-mgmt/common-service/utils/commfunc"
	"github.com/product-mgmt/common-service/utils/jwtauth"
)

func (s *Storage) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Reading token from headers
		tokenString := r.Header.Get("Authorization")

		// checking there is a token or not
		token, err := jwtauth.ValidateJWT(tokenString)
		if err != nil {
			s.logger.Error(fmt.Errorf("jwtauth.ValidateJWT: %s", err))
			commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.ACCESSDENIED})
			return
		}

		if !token.Valid {
			s.logger.Error(fmt.Errorf("token.Valid.Error: %s", messages.ACCESSDENIED))
			commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.ACCESSDENIED})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			s.logger.Error(fmt.Errorf("token.Claims.Error: %s", messages.ACCESSDENIED))
			commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.ACCESSDENIED})
			return
		}

		// userID, err := primitive.ObjectIDFromHex(claims["userID"].(string))
		userID := claims["userID"].(float64)
		if err != nil {
			s.logger.Error(fmt.Errorf("claims[userID].(float64).Error: %s", messages.ACCESSDENIED))
			commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.ACCESSDENIED})
			return
		}

		// create a context to timeout db operation once work end
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// checking user is registered or not, if not then throw USERNOTREGISTERED error
		rows, err := s.mysqlStore.GetRecordByArgs(ctx, procedures.SP_GETRECORD, tables.USERS, "id", userID)
		if err != nil {
			s.logger.Error(fmt.Errorf("checkingUserRegistered.Error: %s", messages.ACCESSDENIED))
			commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.ACCESSDENIED})
			return
		}

		// maping db user user details to user struct
		var user types.User
		if err := sqlscan.Row(&user, rows); err != nil {
			s.logger.Error(fmt.Errorf("sqlscan.Error: %s", messages.ACCESSDENIED))
			commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.ACCESSDENIED})
			return
		}

		// check usr profile recently updated
		iat := claims["iat"].(float64)
		ok, err = user.IsProfileRecentlyUpdated(iat)
		if err != nil {
			s.logger.Error(fmt.Errorf("IsProfileRecentlyUpdated.Error: %s", messages.PROFILEUPDATE))
			commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.PROFILEUPDATE})
			return
		}

		if !ok {
			s.logger.Error(fmt.Errorf("IsProfileRecentlyUpdated.ok: %s", messages.PROFILEUPDATE))
			commfunc.WriteJSON(w, http.StatusUnauthorized, commfunc.ApiError{Error: messages.PROFILEUPDATE})
			return
		}

		// Set user data in the request context
		ctx = context.WithValue(r.Context(), types.CTXKey{Key: "userID"}, userID)

		// Update the request with the new context
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
