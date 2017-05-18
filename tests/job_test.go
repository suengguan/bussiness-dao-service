package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"

	_ "dao-service/bussiness-dao-service/routers"
	"github.com/astaxie/beego/orm"
	"model"
)

const (
	job_base_url = "http://localhost:8080/v1/job"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:corex123@tcp(localhost:3306)/PME?charset=utf8")

	o := orm.NewOrm()
	o.Using("PME")

	var maps []orm.Params
	num, err := o.Raw("SELECT ID FROM PROJECT_T WHERE ID = ?", 1).Values(&maps)

	if err != nil {
		fmt.Println("get project failed!", err)
		return
	}

	if num == 0 {
		// create project
		_, err := o.Raw("insert into PROJECT_T(ID,USER_ID) values(1,1)").Exec()
		if err != nil {
			fmt.Println("insert project failed!", err)
			return
		}
	} else if num == 1 {
		// project is existed, nothing todo
		fmt.Println("project is already exited")
		return
	} else {
		// error
		fmt.Println("get project failed!", err, num)
		return
	}
}

func Test_Job_Create(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte
	var requestBody []byte

	var project model.Project
	var job model.Job
	project.Id = 1
	job.Id = 0
	job.Name = "job"
	job.Description = "this is job test"
	job.Project = &project

	// post create job
	requestBody, err = json.Marshal(&job)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	res, err = http.Post(job_base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_Job_GetAll(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	res, err = http.Get(job_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_Job_GetById(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get all jobs
	res, err = http.Get(job_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var jobs []*model.Project
	json.Unmarshal(([]byte)(response.Result), &jobs)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(jobs) <= 0 {
		t.Log("error : ", "there is no job to operate!")
		return
	}

	// get job by id
	res, err = http.Get(job_base_url + "/id/" + strconv.FormatInt(jobs[0].Id, 10))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_Job_Update(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte
	var requestData []byte

	// get all
	res, err = http.Get(job_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var jobs []*model.Job
	json.Unmarshal(([]byte)(response.Result), &jobs)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(jobs) <= 0 {
		t.Log("error : ", "there is no job to update!")
		return
	}

	var project model.Project
	project.Id = 1
	jobs[0].Project = &project
	jobs[0].Name = "job-update"
	jobs[0].Description = "this is job test-update"

	// put update job
	requestData, err = json.Marshal(&jobs[0])
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	// put
	client := http.Client{}
	req, _ := http.NewRequest("PUT", job_base_url, strings.NewReader(string(requestData)))

	res, err = client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Log("erro : ", err)
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}

func Test_Job_DeleteById(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get all
	res, err = http.Get(job_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	t.Log(string(resBody))

	var response model.Response
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var jobs []*model.Project
	json.Unmarshal(([]byte)(response.Result), &jobs)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(jobs) <= 0 {
		t.Log("error : ", "there is no job to delete!")
		return
	}

	// delete
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", job_base_url+"/id/"+strconv.FormatInt(jobs[0].Id, 10), nil)

	res, err = client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)

	if err != nil {
		t.Log("erro : ", err)
		return
	}

	t.Log(string(resBody))

	response = model.Response{}
	json.Unmarshal(resBody, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if response.Reason == "success" {
		t.Log("PASS OK")
	} else {
		t.Log("ERROR:", response.Reason)
		t.FailNow()
	}
}
