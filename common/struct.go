package common

import "encoding/xml"

// GameXML 游戏XML 数据
type GameXML struct {
	XMLName xml.Name `xml:"configure"`
	Groups ActionGroups `xml:"groups"`
	InitAPP InitApplication `xml:"initApplication"`
	ReplaceFile ReplaceApplication `xml:"replaceApplication"`
	CopyServer CopyApplication `xml:"copyApplication"`
	ReloadFile ReloadApplication `xml:"reloadApplication"`
	RsyncServer RsyncApplication `xml:"rsyncApplication"`
	CleanServer DeleteApplication `xml:"deleteApplication"`
}
// ActionGroup 游戏xml中groups 数据
type ActionGroups struct {
	Group struct{
		Name string `xml:"name,attr"`
		Apps string `xml:"apps,attr"`
	}`xml:"group"`
}
//InitApplication 初始化操作的XML信息
type InitApplication struct {
	AppName string `xml:"name,attr"`
	InitCommand InitSystemrun `xml:"initSystemrun"`
}
// InitSystemrun 初始化操作中的操作名称和命令
type InitSystemrun struct {
	//XMLName xml.Name `xml:"initSystemrun"`
	Command []string `xml:"command,attr"`
}
// ReplaceApplication 游戏XML配置文件替换信息
type ReplaceApplication struct {
	Name string `xml:"name,attr"`
	ReplaceInfo []struct{
		File string `xml:"file,attr"`
		SRC string `xml:"src,attr"`
		Dest string `xml:"dest,attr"`
	}`xml:"replace"`
}
// CopyApplication 游戏XML中拷贝服务器端 信息
type CopyApplication struct {
	Name string `xml:"name,attr"`
	CopyCommand []struct{
		Command string `xml:"command,attr"`
	}`xml:"copySystemrun"`
}
//ReloadApplication 游戏XML中 修改完配置文件后的reload操作信息
type ReloadApplication struct {
	Name string `xml:"name,attr"`
	ReloadCommand []struct{
		Command string `xml:"command,attr"`
	}`xml:"reloadSystemrun"`
}
//RsyncApplication 游戏XML中同步服务器端的操作信息
type RsyncApplication struct {
	Name string `xml:"name,attr"`
	RsyncCommand []struct{
		Command string `xml:"command,attr"`
		Host string `xml:"host,attr"`
	}`xml:"rsyncSystemrun"`
}

//DeleteApplication 游戏XML中的收尾清理操作信息
type DeleteApplication struct {
	Name string `xml:"name,attr"`
	DeleteCommand []struct{
		Command string `xml:"command,attr"`
		Host string `xml:"host,attr"`
	}`xml:"deleteSystemrun"`
}