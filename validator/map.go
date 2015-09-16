package validator

// Map is a map containing all validators
var Map = map[string]Validator{
	"length":  Length,
	"numeric": Numeric,
	"integer": Integer,
	"float":   Float,
}
