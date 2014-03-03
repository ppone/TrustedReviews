package table

import (
	"errors"
	"fmt"
)

type factTable struct {
	tableName string
}

var factTables = map[string]factTable{
	"places":              factTable{"places"},
	"places-edge":         factTable{"places-edge"},
	"place-categories":    factTable{"place-categories"},
	"restaurants-us":      factTable{"restaurants-us"},
	"restaurants-us-edge": factTable{"restaurants-us-edge"},
	"restaurants-gb":      factTable{"restaurants-gb"},
	"hotels-us":           factTable{"hotels-us"},
	"world-geographies":   factTable{"world-geographies"},
	"crosswalk":           factTable{"crosswalk"},
	"products-cpg":        factTable{"products-cpg"},
	"products-crosswalk":  factTable{"products-crosswalk"},
}

func ListAllTables() []string {
	var tableList []string

	for key, _ := range factTables {
		tableList = append(tableList, key)

	}
	return tableList
}

func NewTable(tableName string) (factTable, error) {
	value, ok := factTables[tableName]
	if ok {
		return value, nil
	}

	return factTable{}, errors.New("The table you entered " + tableName + " in not in the list of tables; see ListAllTables ")

}

func (Table factTable) ToJson() string {
	return "/t/" + Table.tableName
}

func main() {
	tab, err := NewTable("places")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tab.ToJson(), ListAllTables())

}
