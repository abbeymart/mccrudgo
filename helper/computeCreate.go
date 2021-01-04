// @Author: abbeymart | Abi Akindele | @Created: 2020-12-08 | @Updated: 2020-12-08
// @Company: mConnect.biz | @License: MIT
// @Description: compute create-SQL script, for bulk/copy insert operation

package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/abbeymart/mctypes"
	"github.com/asaskevich/govalidator"
	"strings"
	"time"
)

func errMessage(errMsg string) (mctypes.CreateQueryResponseType, error) {
	return mctypes.CreateQueryResponseType{
		CreateQuery: "",
		FieldNames:  nil,
		FieldValues: nil,
	}, errors.New(errMsg)
}

// ComputeCreateScript computes insert SQL script. It returns createScripts []string, fieldNames []string and err error
func ComputeCreateQuery(tableName string, tableFields []string, actionParams mctypes.ActionParamsType) ([]string, error) {
	if tableName == "" || len(actionParams) < 1 || len(tableFields) < 1 {
		return nil, errors.New("table-name, action-params and table-fields are required for the create operation")
	}
	var insertQuery []string
	// value-computation for each of the actionParams' records must match the tableFields
	// compute create script for all the records in actionParams
	var itemQuery = fmt.Sprintf("INSERT INTO %v(%v)", tableName, strings.Join(tableFields, ", "))

	// compute create values from actionParams/records
	for recNum, rec := range actionParams {
		// initial item-values-computation variables
		var itemValues = " VALUES("
		//recLength := len(rec)
		fieldLength := len(tableFields)
		recCount := 0
		for _, fieldName := range tableFields {
			fieldValue := rec[fieldName]
			// check for the required field in each record
			if fieldValue == nil {
				return nil, errors.New(fmt.Sprintf("Record #%v [%#v]: required field_name[%v] has field_value of %v ", recNum, rec, fieldName, fieldValue))
			}
			recCount += 1
			// update recFieldValues by fieldValue-type
			var currentFieldValue interface{}
			switch fieldValue.(type) {
			case time.Time:
				if fVal, ok := fieldValue.(time.Time); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					//currentFieldValue = fmt.Sprintf("'%v'", fVal)
					currentFieldValue = "'" + fVal.Format("2006-01-02 15:04:05.000000") + "'"
				}
			case string:
				if fVal, ok := fieldValue.(string); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					if govalidator.IsJSON(fVal) {
						if fValue, err := govalidator.ToJSON(fieldValue); err != nil {
							return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
						} else {
							fmt.Printf("string-toJson-value: %v\n\n", fValue)
							currentFieldValue = "'" + fValue + "'"
						}
					} else {
						currentFieldValue = "'" + fVal + "'"
					}
				}
			case bool:
				if fVal, ok := fieldValue.(bool); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case int8:
				if fVal, ok := fieldValue.(int8); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case int16:
				if fVal, ok := fieldValue.(int16); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case int32:
				if fVal, ok := fieldValue.(int32); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case int64:
				if fVal, ok := fieldValue.(int64); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case int:
				if fVal, ok := fieldValue.(int); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case uint8:
				if fVal, ok := fieldValue.(uint8); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case uint16:
				if fVal, ok := fieldValue.(uint16); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case uint32:
				if fVal, ok := fieldValue.(uint32); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case uint64:
				if fVal, ok := fieldValue.(uint64); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case uint:
				if fVal, ok := fieldValue.(uint); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case float32:
				if fVal, ok := fieldValue.(float32); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case float64:
				if fVal, ok := fieldValue.(float64); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case []string:
				if fVal, ok := fieldValue.([]string); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case []int:
				if fVal, ok := fieldValue.([]int); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case []float32:
				if fVal, ok := fieldValue.([]float32); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case []float64:
				if fVal, ok := fieldValue.([]float64); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			case []struct{}:
				if fVal, ok := fieldValue.([]struct{}); !ok {
					return nil, errors.New(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					currentFieldValue = fVal
				}
			default:
				// json-stringify fieldValue
				if fVal, err := json.Marshal(fieldValue); err != nil {
					return nil, errors.New(fmt.Sprintf("Unknown or Unsupported field-value type: %v", err.Error()))
				} else {
					//fmt.Printf("***is-json-value***: %v\n\n", govalidator.IsJSON(string(fVal)))
					fmt.Printf("default-toJson-value: %v\n\n", string(fVal))
					currentFieldValue = "'" + string(fVal) + "'"
				}
			}
			// add itemValue
			itemValues += fmt.Sprintf("%v", currentFieldValue)
			if fieldLength > 1 && recCount < fieldLength {
				itemValues += ", "
			}
		}
		// close itemValues for the current-record
		itemValues += ")"
		// update insertQuery with the recordItem
		insertQuery = append(insertQuery, itemQuery+itemValues)
		// reset itemValues for the next record iteration
		itemValues = " VALUES("
	}
	// result
	return insertQuery, nil
}

