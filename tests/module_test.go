package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	_ "dao-service/bussiness-dao-service/routers"
	"github.com/astaxie/beego/orm"
	"model"
)

const (
	module_base_url = "http://localhost:8080/v1/module"
)

func init() {
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:corex123@tcp(localhost:3306)/PME?charset=utf8")

	o := orm.NewOrm()
	o.Using("PME")

	var maps []orm.Params
	num, err := o.Raw("SELECT ID FROM MODULE_T WHERE ID = ?", 1).Values(&maps)

	if err != nil {
		fmt.Println("get module failed!", err)
		return
	}

	if num == 0 {
		// create module
		_, err := o.Raw("insert into MODULE_T(ID,JOB_ID) values(1,1)").Exec()
		if err != nil {
			fmt.Println("insert module failed!", err)
			return
		}
		fmt.Println("create module success!")
	} else if num == 1 {
		// module is existed, nothing todo
		fmt.Println("module is already exited")
		return
	} else {
		// error
		fmt.Println("get module failed!", err, num)
		return
	}
}

func Test_Module_GetById(t *testing.T) {
	var res *http.Response
	var err error
	var resBody []byte

	// get job by id
	res, err = http.Get(module_base_url + "/id/1")
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
