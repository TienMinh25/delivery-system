package gateway

import (
	"context"

	"github.com/TienMinh25/delivery-system/internal/products"
	"github.com/TienMinh25/delivery-system/pkg"
	"github.com/pkg/errors"
)

type Service interface {
	// service handles for user management
	signUp(ctx context.Context, payload *signUpRequest) (*signUpResponse, error)
	confirmSignUp(ctx context.Context, payload *confirmSignUpRequest) error
	signIn(ctx context.Context, payload *signInRequest) (*token, error)
	checkToken(ctx context.Context, accessToken string) (*pkg.User, error)
	refreshToken(ctx context.Context, payload *refreshTokenRequest) (*token, error)
	updateUser(ctx context.Context, user *user, payload *updateUserRequest) error
	confirmChangeEmail(ctx context.Context, payload *confirmChangeEmailRequest) error
	updateAvatar(ctx context.Context, user *user, file pkg.File, fileInfo pkg.FileInfo) error
	forgotPassword(ctx context.Context, userID string) error
	resetPassword(ctx context.Context, payload *resetPasswordRequest) error
	changePassword(ctx context.Context, user *user, payload *changePasswordRequest) error
	signOut(ctx context.Context, userID string, refreshTokenPayload *refreshTokenRequest) error
	deleteUser(context.Context, string) error

	// service handles for other services
	getPartnerProducts(ctx context.Context) (*products.GetAllResponse, error)
	checkOrder(ctx context.Context, user *user, payload *checkRequest) error
	confirmOrder(ctx context.Context, user *user, payload *confirmRequest) (*confirmResponse, error)
	payOrder(ctx context.Context, payload *payRequest) (*payResponse, error)
	pickUpOrder(ctx context.Context, payload *pickupRequest) error
	assignOrder(ctx context.Context, user *user, payload *assignRequest) error
}

type service struct {
	authManager pkg.Auth
}

func NewService(
	authManager pkg.Auth,
) Service {
	return &service{
		authManager: authManager,
	}
}

// assignOrder implements Service.
func (s *service) assignOrder(ctx context.Context, user *user, payload *assignRequest) error {
	panic("unimplemented")
}

// changePassword implements Service.
func (s *service) changePassword(ctx context.Context, user *user, payload *changePasswordRequest) error {
	panic("unimplemented")
}

// checkOrder implements Service.
func (s *service) checkOrder(ctx context.Context, user *user, payload *checkRequest) error {
	panic("unimplemented")
}

// checkToken implements Service.
func (s *service) checkToken(ctx context.Context, accessToken string) (*pkg.User, error) {
	user, err := s.authManager.CheckToken(ctx, accessToken)

	if err != nil {
		return nil, errors.Wrap(err, "s.authManager.CheckToken")
	}

	return user, nil
}

// confirmChangeEmail implements Service.
func (s *service) confirmChangeEmail(ctx context.Context, payload *confirmChangeEmailRequest) error {
	err := s.authManager.ConfirmChangeEmail(ctx, payload.AccessToken, &pkg.VerifyEmailRequest{
		UserID: payload.UserID,
		Code:   payload.Code,
	})

	if err != nil {
		return errors.Wrap(err, "s.authManager.ConfirmChangeEmail")
	}

	return nil
}

// confirmOrder implements Service.
func (s *service) confirmOrder(ctx context.Context, user *user, payload *confirmRequest) (*confirmResponse, error) {
	panic("unimplemented")
}

// confirmSignUp implements Service.
func (s *service) confirmSignUp(ctx context.Context, payload *confirmSignUpRequest) error {
	err := s.authManager.ConfirmSignUp(ctx, &pkg.VerifyEmailRequest{
		UserID: payload.UserID,
		Code:   payload.Code,
	})

	if err != nil {
		return errors.Wrap(err, "s.authManager.ConfirmSignUp")
	}

	return nil
}

// deleteUser implements Service.
func (s *service) deleteUser(context.Context, string) error {
	panic("unimplemented")
}

