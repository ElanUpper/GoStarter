package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

// docker typing
// 看着&用起来像就是！！

// 定义接口
type GetInfo interface {
	GetContent (url string) string ;
	// GetCookie  (url string) string ;
}

//------------------------------------ simple ------------------------------------
// 实现接口(复制方式)
type ImpInfo struct {
	Content string ;
}

// 名称是什么 没有关系 最核心的是实现接口全部方法 方法的参数一定要match！！！！
func (Info ImpInfo) GetContent(url string) string {
	return Info.Content ;
}

//------------------------------------ complex ------------------------------------
// 实现接口(引用地址方式)
type WebInfo struct {
	Content   string ;
	UserAgent string ;
	timezone  time.Duration ;
}

func (web *WebInfo) GetContent(url string) string {
	if res, err := http.Get(url); err == nil {
		if result, err := httputil.DumpResponse(res, true); err == nil {
			// 关闭http
			res.Body.Close();
			return string(result) ;
		} else {
			panic(err)
		}
	} else{
		panic(err) ;
	}
}

// 调用接口
func getContent(info GetInfo) string {
	return info.GetContent("http://www.baidu.com") ;
}

// 检查类型
func Inspect_type(Info GetInfo){

	// 打印类型， 值
	fmt.Printf("%T,  %v\n", Info, Info) // , []rune(getContent(Info))[:30])

	// type分类
	switch v := Info.(type) {
	case ImpInfo: //
		fmt.Println(Info, "implemented by ImpInfo", v.Content )
	case *WebInfo: // 指针方式
		fmt.Println(Info, "implemented by WebInfo", v.UserAgent, v.timezone);
	}

}

// interface as parameters
func getCalc(inter1, inter2 interface{})  {
	switch type1 := inter1.(type) {
	case int:
		fmt.Printf("type [%T] value [%d]\n", type1, inter1.(int) + inter2.(int) )
	case float32:
		fmt.Printf("type [%T] value [%f]\n", type1, inter1.(float32) + inter2.(float32) )
	default:
		fmt.Printf("type [%T], not support!\n", type1 )
	}
}


func main() {
	// 接口
	var Info GetInfo ;


	var a interface{}
	v1, ok1 := a.(string)
	fmt.Println(v1, ok1)

	// 查看
	fmt.Println();
	Info = ImpInfo{"type:value "}
	Inspect_type(Info) ; // switch (type)

	fmt.Println();
	Info = &WebInfo{"type:pointer ", "Chrome", time.Minute }
	Inspect_type(Info) ;

	// type assertion
	fmt.Printf("\n------------------------------- type assertion -------------------------------\n");
	webTypeAssert := Info.(*WebInfo)
	fmt.Println(webTypeAssert.UserAgent, webTypeAssert.timezone)

	// interface as input parameter
	fmt.Println();
	getCalc(10, 20)
	getCalc(float32(10.11), float32(20.22))
	getCalc(true, false)


}

