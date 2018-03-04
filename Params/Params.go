/*
获取配置文件里的从客户端到服务器的传递参数
*/

package Params

import (
	"strings"
	"fmt"
)

func GetParams(value map[string]interface{}) (get_params string,imports map[string]string) {

	imports = make(map[string]string)
	for key,v_v := range value {

		if key != "type" {

			get_params += "	var " + key + " " + v_v.(string) + "\n"

			get_params_str ,imports_l := getParams(key,v_v.(string),value["type"].(string))

			get_params += get_params_str

			for k,v := range imports_l {
				imports[k] = v
			}
		}
	}
	fmt.Println("imports:",imports)
	return
}

func getParams(key,value ,getType string) (get_params string,imports map[string]string) {

	imports = make(map[string]string)
	if strings.EqualFold(getType,"POST") {
		get_params = `	{
		temp,ok := c.GetPostForm("` + key + `")
		if !ok {
			result["status"] = false
			result["message"] = "` + key + `参数错误"
			return
		}
`
	} else {
		get_params = `	{
		temp,ok := c.GetQuery("` + key + `")
		if !ok {
			result["status"] = false
			result["message"] = "` + key + `参数错误"
			return
		}
`
	}

	get_params_str ,imports_l := formatParams(key,value)

	imports = imports_l
	get_params += get_params_str
	get_params += `	}
`
	return
}

func formatParams(key,value string) (get_params string,imports map[string]string) {


	imports = make(map[string]string)
	if strings.EqualFold(value,"string") {
		get_params += `		` + key + " = temp\n"
	} else if strings.EqualFold(value,"int64") {
		imports["strconv"] = ""
		get_params += `			
		i ,err := strconv.ParseInt(temp,10,64)
		if err != nil {
			result["status"] = false
			result["message"] = "` + key + `参数错误"
			return
		}
		`		+ key + ` = i
`		} else if strings.EqualFold(value,"float64") {
		imports["strconv"] = ""
		get_params += `		i ,err := strconv.ParseFloat(temp,10)
		if err != nil {
			result["status"] = false
			result["message"] = "` + key + `参数错误"
			return
		}
		`						+ key + ` = i
`			} else if strings.EqualFold(value,"bool") {
		imports["strings"] = ""
		get_params += `		` + key + ` = strings.EqualFold(temp,"true")
`
	}
	return
}