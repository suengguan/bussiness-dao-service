package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "dao-service/bussiness-dao-service/routers"
	"github.com/astaxie/beego/orm"
	"model"
)

const (
	project_base_url = "http://localhost:8080/v1/project"
)

func Test_Project_Create(t *testing.T) {
	// create user
	o := orm.NewOrm()
	o.Using("PME")

	var maps []orm.Params
	num, err := o.Raw("SELECT ID FROM USER_T WHERE ID = ?", 1).Values(&maps)

	if err != nil {
		t.Log("get user failed!", err)
		return
	}

	if num == 0 {
		// create user
		_, err := o.Raw("insert into USER_T(ID) values(1)").Exec()
		if err != nil {
			t.Log("insert user failed!", err)
			return
		}
		t.Log("create project success!")
	} else if num == 1 {
		// user is existed, nothing todo
		t.Log("user is already exited")
		return
	} else {
		// error
		t.Log("get user failed!", err, num)
		return
	}

	// create project
	var project model.Project
	var user model.User
	user.Id = 1
	project.Id = 0
	project.Name = "project"
	project.CreatedBy = user.Id
	project.CreatedAt = time.Now().Unix()
	project.Description = "this is project test"
	project.User = &user

	// post create action
	requestData, err := json.Marshal(&project)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	res, err := http.Post(project_base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
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

func Test_Project_GetAll(t *testing.T) {
	res, err := http.Get(project_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
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

func Test_Project_GetById(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get all project
	res, err = http.Get(project_base_url + "/1")
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

	var projects []*model.Project
	json.Unmarshal(([]byte)(response.Result), &projects)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(projects) <= 0 {
		t.Log("error : ", "there is no project to operate!")
		return
	}

	// get project by id
	res, err = http.Get(project_base_url + "/id/" + strconv.FormatInt(projects[0].Id, 10))
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

func Test_Project_Update(t *testing.T) {
	// get all
	res1, err := http.Get(project_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res1.Body.Close()
	resBody1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	t.Log(string(resBody1))

	var response model.Response
	json.Unmarshal(resBody1, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var projects []*model.Project
	json.Unmarshal(([]byte)(response.Result), &projects)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(projects) <= 0 {
		t.Log("error : ", "there is no project to update!")
		return
	}

	var user model.User
	user.Id = 1
	projects[0].User = &user
	projects[0].Name = "project-update"
	projects[0].Description = "this is project test-update"

	// put update project
	requestData, err := json.Marshal(&projects[0])
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	// put
	client := http.Client{}
	req, _ := http.NewRequest("PUT", project_base_url, strings.NewReader(string(requestData)))

	res, err := client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

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

func Test_Project_DeleteById(t *testing.T) {
	// get all
	res1, err := http.Get(project_base_url + "/1")
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	defer res1.Body.Close()
	resBody1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		t.Log("erro : ", err)
		return
	}
	t.Log(string(resBody1))

	var response model.Response
	json.Unmarshal(resBody1, &response)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	var projects []*model.Project
	json.Unmarshal(([]byte)(response.Result), &projects)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(projects) <= 0 {
		t.Log("error : ", "there is no project to delete!")
		return
	}

	// delete
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", project_base_url+"/id/"+strconv.FormatInt(projects[0].Id, 10), nil)

	res, err := client.Do(req)

	if err != nil {
		// handle error
		t.Log("erro : ", err)
		return
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)

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

func Test_Project_DeleteAll(t *testing.T) {
	var res *http.Response
	var err error
	var requestData []byte
	var resBody []byte

	// create projects
	for i := 0; i < 2; i++ {
		var project model.Project
		var user model.User
		user.Id = 1
		project.Id = 0
		project.Name = "project"
		project.CreatedBy = user.Id
		project.CreatedAt = time.Now().Unix()
		project.Description = "this is project test"
		project.User = &user

		// post create action
		requestData, err = json.Marshal(&project)
		if err != nil {
			t.Log("erro : ", err)
			return
		}

		res, err = http.Post(project_base_url+"/", "application/x-www-form-urlencoded", bytes.NewBuffer(requestData))
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
	}

	// get all projects
	res, err = http.Get(project_base_url + "/1")
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

	var projects []*model.Project
	json.Unmarshal(([]byte)(response.Result), &projects)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	if len(projects) <= 0 {
		t.Log("error : ", "there is no project to delete!")
		return
	}

	// delete
	client := http.Client{}
	req, _ := http.NewRequest("DELETE", project_base_url+"/1", nil)

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
