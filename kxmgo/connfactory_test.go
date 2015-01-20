package kxmgo

import "testing"

//func TestMongoDb_InitMongo(t *testing.T) {
//Init_mongo()
//}

func TestConnFactory_AddConnection(t *testing.T) {
	err := AddConnect("mongodb://test:test@127.0.0.1/dbtest")
	if err != nil {
		t.Fatal(err)
	}

}

func TestConnFactory_GetMgo(t *testing.T) {
	err := AddConnect(testDb)
	if err != nil {
		t.Fatal(err)
	}
	mgo := GetMgo(testDb)
	defer mgo.Close()
	if mgo == nil {
		t.Fatal("mgo should not be nil")
	}
}
