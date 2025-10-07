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


func (adapter *MysqlAdapter) ListAll() ([]entities.User, error) {
	var users []entities.User

	query := "SELECT * FROM users" 

	prepare, err := adapter.conn.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparando consulta: %v", err)
	}
	defer prepare.Close()

	rows, err := prepare.Query()
	if err != nil {
		return nil, fmt.Errorf("error ejecutando consulta: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Password); err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}


func (adapter *MysqlAdapter)  Update(id int32, user *entities.User) error {
	query := "UPDATE users SET username = ?, password = ? WHERE id = ?"

	prepare, err := adapter.conn.Prepare(query)

	if err != nil {
		return err
	}

	_, errResult := prepare.Exec(user.Username, user.Password, id)

	if errResult != nil {
		return errResult
	}

	return nil
}