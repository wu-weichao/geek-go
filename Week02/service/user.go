package service

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"runtime"
	"week02/dao"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	user, err := dao.GetUserByName(name)
	if err != nil {
		// 追加消息测试
		err = errors.WithMessage(err, "This is service error message test")
		// 打印测试
		w.Write([]byte(fmt.Sprintf("stack trace \n%+v\n", err)))
		//stack trace
		//sql: no rows in result set
		//	GetUserByName Is Empty
		// ...
		//	stack info
		// ...
		//	This is service error message test
		return
	}
	res, _ := json.Marshal(user)
	_, _ = w.Write(res)
	return
}
