package kxmgo

import (
	"github.com/Keyang/golibs/kxlog"
	"gopkg.in/mgo.v2"
)

var (
	conns map[string]*Mgo
)

//Connect to a mongodb using connection string.
func AddConnect(connStr string) error {
	if _, ok := conns[connStr]; ok == false {
		kxlog.I("Connect to: %s ", connStr)
		conn, err := mgo.Dial(connStr)
		if err != nil {
			return err
		}
		if conns == nil {
			conns = map[string]*Mgo{}
		}
		conns[connStr] = &Mgo{Conn: conn}
		kxlog.I("Successfully connected to: %s", connStr)
	} else {
		kxlog.I("%s is already existed in Mongodb Connection pool. Skip the operation.", connStr)
	}
	return nil
}

//Get a mgo.Session copy of existing connection from connection pool
func GetMgo(connStr string) *Mgo {
	if m, ok := conns[connStr]; ok == true {
		return m
	} else {
		kxlog.W("%s cannot be found in mongodb connection pool.", connStr)
		err := AddConnect(connStr)
		if err != nil {
			panic(err)
		}
		return GetMgo(connStr)
	}
}
