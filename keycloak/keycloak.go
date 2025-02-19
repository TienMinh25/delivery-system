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
func (a *auth) ChangePassword(ctx context.Context, change *pkg.ChangePasswordRequest) error {
	panic("unimplemented")
}

// CheckToken implements pkg.Auth.
func (a *auth) CheckToken(ctx context.Context, accessToken string) (*pkg.User, error) {
	panic("unimplemented")
}

// ConfirmChangeEmail implements pkg.Auth.
func (a *auth) ConfirmChangeEmail(ctx context.Context, accessToken string, verifyEmail *pkg.VerifyEmailRequest) error {
	panic("unimplemented")
}

// ConfirmSignUp implements pkg.Auth.
func (a *auth) ConfirmSignUp(ctx context.Context, verify *pkg.VerifyEmailRequest) error {
	// TODO: do not understand this
	return nil
}

// DeleteUser implements pkg.Auth.
func (a *auth) DeleteUser(ctx context.Context, userID string) error {
	panic("unimplemented")
}

// ForgotPassword implements pkg.Auth.
func (a *auth) ForgotPassword(ctx context.Context, userID string) error {
	panic("unimplemented")
}

// RefreshToken implements pkg.Auth.
func (a *auth) RefreshToken(ctx context.Context, userID string, refreshToken string) (*pkg.Token, error) {
	panic("unimplemented")
}

// ResetPassword implements pkg.Auth.
func (a *auth) ResetPassword(ctx context.Context, reset *pkg.ResetPasswordRequest) error {
	panic("unimplemented")
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
	panic("unimplemented")
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
func (a *auth) UpdateUser(ctx context.Context, accessToken string, user *pkg.UpdateUserRequest) error {
	panic("unimplemented")
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
