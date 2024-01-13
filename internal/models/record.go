package models

type Record struct {
	PK   string `dynamodbav:"pk"  json:"pk"`
	SK   string `dynamodbav:"sk"  json:"sk"`
	Type string `dynamodbav:"typ" json:"typ"`
}

