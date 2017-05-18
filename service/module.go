package service

import (
	"fmt"
	"model"

	"dao-service/bussiness-dao-service/dao"

	"github.com/astaxie/beego"
)

type ModuleService struct {
}

func (this *ModuleService) Create(module *model.Module, job *model.Job) error {
	var err error
	var moduleDao = dao.NewModuleDao()
	var podService PodService

	err = moduleDao.Create(module)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "create module failed!", "reason:"+err.Error())
		return err
	}

	// create pod
	var m model.Module
	m.Id = module.Id
	for _, p := range module.Pods {
		p.Module = &m
		err = podService.Create(p, module, job)
		if err != nil {
			beego.Debug(err)
			err = fmt.Errorf("%s", "create module failed!", "reason:"+err.Error())
			return err
		}
	}

	return err
}

func (this *ModuleService) GetAll(jobId int64) ([]*model.Module, error) {
	var err error
	var moduleDao = dao.NewModuleDao()
	var modules []*model.Module
	var podService PodService

	// get modules
	modules, err = moduleDao.GetByJobId(jobId)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	// get pod
	var pods []*model.Pod
	for _, m := range modules {
		m.Job = nil
		pods, err = podService.GetAll(m.Id)

		for _, p := range pods {
			m.Pods = append(m.Pods, p)
		}
	}

	return modules, err
}

func (this *ModuleService) GetById(id int64) (*model.Module, error) {
	var err error
	var moduleDao = dao.NewModuleDao()
	var module *model.Module
	var podService PodService

	// get module
	module, err = moduleDao.GetById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "module is not existed!")
		return nil, err
	}

	// get pod
	var pods []*model.Pod
	pods, err = podService.GetAll(module.Id)
	for _, p := range pods {
		module.Pods = append(module.Pods, p)
	}

	return module, err
}

func (this *ModuleService) Update(module *model.Module) error {
	var err error
	var moduleDao = dao.NewModuleDao()

	err = moduleDao.Update(module)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "update module failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *ModuleService) DeleteById(id int64) error {
	var err error
	var moduleDao = dao.NewModuleDao()
	var podService PodService

	// delete module
	err = moduleDao.DeleteById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete module failed!", "reason:"+err.Error())
		return err
	}

	// delete pod
	podService.DeleteAll(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete module failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *ModuleService) DeleteAll(jobId int64) error {
	var err error
	var modules []*model.Module

	modules, err = this.GetAll(jobId)
	if err != nil {
		return err
	}

	for _, m := range modules {
		err = this.DeleteById(m.Id)
		if err != nil {
			return err
		}
	}

	return err
}
