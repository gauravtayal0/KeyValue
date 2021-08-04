package internal

import "fmt"

func updateSecondaryIndexes(key string, value Value){
	for k, v:= range value{
		attribute := GetAttributeSchema(k)

		if attribute.Indexer{
			indexName := fmt.Sprint(v)
			attribute.SecondaryIndex[indexName] = append(attribute.SecondaryIndex[indexName], key)
		}
	}
}
