package object

type ObjectType string

const (
	INTEGER_TYPE  ObjectType = "INTEGER"
	BOOLEAN_TYPE  ObjectType = "BOOLEAN"
	NULL_TYPE     ObjectType = "NULL"
	RETURN_TYPE   ObjectType = "RETURN"
	ERROR_TYPE    ObjectType = "ERROR"
	FUNCTION_TYPE ObjectType = "FUNCTION"
	STRING_TYPE   ObjectType = "STRING"
)
