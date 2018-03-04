/*
函数内的成员变量解析
*/
package Variable

import (
	"CreateProject/InterfaceFormat"
	"strings"
)

func GetVariable(value map[string]interface{},name string) (result string) {
	switch value["value"].(type) {

		case map[string]interface{}:{
			result += "	var " + name + " map[string]interface{}" + "\n"
			result +=  InterfaceFormat.Interface2String(name,value["value"])
		}

		case []interface{}:{
			result += "	var " + name + " []interface{}" + "\n"
			result += InterfaceFormat.Interface2String(name,value["value"])
		}

		case string:{
			result += "	var " + name + " string" + "\n"
			result += InterfaceFormat.Interface2String(name,value["value"])
		}

		case float64:{
			if value["type"] != nil && strings.EqualFold(value["type"].(string), "int64") {
				result += "	var " + name + " int64" + "\n"
				result += InterfaceFormat.Interface2String(name,int64(value["value"].(float64)))
			} else {
				result += "	var " + name + " float64" + "\n"
				result += InterfaceFormat.Interface2String(name,value["value"])
			}

		}
		case bool:{
			result += "	var " + name + " bool" + "\n"
			result += InterfaceFormat.Interface2String(name,value["value"])

		}
	}
	return
}
