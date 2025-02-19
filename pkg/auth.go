package pkg

import "context"

type Auth interface {
	// Registers a new user and returns a user ID or error if the process fails.
	SignUp(ctx context.Context, signUp *SignUpRequest) (string, error)

	// Verifies a user's email address using a confirmation code.
	ConfirmSignUp(ctx context.Context, verify *VerifyEmailRequest) error

	// Authenticates a user and returns user details, tokens, or an error.
	SignIn(ctx context.Context, signIn *SignInRequest) (*User, *Token, error)

	// Validates an access token and returns the associated user details.
	CheckToken(ctx context.Context, accessToken string) (*User, error)

	// Generates a new access token using a refresh token for an authenticated user.
	RefreshToken(ctx context.Context, refreshToken string) (*Token, error)

	// Updates user information (e.g., profile details) using an access token.
	UpdateUser(ctx context.Context, accessToken string, user *UpdateUserRequest) error

	// Confirms a user's email change request with a verification code.
	ConfirmChangeEmail(ctx context.Context, accessToken string, verifyEmail *VerifyEmailRequest) error

	// Initiates a password reset process for a user by their ID.
	ForgotPassword(ctx context.Context, userID string) error

	// Resets a user's password using a reset token or code.
	ResetPassword(ctx context.Context, reset *ResetPasswordRequest) error

	// Allows a user to change their password after providing the current one.
	ChangePassword(ctx context.Context, change *ChangePasswordRequest) error

	// Logs out a user by invalidating their refresh token.
	SignOut(ctx context.Context, userID string, refreshToken string) error

	// Deletes a user account by their ID.
	DeleteUser(ctx context.Context, userID string) error
}

type SignUpRequest struct {
	FirstName  string `bson:"first_name"`
	LastName   string `bson:"last_name"`
	Email      string `bson:"email"`
	Password   string `bson:"password"`
	Phone      string `bson:"phone"`
	BirthDate  string `bson:"birth_date"`
	Role       string `bson:"role"`
	NotifToken string `bson:"notif_token"`
}

type VerifyEmailRequest struct {
	UserID string
	Code   string
}

type SignInRequest struct {
	Email    string
	Password string
}

type User struct {
	ID         string
	FirstName  string
	LastName   string
	Email      string
	Phone      string
	PhotoURL   string
	BirthDate  string
	Roles      []string
	NotifToken string
}

type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int
}

type UpdateUserRequest struct {
	ID         string
	FirstName  *string
	LastName   *string
	Email      *string
	Phone      *string
	PhotoURL   *string
	BirthDate  *string
	NotifToken *string
}

type ResetPasswordRequest struct {
	UserID   string
	Code     string
	Password string
}

type ChangePasswordRequest struct {
	UserID      string
	AccessToken string
	Email       string
	OldPassword string
	NewPassword string
}
