package models

type Record struct {
	PK   string `dynamodbav:"PK"  json:"pk"`
	SK   string `dynamodbav:"SK"  json:"sk"`
	Type string `dynamodbav:"Typ" json:"typ"`
}

