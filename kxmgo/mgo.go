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
func (m *Mgo) Close() {
	if m.Conn != nil {
		m.Conn.Close()
	}
}

//Get Collection based on its name
func (m *Mgo) GetCol(colName string) *mgo.Collection {
	ses := m.Conn
	return ses.DB("").C(colName)
}

//Remove a document based on its id from a collection
func (m *Mgo) RemoveById(colName string, id interface{}) error {
	col := m.GetCol(colName)
	return col.RemoveId(id)
}

//Add a document to a collection
func (m *Mgo) AddNew(colName string, d ...interface{}) error {
	col := m.GetCol(colName)
	return col.Insert(d...)
}

//Retrieve multiple documents based on filter
func (m *Mgo) FindAll(colName string, filter bson.M, d interface{}) error {
	query := m.FindBy(colName, filter)
	return query.All(d)
}

//Retrieve one documents based on filter
func (m *Mgo) FindOne(colName string, filter bson.M, d interface{}) error {
	query := m.FindBy(colName, filter)
	return query.One(d)
}

//Retrieve number of documents meeting the filter.
func (m *Mgo) Count(colName string, filter bson.M) (int, error) {
	query := m.FindBy(colName, filter)
	return query.Count()
}

//Retrieve a *mgo.Query based on filter
func (m *Mgo) FindBy(colName string, filter bson.M) *mgo.Query {
	col := m.GetCol(colName)
	return col.Find(filter)
}

//Find a document by its id
func (m *Mgo) FindById(colName string, id interface{}, d interface{}) error {
	col := m.GetCol(colName)
	return col.FindId(id).One(d)
}
