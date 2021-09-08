package util

import (
	"encoding/xml"
	"gameredist/common"
)
//ParseXML 解析游戏XNL
func ParseXML(b []byte,data *common.GameXML)error{
	err := xml.Unmarshal(b,data)
	if err != nil{
		return err
	}
	return nil
}
