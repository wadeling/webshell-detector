package main

import (
	"fmt"
	"io/ioutil"
	"github.com/chaitin/cloudwalker/tool/webshell-detector/src"
	"github.com/wadeling/webshell-test/php"
)

func main() {
	fmt.Println("hello")

	php.Start()
	detector,err := WebshellDetector.NewDefaultDetector(php.Stdin,php.Stdout)
	if err != nil {
		fmt.Println("new detector err",err)
		return
	}

	path:= "xxx.txt"
	src, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("read file err",err)
		return
	}
	score, err := detector.Predict(src)
	if err != nil {
		fmt.Println("predict err",err)
		return
	}
	fmt.Println("score ",score)

	fmt.Println("end")
}
