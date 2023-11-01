package types

type CTXKey struct {
	Key string
}

type SigninRequest struct {
	Email    string `json:"email"`
	Passowrd string `json:"password"`
}

type SigninResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type SignupRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	ReferralCode  string `json:"referral_code"`
	RefferredCode string `json:"refferred_code"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

type RegisterOutput struct {
	Message    string `db:"message"`
	InsertedID int    `db:"insertedID"`
}

type ProfileResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}
