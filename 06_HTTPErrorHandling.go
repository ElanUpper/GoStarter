package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func func_writefile(file_name string) {
	if f, err := os.Create(file_name); err == nil {
		defer f.Close()

		fpwriter := bufio.NewWriter(f)
		defer fpwriter.Flush()

		fmt.Fprintln(fpwriter, "hello elan, this's tmp file")
	}
}

type CusError struct {
	error
	Msg string
}

func (cus CusError) Error() string {
	return cus.Msg
}

const prefix = "/list/"

func handle_file_list(writer http.ResponseWriter, request *http.Request) error {

	// 判断请求必须是/list/开头的 否则返回自定义err
	if ok := strings.Contains(request.URL.Path, prefix); !ok {
		return CusError{Msg: "your input must contain /list/"}
	}

	// get file name
	file_name := request.URL.Path[len(prefix):]
	if f, err := os.Open(file_name); err != nil {
		return err
		// http.Error(writer, err.Error(), 404)
	} else {
		defer f.Close()
		if content, err := ioutil.ReadAll(f); err != nil {
			return err
		} else {
			writer.Write(content)
			return nil
		}
	}
}

type func_type_handle_file_list func(writer http.ResponseWriter, request *http.Request) error

// 中心处理报错信息
func handle_file_wrapper(i_func func_type_handle_file_list) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := i_func(writer, request); err != nil {

			// 保护服务器
			defer func() {
				if CatchErr := recover(); CatchErr != nil {
					http.Error(writer, http.StatusText(http.StatusInternalServerError),
						http.StatusInternalServerError)
					log.Printf("the value %v\n", CatchErr)
				}
			}()

			//writer.Write([]byte(err.Error()))
			code := http.StatusOK

			// 处理客户error
			if CustomErr, ok := err.(CusError); ok {
				http.Error(writer, CustomErr.Msg, http.StatusInternalServerError)
				return
			}

			// 处理具体code对外信息
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusUnauthorized
			default:
				code = http.StatusInternalServerError
			}
			// 打印对外信息
			http.Error(writer, http.StatusText(code), code)
			// 打印对内信息
			log.Print(err.Error())

		}
	}
}

func main() {

	// write file
	file_name := "temp.txt"
	// 不等于空 那么就创建文件
	if fInfo, err := os.Stat(file_name); err != nil {
		func_writefile(file_name)
	} else {
		fmt.Println(fInfo.Name())
	}

	http.HandleFunc("/", handle_file_wrapper(handle_file_list))
	http.ListenAndServe(":10000", nil)
}
