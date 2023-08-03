package repository

type MysqlRepository struct {
	// Db
}

func (my *MysqlRepository) Save() string {
	return "I am the mysql save method"
}
