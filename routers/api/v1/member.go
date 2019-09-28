package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"golang_demo/models"
	"golang_demo/pkg/e"
	"golang_demo/pkg/setting"
	"golang_demo/pkg/util"
	"net/http"
)

func GetMembers(c *gin.Context) {
	param := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetMembers(util.GetRestFulPage(c), setting.PageSize, param)
	data["total"] = models.GetMembersTotal(param)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetMember(c *gin.Context) {
	id := c.Param("member_id")

	param := make(map[string]interface{})
	data := make(map[string]interface{})

	if id != "" {
		param["member_id"] = id
	}

	code := e.SUCCESS

	data["lists"] = models.GetMembers(util.GetRestFulPage(c), setting.PageSize, param)
	data["total"] = models.GetMembersTotal(param)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddMember(c *gin.Context) {
	account := c.PostForm("account")
	valid := validation.Validation{}
	valid.Required(account, "account").Message("名稱不能為空")
	valid.MaxSize(account, 100, "account").Message("名稱最100字元")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistMemberByName(account) {
			code = e.SUCCESS
			models.AddMember(account)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func EditMember(c *gin.Context) {
	email := c.PostForm("email")
	real_name := c.PostForm("real_name")
	id := c.Param("member_id")

	param := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(id, "member_id").Message("member_id 不能為空")
	valid.Required(email, "email").Message("email 不能為空")
	valid.Email(email, "email").Message("email 格式錯誤")

	param["email"] = email
	param["real_name"] = real_name

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistMemberById(id) {
			code = e.SUCCESS
			models.EditMember(id, param)
			print(id)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

func DeleteMember(c *gin.Context) {
	id := c.Param("member_id")

	valid := validation.Validation{}
	valid.Required(id, "member_id").Message("member_id 不能為空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistMemberById(id) {
			code = e.SUCCESS
			models.DeleteMemberByID(id)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
