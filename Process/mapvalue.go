/*
配置文件里的工作流
*/
package Process

import (
	"fmt"
)

func mapValue(list map[string]interface{},operator string,imports_list map[string]string) (result string,ok bool,imports map[string]string) {

	ok = false
	if list["left"] == nil || list["key"] == nil {
		fmt.Println(`list["left"] == nil || list["right"] == nil`)
		fmt.Println(`list:`,list,operator)
		return
	}

	imports = make(map[string]string)
	switch list["left"].(type) {
	case string:{
		var variable bool = true
		if list["variable"] != nil {
			variable = list["variable"].(bool)
		}

		if variable {
			result = list["left"].(string)
		} else {
			return
		}
	}
	case map[string]interface{}:{
		r ,ok1,imports_l := Process(list["left"].(map[string]interface{}),imports_list)
		if !ok1 {
			fmt.Println(`mapValue(list map[string]interface{},operator string) Process(list["left"].(map[string]interface{}))`)
			fmt.Println(`list:`,list,operator)
			return
		}
		for k,v := range imports_l {
			imports[k] = v
			imports_list[k] = v
		}
		result = r
	}
	default:
		return
	}

	result += "["

	r , ok ,imports_l:= Process(list["key"].(map[string]interface{}),imports_list)

	if !ok {
		fmt.Println(`mapValue(list map[string]interface{},operator string) (result string,ok bool)`)
		fmt.Println(list["key"])
		return
	}
	for k,v := range imports_l {
		imports[k] = v
		imports_list[k] = v
	}
	result += r
	result += "]"

	if list["convert"] != nil {
		result += ".(" + list["convert"].(string) + ")"
	}

	ok = true
	return
}
