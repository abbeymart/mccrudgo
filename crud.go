// @Author: abbeymart | Abi Akindele | @Created: 2020-12-01 | @Updated: 2020-12-01
// @Company: mConnect.biz | @License: MIT
// @Description: Base type/function CRUD operations for PgDB

package mccrud

import (
	"encoding/json"
	"fmt"
	"github.com/abbeymart/mcauditlog"
	"github.com/abbeymart/mctypes"
)

// Crud object / struct
type Crud struct {
	mctypes.CrudParamsType
	mctypes.CrudOptionsType
	CurrentRecords []interface{}
	TransLog       mcauditlog.PgxLogParam
	HashKey        string // Unique for exactly the same query
}

// NewCrud constructor returns a new crud-instance
func NewCrud(params mctypes.CrudParamsType, options mctypes.CrudOptionsType) (crudInstance Crud) {
	// crudInstance = Crud{}
	// compute crud params
	crudInstance.AppDb = params.AppDb
	crudInstance.TableName = params.TableName
	crudInstance.UserInfo = params.UserInfo
	crudInstance.ActionParams = params.ActionParams
	crudInstance.RecordIds = params.RecordIds
	crudInstance.QueryParams = params.QueryParams
	crudInstance.SortParams = params.SortParams
	crudInstance.ProjectParams = params.ProjectParams
	crudInstance.ExistParams = params.ExistParams
	crudInstance.Token = params.Token
	crudInstance.TaskName = params.TaskName
	crudInstance.Skip = params.Skip
	crudInstance.Limit = params.Limit

	// crud options
	crudInstance.MaxQueryLimit = options.MaxQueryLimit
	crudInstance.AuditTable = options.AuditTable
	crudInstance.AccessTable = options.AccessTable
	crudInstance.RoleTable = options.RoleTable
	crudInstance.UserTable = options.UserTable
	crudInstance.AuditDb = options.AuditDb
	crudInstance.AccessDb = options.AccessDb
	crudInstance.LogAll = options.LogAll
	crudInstance.LogRead = options.LogRead
	crudInstance.LogCreate = options.LogCreate
	crudInstance.LogUpdate = options.LogUpdate
	crudInstance.LogDelete = options.LogDelete
	crudInstance.CheckAccess = options.CheckAccess // Dec 09/2020: user to implement auth as a middleware
	// Compute HashKey from TableName, QueryParams, SortParams, ProjectParams and RecordIds
	qParam, _ := json.Marshal(params.QueryParams)
	sParam, _ := json.Marshal(params.SortParams)
	pParam, _ := json.Marshal(params.ProjectParams)
	dIds, _ := json.Marshal(params.RecordIds)
	crudInstance.HashKey = params.TableName + string(qParam) + string(sParam) + string(pParam) + string(dIds)

	// Audit/TransLog instance
	crudInstance.TransLog = mcauditlog.NewAuditLogPgx(crudInstance.AuditDb, crudInstance.AuditTable)

	// Default values
	if crudInstance.AuditTable == "" {
		crudInstance.AuditTable = "audits"
	}
	if crudInstance.AccessTable == "" {
		crudInstance.AccessTable = "accesskeys"
	}
	if crudInstance.RoleTable == "" {
		crudInstance.RoleTable = "roles"
	}
	if crudInstance.UserTable == "" {
		crudInstance.UserTable = "users"
	}
	if crudInstance.ServiceTable == "" {
		crudInstance.AuditTable = "services"
	}
	if crudInstance.AuditDb == nil {
		crudInstance.AuditDb = crudInstance.AppDb
	}
	if crudInstance.AccessDb == nil {
		crudInstance.AccessDb = crudInstance.AppDb
	}
	if crudInstance.Skip < 0 {
		crudInstance.Skip = 0
	}

	if crudInstance.MaxQueryLimit == 0 {
		crudInstance.MaxQueryLimit = 10000
	}

	if crudInstance.Limit > crudInstance.MaxQueryLimit && crudInstance.MaxQueryLimit != 0 {
		crudInstance.Limit = crudInstance.MaxQueryLimit
	}

	return crudInstance
}

// String() function implementation
func (crud Crud) String() string {
	//appDb := fmt.Sprintf("Application DB: %v", crud.AppDb)
	return fmt.Sprintf(`
	Application DB: %v \n Table Name: %v \n
	`,
		crud.AppDb,
		crud.TableName)
}

// Methods => separate go-files: auth.go...
