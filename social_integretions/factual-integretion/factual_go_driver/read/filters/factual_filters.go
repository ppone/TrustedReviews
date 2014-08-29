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

type Filter interface {
	String() string
	Raw() string

	/*
		And(f1, f2 *filter) (*filter, error)
		Or(f1, f2 *filter) (*filter, error)
		Blank(keyword string, b bool) *filter
		BeginsWith(keyword string, value string) (*filter, error)
		BeginsWithAny(keyword string, values ...string) (*filter, error)
		EqualTo(keyword string, value interface{}) (*filter, error)
		Excludes(keyword string, value interface{}) (*filter, error)
		ExcludesAny(keyword string, values ...interface{}) (*filter, error)
		GreaterThan(keyword string, value interface{}) (*filter, error)
		GreaterThanEqual(keyword string, value interface{}) (*filter, error)
		EqualsAnyOf(keyword string, values ...interface{}) (*filter, error)
		Includes(keyword string, values interface{}) (*filter, error)
		IncludesAny(keyword string, values ...interfasce{}) (*filter, error)
		LessThan(keyword string, value interface{}) (*filter, error)
		LessThanEqual(keyword string, values interface{}) (*filter, error)
		NotBeginWith(keyword string, value string) (*filter, error)
		NotBeginWithAny(keyword string, values ...string) (*filter, error)
		NotEqualTo(keyword string, value interface{}) (*filter, error)
		NotEqualAnyOf(keyword string, values ...interface{}) (*filter, error)
		Search(keyword string, value string) (*filter, error)*/
}

type filter struct {
	filterValue string
}

func (F *filter) String() string {
	return "filters=" + F.filterValue
}

func (F *filter) Raw() string {
	return F.filterValue
}

func truncateLastZeroInFloatString(s string) string {
	splitedString := strings.Split(s, ".")

	if len(splitedString) == 2 && len(splitedString[1]) == 2 && string(splitedString[1][1]) == "0" {
		return splitedString[0] + "." + string(splitedString[1][0])
	}

	return s
}

func floatToString(f float64) string {
	return truncateLastZeroInFloatString(strconv.FormatFloat(f, 'f', floatPrecision, floatBits))
}

func intToString(i int) string {
	return strconv.Itoa(i)
}

func covertSliceToInterfaceSlice(t interface{}) ([]interface{}, error) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		rl := make([]interface{}, s.Len())

		for i := 0; i < s.Len(); i++ {
			rl[i] = s.Index(i).Interface()
		}

		return rl, nil
	}

	return nil, errors.New("Input to function needs to be a slice")
}

func NewInterfaceSlice(n ...interface{}) []interface{} {
	return n
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

func returnFilter(keyword, operator string, value interface{}) (*filter, error) {

	switch v := value.(type) {
	case int:
		return &filter{"{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + intToString(v) + "}}"}, nil
	case float64:
		return &filter{"{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + floatToString(v) + "}}"}, nil
	case string:
		return &filter{"{" + "\"" + keyword + "\":" + "{\"" + operator + "\":\"" + v + "\"}}"}, nil
	case []string:

		unboxedValue, err := returnJsonArray(value)

		if err != nil {
			return nil, err
		}
		return &filter{"{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + unboxedValue + "}}"}, nil
	case []int:
		unboxedValue, err := returnJsonArray(value)
		if err != nil {
			return nil, err
		}
		return &filter{"{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + unboxedValue + "\"}}"}, nil
	case []float64:
		unboxedValue, err := returnJsonArray(value)
		if err != nil {
			return nil, err
		}
		return &filter{"{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + unboxedValue + "\"}}"}, nil
	case []interface{}:
		unboxedValue, err := returnJsonArray(value)
		if err != nil {
			return nil, err
		}
		return &filter{"{" + "\"" + keyword + "\":" + "{\"" + operator + "\":" + unboxedValue + "}}"}, nil

	default:
		return nil, errors.New("Error: Accepts only Text or Numeric Types")

	}

}

func And(f1, f2 Filter) (Filter, error) {
	if f1 == nil || f2 == nil {
		return nil, errors.New("filter1 or filter2 cannot be nil")
	}

	s := "{\"$and\":[" + f1.Raw() + "," + f2.Raw() + "]}"

	return &filter{s}, nil

}

func Or(f1, f2 Filter) (Filter, error) {
	if f1 == nil || f2 == nil {
		return nil, errors.New("filter1 or filter2 cannot be nil")
	}

	s := "{\"$or\":[" + f1.Raw() + "," + f2.Raw() + "]}"

	return &filter{s}, nil

}

func Blank(keyword string, b bool) Filter {
	if b == true {
		return &filter{"{\"" + keyword + "\":" + "{\"$blank\":true}}"}
	} else {
		return &filter{"{\"" + keyword + "\":" + "{\"$blank\":false}}"}
	}
}

func BeginsWith(keyword string, value string) (Filter, error) {
	filterPointer, err := returnFilter(keyword, "$bw", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil
}

func BeginsWithAny(keyword string, values ...string) (Filter, error) {
	filterPointer, err := returnFilter(keyword, "$bwin", values)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func EqualTo(keyword string, value interface{}) (Filter, error) {
	_, errA := checkTypes(value, stringAndNumericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$eq", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func Excludes(keyword string, value interface{}) (Filter, error) {
	_, errA := checkTypes(value, stringAndNumericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$excludes", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func ExcludesAny(keyword string, values ...interface{}) (Filter, error) {
	_, errA := checkTypesInArray(values, stringAndNumericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$excludes_any", values)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func GreaterThan(keyword string, value interface{}) (Filter, error) {

	_, errA := checkTypes(value, numericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$gt", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func GreaterThanEqual(keyword string, value interface{}) (Filter, error) {

	_, errA := checkTypes(value, numericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$gte", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func EqualsAnyOf(keyword string, values ...interface{}) (Filter, error) {
	_, errA := checkTypesInArray(values, stringAndNumericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$in", values)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func Includes(keyword string, values interface{}) (Filter, error) {
	_, errA := checkTypes(values, stringAndNumericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$includes", values)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func IncludesAny(keyword string, values ...interface{}) (Filter, error) {
	_, errA := checkTypesInArray(values, stringAndNumericT)

	if errA != nil {
		return nil, errA
	}
	filterPointer, err := returnFilter(keyword, "$includes_any", values)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func LessThan(keyword string, value interface{}) (Filter, error) {
	_, errA := checkTypes(value, numericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$lt", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func LessThanEqual(keyword string, values interface{}) (Filter, error) {
	_, errA := checkTypes(values, numericT)

	if errA != nil {
		return nil, errA
	}
	filterPointer, err := returnFilter(keyword, "$lte", values)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func NotBeginWith(keyword string, value string) (Filter, error) {

	filterPointer, err := returnFilter(keyword, "$nbw", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func NotBeginWithAny(keyword string, values ...string) (Filter, error) {

	filterPointer, err := returnFilter(keyword, "$nbwin", values)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func NotEqualTo(keyword string, value interface{}) (Filter, error) {

	_, errA := checkTypes(value, stringAndNumericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$neq", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func NotEqualAnyOf(keyword string, values ...interface{}) (Filter, error) {

	_, errA := checkTypesInArray(values, stringAndNumericT)

	if errA != nil {
		return nil, errA
	}

	filterPointer, err := returnFilter(keyword, "$nin", values)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}

func Search(keyword string, value string) (Filter, error) {

	filterPointer, err := returnFilter(keyword, "$search", value)
	if err != nil {
		return nil, err
	}
	return filterPointer, nil

}
