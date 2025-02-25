package handler

import (
	"douyin-app/domain"
	"douyin-app/repository"
	"douyin-app/service"
	"douyin-app/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"log"
	"strings"
)

type RegisterReq struct {
	Username string `validate:"required,max=32"`
	Password string `validate:"required,min=6,max=32"`
}

func RegisterHandler(c *gin.Context) {
	//解析,处理参数
	var req RegisterReq
	var resp util.HttpResponse
	req.Username = c.Query("username")
	req.Password = c.Query("password")
	req.Username = strings.TrimSpace(req.Username) //去首尾空格

	err := util.Validate.Struct(req) // 执行验证
	if err != nil {
		resp.StatusCode = util.ParamError
		util.MakeResponse(c, &resp)
		return
	}
	userId, token, err := service.Register(c, req.Username, req.Password)
	resp.ReturnVal = map[string]interface{}{
		"user_id": userId,
		"token":   token,
	}
	if err != nil {
		if errors.Is(err, util.ErrUserExisted) { //用户已存在
			resp.StatusCode = util.UserExisted
			util.MakeResponse(c, &resp)
			return
		}

		resp.StatusCode = util.InternalServerError
		util.MakeResponse(c, &resp)
		return
	}
	resp.StatusCode = util.Success
	util.MakeResponse(c, &resp)
	return
}

func LoginHandler(c *gin.Context) {
	//登录请求和注册请求是一样的
	var req RegisterReq
	var resp util.HttpResponse
	req.Username = c.Query("username")
	req.Password = c.Query("password")
	//去username首尾空格
	req.Username = strings.TrimSpace(req.Username)

	validate := validator.New() // 创建验证器
	err := validate.Struct(req) // 执行验证
	if err != nil {
		resp.StatusCode = util.ParamError
		util.MakeResponse(c, &resp)
		return
	}

	userId, token, err := service.Login(c, req.Username, req.Password)
	resp.ReturnVal = map[string]interface{}{
		"user_id": userId,
		"token":   token,
	}
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { //用户未注册
			resp.StatusCode = util.UserNotExist
			util.MakeResponse(c, &resp)
			return
		}

		if errors.Is(err, util.ErrWrongPassword) { //用户密码错误
			resp.StatusCode = util.WrongPassword
			util.MakeResponse(c, &resp)
			return
		}

		resp.StatusCode = util.InternalServerError
		util.MakeResponse(c, &resp)
		return
	}
	resp.StatusCode = util.Success
	util.MakeResponse(c, &resp)
	return
}

func GetUserInfoHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		log.Println("GetUserInfoHandler Token <nil>")
		util.MakeResponse(c, &util.HttpResponse{
			StatusCode: util.ParamError,
		})
	}

	userId, err := util.ParseToken(token)
	if err != nil {
		log.Println("GetUserInfoHandler ParseToken Failed", err)
		util.MakeResponse(c, &util.HttpResponse{
			StatusCode: util.WrongAuth,
		})
	}

	userIdStr := c.Query("user_id")
	if userIdStr != fmt.Sprintf("%d", userId) {
		log.Println("user_id not correct")
		util.MakeResponse(c, &util.HttpResponse{
			StatusCode: util.WrongAuth,
		})
	}

	user, err := repository.GetUserRepository().FindByUserId(c, userId)
	if err != nil {
		log.Println("GetUserRepository().FindByUserId Failed", err)
		util.MakeResponse(c, &util.HttpResponse{
			StatusCode: util.InternalServerError,
		})
	}

	author := domain.FillAuthor(user)
	util.MakeResponse(c, &util.HttpResponse{
		StatusCode: util.Success,
		ReturnVal: map[string]interface{}{
			"user": author,
		},
	})
}
