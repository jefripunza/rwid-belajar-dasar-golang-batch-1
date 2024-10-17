package modules

type Auth struct{}

func (ref Auth) Route() {
	handler := AuthHandler{}

	handler.Login("rwid", "my-password")
	handler.Logout("rwid")
	handler.TokenValidation("dufgdksjhgfdshgfih")

}

type AuthHandler struct{}

func (ref AuthHandler) Login(username string, password string) bool {
	return true
}

func (ref AuthHandler) Logout(username string) bool {
	return true
}

func (ref AuthHandler) TokenValidation(token string) bool {
	return true
}
