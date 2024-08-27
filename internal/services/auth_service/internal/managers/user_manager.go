package managers

type DatabaseManager interface {
	Add()
	GetById()
	UpdateById()
	DeleteById()
}

type UserManager struct {

}

type UserRecord struct {
	UserName string
	Password string
}