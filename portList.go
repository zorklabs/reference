package main

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/beevik/etree"
	"github.com/bitwurx/jrpc2"
)

type SwitchIP struct {
	IP *string `json:"ip"`
}

func (swIp *SwitchIP) FromPositional(params []interface{}) error {
	if len(params) != 1 {
		return errors.New("exactly onet IP string is required")
	}

	x := params[0].(string)
	swIp.IP = &x

	return nil
}

func ReadPortsFromSwitch() map[string]int {
	portList := make(map[string]int)

	doc := etree.NewDocument()
	if err := doc.ReadFromFile("data.xml"); err != nil {
		panic(err)
	}
	root := doc.FindElement("/rpc-reply/data/top/Ifmgr/Interfaces")

	for _, swIf := range root.SelectElements("Interface") {
		idx, _ := strconv.Atoi(swIf.SelectElement("IfIndex").Text())
		portList[swIf.SelectElement("AbbreviatedName").Text()] = idx
	}

	return portList
}

func GetPortList(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(SwitchIP)

	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}

	return ReadPortsFromSwitch(), nil
}
