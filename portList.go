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
		return errors.New("exactly one IP string is required")
	}

	x := params[0].(string)
	swIp.IP = &x

	return nil
}

type PortList map[string]int

func ReadPortsFromSwitch() PortList {
	portsIdxList := make(PortList)

	doc := etree.NewDocument()
	if err := doc.ReadFromFile("data.xml"); err != nil {
		panic(err)
	}
	interfaceData := doc.FindElement("/rpc-reply/data/top/Ifmgr/Interfaces")

	for _, swIface := range interfaceData.SelectElements("Interface") {
		idxNum, _ := strconv.Atoi(swIface.SelectElement("IfIndex").Text())
		portsIdxList[swIface.SelectElement("AbbreviatedName").Text()] = idxNum
	}

	return portsIdxList
}

func GetPortList(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(SwitchIP)

	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}

	return ReadPortsFromSwitch(), nil
}
