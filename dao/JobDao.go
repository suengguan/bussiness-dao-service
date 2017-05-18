package dao

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"model"
)

//var cfg = beego.AppConfig

type JobDao struct {
	m_Orm        orm.Ormer
	m_QuerySeter orm.QuerySeter
	m_QueryTable *model.Job
}

func NewJobDao() *JobDao {
	d := new(JobDao)

	d.m_Orm = orm.NewOrm()
	d.m_Orm.Using(cfg.String("dbname"))

	d.m_QuerySeter = d.m_Orm.QueryTable(d.m_QueryTable)
	d.m_QuerySeter.Limit(-1)

	return d
}

//add
func (this *JobDao) Create(job *model.Job) error {
	num, err := this.m_Orm.Insert(job)
	if err != nil {
		beego.Debug(num, err)
		return err
	}

	return err
}

//delete
func (this *JobDao) DeleteById(id int64) error {
	num, err := this.m_QuerySeter.Filter("ID", id).Delete()

	if err != nil {
		return err
	}

	if num < 1 {
		err = fmt.Errorf("%s", "there is no job to delete")
		return err
	}

	return err
}

// update
func (this *JobDao) Update(job *model.Job) error {
	num, err := this.m_Orm.Update(job)

	if err != nil {
		return err
	}

	if num < 1 {
		beego.Debug("there is no job to update")
	}

	return err
}

// find
func (this *JobDao) GetByProjectId(projectId int64) ([]*model.Job, error) {
	var jobs []*model.Job

	num, err := this.m_QuerySeter.Filter("PROJECT_ID", projectId).RelatedSel().All(&jobs)

	if err != nil {
		beego.Debug(num, err)
		return nil, err
	}

	return jobs, err
}

func (this *JobDao) GetById(Id int64) (*model.Job, error) {
	var job model.Job

	err := this.m_QuerySeter.Filter("ID", Id).One(&job)

	if err != nil {
		//beego.Debug(err)
		return nil, err
	}

	return &job, err
}
