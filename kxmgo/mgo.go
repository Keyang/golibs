package kxmgo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Retrieve a new object id in string
func NewId() string {
	return bson.NewObjectId().Hex()
}

/**
*
* A basic wrapper for Mgo to provide fast and convenient functions.
 */

type Mgo struct {
	Conn *mgo.Session //the actual copy of the connection session
}

//Close the DB Connection
//func (m *Mgo) Close() {
//if m.Conn != nil {
//m.Conn.Close()
//}
//}

//update a doc by its ip
//TODO untested
func (m *Mgo) Update(colName string, id interface{}, d interface{}) error {
	col, ses := m.GetCol(colName)
	defer ses.Close()
	return col.UpdateId(id, bson.M{
		"$set": d,
	})
}

//Get Collection based on its name
func (m *Mgo) GetCol(colName string) (*mgo.Collection, *mgo.Session) {
	ses := m.Conn.Copy()
	return ses.DB("").C(colName), ses
}

//Remove a document based on its id from a collection
func (m *Mgo) RemoveById(colName string, id interface{}) error {
	col, ses := m.GetCol(colName)
	defer ses.Close()
	return col.RemoveId(id)
}

//Add a document to a collection
func (m *Mgo) AddNew(colName string, d ...interface{}) error {
	col, ses := m.GetCol(colName)
	defer ses.Close()
	return col.Insert(d...)
}

//Retrieve multiple documents based on filter
func (m *Mgo) FindAll(colName string, filter bson.M, d interface{}) error {
	query, ses := m.FindBy(colName, filter)
	defer ses.Close()
	return query.All(d)
}

//Retrieve one documents based on filter
func (m *Mgo) FindOne(colName string, filter bson.M, d interface{}) error {
	query, ses := m.FindBy(colName, filter)
	defer ses.Close()
	return query.One(d)
}

//Retrieve number of documents meeting the filter.
func (m *Mgo) Count(colName string, filter bson.M) (int, error) {
	query, ses := m.FindBy(colName, filter)
	defer ses.Close()
	return query.Count()
}

//Retrieve a *mgo.Query based on filter
func (m *Mgo) FindBy(colName string, filter bson.M) (*mgo.Query, *mgo.Session) {
	col, ses := m.GetCol(colName)
	return col.Find(filter), ses
}

//Find a document by its id
func (m *Mgo) FindById(colName string, id interface{}, d interface{}) error {
	col, ses := m.GetCol(colName)
	defer ses.Close()
	return col.FindId(id).One(d)
}

//Remove all doc match filter
//TODO untested
func (m *Mgo) RemoveAll(colName string, filter interface{}) error {
	col, ses := m.GetCol(colName)
	defer ses.Close()
	_, err := col.RemoveAll(filter)
	return err
}
