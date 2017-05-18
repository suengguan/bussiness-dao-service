package controllers

import (
	"dao-service/bussiness-dao-service/models"
	"dao-service/bussiness-dao-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Pod
type PodController struct {
	beego.Controller
}

// @Title GetById
// @Description get pod by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @Failure 403 :id is empty
// @router /id/:id [get]
func (this *PodController) GetById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("GetById", id)
	if id > 0 && err == nil {
		var svc service.PodService
		var pod *model.Pod
		var result []byte
		pod, err = svc.GetById(id)
		if err == nil {
			result, err = json.Marshal(pod)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "pod id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title Update
// @Description update the pod
// @Param	body		body 	models.Pod	true		"body for pod content"
// @Success 200 {object} models.Response
// @Failure 403 :id is invalid
// @router / [put]
func (this *PodController) Update() {
	var err error
	var pod model.Pod
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &pod)
	if err == nil {
		var svc service.PodService
		var result []byte
		var newPod *model.Pod
		newPod, err = svc.Update(&pod)
		if err == nil {
			result, err = json.Marshal(newPod)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug("Unmarshal data failed")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}

	this.Data["json"] = &response

	this.ServeJSON()
}
