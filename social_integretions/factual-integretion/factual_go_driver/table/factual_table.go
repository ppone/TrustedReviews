package table

import (
	"errors"
	"fmt"
)

type FactTable struct {
	tableName string
}

var factTables = map[string]FactTable{
	"places":              FactTable{"places"},
	"places-edge":         FactTable{"places-edge"},
	"place-categories":    FactTable{"place-categories"},
	"restaurants-us":      FactTable{"restaurants-us"},
	"restaurants-us-edge": FactTable{"restaurants-us-edge"},
	"restaurants-gb":      FactTable{"restaurants-gb"},
	"hotels-us":           FactTable{"hotels-us"},
	"world-geographies":   FactTable{"world-geographies"},
	"crosswalk":           FactTable{"crosswalk"},
	"products-cpg":        FactTable{"products-cpg"},
	"products-crosswalk":  FactTable{"products-crosswalk"},
}

func ListAllTables() []string {
	var tableList []string

	for key, _ := range factTables {
		tableList = append(tableList, key)

	}
	return tableList
}

func NewTable(tableName string) (FactTable, error) {
	value, ok := factTables[tableName]
	if ok {
		return value, nil
	}

	return FactTable{}, errors.New("The table you entered " + tableName + " in not in the list of tables; see ListAllTables ")

}

func (Table FactTable) ToJson() string {
	return "/t/" + Table.tableName
}

func main() {
	tab, err := NewTable("places")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tab.ToJson(), ListAllTables())

}
