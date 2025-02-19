package gateway

// registerUser represents the details of a user for registration/
type signUpRequest struct {
	FirstName  string `json:"first_name" validate:"required,alpha,min=2"` // The first name of the user.
	LastName   string `json:"last_name" validate:"required,alpha,min=2"`  // The last name of the user.
	Email      string `json:"email" validate:"required,email"`            // The email address of the user.
	Password   string `json:"password" validate:"required,min=6"`         // The hashed password of the user. Note: It's important to store passwords securely by hashing them.
	Phone      string `json:"phone" validate:"omitempty,e164"`            // optional: E.164 formatted phone number: [+] [country code] [subscriber number including area code] and can have a maximum of fifteen digits.
	BirthDate  string `json:"birth_date" validate:"omitempty,dateonly"`   // optional: YYYY-MM-DD
	Role       string `json:"role" validate:"required,oneof=customer deliverer parter admin"`
	NotifToken string `json:"notif_token" validate:"omitempty"` // Notification Token ID
}

type user struct {
	ID          string   `json:"id"`         // A unique identifier for the user.
	FirstName   string   `json:"first_name"` // The first name of the user.
	LastName    string   `json:"last_name"`  // The last name of the user.
	Email       string   `json:"email"`      // The email address of the user.
	Phone       string   `json:"phone"`
	PhotoURL    string   `json:"phone"`
	BirthDate   string   `json:"birth_date"` // YYYY-MM-DD
	Roles       []string `json:"roles"`
	NotifToken  string   `json:"notif_token"`
	AccessToken string   `json:"-"`
}

type signUpResponse struct {
	UserID string `json:"user_id"`
}

type confirmSignUpRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Code   string `json:"code" validate:"required,number"`
}

type confirmChangeEmailRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	AccessToken string `json:"access_token" validate:"required"`
	Code        string `json:"code" validate:"required,number"`
}

type signInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type updateUserRequest struct {
	FirstName  string `json:"first_name" validate:"omitempty,alpha,min=2"` // The updated first name of the user.
	LastName   string `json:"last_name" validate:"omitempty,alpha,min=2"`  // The updated last name of the user.
	Email      string `json:"email" validate:"omitempty,email"`            // The updated email address of the user.
	Phone      string `json:"phone" validate:"omitempty,e164"`             // E.164 formatted phone number: [+] [country code] [subscriber number including area code] and can have a maximum of fifteen digits.
	BirthDate  string `json:"birth_date" validate:"omitempty,dateonly"`    // YYYY-MM-DD
	NotifToken string `json:"notif_token" validate:"omitempty"`
}

type resetPasswordRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Code     string `json:"code" validate:"required,number"`
	Password string `json:"password" validate:"required,min=6"`
}

type changePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=6"`
	NewPassword string `json:"new_password" validate:"required,min=6"`
}

type checkRequest struct {
	OrderID         string     `json:"order_id" validate:"required"`
	PartnerID       int32      `json:"partner_id" validate:"required,gt=0"`
	CustomerPhone   string     `json:"customer_phone" validate:"omitempty,e164"`
	DeliveryAddress string     `json:"delivery_address" validate:"required,min=6"`
	Products        []*product `json:"products" validate:"required"`
	TotalAmount     int64      `json:"total_amount" validate:"required,gt=0"`
	Paytype         string     `json:"paytype" validate:"required"`
}

type product struct {
	ID       int `json:"id" validate:"required"`
	Quantity int `json:"quantity" validate:"required"`
}

type confirmRequest struct {
	OrderID string `json:"order_id" validate:"required"`
}

type confirmResponse struct {
	OrderID        int64  `json:"order_id"`
	TotalAmount    int64  `json:"total_amount"`
	PartnerTitle   string `json:"partner_title"`
	PartnerBrand   string `json:"partner_brand"`
	WebcheckoutURL string `json:"webcheckout_url"`
	CallbackURL    string `json:"callback_url"`
}

type payRequest struct {
	OrderID    int64  `json:"order_id" validate:"required"`
	PaymentID  string `json:"payment_id" validate:"required"`
	PaidAmount int64  `json:"paid_amount" validate:"required"`
}

type payResponse struct {
	PaymentID string `json:"payment_id"`
}

type pickupRequest struct {
	OrderID       int64  `json:"order_id" validate:"required"`
	PickupAddress string `json:"pickup_address" validate:"required"`
}

type assignRequest struct {
	OrderID int64 `json:"order_id" validate:"required"`
}
