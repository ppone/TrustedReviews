package filters

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func floatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 8, 64)
}

func intToString(i int) string {
	return strconv.Itoa(i)
}

func returnArrayString(valueBox interface{}) (string, error) {

	jsonType, err := checkTypesInArrayAreStringOrNumeric(valueBox)
	if err != nil {
		return "", err
	}
	values := valueBox.([]interface{})
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

func checkTypesInArrayAreStringOrNumeric(valueBox interface{}) (string, error) {
	//Get the first type of the first value
	values := valueBox.([]interface{})
	firstValueType := reflect.TypeOf(values[0]).String()

	for _, value := range values {
		if firstValueType != reflect.TypeOf(value).String() {
			return "", errors.New("Types in array are inconsistent")
		}

	}

	if firstValueType == "string" || firstValueType == "int" || firstValueType == "float64" {
		return firstValueType, nil
	} else {
		return "", errors.New("Type is neither String or Numeric(int or float64)")
	}

}

func returnJsonString(keyword, operator string, value interface{}) (string, error) {

	switch v := value.(type) {
	case int:
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + intToString(v) + "}}", nil
	case float64:
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + floatToString(v) + "}}", nil
	case string:
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + v + "\"}}", nil
	case []string:
		unboxedValue, err := returnArrayString(value)
		if err != nil {
			return "", err
		}
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + unboxedValue + "\"}}", nil
	case []int:
		unboxedValue, err := returnArrayString(value)
		if err != nil {
			return "", err
		}
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + unboxedValue + "\"}}", nil
	case []float64:
		unboxedValue, err := returnArrayString(value)
		if err != nil {
			return "", err
		}
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + unboxedValue + "\"}}", nil
	default:
		return "", errors.New("Error: Accepts only Text or Numeric Types")

	}

}

func Blank(keyword string, b bool) string {
	if b == true {
		return keyword + ":" + "{\"$blank\":true}"
	} else {
		return keyword + ":" + "{\"$blank\":false}"
	}
}

func BeginsWith(keyword string, value string) (string, error) {
	jsonString, err := returnJsonString(keyword, "bw", value)
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

func ExcludesAny(keyword string, values interface{}) (string, error) {

	jsonString, err := returnJsonString(keyword, "$excludes_any", values)
	if err != nil {
		return "", err
	}
	return jsonString, nil

}

func GreaterThan(keyword string, values interface{}) (string, error) {

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

	valueType := reflect.TypeOf(values).String()

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
