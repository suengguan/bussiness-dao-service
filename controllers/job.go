package controllers

import (
	"dao-service/bussiness-dao-service/models"
	"dao-service/bussiness-dao-service/service"
	"encoding/json"
	"fmt"
	"model"

	"github.com/astaxie/beego"
)

// Operations about Job
type JobController struct {
	beego.Controller
}

// @Title Create
// @Description create job
// @Param	body		body 	models.Job	true		"body for job content"
// @Success 200 {object} models.Response
// @Failure 403 body is empty
// @router / [post]
func (this *JobController) Create() {
	var err error
	var job model.Job
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &job)
	if err == nil {
		var svc service.JobService
		var result []byte
		err = svc.Create(&job)
		if err == nil {
			result, err = json.Marshal(&job)
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

// @Title GetAll
// @Description get all project's jobs
// @Param	projectId		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @router /:projectId [get]
func (this *JobController) GetAll() {
	var err error
	var response models.Response

	var projectId int64
	projectId, err = this.GetInt64(":projectId")
	beego.Debug("GetAll", projectId)
	if projectId > 0 && err == nil {
		var svc service.JobService
		var jobs []*model.Job
		var result []byte
		jobs, err = svc.GetAll(projectId)
		if err == nil {
			result, err = json.Marshal(jobs)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "projectId is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}

// @Title GetById
// @Description get job by id
// @Param	id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.Response
// @Failure 403 :id is empty
// @router /id/:id [get]
func (this *JobController) GetById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("GetById", id)
	if id > 0 && err == nil {
		var svc service.JobService
		var job *model.Job
		var result []byte
		job, err = svc.GetById(id)
		if err == nil {
			result, err = json.Marshal(job)
			if err == nil {
				response.Status = model.MSG_RESULTCODE_SUCCESS
				response.Reason = "success"
				response.Result = string(result)
			}
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "job id is invalid")
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
// @Description update the job
// @Param	body		body 	models.Job	true		"body for user content"
// @Success 200 {object} models.Response
// @Failure 403 :id is invalid
// @router / [put]
func (this *JobController) Update() {
	var err error
	var job model.Job
	var response models.Response

	// unmarshal
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &job)
	if err == nil {
		var svc service.JobService
		var result []byte
		err = svc.Update(&job)
		if err == nil {
			result, err = json.Marshal(&job)
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

// @Title DeleteById
// @Description delete the job by id
// @Param	id		path 	int64	true		"The int you want to delete"
// @Success 200 {object} models.Response
// @Failure 403 :id is invalid
// @router /id/:id [delete]
func (this *JobController) DeleteById() {
	var err error
	var response models.Response

	var id int64
	id, err = this.GetInt64(":id")
	beego.Debug("DeleteById", id)
	if id > 0 && err == nil {
		var svc service.JobService
		err = svc.DeleteById(id)
		if err == nil {
			response.Status = model.MSG_RESULTCODE_SUCCESS
			response.Reason = "success"
			response.Result = ""
		}
	} else {
		beego.Debug(err)
		err = fmt.Errorf("%s", "job id is invalid")
	}

	if err != nil {
		response.Status = model.MSG_RESULTCODE_FAILED
		response.Reason = err.Error()
		response.RetryCount = 3
	}
	this.Data["json"] = &response

	this.ServeJSON()
}
