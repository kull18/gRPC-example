package infraestructure

import (
	"database/sql"
	"fmt"
	"gRPC-Example/core"
	"gRPC-Example/src/domain/entities"
)

type MysqlAdapter struct {
	conn *sql.DB
}

func NewMysqlAdapter() *MysqlAdapter {
	conn := core.GetConn(); 

	return &MysqlAdapter{
		conn: conn,
	}
}

func (adapter *MysqlAdapter) Save(user *entities.User) error {
	query := "INSERT INTO users(username, password) VALUES (?,?)"

	prepare, err := adapter.conn.Prepare(query)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	defer prepare.Close()

	_, errResult := adapter.conn.Exec(query, user.Username, user.Password)

	if errResult != nil {
		fmt.Printf("Error to create user")
		return errResult
	}

	return nil
}