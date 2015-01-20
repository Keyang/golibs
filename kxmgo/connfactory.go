package kxmgo

import (
	"gopkg.in/mgo.v2"
	"keyangxiang.com/libs/kxlog"
)

var (
	conns map[string]*mgo.Session
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
			conns = map[string]*mgo.Session{}
		}
		conns[connStr] = conn
		kxlog.I("Successfully connected to: %s", connStr)
	} else {
		kxlog.I("%s is already existed in Mongodb Connection pool. Skip the operation.", connStr)
	}
	return nil
}

//Get a mgo.Session copy of existing connection from connection pool
func GetMgo(connStr string) *Mgo {
	if conn, ok := conns[connStr]; ok == true {
		conn_copy := conn.Copy()
		return &Mgo{
			Conn: conn_copy,
		}
	} else {
		kxlog.W("%s cannot be found in mongodb connection pool.", connStr)
		err := AddConnect(connStr)
		if err != nil {
			panic(err)
		}
		return GetMgo(connStr)
	}
}
