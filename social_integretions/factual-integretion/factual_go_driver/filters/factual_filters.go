/*
TODO: Need to consider making external functions more strongly typed.  Right now we infer the types at runtime, might be better
to change arguments to specifc types instead of interface{}.   This will catch certain errors at compile time instead of waiting for
runtime.


*/
package filters

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

const (
	floatPrecision = 2
	floatBits      = 64
)

type jsType int

const (
	stringT           = iota // 1
	numericT                 // = iota
	stringAndNumericT        // = iota
)

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', floatPrecision, floatBits)
}

func intToString(i int) string {
	return strconv.Itoa(i)
}

func returnJsonArray(valueBox interface{}) (string, error) {

	var values []interface{}

	var jsonType string

	switch v := valueBox.(type) {
	case []string:

		values = make([]interface{}, len(v))
		jsonType = "string"
		for i, j := range v {

			values[i] = j
		}
	case []int:
		values = make([]interface{}, len(v))
		jsonType = "int"
		for i, j := range v {
			values[i] = j
		}
	case []float64:
		values = make([]interface{}, len(v))
		jsonType = "float64"
		for i, j := range v {
			values[i] = j
		}
	case []interface{}:
		var err error
		jsonType, err = checkTypesInArray(v, stringAndNumericT)
		if err != nil {
			return "", err
		}
		values = v

	default:
		return "", errors.New("valueBox is not []string,[]int, or []float64")
	}

	toArray := "["
	for _, s := range values {
		if jsonType == "string" {
			toArray += "\""
		}

		if jsonType == "string" {
			toArray += s.(string)
		} else if jsonType == "float64" {
			toArray += floatToString(s.(float64))
		} else if jsonType == "int" {
			toArray += intToString(s.(int))
		} else {
			return "", errors.New("Type is neither String or Numeric(int or float64)")
		}

		if jsonType == "string" {
			toArray += "\""
		}
		toArray += ","
	}
	toArray = strings.TrimRight(toArray, ",")
	toArray += "]"
	return toArray, nil
}

func verifyType(valueType string, acceptedTypes jsType) error {
	if acceptedTypes == stringT {

		if valueType == "string" {
			return nil
		} else {
			return errors.New("Error: Expected Type was a string, but instead got a " + valueType)
		}
	}

	if acceptedTypes == numericT {

		if valueType == "int" || valueType == "float64" {
			return nil
		} else {
			return errors.New("Error: Expected Type was a numeric type (float64 or int), but instead got a " + valueType)
		}
	}

	if acceptedTypes == stringAndNumericT {

		if valueType == "int" || valueType == "float64" || valueType == "string" {
			return nil
		} else {
			return errors.New("Error: Expected Type was a numeric type (float64 or int), or string, but instead got a " + valueType)
		}
	}

	return errors.New("Unknown valueType given or unknown jsType given")

}

func checkTypes(values interface{}, acceptedTypes jsType) (string, error) {
	//Get the first type of the first value
	valueType := reflect.TypeOf(values).String()
	err := verifyType(valueType, acceptedTypes)

	if err != nil {
		return "", err
	}

	return valueType, nil

}

func checkTypesInArray(values []interface{}, acceptedTypes jsType) (string, error) {
	//Get the first type of the first value
	valueType := reflect.TypeOf(values[0]).String()

	for _, value := range values {
		if valueType != reflect.TypeOf(value).String() {
			return "", errors.New("Types in array are inconsistent")
		}

	}

	err := verifyType(valueType, acceptedTypes)

	if err != nil {
		return "", err
	}

	return valueType, nil

}

func returnJsonString(keyword, operator string, value interface{}) (string, error) {

	switch v := value.(type) {
	case int:
		return "{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + intToString(v) + "}}", nil
	case float64:
		return "{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + floatToString(v) + "}}", nil
	case string:
		return "{" + "\"" + keyword + "\":" + "{\"" + operator + "\":\"" + v + "\"}}", nil
	case []string:

		unboxedValue, err := returnJsonArray(value)

		if err != nil {
			return "", err
		}
		return "{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + unboxedValue + "}}", nil
	case []int:
		unboxedValue, err := returnJsonArray(value)
		if err != nil {
			return "", err
		}
		return "{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + unboxedValue + "\"}}", nil
	case []float64:
		unboxedValue, err := returnJsonArray(value)
		if err != nil {
			return "", err
		}
		return "{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + unboxedValue + "\"}}", nil
	case []interface{}:
		unboxedValue, err := returnJsonArray(value)
		if err != nil {
			return "", err
		}
		return "{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + unboxedValue + "}}", nil

	default:
		return "", errors.New("Error: Accepts only Text or Numeric Types")

	}

}

func Blank(keyword string, b bool) string {
	if b == true {
		return "{\"" + keyword + "\":" + "{\"$blank\":true}}"
	} else {
		return "{\"" + keyword + "\":" + "{\"$blank\":false}}"
	}
}

func BeginsWith(keyword string, value string) (string, error) {
	jsonString, err := returnJsonString(keyword, "$bw", value)
	if err != nil {
		return "", err
	}
	return jsonString, nil
}

func BeginsWithAny(keyword string, values ...string) (string, error) {
	jsonString, err := returnJsonString(keyword, "$bwin", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func EqualTo(keyword string, value interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$eq", value)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func Excludes(keyword string, values interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$excludes", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func ExcludesAny(keyword string, values ...interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$excludes_any", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func GreaterThan(keyword string, values interface{}) (string, error) {
	jsonType, err := checkTypes(values, numericT)
	if err != nil {
		return "", err
	}
	if jsonType == "string" {
		return "", errors.New("Can only accept numeric types")
	}

	jsonString, err := returnJsonString(keyword, "$gt", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func GreaterThanEqual(keyword string, values interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$gte", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func EqualsAnyOf(keyword string, values interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$in", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func Includes(keyword string, values interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$includes", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func IncludesAny(keyword string, values interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$includes", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func LessThan(keyword string, values interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$lt", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func LessThanEqual(keyword string, values interface{}) (string, error) {
	jsonString, err := returnJsonString(keyword, "$lte", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func NotBeginWith(keyword string, values interface{}) (string, error) {

	jsonString, err := returnJsonString(keyword, "$nbw", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func NotBeginWithAny(keyword string, values interface{}) (string, error) {

	jsonString, err := returnJsonString(keyword, "$nbwin", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func NotEqualTo(keyword string, values interface{}) (string, error) {

	jsonString, err := returnJsonString(keyword, "$nbwin", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func NotEqualAnyOf(keyword string, values interface{}) (string, error) {

	jsonString, err := returnJsonString(keyword, "$nbwin", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func Search(keyword string, values interface{}) (string, error) {

	jsonString, err := returnJsonString(keyword, "$search", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}