// ComputeCreateScript computes insert SQL script. It returns createScripts []string, fieldNames []string and err error
func ComputeCreateBatchQuery(tableName string, tableFields []string, actionParams mctypes.ActionParamsType) (mctypes.CreateQueryResponseType, error) {
	if tableName == "" || len(actionParams) < 1 || len(tableFields) < 1 {
		return errMessage("table-name, action-params and table-fields are required for the create operation")
	}
	var insertQuery string
	var fValues [][]interface{} // fieldValues array of ValueParamType
	// value-computation for each of the actionParams' records must match the tableFields
	// compute create script for all the create-task, with value-placeholders
	var itemQuery = fmt.Sprintf("INSERT INTO %v(", tableName)
	var itemValuePlaceholder = " VALUES("
	fieldsLength := len(tableFields)
	for fieldIndex, fieldName := range tableFields {
		itemQuery += fmt.Sprintf(" %v", fieldName)
		itemValuePlaceholder += fmt.Sprintf(" $%v", fieldIndex+1)
		if fieldsLength > 1 && fieldIndex < fieldsLength-1 {
			itemQuery += ", "
			itemValuePlaceholder += ", "
		}
	}
	// close item-script/value-placeholder
	itemQuery += " )"
	itemValuePlaceholder += " )"
	// add/append item-script & value-placeholder to the createScripts
	insertQuery = itemQuery + itemValuePlaceholder

	// compute create values from actionParams
	for recNum, rec := range actionParams {
		// initial item-values-computation variables
		var recFieldValues []interface{}
		for _, fieldName := range tableFields {
			fieldValue := rec[fieldName]
			// check for required field in each record
			if fieldValue == nil {
				return errMessage(fmt.Sprintf("Record #%v [%#v]: required field_name[%v] has field_value of %v ", recNum, rec, fieldName, fieldValue))
			}
			// update recFieldValues by fieldValue-type
			switch fieldValue.(type) {
			case time.Time:
				if fVal, ok := fieldValue.(time.Time); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, "'" + fVal.Format("2006-01-02 15:04:05.000000") + "'")
				}
			case string:
				if fVal, ok := fieldValue.(string); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					if govalidator.IsJSON(fVal) {
						//fmt.Printf("string-toJson-value: %v\n\n", fVal)
						//recFieldValues = append(recFieldValues, "'" + fVal + "'")
						if fValue, err := govalidator.ToJSON(fieldValue); err != nil {
							return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
						} else {
							fmt.Printf("string-toJson-value: %v\n\n", fValue)
							recFieldValues = append(recFieldValues, "'" + fValue + "'")
						}
					} else {
						recFieldValues = append(recFieldValues, "'" + fVal + "'")
					}
				}
			case bool:
				if fVal, ok := fieldValue.(bool); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int8:
				if fVal, ok := fieldValue.(int8); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int16:
				if fVal, ok := fieldValue.(int16); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int32:
				if fVal, ok := fieldValue.(int32); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int64:
				if fVal, ok := fieldValue.(int64); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int:
				if fVal, ok := fieldValue.(int); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint8:
				if fVal, ok := fieldValue.(uint8); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint16:
				if fVal, ok := fieldValue.(uint16); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint32:
				if fVal, ok := fieldValue.(uint32); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint64:
				if fVal, ok := fieldValue.(uint64); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint:
				if fVal, ok := fieldValue.(uint); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case float32:
				if fVal, ok := fieldValue.(float32); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case float64:
				if fVal, ok := fieldValue.(float64); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []string:
				if fVal, ok := fieldValue.([]string); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []int:
				if fVal, ok := fieldValue.([]int); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []float32:
				if fVal, ok := fieldValue.([]float32); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []float64:
				if fVal, ok := fieldValue.([]float64); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []struct{}:
				if fVal, ok := fieldValue.([]struct{}); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			default:
				// json-stringify fieldValue
				if fVal, err := json.Marshal(fieldValue); err != nil {
					return errMessage(fmt.Sprintf("Unknown or Unsupported field-value type: %v", err.Error()))
				} else {
					//fmt.Printf("***is-json-value***: %v\n\n", govalidator.IsJSON(string(fVal)))
					//fmt.Printf("toJson-value: %v\n\n", string(fVal))
					recFieldValues = append(recFieldValues, "'" + string(fVal) + "'" )
				}
			}
		}
		// update fieldValues
		fValues = append(fValues, recFieldValues)
		// re-initialise recFieldValues, for next update
		recFieldValues = []interface{}{}
	}
	// result
	return mctypes.CreateQueryResponseType{
		CreateQuery: insertQuery,
		FieldNames:  tableFields,
		FieldValues: fValues,
	}, nil
}

