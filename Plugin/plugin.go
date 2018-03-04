package Plugin

import (
	"strings"
	"go/format"
	"github.com/NiuStar/utils"
	"time"
)


func UpdatePlugin(data string) *Plugin {


	body,err := format.Source([]byte(data))
	if err != nil {
		panic(err)
	}

	data = formatCodeQuote(string(body))

	//fmt.Println("data",data)
	d := formatCodeBrackets(data)

	var classes map[string]interface{} = make(map[string]interface{})
	list := getFuncs(d,classes)

	var clist []string

	for key,_ := range classes {
		clist = append(clist,key)
	}

	return &Plugin{Classes:clist,Method:list,Time:utils.FormatTimeAll(time.Now().Unix()),Price:0}
}

func GetPackageName(data string) string {

	body,err := format.Source([]byte(data))
	if err != nil {
		panic(err)
	}

	data = formatCodeQuote(string(body))

	data = strings.Split(data,"package ")[1]
	return strings.Split(data,"\n")[0]
}

type Plugin struct {
	Classes []string
	Method []Func
	Descirble string
	Time string
	Name string
	Price float64
	Had bool
}

type Params struct {
	Name string //参数名称
	Type string //参数属性
	Key string //参数在工具里显示的名称
}

type Func struct {
	Class string
	Name string
	Param []Params
	Results []Params
	Key string //方法在工具里显示的名称
	Static bool//是否为静态方法
	Descirble string//方法介绍
}

func getFuncs(data string,classes map[string]interface{}) (list []Func) {

	list1 := strings.Split(data,"func ")

	for i := 1; i < len(list1) ; i++ {

		var text string = list1[i]
		var static bool = true
		var class string = ""
		if list1[i][0] == '(' {

			list2 := strings.SplitN(list1[i],"(",2)

			list2 = strings.SplitN(list2[1],")",2)

			class = strings.Split(list2[0],"*")[1]
			//class = list2[0]


			classes[class] = ""
			text = strings.TrimLeft(list2[1]," ")


			static = false
		}

		list2 := strings.SplitN(text,"(",2)
		name := list2[0]


		if len(name) > 0 && utils.IsCapital(name[0]) {

			list2 = strings.SplitN(list2[1],")",2)

			p := getParams(list2[0])

			//fmt.Println("params:",list2[0])
			//fmt.Println("result:",list2[1])

			r := getResults(strings.Split(list2[1],"{")[0])
			f := Func{Class:class,Name:name,Param:p,Results:r,Static:static}

			list = append(list,f)
		}
	}

	return
}

func getParams(data string) (list []Params) {


	list1 := strings.Split(data,", ")

	for i := 0;i < len(list1) ; i ++ {
		list2 := strings.Split(list1[i]," ")
		var type_ string
		name := list2[0]
		if len(list2) > 1 {
			type_ = list2[1]
			for j := len(list) - 1;j >= 0 ; j-- {

				if len(list[j].Type) > 0 {
					break
				} else {

					list[j].Type = type_
				}
			}
		}
		p := Params{Name:name,Type:type_}
		list = append(list,p)
	}

	return
}

func getResults(data string) (list []Params) {

	list1 := strings.Split(data,"(")
	if len(list1) == 2 {

		l := strings.Split(list1[1],")")
		list = getParams(l[0])

		for i := 0 ; i < len(list) ; i++ {
			if list[i].Type == "" {
				list[i].Type = list[i].Name
			}
		}
	} else {
		if len(data) == 0 {
			return
		}

		p := Params{Name:data,Type:data}
		list = append(list,p)

	}

	return
}

func formatFile(name string) string {
	var data string

	data = formatCode(utils.ReadFileFullPath(name))

	body,err := format.Source([]byte(data))
	if err != nil {
		panic(err)
	}
	return string(body)
	//fmt.Println("data:",string(body))
}


func formatCodeBrackets(s string) (result string) {//去除括号内的内容


	var left uint8 = '{'
	var right uint8 = '}'

	var char uint8 = 0      //当前字符, 0 为没有任何值  1为`  2为" 3为{

	var count int = 0

	for i := 0;i < len(s) ; i++ {

		//

		switch char {
		case 0:{

			if s[i] == left {
				count++
				char = 3
				result += string(s[i])
			} else {

				result += string(s[i])
			}
		}
			break

		case 3:{

			if s[i] == right {

				count--
				if count == 0 {
					char = 0
					result += string(s[i])
				} else if count < 0 {
					count = 0
				}
			} else if s[i] == left {
				count++
			}

		}
			break

		}

	}

	return
}

