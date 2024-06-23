package domain

type LoginCredential struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

type TokenData struct {
	Id int64
}

type UserData struct {
	LoginCredential
	FullName string
}
type UserDataWithID struct {
	Id int64
	UserData
}
