package keycloak

import (
	"context"
	"os"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/TienMinh25/delivery-system/pkg"
	"github.com/pkg/errors"
)

type auth struct {
	client       *gocloak.GoCloak // keycloak client
	clientID     string           // clientID specified in Keycloak
	clientSecret string           // client secret specified in Keycloak
	realm        string           // realm specified in Keycloak
}

func NewAuth() pkg.Auth {
	return &auth{
		realm:        os.Getenv("KEYCLOAK_REALM"),
		clientID:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		clientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		client:       gocloak.NewClient(os.Getenv("KEYCLOAK_BASE_URL")),
	}
}

// ChangePassword implements pkg.Auth.
func (a *auth) ChangePassword(ctx context.Context, payload *pkg.ChangePasswordRequest) error {
	// login to check the old password
	_, err := a.client.Login(ctx, a.clientID, a.clientSecret, a.realm, payload.Email, payload.OldPassword)

	if err != nil {
		return errors.Wrap(err, "a.client.Login")
	}

	// get admin token to set new password
	adminToken, err := a.client.LoginClient(ctx, a.clientID, a.clientSecret, a.realm)

	if err != nil {
		return errors.Wrap(err, "a.client.LoginClient")
	}

	// password will be reset, so all released token should be revoked
	err = a.client.LogoutAllSessions(ctx, adminToken.AccessToken, a.realm, payload.UserID)

	if err != nil {
		return errors.Wrap(err, "a.client.LogoutAllSessions")
	}

	// change password
	return a.client.SetPassword(ctx, adminToken.AccessToken, payload.UserID, a.realm, payload.NewPassword, false)
}

// CheckToken implements pkg.Auth. (using for authorize)
func (a *auth) CheckToken(ctx context.Context, accessToken string) (*pkg.User, error) {
	result, err := a.client.RetrospectToken(ctx, accessToken, a.clientID, a.clientSecret, a.realm)

	if err != nil {
		return nil, errors.Wrap(err, "a.client.RetrospectToken")
	}

	if !gocloak.PBool(result.Active) {
		return nil, errors.New("token is invalid")
	}

	return a.getUser(ctx, accessToken)
}

// ConfirmChangeEmail implements pkg.Auth.
func (a *auth) ConfirmChangeEmail(ctx context.Context, accessToken string, verifyEmail *pkg.VerifyEmailRequest) error {
	// TODO: modify it
	return nil
}

// ConfirmSignUp implements pkg.Auth.
func (a *auth) ConfirmSignUp(ctx context.Context, verify *pkg.VerifyEmailRequest) error {
	// TODO: modify it
	return nil
}

// DeleteUser implements pkg.Auth.
func (a *auth) DeleteUser(ctx context.Context, userID string) error {
	adminToken, err := a.client.LoginClient(ctx, a.clientID, a.clientSecret, a.realm)

	if err != nil {
		return errors.Wrap(err, "a.client.LoginClient")
	}

	// ! Do NOT delete users, just inactive them.
	// a.client.DeleteUser(ctx, token.AccessToken, a.realm, userID)

	// in here decide to use soft delete
	// soft delete

	return a.client.UpdateUser(ctx, adminToken.AccessToken, a.realm, gocloak.User{
		ID:      gocloak.StringP(userID),
		Enabled: gocloak.BoolP(false),
	})
}

// ForgotPassword implements pkg.Auth.
func (a *auth) ForgotPassword(ctx context.Context, userID string) error {
	adminToken, err := a.client.LoginClient(ctx, a.clientID, a.clientSecret, a.realm)

	if err != nil {
		return errors.Wrap(err, "a.client.LoginClient")
	}

	return a.client.ExecuteActionsEmail(ctx, adminToken.AccessToken, a.realm, gocloak.ExecuteActionsEmail{
		UserID:   gocloak.StringP(userID),
		ClientID: gocloak.StringP(a.clientID),
		Actions:  &([]string{"UPDATE_PASSWORD"}),
	})
}

