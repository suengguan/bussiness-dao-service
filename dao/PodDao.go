package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"model"
)

//var cfg = beego.AppConfig

type PodDao struct {
	m_Orm        orm.Ormer
	m_QuerySeter orm.QuerySeter
	m_QueryTable *model.Pod
}

func NewPodDao() *PodDao {
	d := new(PodDao)

	d.m_Orm = orm.NewOrm()
	d.m_Orm.Using(cfg.String("dbname"))

	d.m_QuerySeter = d.m_Orm.QueryTable(d.m_QueryTable)
	d.m_QuerySeter.Limit(-1)

	return d
}

//add
func (this *PodDao) Create(pod *model.Pod) error {
	num, err := this.m_Orm.Insert(pod)
	if err != nil {
		beego.Debug(num, err)
		return err
	}

	return err
}

//delete
func (this *PodDao) DeleteById(id int64) error {
	num, err := this.m_QuerySeter.Filter("ID", id).Delete()

	if err != nil {
		return err
	}

	if num < 1 {
		err = fmt.Errorf("%s", "there is no pod to delete")
		return err
	}

	return err
}

// update
func (this *PodDao) Update(pod *model.Pod) error {
	num, err := this.m_Orm.Update(pod)

	if err != nil {
		return err
	}

	if num < 1 {
		beego.Debug("there is no pod to update")
	}

	return err
}

// find
func (this *PodDao) GetByModuleId(moduleId int64) ([]*model.Pod, error) {
	var pods []*model.Pod

	num, err := this.m_QuerySeter.Filter("MODULE_ID", moduleId).RelatedSel().All(&pods)

	if err != nil {
		beego.Debug(num, err)
		return nil, err
	}

	return pods, nil
}

func (this *PodDao) GetById(Id int64) (*model.Pod, error) {
	var pod model.Pod

	err := this.m_QuerySeter.Filter("ID", Id).One(&pod)

	if err != nil {
		//beego.Debug(err)
		return nil, err
	}

	return &pod, nil
}
