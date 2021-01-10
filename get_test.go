// @Author: abbeymart | Abi Akindele | @Created: 2020-12-24 | @Updated: 2020-12-24
// @Company: mConnect.biz | @License: MIT
// @Description: get/read records test cases

package mccrud

import (
	"fmt"
	"github.com/abbeymart/mcdb"
	"github.com/abbeymart/mctest"
	"github.com/abbeymart/mctypes"
	"testing"
	"time"

	//"time"
)

func TestGet(t *testing.T) {
	myDb := mcdb.DbConfig{
		DbType:   "postgres",
		Host:     "localhost",
		Username: "postgres",
		Password: "ab12testing",
		Port:     5432,
		DbName:   "mcdev",
		Filename: "testdb.db",
		PoolSize: 20,
		Url:      "localhost:5432",
	}
	myDb.Options = mcdb.DbConnectOptions{}

	// db-connection
	dbc, err := myDb.OpenPgxDbPool()
	// defer dbClose
	defer myDb.ClosePgxDbPool()

	// check db-connection-error
	if err != nil {
		fmt.Printf("*****db-connection-error: %v\n", err.Error())
		return
	}

	getCrudParams := mctypes.CrudParamsType{
		AppDb:       dbc.DbConn,
		TableName:   TestTable,
		UserInfo:    TestUserInfo,
		RecordIds:   GetIds,
		QueryParams: GetParams,
	}

	var getCrud = NewCrud(getCrudParams, TestCrudParamOptions)

	mctest.McTest(mctest.OptionValue{
		Name: "should get records by Ids and return success:",
		TestFunc: func() {
			var getResults []GetRecordType
			getChan := make(chan int, 1)
			var (
				id            string
				tableName     string
				logRecords    interface{}
				newLogRecords interface{}
				logBy         string
				logType       string
				logAt         time.Time
			)
			tableFieldPointers := []interface{}{&id, &tableName, &logRecords, &newLogRecords, &logBy, &logType, &logAt}
			res := getCrud.GetById(GetTableFields, getChan, tableFieldPointers)
			// compute get-records
			for <-getChan >= 0 {
				getResult := GetRecordType{
					Id:            id,
					TableName:     tableName,
					LogRecords:    logRecords,
					NewLogRecords: newLogRecords,
					LogBy:         logBy,
					LogType:       logType,
					LogAt:         logAt,
				}
				getResults = append(getResults, getResult)
			}
			//fmt.Println(res.Message, res.ResCode)
			value, _ := res.Value.(int)
			mctest.AssertEquals(t, res.Code, "success", "get-task should return code: success")
			mctest.AssertEquals(t, value, 2, "get-task-count should be: 2")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should get records by query-params and return success:",
		TestFunc: func() {
			var getResults []GetRecordType
			getChan := make(chan int, 1)
			var (
				id            string
				tableName     string
				logRecords    interface{}
				newLogRecords interface{}
				logBy         string
				logType       string
				logAt         time.Time
			)
			tableFieldPointers := []interface{}{&id, &tableName, &logRecords, &newLogRecords, &logBy, &logType, &logAt}
			res := getCrud.GetByParam(GetTableFields, getChan, tableFieldPointers)
			// compute get-records
			for <-getChan >= 0 {
				getResult := GetRecordType{
					Id:            id,
					TableName:     tableName,
					LogRecords:    logRecords,
					NewLogRecords: newLogRecords,
					LogBy:         logBy,
					LogType:       logType,
					LogAt:         logAt,
				}
				getResults = append(getResults, getResult)
			}
			//fmt.Println(res.Message, res.ResCode)
			value, _ := res.Value.(int)
			mctest.AssertEquals(t, res.Code, "success", "get-task should return code: success")
			mctest.AssertEquals(t, value > 1, true, "get-count should be > 1")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should get all records and return success:",
		TestFunc: func() {
			var getResults []GetRecordType
			getChan := make(chan int, 1)
			var (
				id            string
				tableName     string
				logRecords    interface{}
				newLogRecords interface{}
				logBy         string
				logType       string
				logAt         time.Time
			)
			tableFieldPointers := []interface{}{&id, &tableName, &logRecords, &newLogRecords, &logBy, &logType, &logAt}
			res := getCrud.GetAll(GetTableFields, getChan, tableFieldPointers)
			// compute get-records
			for <-getChan >= 0 {
				getResult := GetRecordType{
					Id:            id,
					TableName:     tableName,
					LogRecords:    logRecords,
					NewLogRecords: newLogRecords,
					LogBy:         logBy,
					LogType:       logType,
					LogAt:         logAt,
				}
				getResults = append(getResults, getResult)
			}
			//fmt.Println(res.Message, res.ResCode)
			value, _ := res.Value.(int)
			mctest.AssertEquals(t, res.Code, "success", "get-task should return code: success")
			mctest.AssertEquals(t, value > 2, true, "get-task-count should > 2")
		},
	})
	mctest.McTest(mctest.OptionValue{
		Name: "should get all records by limit/skip(offset) and return success:",
		TestFunc: func() {
			var getResults []GetRecordType
			getChan := make(chan int, 1)
			var (
				id            string
				tableName     string
				logRecords    interface{}
				newLogRecords interface{}
				logBy         string
				logType       string
				logAt         time.Time
			)
			getCrud.Skip = 0
			getCrud.Limit = 10
			tableFieldPointers := []interface{}{&id, &tableName, &logRecords, &newLogRecords, &logBy, &logType, &logAt}
			res := getCrud.GetAll(GetTableFields, getChan, tableFieldPointers)
			// compute get-records
			for <-getChan >= 0 {
				getResult := GetRecordType{
					Id:            id,
					TableName:     tableName,
					LogRecords:    logRecords,
					NewLogRecords: newLogRecords,
					LogBy:         logBy,
					LogType:       logType,
					LogAt:         logAt,
				}
				getResults = append(getResults, getResult)
			}
			//fmt.Println(res.Message, res.ResCode)
			value, _ := res.Value.(int)
			mctest.AssertEquals(t, res.Code, "success", "get-task should return code: success")
			mctest.AssertEquals(t, value > 2, true, "get-task-count should > 2")
		},
	})

	mctest.PostTestResult()

}