// RefreshToken implements pkg.Auth.
// implement flow revoke access token
func (a *auth) RefreshToken(ctx context.Context, userID string, refreshToken string) (*pkg.Token, error) {
	token, err := a.client.RefreshToken(ctx, refreshToken, a.clientID, a.clientSecret, a.realm)

	if err != nil {
		return nil, errors.Wrap(err, "a.client.RefreshToken")
	}

	return &pkg.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}

// ResetPassword implements pkg.Auth.
func (a *auth) ResetPassword(ctx context.Context, reset *pkg.ResetPasswordRequest) error {
	// TODO: implement after
	return nil
}

// SignIn implements pkg.Auth.
func (a *auth) SignIn(ctx context.Context, payload *pkg.SignInRequest) (*pkg.User, *pkg.Token, error) {
	userToken, err := a.client.Login(ctx, a.clientID, a.clientSecret, a.realm, payload.Email, payload.Password)

	if err != nil {
		return nil, nil, errors.Wrap(err, "a.client.Login")
	}

	user, err := a.getUser(ctx, userToken.AccessToken)

	if err != nil {
		return nil, nil, errors.Wrap(err, "a.getUser")
	}

	token := &pkg.Token{
		AccessToken:  userToken.AccessToken,
		ExpiresIn:    userToken.ExpiresIn,
		RefreshToken: userToken.RefreshToken,
	}

	return user, token, nil
}

// SignOut implements pkg.Auth.
func (a *auth) SignOut(ctx context.Context, userID string, refreshToken string) error {
	return a.client.Logout(ctx, a.clientID, a.clientSecret, a.realm, refreshToken)
}

// SignUp implements pkg.Auth.
func (a *auth) SignUp(ctx context.Context, payload *pkg.SignUpRequest) (string, error) {
	// get admin token to create the user
	adminToken, err := a.client.LoginClient(ctx, a.clientID, a.clientSecret, a.realm)

	if err != nil {
		return "", errors.Wrap(err, "a.client.LoginClient")
	}

	// create user
	user := gocloak.User{
		FirstName:     gocloak.StringP(payload.FirstName),
		LastName:      gocloak.StringP(payload.LastName),
		Email:         gocloak.StringP(payload.Email),
		EmailVerified: gocloak.BoolP(false),
		Enabled:       gocloak.BoolP(true),
	}

	attributes := make(map[string][]string)

	// kha nang cao error
	if payload.Phone != "" {
		attributes["phoneNumber"] = []string{payload.Phone}
	}

	if payload.BirthDate != "" {
		attributes["birthDate"] = []string{payload.BirthDate}
	}

	if payload.NotifToken != "" {
		attributes["notifToken"] = []string{payload.NotifToken}
	}

	user.Attributes = &attributes

	userID, err := a.client.CreateUser(ctx, adminToken.AccessToken, a.realm, user)

	if err != nil {
		return "", errors.Wrap(err, "a.client.CreateUser")
	}

	// set password
	err = a.client.SetPassword(ctx, adminToken.AccessToken, userID, a.realm, payload.Password, false)

	if err != nil {
		return "", errors.Wrap(err, "a.client.SetPassword")
	}

	// add role to user
	role, err := a.client.GetRealmRole(ctx, adminToken.AccessToken, a.realm, strings.ToLower(payload.Role))

	if err != nil {
		return "", errors.Wrap(err, "a.client.getRealmRole")
	}

	err = a.client.AddRealmRoleToUser(ctx, adminToken.AccessToken, a.realm, userID, []gocloak.Role{*role})

	if err != nil {
		return "", errors.Wrap(err, "a.client.AddRealmRoleToUser")
	}

	// send email verification link
	err = a.client.SendVerifyEmail(ctx, adminToken.AccessToken, userID, a.realm)

	if err != nil {
		return "", errors.Wrap(err, "a.client.SendVerifyEmail")
	}

	return userID, nil
}

