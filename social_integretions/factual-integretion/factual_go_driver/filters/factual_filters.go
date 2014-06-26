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
	return strconv.Itoa(v)
}

func returnArrayString(values []interface{}, jsonType string) (string, error) {

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

func checkTypesInArrayAreStringOrNumeric(values []interface{}) (string, error) {
	//Get the first type of the first value
	firstValueType := reflect.TypeOf(values[0]).String()
	fmt.Println("first type is ", firstValueType)
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
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + strconv.Itoa(v) + "}}", nil
	case float64:
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + strconv.FormatFloat(v, 'f', 8, 64) + "}}", nil
	case string:
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + v + "\"}}", nil
	case []string:
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + returnArrayString(v, true) + "\"}}", nil
	case []int:
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + returnArrayString(v, false) + "\"}}", nil
	case []float64:
		return "{" + "\"" + keyword + ":" + "{\"" + operator + "\":" + returnArrayString(v, false) + "\"}}", nil
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
	if err {
		return nil, err
	}
	return jsonString, nil
}

func BeginsWithAny(keyword string, values ...string) string {

	return "{" + "\"" + keyword + ":" + "{\"$bwin\":" + returnArrayString(values) + "}}"

}

func EqualTo(keyword string, value interface{}) (string, error) {

	switch v := value.(type) {
	case int:
		return "{" + "\"" + keyword + ":" + "{\"$eq\":" + strconv.Itoa(v) + "}}", nil
	case float64:
		return "{" + "\"" + keyword + ":" + "{\"$eq\":" + strconv.FormatFloat(v, 'f', 8, 64) + "}}", nil
	case string:
		return "{" + "\"" + keyword + ":" + "{\"$eq\":\"" + v + "\"}}", nil
	default:
		return "", errors.New("Error: Accepts only Text or Numeric Types")

	}

}

func Excludes(keyword string, values interface{}) string {

	switch v := value.(type) {
	case int:
		return "{" + "\"" + keyword + ":" + "{\"$excludes\":" + strconv.Itoa(v) + "}}", nil
	case float64:
		return "{" + "\"" + keyword + ":" + "{\"$excludes\":" + strconv.FormatFloat(v, 'f', 8, 64) + "}}", nil
	case string:
		return "{" + "\"" + keyword + ":" + "{\"$eq\":\"" + v + "\"}}", nil
	default:
		return "", errors.New("Error: Accepts only Text or Numeric Types")

	}

	return "{" + "\"" + keyword + ":" + "{\"$bwin\":" + returnArrayString(values) + "}}"

}
