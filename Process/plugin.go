/*
配置文件里的工作流
*/
package Process

import (
	"fmt"
	"strings"
	"strconv"
)

func Plugin(list map[string]interface{},operator string,imports_list map[string]string) (result string,ok bool,imports map[string]string) {

	ok = false
	if list["func"] == nil || list["key"] == nil {
		fmt.Println(`list["func"] == nil || list["key"] == nil`)
		fmt.Println(`list:`,list,operator)
		return
	}

	imports = make(map[string]string)

	key := list["key"].(string)

	key_s := strings.Replace(key,"_","/",-1)

	list1 := strings.Split(key_s,"/")
	import_s := list1[len(list1) - 1]

	var count int64 = 0
	var had bool = false
	for key1,v := range imports_list {
		l := strings.Split(key1,"/")
		if strings.EqualFold(l[len(l) - 1],import_s) {
			count++
		}
		if strings.EqualFold(key_s,key1) {
			import_s = v
			had = true
		}
	}

	if !had {
		if count > 0 {
			import_s += "_" + strconv.FormatInt(count,10)
		}
		imports[key_s] = import_s
		imports_list[key_s] = import_s
	}

	if list["class"] != nil {
		result += list["class"].(string) + "."
	} else {
		result += import_s + "."
	}

	result += list["func"].(string) + "("

	if list["params"] != nil {

		for index,value := range list["params"].([]interface{}) {
			r ,ok1,imports_l := Process(value.(map[string]interface{}),imports_list)
			if !ok1 {
				fmt.Println(`mapValue(list map[string]interface{},operator string) Process(list["left"].(map[string]interface{}))`)
				fmt.Println(`list:`,list,operator)
				return
			}

			for k,v := range imports_l {
				imports[k] = v
				imports_list[k] = v
			}

			result += r

			if index != len(list["params"].([]interface{})) - 1 {
				result += " , "
			}
		}

	}

	result += ")"
	ok = true
	return
}
