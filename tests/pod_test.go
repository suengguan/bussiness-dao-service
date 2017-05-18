package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	_ "dao-service/bussiness-dao-service/routers"
	"github.com/astaxie/beego/orm"
	"model"
)

const (
	pod_base_url = "http://localhost:8080/v1/pod"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:corex123@tcp(localhost:3306)/PME?charset=utf8")

	o := orm.NewOrm()
	o.Using("PME")

	var maps []orm.Params
	num, err := o.Raw("SELECT ID FROM POD_T WHERE ID = ?", 1).Values(&maps)

	if err != nil {
		fmt.Println("get pod failed!", err)
		return
	}

	if num == 0 {
		// create pod
		_, err := o.Raw("insert into POD_T(ID, MODULE_ID) values(1, 1)").Exec()
		if err != nil {
			fmt.Println("insert pod failed!", err)
			return
		}
		fmt.Println("create pod success!")
	} else if num == 1 {
		// pod is existed, nothing todo
		fmt.Println("pod is already exited")
		return
	} else {
		// error
		fmt.Println("get pod failed!", err, num)
		return
	}
}

func Test_Pod_GetById(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get pod by id
	res, err = http.Get(pod_base_url + "/id/1")
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

func Test_Pod_Update(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte
	var requestData []byte

	var module model.Module
	module.Id = 1
	var pod model.Pod
	pod.Module = &module
	pod.Name = "pod-update"

	// put update pod
	requestData, err = json.Marshal(&pod)
	if err != nil {
		t.Log("erro : ", err)
		return
	}

	// put
	client := http.Client{}
	req, _ := http.NewRequest("PUT", pod_base_url, strings.NewReader(string(requestData)))

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
