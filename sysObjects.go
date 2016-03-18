package models

import ()

type SystemLogging struct {
	BaseObj
	SRLogger      string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"1", DESCRIPTION: "Global logging"`
	SystemLogging string `DESCRIPTION: "Global logging", DEFAULT: "on"`
}

type ComponentLogging struct {
	BaseObj
	Module string `SNAPROUTE: "KEY", ACCESS:"w",  MULTIPLICITY:"*", DESCRIPTION: "Module name to set logging level"`
	Level  string `DESCRIPTION: "Logging level", DEFAULT: "info"`
}

type IpTableAcl struct {
	BaseObj
	Name         string `SNAPROUTE: "KEY", ACCESS:"w", MULTIPLICITY:"*", DESCRIPTION: "Ip Table ACL rule name"`
	PhysicalPort string `DESCRIPTION: "IfIndex where the acl rule is to be applied", DEFAULT: "all"`
	Action       string `DESCRIPTION: "ACCEPT or DROP"`
	IpAddr       string `DESCRIPTION: "ip address of subnet or host, e.g: 192.168.1.0/24, 192.168.1.1"`
	Protocol     string `DESCRITION: "protocol for which rule is to be applied, e.g TCP, UDP"`
	Port         string `DESCRITION: "port for protocol, e.g for dhcprelay port is 68", DEFAULT: "all"`
}

/*
type IpTableAclState struct {
	BaseObj
	Name         string `SNAPROUTE: "KEY", ACCESS:"r", MULTIPLICITY:"*", DESCRIPTION: "Ip Table ACL rule name"`
	PhysicalPort string `DESCRIPTION: "IfIndex where the acl rule is to be applied", DEFAULT: "all"`
	Action       string `DESCRIPTION: "ACCEPT or DROP"`
	IpAddr       string `DESCRIPTION: "ip address of subnet or host, e.g: 192.168.1.0/24, 192.168.1.1"`
	Protocol     string `DESCRITION: "protocol for which rule is to be applied, e.g TCP, UDP"`
	Port         string `DESCRITION: "port for protocol, e.g for dhcprelay port is 68", DEFAULT: "all"`
}
*/
