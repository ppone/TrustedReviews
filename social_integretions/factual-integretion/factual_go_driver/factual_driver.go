package main

import "oauth1"
import "fmt"
import "net/http"
import "net/url"
import "io/ioutil"

//import "strings"

const VERSION = "V3"

type FactualQuery struct {
	query string
	limit int
}

func EncodeURLParameters(urlstring string) string {
	encodedURLValue, err := url.Parse(urlstring)
	if err != nil {
		panic("Could not parse the url")
	}
	encodedURLValue.RawQuery = encodedURLValue.Query().Encode()

	return encodedURLValue.String()
}

func AddQuery(urlstring *string, query string) string {

	query = "q=" + query
	return *urlstring + query

}

func Read() {

}

func main() {
	token := oauth1.NewAccessToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO", "", "")
	urlstring := EncodeURLParameters("http://api.v3.factual.com/t/restaurants-us/schema")
	fmt.Println("TESTING QUERY STRING", AddQuery(&urlstring, "Los gatos"))
	request, err := http.NewRequest("GET", urlstring, nil)
	client := &http.Client{}

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	token.SignRequestHeader(request)

	result, err := client.Do(request)

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	content, err := ioutil.ReadAll(result.Body)

	defer result.Body.Close()

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	//fmt.Println(request)
	fmt.Println(result)

	fmt.Println(request)
	fmt.Printf("%s", content)

}
