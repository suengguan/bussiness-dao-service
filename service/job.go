package service

import (
	"fmt"
	"model"

	"dao-service/bussiness-dao-service/dao"

	"github.com/astaxie/beego"
)

type JobService struct {
}

func (this *JobService) Create(job *model.Job) error {
	var err error
	var jobDao = dao.NewJobDao()
	var moduleService ModuleService

	err = jobDao.Create(job)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "create job failed!", "reason:"+err.Error())
		return err
	}

	// create module
	var j model.Job
	j.Id = job.Id
	for _, m := range job.Modules {
		m.Job = &j
		err = moduleService.Create(m, job)
		if err != nil {
			beego.Debug(err)
			err = fmt.Errorf("%s", "create job failed!", "reason:"+err.Error())
			return err
		}
	}

	return err
}

func (this *JobService) GetAll(projectId int64) ([]*model.Job, error) {
	var err error
	var jobDao = dao.NewJobDao()
	var jobs []*model.Job
	var moduleService ModuleService

	// get jobs
	jobs, err = jobDao.GetByProjectId(projectId)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	// get module
	var modules []*model.Module
	for _, j := range jobs {
		j.Project = nil
		modules, err = moduleService.GetAll(j.Id)

		for _, m := range modules {
			j.Modules = append(j.Modules, m)
		}
	}

	return jobs, err
}

func (this *JobService) GetById(id int64) (*model.Job, error) {
	var err error
	var jobDao = dao.NewJobDao()
	var job *model.Job
	var moduleService ModuleService

	// get job
	job, err = jobDao.GetById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "job is not existed!")
		return nil, err
	}

	// get module
	var modules []*model.Module
	modules, err = moduleService.GetAll(job.Id)
	for _, m := range modules {
		job.Modules = append(job.Modules, m)
	}

	return job, err
}

func (this *JobService) Update(job *model.Job) error {
	var err error
	var jobDao = dao.NewJobDao()
	var moduleService ModuleService

	// update job
	err = jobDao.Update(job)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "update job failed!", "reason:"+err.Error())
		return err
	}

	// delete all module
	err = moduleService.DeleteAll(job.Id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "update job failed!", "reason:"+err.Error())
		return err
	}

	// create new module
	var j model.Job
	j.Id = job.Id
	for _, m := range job.Modules {
		m.Job = &j
		err = moduleService.Create(m, job)
		if err != nil {
			beego.Debug(err)
			err = fmt.Errorf("%s", "update job failed!", "reason:"+err.Error())
			return err
		}
	}

	return err
}

func (this *JobService) DeleteById(id int64) error {
	var err error
	var jobDao = dao.NewJobDao()
	var moduleService ModuleService

	// delete job
	err = jobDao.DeleteById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete job failed!", "reason:"+err.Error())
		return err
	}

	// delete module
	moduleService.DeleteAll(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete job failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *JobService) DeleteAll(projectId int64) error {
	var err error
	var jobs []*model.Job

	jobs, err = this.GetAll(projectId)
	if err != nil {
		return err
	}

	for _, j := range jobs {
		err = this.DeleteById(j.Id)
		if err != nil {
			return err
		}
	}

	return err
}
