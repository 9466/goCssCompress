package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	sf := "test.css"
	df := "output.css"
	var in, out []byte
	var err error

	in, err = ioutil.ReadFile(sf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	out, err = compress(in)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = ioutil.WriteFile(df, out, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("ok")
}

func compress(in []byte) ([]byte, error) {
	var c byte
	out := make([]byte, len(in))
	j := 0
	n := len(in)
	// 在结尾插入一个空格，防止后面+1判断时溢出
	in = append(in, ' ')
	for i := 0; i < n; i++ {
		c = in[i]
		// 注释处理
		if c == '/' {
			if in[i+1] == '*' {
				// 这里是注释
				// 开始寻找注释的结尾
				i++
				for {
					i++
					if i >= n || (in[i] == '*' && in[i+1] == '/') {
						i++
						break
					}
				}
				continue
			}
		}
		// 换行处理
		if c == '\n' {
			continue
		}
		// 处理tab
		if c == '\t' {
			c = ' '
		}
		// 干掉第一个空格
		if j == 0 && c == ' ' {
			continue
		}
		// 处理,:;后面的空格，同时处理连续空格问题
		if c == ' ' || c == ',' || c == ':' || c == ';' || c == '{' || c == '}' {
			// 后面所有的空格都不要啦
			for {
				i++
				if i >= n || (in[i] != ' ' && in[i] != '\n' && in[i] != '\t') {
					i--
					break
				}
			}
		}
		// 处理{前面的空格
		if c == '{' && out[j-1] == ' ' {
			j--
		}
		// 处理}前面的空格和分号
		if c == '}' && (out[j-1] == ' ' || out[j-1] == ';') {
			j--
		}
		out[j] = c
		j++
	}
	return out[0:j], nil
}
