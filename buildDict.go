package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var Profile = flag.String("c", "conf.ini", "build password detection Profile.")
var max  = flag.Int("f", 3, "The maximum number of stitches.")

const Dictfile = "pass.txt"
var selfwriter *bufio.Writer
func main(){
	flag.Parse()
	Profileinfo := loadProfile(*Profile)

	file, err := os.OpenFile(Dictfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	selfwriter = bufio.NewWriter(file)

	permutation(Profileinfo,"",min(*max,len(Profileinfo)))
	//Flush将缓存的文件真正写入到文件中
	selfwriter.Flush()
}


func permutation(S []string,tmp string,flag int){
	if flag == 0 {
		return
	}
	for _,v := range S {
		fmt.Println(v+tmp)
		selfwriter.WriteString(v+tmp+"\n")
		permutation(S,v+tmp,flag-1)
	}
}

func loadProfile(profile string) []string {
	file, err := os.OpenFile(profile, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open ",profile," file error!", err)
		return nil
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("Profile size=", size)

	var linelist []string
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		linelist = append(linelist, line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Profile read ok!")
				break
			} else {
				fmt.Println("Read Profile error!", err)
				return nil
			}
		}
	}
	return linelist
}

func min(a,b int)int{
	if a < b{
		return a
	}
	return b
}