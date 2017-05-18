package service

import (
	"fmt"
	"model"
	"strconv"

	"dao-service/bussiness-dao-service/dao"

	"github.com/astaxie/beego"
)

type PodService struct {
}

func (this *PodService) Create(pod *model.Pod, module *model.Module, job *model.Job) error {
	var err error
	var podDao = dao.NewPodDao()

	err = podDao.Create(pod)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "create pod failed!", "reason:"+err.Error())
		return err
	}

	// update rc and svc name
	pod.RcName = job.Name + "-" + module.Name + "-" + pod.Name + "-" + strconv.FormatInt(pod.Id, 36)
	pod.SvcName = pod.Name + "-" + strconv.FormatInt(pod.Id, 36)
	err = podDao.Update(pod)
	if err != nil {
		beego.Debug("update pod rc,svc name failed")
		beego.Debug(err)
		return err
	}

	return err
}

func (this *PodService) GetAll(moduleId int64) ([]*model.Pod, error) {
	var err error
	var podDao = dao.NewPodDao()
	var pods []*model.Pod

	// get pods
	pods, err = podDao.GetByModuleId(moduleId)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	for _, p := range pods {
		p.Module = nil
	}

	return pods, err
}

func (this *PodService) GetById(id int64) (*model.Pod, error) {
	var err error
	var podDao = dao.NewPodDao()
	var pod *model.Pod

	// get pod
	pod, err = podDao.GetById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "pod is not existed!")
		return nil, err
	}

	return pod, err
}

func (this *PodService) Update(pod *model.Pod) (*model.Pod, error) {
	var err error
	var podDao = dao.NewPodDao()

	err = podDao.Update(pod)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "update pod failed!", "reason:"+err.Error())
		return nil, err
	}

	return pod, err
}

func (this *PodService) DeleteById(id int64) error {
	var err error
	var podDao = dao.NewPodDao()

	// delete pod
	err = podDao.DeleteById(id)
	if err != nil {
		beego.Debug(err)
		err = fmt.Errorf("%s", "delete pod failed!", "reason:"+err.Error())
		return err
	}

	return err
}

func (this *PodService) DeleteAll(moduleId int64) error {
	var err error
	var pods []*model.Pod

	pods, err = this.GetAll(moduleId)
	if err != nil {
		return err
	}

	for _, p := range pods {
		err = this.DeleteById(p.Id)
		if err != nil {
			return err
		}
	}

	return err
}
