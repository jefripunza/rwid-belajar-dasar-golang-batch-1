package modules

type User struct{}

func (ref User) Route() {

	handler := UserHandler{}

	handler.CreateClient("rwid", "my-password")
	handler.Paginate(1, 10)
	handler.List()
	handler.Update("rwid", "my-password")
	handler.Block("rwid")

}

type UserHandler struct{}

func (ref UserHandler) CreateClient(username string, password string) bool {
	return true
}

func (ref UserHandler) Paginate(page int, limit int) bool {
	return true
}

func (ref UserHandler) List() bool {
	return true
}

func (ref UserHandler) Update(username string, body interface{}) bool {
	return true
}

func (ref UserHandler) Block(username string) bool {
	return true
}
