/*
Package util 主要时用于执行XML中的ShellCommand
 */
package util

import (
	"log"
	"os/exec"
	"strings"
	"sync"
)

var lock sync.Mutex
// Work 限制goroutine
type Work struct {
	Pool int //限制并发数量
	Limit chan struct{} // 信号
	Type string // 执行shell时的类型（replace,rsync...）
}

// NewWork 初始化Work
func NewWork(p int,t string)*Work {
	return &Work{
		Pool:  p,
		Limit: make(chan struct{},p),
		Type:  t,
	}
}

// Run 正式执行shellCommand
func (w *Work) Run(f func()){
	w.Limit <- struct{}{}
	go func(){
		f()
		<- w.Limit
	}()
}

// NewFunc 根据XML中的systemRun，构造每条命令的函数
func NewFunc(s string,w *Work) func(){
	lockShell := func(){
		lock.Lock()
		cmd := exec.Command("sh","-c",s)
		_, err := cmd.CombinedOutput()
		if err != nil{
			log.Println(err)
			return
		}
		lock.Unlock()
	}
	unlockShell := func(){
		cmd := exec.Command("sh","-c",s)
		_, err := cmd.CombinedOutput()
		if err != nil{
			log.Println(err)
			return
		}
	}
	if lowerStr := strings.ToLower(w.Type); lowerStr == "replace"{
		return lockShell
	}else{
		return unlockShell
	}
}