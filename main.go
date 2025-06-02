package main

import (
	"fmt"
	"learn/pkg"

	"gorm.io/gorm"
)

type QueryBuilder struct {
	Operator   string
	Conditions *[]QueryCondition
}

type QueryCondition struct {
	Field    string
	Operator string
	Value    any
}

func addOrConditions(db *gorm.DB, queryBuilder *[]QueryCondition) *gorm.DB {
	if len(*queryBuilder) == 0 {
		return db
	}

	var query string
	var args []any

	for i, cond := range *queryBuilder {
		if i > 0 {
			query += " OR "
		}
		query += fmt.Sprintf("%s %s ?", cond.Field, cond.Operator)
		args = append(args, cond.Value)
	}

	return db.Where(query, args...)
}

func ConvertMapToQuery(db *gorm.DB, queryBuilders *[]QueryBuilder) *gorm.DB {
	if queryBuilders == nil || len(*queryBuilders) == 0 {
		return db
	}

	for _, queryBuilder := range *queryBuilders {
		fmt.Println("queryBuilder", queryBuilder)
	}

	return db
}

func main() {
	// conditions := []QueryBuilder{
	// 	{Operator: "or", Conditions: &[]QueryCondition{
	// 		{Field: "national_id", Operator: "=", Value: 1},
	// 	}},
	// }

	// val, err := json.Marshal(conditions)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("val", string(val))

	pkg.RunForLoop()

}
