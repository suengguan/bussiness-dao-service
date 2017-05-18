package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"model"
)

//var cfg = beego.AppConfig

type ModuleDao struct {
	m_Orm        orm.Ormer
	m_QuerySeter orm.QuerySeter
	m_QueryTable *model.Module
}

func NewModuleDao() *ModuleDao {
	d := new(ModuleDao)

	d.m_Orm = orm.NewOrm()
	d.m_Orm.Using(cfg.String("dbname"))

	d.m_QuerySeter = d.m_Orm.QueryTable(d.m_QueryTable)
	d.m_QuerySeter.Limit(-1)

	return d
}

//add
func (this *ModuleDao) Create(module *model.Module) error {
	num, err := this.m_Orm.Insert(module)
	if err != nil {
		beego.Debug(num, err)
		return err
	}

	return err
}

//delete
func (this *ModuleDao) DeleteById(id int64) error {
	num, err := this.m_QuerySeter.Filter("ID", id).Delete()

	if err != nil {
		return err
	}

	if num < 1 {
		err = fmt.Errorf("%s", "there is no module to delete")
		return err
	}

	return err
}

// update
func (this *ModuleDao) Update(module *model.Module) error {
	num, err := this.m_Orm.Update(module)

	if err != nil {
		return err
	}

	if num < 1 {
		beego.Debug("there is no module to update")
	}

	return err
}

// find
func (this *ModuleDao) GetByJobId(jobId int64) ([]*model.Module, error) {
	var modules []*model.Module

	num, err := this.m_QuerySeter.Filter("JOB_ID", jobId).RelatedSel().All(&modules)

	if err != nil {
		beego.Debug(num, err)
		return nil, err
	}

	return modules, nil
}

func (this *ModuleDao) GetById(Id int64) (*model.Module, error) {
	var module model.Module

	err := this.m_QuerySeter.Filter("ID", Id).One(&module)

	if err != nil {
		//beego.Debug(err)
		return nil, err
	}

	return &module, nil
}
