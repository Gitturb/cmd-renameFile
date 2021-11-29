package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

var pathSeparator = string(os.PathSeparator)

var wg = sync.WaitGroup{}

func main() {
	func2()
}

func func2() {
	wg.Add(1)
	// 1.获取要被重命名目录(文件)的绝对目录
	fmt.Print("请输入文件的绝对路径：")
	reader := bufio.NewReader(os.Stdin)
	filePath, _ := reader.ReadString('\n')
	filePath = strings.Replace(filePath, "\n", "", -1)

	// 获取新的文件名称
	fmt.Print("请输入文件的新名称")
	nameReader := bufio.NewReader(os.Stdin)
	newName, _ := nameReader.ReadString('\n')
	newName = strings.Replace(newName, "\n", "", -1)

	if newName == "" {
		log.Fatalln("文件新名称必须！！")
	}
	go rename2(filePath, newName)
	wg.Wait()
}

func rename2(path2 string, new string) {
	fileInfos, err := ioutil.ReadDir(path2)
	if err != nil {
		return
	}
	for i, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			wg.Add(1)
			go rename2(path.Join(path2, fileInfo.Name()), new)
		} else {
			// 获取新的文件名称
			var frame = replace(fileInfo.Name(), new, i)
			newFileName := path.Join(path2, frame)
			err := os.Rename(path.Join(path2, fileInfo.Name()), newFileName)
			if err != nil {
				fmt.Println("重命名失败", err)
			}
		}
		fmt.Println(path.Join(path2, fileInfo.Name()))
	}
	wg.Done()
}

func replace(old string, new string, i int) string {
	// 先获取后缀
	ext := path.Ext(old)
	var pre string
	if i < 10 {
		pre = "000"
	} else if i < 100 {
		pre = "00"
	} else if i < 1000 {
		pre = "0"
	} else if i < 10000 {
		pre = ""
	}
	return new + pre + strconv.Itoa(i) + ext
}

func func1() {
	//1.获取要被重命名目录(文件)的绝对路径
	fmt.Print("请输入文件的绝对路径：")
	reader := bufio.NewReader(os.Stdin)
	filePath, _ := reader.ReadString('\n')
	filePath = strings.Replace(filePath, "\n", "", -1)
	if filePath == "" {
		log.Fatalln("目标文件绝对路径必须填写！！")
	}

	//2.获取要被替换掉的名称
	fmt.Print("请输入你想修改的文件名前缀：")
	reader = bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	if name == "" {
		fmt.Print("请输入你想修改的文件名前缀：")
		reader = bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		name = strings.Replace(name, "\n", "", -1)
	}

	if name == "" {
		log.Fatalf("文件名前缀必须填写！！")
	}

	//3.递归调用重命名
	err := rename(filePath, name)
	if err != nil {
		log.Fatalf("发生错误，错误为：%v\n", err)
	}
	fmt.Println("success")
}

func rename(path1 string, new string) (err error) {
	fmt.Println(path1)
	files, err := ioutil.ReadDir(path1)
	if err != nil {
		return err
	}
	for i, fileInfo := range files {
		if fileInfo.IsDir() {
			err = rename(path1+pathSeparator+fileInfo.Name(), new)
			if err != nil {
				return err
			}
		} else {
			ext := path.Ext(fileInfo.Name())
			if i < 10 {
				err = os.Rename(path1+pathSeparator+fileInfo.Name(), path1+pathSeparator+strings.Replace(fileInfo.Name(), fileInfo.Name(), new+"000"+strconv.Itoa(i+1)+ext, -1))
			} else if i < 100 {
				err = os.Rename(path1+pathSeparator+fileInfo.Name(), path1+pathSeparator+strings.Replace(fileInfo.Name(), fileInfo.Name(), new+"00"+strconv.Itoa(i+1)+ext, -1))
			} else if i < 1000 {
				err = os.Rename(path1+pathSeparator+fileInfo.Name(), path1+pathSeparator+strings.Replace(fileInfo.Name(), fileInfo.Name(), new+"0"+strconv.Itoa(i+1)+ext, -1))
			} else if i < 10000 {
				err = os.Rename(path1+pathSeparator+fileInfo.Name(), path1+pathSeparator+strings.Replace(fileInfo.Name(), fileInfo.Name(), strconv.Itoa(i+1)+new+ext, -1))
			}
		}
	}
	return err
}
