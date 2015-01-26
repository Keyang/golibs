package kxmgo

import (
	"fmt"
	"testing"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	testDb = "mongodb://test:test@127.0.0.1/dbtest"
)

type TestObj struct {
	Id     string `bson:"_id"`
	FieldA string
	FieldB int
	FieldC time.Time
}

func init() {
	AddConnect(testDb)
}
func TestMongoDb_AddFindOneFindByIdRemove(t *testing.T) {
	m := GetMgo(testDb)
	obj := TestObj{
		Id:     NewId(),
		FieldA: "Hello",
		FieldB: 10,
		FieldC: time.Now(),
	}
	err := m.AddNew("test", &obj)
	if err != nil {
		t.Fatal(err)
	}
	filter := bson.M{"fielda": "Hello", "fieldb": 10}
	c, err := m.Count("test", filter)
	if err != nil {
		t.Fatal(err)
	}
	if c != 1 {
		t.Log("Should contain 1 entry but %d", c)
		t.Fail()
	}
	obj1 := TestObj{}
	err = m.FindOne("test", filter, &obj1)
	if err != nil {
		t.Fatal(err)
	}
	if obj1.Id != obj.Id {
		t.Log("Should retrieve object %#v but %#v", obj, obj1)
		t.Fail()
	}
	obj2 := TestObj{}
	err = m.FindById("test", obj.Id, &obj2)
	if err != nil {
		t.Fatal(err)
	}
	if obj2.Id != obj.Id {
		t.Log("Should retrieve object %#v but %#v", obj, obj2)
		t.Fail()
	}
	err = m.RemoveById("test", obj.Id)
	if err != nil {
		t.Fatal(err)
	}
	c, err = m.Count("test", filter)
	if err != nil {
		t.Fatal(err)
	}
	if c != 0 {
		t.Fatal(fmt.Sprintf("Should contain 0 entry but %d", c))
	}
}

func TestMgo_FindAllCountFindBy(t *testing.T) {
	m := GetMgo(testDb)
	objA := TestObj{
		Id:     NewId(),
		FieldA: "Hello",
		FieldB: 10,
		FieldC: time.Now(),
	}
	objB := TestObj{
		Id:     NewId(),
		FieldA: "Hello",
		FieldB: 10,
		FieldC: time.Now(),
	}
	objC := TestObj{
		Id:     NewId(),
		FieldA: "Hello",
		FieldB: 10,
		FieldC: time.Now(),
	}
	defer m.RemoveById("test", objA.Id)
	defer m.RemoveById("test", objB.Id)
	defer m.RemoveById("test", objC.Id)
	err := m.AddNew("test", objA, objB, objC)
	if err != nil {
		t.Fatal(err)
	}
	objs := []TestObj{}
	err = m.FindAll("test", bson.M{"fielda": "Hello", "fieldb": 10}, &objs)
	if err != nil {
		t.Fatal(err)
	}
	if len(objs) != 3 {
		t.Log("Should return 3 items but returned %d", len(objs))
		t.Fail()
	}
}
