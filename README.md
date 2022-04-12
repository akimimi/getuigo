getuigo
==========

[![Go Reference](https://pkg.go.dev/badge/github.com/akimimi/getuigo.svg)](https://pkg.go.dev/github.com/akimimi/getuigo)
[![Build Status](https://app.travis-ci.com/akimimi/getuigo.svg?branch=master)](https://app.travis-ci.com/akimimi/getuigo)
[![Coverage Status](https://coveralls.io/repos/github/akimimi/getuigo/badge.svg?branch=master)](https://coveralls.io/github/akimimi/getuigo?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/akimimi/getuigo)](https://goreportcard.com/report/github.com/akimimi/getuigo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

### 简介
项目在github.com/geek-go/getui的基础上封装了个推的go语言SDK。
项目定义了Payload数据结构和接口，并且封装了发送流程接口。

### 数据结构说明
#### GetuiConfig
```
type GetuiConfig struct {
     AppId        string `yaml:"appid"`
     AppKey       string `yaml:"appkey"`
     AppSecret    string `yaml:"appsecret"`
     MasterSecret string `yaml:"mastersecret"`
 }
```
可以通过YAML或者其他方式填充个推的配置，注意请不要将秘钥包含在版本库中。

#### GetuiPush
GetuiPush是外部调用的数据结构，封装了发送推送的接口。

#### BasicPayload
```
type BasicPayload struct {
    PushTitle    string `json:"push_title"`
    PushBody     string `json:"push_body"`
    IsShowNotify int    `json:"is_show_notify"`
    Ext          string `json:"ext"`
}
```
可以通过组合BasicPayload并实现Payload接口的方式定义属于业务的Payload。

### 示例
sample 中实现了调用示例。


