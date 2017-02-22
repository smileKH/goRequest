package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "huangkassssixi", ServerIP: "192.168.1.1"})
	s.Servers = append(s.Servers, Server{ServerName: "huangkaixi", ServerIP: "192.168.1.2"})
	this.Data["json"] = s
	this.ServeJSON()
}

type Server struct {
	ServerName string `json :"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}
