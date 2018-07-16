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
// 实现接口
type ImpInfo struct {
	Content string ;
}

// 名称是什么 没有关系 最核心的是实现接口全部方法 方法的参数一定要match！！！！
func (Info ImpInfo) GetContent(url string) string {
	return Info.Content ;
}

//------------------------------------ complex ------------------------------------
// 实现接口
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

	// 打印值
	fmt.Printf("%T,  %v,  %s\n", Info, Info, []rune(getContent(Info))[:30])

	//
	switch Info.(type) {

	case ImpInfo:
		fmt.Println("ImpInfo")
	case *WebInfo:
		fmt.Println("WebInfo");
	}
}

func main() {
	var Info GetInfo ;
	Info = ImpInfo{"hello world"}
	fmt.Printf("%T,  %v,  %s\n", Info, Info, getContent(Info))


	var Web WebInfo ;
	Web = WebInfo{"hello world", "Chrome", time.Minute }
	Inspect_type(&web);



}

