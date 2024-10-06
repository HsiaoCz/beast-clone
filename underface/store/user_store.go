package store

type UserStoreInter interface {
	CreateUser()
}

type UserStore struct{}

func UserStoreInit() *UserStore {
	return &UserStore{}
}
