package dao

import "github.com/pkg/errors"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserByName(name string) (*User, error) {
	user := &User{}
	row := db.QueryRow("SELECT id, name from users WHERE name LIKE ?", "%"+name+"%")
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		// 数据为空时 warp error
		return nil, errors.Wrap(err, "GetUserByName Is Empty")
	}
	return user, nil
}
