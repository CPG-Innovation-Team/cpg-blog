package qo

import "cpg-blog/global/cpgConst"

type Permission struct {
	Name    string
	Uri     string
	Operate string
}
type GroupAddPermission struct {
	GName string `json:"gName"` //group name
	PName string `json:"pName"` //policy name
}

func GetNewPermission() (p Permission) {
	p.Operate = cpgConst.Operate
	return
}
