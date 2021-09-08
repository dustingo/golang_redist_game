package main

import (
	"fmt"
	"gameredist/common"
	"gameredist/util"
	"io/ioutil"
)

func main(){
	data := common.GameXML{}
	byteData, _ := ioutil.ReadFile("server.xml")
	err := util.ParseXML(byteData,&data)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(data.Groups)
	fmt.Println(data.InitAPP)
	fmt.Println(data.ReplaceFile)
	fmt.Println(data.CopyServer)
	fmt.Println(data.ReloadFile)
	fmt.Println(data.RsyncServer.RsyncCommand)
	fmt.Println(data.CleanServer)
}
