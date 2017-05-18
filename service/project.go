package service

import (
	"fmt"
	"model"

	"dao-service/bussiness-dao-service/dao"

	"github.com/astaxie/beego"
)

type ProjectService struct {
}

func (this *ProjectService) Create(project *model.Project) error {
	var err error
	var projectDao = dao.NewProjectDao()

	err = projectDao.Create(project)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "create project failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *ProjectService) GetAll(userId int64) ([]*model.Project, error) {
	var err error
	var projectDao = dao.NewProjectDao()
	var projects []*model.Project
	var jobService JobService

	// get project
	projects, err = projectDao.GetByUserId(userId)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	// get job
	var jobs []*model.Job
	for _, p := range projects {
		p.User = nil
		jobs, err = jobService.GetAll(p.Id)

		for _, j := range jobs {
			p.Jobs = append(p.Jobs, j)
		}
	}

	return projects, err
}

func (this *ProjectService) GetById(id int64) (*model.Project, error) {
	var err error
	var projectDao = dao.NewProjectDao()
	var project *model.Project
	var jobService JobService

	// get project
	beego.Debug("->get project")
	project, err = projectDao.GetById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "project is not existed!")
		return nil, err
	}

	// get job
	beego.Debug("->get all jobs")
	var jobs []*model.Job
	jobs, err = jobService.GetAll(project.Id)
	for _, j := range jobs {
		project.Jobs = append(project.Jobs, j)
	}

	beego.Debug("result:", *project)

	return project, err
}

func (this *ProjectService) Update(project *model.Project) error {
	var err error
	var projectDao = dao.NewProjectDao()

	err = projectDao.Update(project)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "update project failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *ProjectService) DeleteById(id int64) error {
	var err error
	var projectDao = dao.NewProjectDao()
	var jobService JobService

	// delete project
	err = projectDao.DeleteById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete project failed!", "reason:"+err.Error())
		return err
	}

	// delete job
	jobService.DeleteAll(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete project failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *ProjectService) DeleteAll(userId int64) error {
	var err error
	var projects []*model.Project

	projects, err = this.GetAll(userId)
	if err != nil {
		return err
	}

	for _, p := range projects {
		err = this.DeleteById(p.Id)
		if err != nil {
			return err
		}
	}

	return err
}
