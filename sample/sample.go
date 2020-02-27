package main

import (
	"fmt"
	"gitee.com/akimimi/getuigo"
	"github.com/go-yaml/yaml"
	"github.com/gogap/logs"
	"io/ioutil"
	"time"
)

type SampleGetuiConfig struct {
	Getui getuigo.GetuiConfig `yaml:"getui"`
	Cid   string              `yaml:"cid"`
}

type SamplePayload struct {
	getuigo.BasicPayload
}

func main() {
	config := SampleGetuiConfig{}
	if file, e := ioutil.ReadFile("sample/conf.yaml"); e != nil {
		panic(e)
	} else {
		if e := yaml.Unmarshal(file, &config); e != nil {
			panic(e)
		}
	}
	logs.Info(config)

	payload := SamplePayload{
		BasicPayload: getuigo.BasicPayload{
			PushTitle:    "title",
			PushBody:     "body",
			IsShowNotify: 1,
			Ext:          "",
		},
	}
	if igetui, e := getuigo.NewGeTui(&config.Getui); e != nil {
		panic(e)
	} else {
		if e := igetui.SendTransmissionByCid(config.Cid, &payload); e != nil {
			logs.Error("Send to cid failed! cid=", config.Cid)
		} else {
			fmt.Println("Successfully send to cid ", config.Cid)
			logs.Info("Successfully send to cid ", config.Cid)
			logs.Info("info")
		}
	}
	time.Sleep(50)
}
