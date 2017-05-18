package controllers

import (
	"dao-service/bussiness-dao-service/models"
	"dao-service/bussiness-dao-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Module
type ModuleController struct {
	beego.Controller
}

// @Title GetById
// @Description get module by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @Failure 403 :id is empty
// @router /id/:id [get]
func (this *ModuleController) GetById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("GetById", id)
	if id > 0 && err == nil {
		var svc service.ModuleService
		var module *model.Module
		var result []byte
		module, err = svc.GetById(id)
		if err == nil {
			result, err = json.Marshal(module)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "module id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