func formatCodeQuote(s string) (result string) {//去引号内的内容

	var single uint8 = '`'
	var double uint8 = '"'

	var ln uint8 = '"'

	var single_ uint8 = '\''
	var convert_ uint8 = '\\' //转移符
	var char uint8 = 0      //当前字符, 0 为没有任何值  1为`  2为"

	for i := 0;i < len(s) ; i++ {

		//

		switch char {
		case 0:{

			if s[i] == single {
				char = 1
			} else if s[i] == double {
				char = 2
			} else if s[i] == single_ {

				var convert bool = false

				for j := i + 1;j < len(s) ; j++ {
					i++

					if s[j] == single_ && !convert {

						break
					} else if s[j] == convert_ {
						convert = !convert
					} else {
						convert = false
					}

				}

			} else {

				result += string(s[i])
			}
		}
			break
		case 1:{

			if s[i] == single {

				var convert bool = false
				for j := i - 1;j > 0 ; j-- {
					if s[j] == convert_ {
						convert = !convert
					} else {
						convert = false
						break
					}
				}

				if !convert {
					char = 0
				}
			}
		}
			break
		case 2:{

			if s[i] == double {
				var convert bool = false
				for j := i - 1;j > 0 ; j-- {
					if s[j] == convert_ {
						convert = !convert
					} else {
						convert = false
						break
					}
				}

				if !convert {
					char = 0
				}

			} else if s[i] == ln {
				char = 0

			}
		}
			break

		}

	}

	return
}

func formatCode(s string) (result string) {

	var double_ uint8 = '/'
	var single uint8 = '`'
	var double uint8 = '"'

	var star uint8 = '*' //123456
	var ln uint8 = '\n' //换行符
	var single_ uint8 = '\''
	var convert_ uint8 = '\\' //转移符
	var char uint8 = 0      //当前字符, 0 为没有任何值  1为`  2为" 3为// 4为/*

	for i := 0;i < len(s) ; i++ {
		//fmt.Println(s[i])

		switch char {
		case 0:{
			if s[i] == single {
				char = 1
				result += string(s[i])
			} else if s[i] == double {
				char = 2
				result += string(s[i])
			} else if s[i] == double_ {
				if i < len(s) - 1 && s[i+1] == double_ {
					char = 3
					i++
				} else if i < len(s) - 1 && s[i+1] == star {
					char = 4
					i++
				} else {
					result += string(s[i])
				}
			} else if s[i] == single_ {
				result += string(s[i])

				var convert bool = false

				for j := i + 1;j < len(s) ; j++ {
					i++
					result += string(s[j])

					if s[j] == single_ && !convert {

						break
					} else if s[j] == convert_ {
						convert = !convert
					} else {
						convert = false
					}

				}

			} else {

				result += string(s[i])
			}
		}
			break
		case 1:{
			if s[i] == single {

				var convert bool = false
				for j := i - 1;j > 0 ; j-- {
					if s[j] == convert_ {
						convert = !convert
					} else {
						convert = false
						break
					}
				}

				if !convert {
					char = 0
				}
				result += string(s[i])
			} else {
				result += string(s[i])
			}
		}
			break
		case 2:{
			if s[i] == double {
				var convert bool = false
				for j := i - 1;j > 0 ; j-- {
					if s[j] == convert_ {
						convert = !convert
					} else {
						convert = false
						break
					}
				}

				if !convert {
					char = 0
				}
				result += string(s[i])
			} else if s[i] == ln {
				char = 0
				result += string(s[i])
			} else {
				result += string(s[i])
			}
		}
			break
		case 3:{
			if s[i] == ln {
				char = 0
				result += string(s[i])
			}

		}
			break
		case 4:{

			if s[i] == star {
				if i < len(s) - 1 && s[i+1] == double_ {
					char = 0
					i++
				}
			}
		}
			break

		}

	}

	return
}