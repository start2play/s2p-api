package user

type User struct {
	ID        string `bson:"_id,omitempty" gql:"id"`
	Name      string `bson:"name"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`
	Password  string `bson:"password"`
	Birthdate string `bson:"birthdate"`
	Verified  bool   `bson:"verified"`
	Type      string `bson:"type"`
}

var UserInstance = User{}

type LoginRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

var LoginRequestInstance = LoginRequest{}

type LoginResponse struct {
	Token string `bson:"token"`
	Email string `bson:"email"`
}

var LoginResponseInstance = LoginResponse{}

type DeleteResponse struct {
	DeletedCount int64  `bson:"deletedCount"`
	ID           string `bson:"id"`
}

var DeleteResponseInstance = DeleteResponse{}

type ResetPasswordRequest struct {
	Password string `bson:"password"`
	Token    string `bson:"token"`
}

var ResetPasswordInstance = ResetPasswordRequest{}

type EmailConfirmationRequest struct {
	Token string `bson:"token"`
}

var EmailConfirmationInstance = EmailConfirmationRequest{}
