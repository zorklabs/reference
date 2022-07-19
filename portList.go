package main

import (
	"encoding/json"
	"errors"
	"fmt"

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

func GetPortList(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	p := new(SwitchIP)

	if err := jrpc2.ParseParams(params, p); err != nil {
		return nil, err
	}

	fmt.Println(*p.IP)

	portList := make(map[int]string)
	portList[1] = "GE1/0/1"
	portList[2] = "GE1/0/2"
	portList[3] = "GE1/0/3"
	portList[4] = "GE1/0/4"
	portList[5] = "GE1/0/5"

	// portList := make(map[string]int)
	// portList["GE1/0/1"] = 1
	// portList["GE1/0/2"] = 2
	// portList["GE1/0/3"] = 3
	// portList["GE1/0/4"] = 4
	// portList["GE1/0/5"] = 5

	return portList, nil
}
