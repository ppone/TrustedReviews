package main

import "oauth1"
import "fmt"
import "net/http"
import "io/ioutil"

func main() {
	token := oauth1.NewAccessToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO", "", "")

	request, err := http.NewRequest("GET", "http://api.v3.factual.com/t/places/03c26917-5d66-4de9-96bc-b13066173c65", nil)

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

	fmt.Printf("%s", content)

}
