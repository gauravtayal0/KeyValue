package internal

import (
	"errors"
	"fmt"
)

type DBSchema struct{
	Items map[string]Value
}

type Value map[string]interface{}

func NewDBSchema() *DBSchema {
	AttributeIndexer = map[string]AttributeSchema{}

	return &DBSchema{
		Items: map[string]Value{},
	}
}

func (dbs *DBSchema) Set(key string, value interface{}) error{

	val := value.(Value)
	for k, v:= range val{
		attribute:= GetOrCreateAttributeSchema(k, v)

		// validate attribute passed
		if vErr := attribute.Validate(v); vErr != nil{
			fmt.Println(fmt.Sprintf("%s", vErr))
			return vErr
		}
	}

	dbs.Items[key] = val
	fmt.Println(fmt.Sprintf("updated db map is: %v", dbs.Items))

	// update secondary indexes
	updateSecondaryIndexes(key, val)

	return nil
}

func (dbs *DBSchema) Get(key string) (interface{}, error){

	object, found := dbs.Items[key]

	fmt.Sprintf("all items in cache are %s", dbs.Items)

	if !found {
		err := errors.New("key not found")
		fmt.Println(fmt.Sprintf("error for get operation on key %s is: %s", key, err))
		return nil, err
	}

	fmt.Println(fmt.Sprintf("value of key %s is: %s", key, object))
	return object, nil
}

func (dbs *DBSchema) GetForColumnValue(field string, key interface{}) ([]string, error){

	attribute := GetAttributeSchema(field)
	IndexName := fmt.Sprint(key)

	object, found := attribute.SecondaryIndex[IndexName]
	if !found {
		err := errors.New("key not found in secondary indexes")
		fmt.Println(fmt.Sprintf("error finding value in secondary index for key %s is: %s", IndexName, err))
		return nil, err
	}

	for _, v := range object{
		if object, itemfound := dbs.Items[v]; itemfound{
			fmt.Printf("value for key %s are %v", v, object)
		}
	}
	return object, nil

}

func (dbs *DBSchema) Delete(key string){
	// todo: delete from secondary index
	delete(dbs.Items, key)

	fmt.Println(fmt.Sprintf("updated map after delete is: %v", dbs.Items))

	// update secondary indexes

}