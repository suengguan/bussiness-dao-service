package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"model"
)

var cfg = beego.AppConfig

type ProjectDao struct {
	m_Orm        orm.Ormer
	m_QuerySeter orm.QuerySeter
	m_QueryTable *model.Project
}

func NewProjectDao() *ProjectDao {
	d := new(ProjectDao)

	d.m_Orm = orm.NewOrm()
	d.m_Orm.Using(cfg.String("dbname"))

	d.m_QuerySeter = d.m_Orm.QueryTable(d.m_QueryTable)
	d.m_QuerySeter.Limit(-1)

	return d
}

//add
func (this *ProjectDao) Create(project *model.Project) error {
	num, err := this.m_Orm.Insert(project)
	if err != nil {
		beego.Debug(num, err)
		return err
	}

	return err
}

//delete
func (this *ProjectDao) DeleteById(id int64) error {
	num, err := this.m_QuerySeter.Filter("ID", id).Delete()

	if err != nil {
		return err
	}

	if num < 1 {
		err = fmt.Errorf("%s", "there is no project to delete")
		return err
	}

	return err
}

// update
func (this *ProjectDao) Update(project *model.Project) error {
	num, err := this.m_Orm.Update(project)

	if err != nil {
		return err
	}

	if num < 1 {
		beego.Debug("there is no project to update")
	}

	return err
}

// find
func (this *ProjectDao) GetByUserId(userId int64) ([]*model.Project, error) {
	var projects []*model.Project

	num, err := this.m_QuerySeter.Filter("USER_ID", userId).RelatedSel().All(&projects)

	if err != nil {
		beego.Debug(num, err)
		return nil, err
	}

	return projects, nil
}

func (this *ProjectDao) GetById(Id int64) (*model.Project, error) {
	var project model.Project

	err := this.m_QuerySeter.Filter("ID", Id).One(&project)

	if err != nil {
		//beego.Debug(err)
		return nil, err
	}

	return &project, nil
}
