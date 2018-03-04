/*
JSON字符串中的map与list转换为代码生成
*/
package InterfaceFormat

import "strconv"

func MapInterface2String(list,name string,value interface{},index int64) (result string) {//map[string]interface 互相嵌套

	tempName := "al" + strconv.FormatInt(index,10)
	switch value.(type) {

	case map[string]interface{}:{


		result = " {\n		var " + tempName + " map[string]interface{} = make(map[string]interface{})\n"

		for key,v := range value.(map[string]interface{}) {
			result += MapInterface2String(tempName,key,v,index + 1) + ")"
		}
		result += "		" + list + "[\"" + name + "\"]" + " = " + tempName
		result += "\n}\n"
	}
		break
	case []interface{}:{

		result = " {\n		var " + tempName + " []interface{}\n"

		for _,v := range value.([]interface{}) {
			result +=  MultipleInterface2String(tempName,v,index + 1)
		}
		result += "		" + list + "[\"" + name + "\"]" + " = " + tempName
		result += "\n}\n"
	}
		break
	case string:{
		result += "		" + list + "[\"" + name + "\"] = " + "\"" + value.(string) + "\""
		//method_value += "	var " + name + " string" + "\n"
	}
		break
	case float64:{
		result += "		" + list + "[\"" + name + "\"] = " + strconv.FormatFloat(value.(float64),'f',6,64)
		//result =
	}
		break
	case float32:{
		result += "		" + list + "[\"" + name + "\"] = " + strconv.FormatFloat(float64(value.(float32)),'f',6,32)
	}
		break
	case int64:{
		result += "		" + list + "[\"" + name + "\"] = " + strconv.FormatInt(value.(int64),10)
		//result =
	}

		break
	case bool:{
		result += "		" + list + "[\"" + name + "\"] = " +  strconv.FormatBool(value.(bool))

	}
		break
	}

	result += "\n"
	return
}

func MultipleInterface2String(list string,value interface{},index int64) (result string) {//[]interface 互相嵌套

	tempName := "al" + strconv.FormatInt(index,10)
	switch value.(type) {

	case map[string]interface{}:{

		result = " {\n		var " + tempName + " map[string]interface{} = make(map[string]interface{})\n"

		for key,v := range value.(map[string]interface{}) {
			result += "		"   + MapInterface2String(tempName,key,v,index + 1) + ")"
		}
		result += "		" + list + " = append(" + list + " , " + tempName + ")"
		result += "\n}\n"
	}
		break
	case []interface{}:{

		result = " {\n		var " + tempName + " []interface{}\n"

		for _,v := range value.([]interface{}) {
			result +=  MultipleInterface2String(tempName,v,index + 1)
		}
		result += "		" + list + " = append(" + list + " , " + tempName + ")"
		result += "\n}\n"
	}
		break
	case string:{
		result += "		" + list + " = append(" + list + " , " + "\"" + value.(string) + "\")"
		//method_value += "	var " + name + " string" + "\n"
	}
		break
	case float64:{
		result += "		" + list + " = append(" + list + " , " + strconv.FormatFloat(value.(float64),'f',6,64) + ")"
		//result =
	}
		break
	case float32:{
		result += "		" + list + " = append(" + list + " , " + strconv.FormatFloat(float64(value.(float32)),'f',6,32) + ")"
	}
		break
	case int64:{
		result += "		" + list + " = append(" + list + " , " + strconv.FormatInt(value.(int64),10)+ ")"
		//result =
	}
		break
	case bool:{
		result += "		" + list + " = append(" + list + " , " + strconv.FormatBool(value.(bool)) + ")"

	}
		break
	}

	result += "\n"
	return
}

func Interface2String(name string,value interface{}) (result string) {



	switch value.(type) {

	case map[string]interface{}:{

		result = " {\n		var al map[string]interface{} = make(map[string]interface{})\n"

		for key,v := range value.(map[string]interface{}) {
			result += MapInterface2String("al",key,v,0)
		}
		result += "		" + name + " = al"
		result += "\n}\n"
	}
		break
	case []interface{}:{

		result = " {\n		var al []interface{}\n"

		for _,v := range value.([]interface{}) {
			result +=  MultipleInterface2String("al",v,0)
		}
		result += "		" + name + " = al"
		result += "\n}\n"
	}
		break
	case string:{
		result += "		" + name + " =  \"" + value.(string) + "\""
		//method_value += "	var " + name + " string" + "\n"
	}
		break
	case int64:{
		result += "		" + name + " = " + strconv.FormatInt(value.(int64),10)
		//result =
	}
	case float64:{
		result += "		" + name + " = " + strconv.FormatFloat(value.(float64),'f',6,64)
		//result =
	}
		break
	case float32:{
		result += "		" + name + " = " + strconv.FormatFloat(float64(value.(float32)),'f',6,32)
	}
		break
	case bool:{
		result += "		" + name + " = " +  strconv.FormatBool(value.(bool))

	}
		break
	}
	result += "\n"
	return
}

/*
仅把value部分转换为string
*/
func Interface2StringValue(value interface{}) (result string) {



	switch value.(type) {

	case string:{
		result = value.(string)
		//method_value += "	var " + name + " string" + "\n"
	}
		break
	case int64:{
		result = strconv.FormatInt(value.(int64),10)
		//result =
	}
	case float64:{
		result = strconv.FormatFloat(value.(float64),'f',6,64)
		//result =
	}
		break
	case float32:{
		result = strconv.FormatFloat(float64(value.(float32)),'f',6,32)
	}
		break
	case bool:{
		result = strconv.FormatBool(value.(bool))

	}
		break
	}
	return
}