// ComputeCreateScript computes insert SQL script. It returns createScripts []string, fieldNames []string and err error
func ComputeCreateCopyQuery(tableName string, tableFields []string, actionParams mctypes.ActionParamsType) (mctypes.CreateQueryResponseType, error) {
	if tableName == "" || len(actionParams) < 1 || len(tableFields) < 1 {
		return errMessage("table-name, action-params and table-fields are required for the create operation")
	}
	var insertQuery string
	var fValues [][]interface{} // fieldValues array of ValueParamType
	// value-computation for each of the actionParams' records must match the tableFields
	// compute create script for all the create-task, with value-placeholders
	var itemQuery = fmt.Sprintf("INSERT INTO %v(", tableName)
	var itemValuePlaceholder = " VALUES("
	fieldsLength := len(tableFields)
	for fieldIndex, fieldName := range tableFields {
		itemQuery += fmt.Sprintf(" %v", fieldName)
		itemValuePlaceholder += fmt.Sprintf(" $%v", fieldIndex+1)
		if fieldsLength > 1 && fieldIndex < fieldsLength-1 {
			itemQuery += ", "
			itemValuePlaceholder += ", "
		}
	}
	// close item-script/value-placeholder
	itemQuery += " )"
	itemValuePlaceholder += " )"
	// add/append item-script & value-placeholder to the createScripts
	insertQuery = itemQuery + itemValuePlaceholder

	// compute create values from actionParams
	for recNum, rec := range actionParams {
		// initial item-values-computation variables
		var recFieldValues []interface{}
		for _, fieldName := range tableFields {
			fieldValue := rec[fieldName]
			// check for required field in each record
			if fieldValue == nil {
				return errMessage(fmt.Sprintf("Record #%v [%#v]: required field_name[%v] has field_value of %v ", recNum, rec, fieldName, fieldValue))
			}
			// update recFieldValues by fieldValue-type
			switch fieldValue.(type) {
			case time.Time:
				if fVal, ok := fieldValue.(time.Time); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, "'" + fVal.Format("2006-01-02 15:04:05.000000") + "'")
				}
			case string:
				if fVal, ok := fieldValue.(string); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					if govalidator.IsJSON(fVal) {
						//fmt.Printf("string-toJson-value: %v\n\n", fVal)
						//recFieldValues = append(recFieldValues, "'" + fVal + "'")
						if fValue, err := govalidator.ToJSON(fieldValue); err != nil {
							return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
						} else {
							fmt.Printf("string-toJson-value: %v\n\n", fValue)
							//recFieldValues = append(recFieldValues, "'" + fVal + "'")
							recFieldValues = append(recFieldValues, "")
						}
					} else {
						recFieldValues = append(recFieldValues, "'" + fVal + "'")
					}
				}
			case bool:
				if fVal, ok := fieldValue.(bool); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int8:
				if fVal, ok := fieldValue.(int8); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int16:
				if fVal, ok := fieldValue.(int16); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int32:
				if fVal, ok := fieldValue.(int32); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int64:
				if fVal, ok := fieldValue.(int64); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case int:
				if fVal, ok := fieldValue.(int); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint8:
				if fVal, ok := fieldValue.(uint8); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint16:
				if fVal, ok := fieldValue.(uint16); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint32:
				if fVal, ok := fieldValue.(uint32); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint64:
				if fVal, ok := fieldValue.(uint64); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case uint:
				if fVal, ok := fieldValue.(uint); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case float32:
				if fVal, ok := fieldValue.(float32); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case float64:
				if fVal, ok := fieldValue.(float64); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []string:
				if fVal, ok := fieldValue.([]string); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []int:
				if fVal, ok := fieldValue.([]int); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []float32:
				if fVal, ok := fieldValue.([]float32); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []float64:
				if fVal, ok := fieldValue.([]float64); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			case []struct{}:
				if fVal, ok := fieldValue.([]struct{}); !ok {
					return errMessage(fmt.Sprintf("field_name: %v | field_value: %v error: ", fieldName, fieldValue))
				} else {
					recFieldValues = append(recFieldValues, fVal)
				}
			default:
				// json-stringify fieldValue
				if fVal, err := json.Marshal(fieldValue); err != nil {
					return errMessage(fmt.Sprintf("Unknown or Unsupported field-value type: %v", err.Error()))
				} else {
					//fmt.Printf("***is-json-value***: %v\n\n", govalidator.IsJSON(string(fVal)))
					fmt.Printf("default-value: %v\n\n", string(fVal))
					recFieldValues = append(recFieldValues, "'" + string(fVal) + "'" )
				}
			}
		}
		// update fieldValues
		fValues = append(fValues, recFieldValues)
		// re-initialise recFieldValues, for next update
		recFieldValues = []interface{}{}
	}
	// result
	return mctypes.CreateQueryResponseType{
		CreateQuery: insertQuery,
		FieldNames:  tableFields,
		FieldValues: fValues,
	}, nil
}
