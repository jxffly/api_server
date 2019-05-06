package user

import (
	"strconv"

	. "apiserver_demos/demo07/handler"
	"apiserver_demos/demo07/model"
	"apiserver_demos/demo07/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	log.Infof("the delete id is %d", userId)
	if userId == 0 {
		SendResponse(c, errno.ErrNeedIdByDelete, nil)
		return
	}
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