// UpdateUser implements pkg.Auth.
func (a *auth) UpdateUser(ctx context.Context, accessToken string, payload *pkg.UpdateUserRequest) error {
	// get admin token for updating the user
	adminToken, err := a.client.LoginClient(ctx, a.clientID, a.clientSecret, a.realm)

	if err != nil {
		return errors.Wrap(err, "a.client.LoginClient")
	}

	// get user to set new fields to update
	user, err := a.client.GetUserByID(ctx, adminToken.AccessToken, a.realm, payload.ID)

	if err != nil {
		return errors.Wrap(err, "a.client.GetUserByID")
	}

	if user.Attributes == nil {
		user.Attributes = &map[string][]string{}
	}

	// set user fields to update

	if payload.FirstName != nil {
		user.FirstName = payload.FirstName
	}

	if payload.LastName != nil {
		user.LastName = payload.LastName
	}

	var emailChanged = false

	if payload.Email != nil {
		if gocloak.PString(user.Email) != gocloak.PString(payload.Email) {
			emailChanged = true
			user.EmailVerified = gocloak.BoolP(false)
			user.Email = payload.Email
		}
	}

	if payload.Phone != nil {
		(*user.Attributes)["phoneNumber"] = []string{gocloak.PString(payload.Phone)}
	}

	if payload.BirthDate != nil {
		(*user.Attributes)["birthDate"] = []string{gocloak.PString(payload.BirthDate)}
	}

	if payload.NotifToken != nil {
		(*user.Attributes)["notifToken"] = []string{gocloak.PString(payload.NotifToken)}
	}

	// update the user
	err = a.client.UpdateUser(ctx, accessToken, a.realm, *user)

	if err != nil {
		return errors.Wrap(err, "a.client.UpdateUser")
	}

	// send email verification

	if emailChanged {
		err = a.client.LogoutAllSessions(ctx, adminToken.AccessToken, a.realm, payload.ID)

		if err != nil {
			return errors.Wrap(err, "a.client.LogoutAllSessions")
		}

		err = a.client.SendVerifyEmail(ctx, adminToken.AccessToken, payload.ID, a.realm)

		if err != nil {
			return errors.Wrap(err, "a.client.SendVerifyEmail")
		}
	}

	return nil
}

// helper function to get user
func (a *auth) getUser(ctx context.Context, accessToken string) (*pkg.User, error) {
	// using raw user info instead, because GetUserInfo() does not return
	// custom attribute
	info, err := a.client.GetRawUserInfo(ctx, accessToken, a.realm)

	if err != nil {
		return nil, errors.Wrap(err, "a.client.GetRawUserInfo")
	}

	// decode access token for missing field (like user role,...)
	_, claims, err := a.client.DecodeAccessToken(ctx, accessToken, a.realm)

	if err != nil {
		return nil, errors.Wrap(err, "a.client.DecodeAccessToken")
	}

	roles := make([]string, 0)

	for _, value := range parseAny[[]any]((parseAny[map[string]any]((*claims)["realm_access"]))["roles"]) {
		roles = append(roles, parseAny[string](value))
	}

	return &pkg.User{
		ID:         parseAny[string](info["sub"]),
		FirstName:  parseAny[string](info["given_name"]),
		LastName:   parseAny[string](info["family_name"]),
		Email:      parseAny[string](info["email"]),
		Phone:      parseAny[string](info["phone_number"]),
		PhotoURL:   parseAny[string](info["picture"]),
		BirthDate:  parseAny[string](info["birthdate"]),
		Roles:      roles,
		NotifToken: parseAny[string](info["notif_token"]),
	}, nil
}

func parseAny[T any](v any) (value T) {
	if v == nil {
		return
	}
	value, _ = v.(T)
	return
}
