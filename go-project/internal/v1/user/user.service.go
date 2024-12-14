package user

import (
	database "backend/database/pg"
	"errors"
	"fmt"
)

// Create user
func Singup(body TypeSingup) (interface{}, error) {
	checkQuery := fmt.Sprintf(`Select id,email,phone from %s where email=%s`, database.UserTable.TableName, body.Email)
	checkUser, err := database.PgService.Query(checkQuery)
	fmt.Print(checkQuery)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("something went wrong")
	}
	if checkUser != nil {
		return nil, errors.New("email/phone number already exists")
	}
	return "", nil
}