// forgotPassword implements Service.
func (s *service) forgotPassword(ctx context.Context, userID string) error {
	panic("unimplemented")
}

// getPartnerProducts implements Service.
func (s *service) getPartnerProducts(ctx context.Context) (*products.GetAllResponse, error) {
	panic("unimplemented")
}

// payOrder implements Service.
func (s *service) payOrder(ctx context.Context, payload *payRequest) (*payResponse, error) {
	panic("unimplemented")
}

// pickUpOrder implements Service.
func (s *service) pickUpOrder(ctx context.Context, payload *pickupRequest) error {
	panic("unimplemented")
}

// refreshToken implements Service.
func (s *service) refreshToken(ctx context.Context, payload *refreshTokenRequest) (*token, error) {
	authToken, err := s.authManager.RefreshToken(ctx, payload.RefreshToken)

	if err != nil {
		return nil, errors.Wrap(err, "s.authManager.RefreshToken")
	}

	return &token{
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
		ExpiresIn:    authToken.ExpiresIn,
	}, nil
}

// resetPassword implements Service.
func (s *service) resetPassword(ctx context.Context, payload *resetPasswordRequest) error {
	panic("unimplemented")
}

// signIn implements Service.
func (s *service) signIn(ctx context.Context, payload *signInRequest) (*token, error) {
	_, authToken, err := s.authManager.SignIn(ctx, &pkg.SignInRequest{
		Email:    payload.Email,
		Password: payload.Password,
	})

	if err != nil {
		return nil, errors.Wrap(err, "s.authManager.SignIn")
	}

	return &token{
		AccessToken:  authToken.AccessToken,
		RefreshToken: authToken.RefreshToken,
		ExpiresIn:    authToken.ExpiresIn,
	}, nil
}

// signOut implements Service.
func (s *service) signOut(ctx context.Context, userID string, refreshTokenPayload *refreshTokenRequest) error {
	err := s.authManager.SignOut(ctx, userID, refreshTokenPayload.RefreshToken)

	if err != nil {
		return errors.Wrap(err, "s.authManager.SignOut")
	}

	return nil
}

// signUp implements Service.
func (s *service) signUp(ctx context.Context, payload *signUpRequest) (*signUpResponse, error) {
	userID, err := s.authManager.SignUp(ctx, &pkg.SignUpRequest{
		FirstName:  payload.FirstName,
		LastName:   payload.LastName,
		Email:      payload.Email,
		Password:   payload.Password,
		Phone:      payload.Phone,
		BirthDate:  payload.BirthDate,
		Role:       payload.Role,
		NotifToken: payload.NotifToken,
	})

	if err != nil {
		return nil, errors.Wrap(err, "s.authManager.SignUp")
	}

	return &signUpResponse{UserID: userID}, nil
}

// updateAvatar implements Service.
func (s *service) updateAvatar(ctx context.Context, user *user, file pkg.File, fileInfo pkg.FileInfo) error {
	panic("unimplemented")
}

// updateUser implements Service.
func (s *service) updateUser(ctx context.Context, user *user, payload *updateUserRequest) error {
	userToUpdate := &pkg.UpdateUserRequest{
		ID: user.ID,
	}

	if payload.FirstName != "" {
		userToUpdate.FirstName = &payload.FirstName
	}

	if payload.LastName != "" {
		userToUpdate.LastName = &payload.LastName
	}

	if payload.Email != "" {
		userToUpdate.Email = &payload.Email
	}

	if payload.Phone != "" {
		userToUpdate.Phone = &payload.Phone
	}

	if payload.BirthDate != "" {
		userToUpdate.BirthDate = &payload.BirthDate
	}

	if payload.NotifToken != "" {
		userToUpdate.NotifToken = &payload.NotifToken
	}

	err := s.authManager.UpdateUser(ctx, user.AccessToken, userToUpdate)

	if err != nil {
		return errors.Wrap(err, "s.authManager.UpdateUser")
	}

	return nil
}
