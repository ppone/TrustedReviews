package main

import "oauth1"
import "fmt"
import "net/http"

func main() {
	token := oauth1.NewAccessToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO", "", "")
	fmt.Println(token)
	request, err := http.NewRequest("GET", "http://api.v3.factual.com/t/places/03c26917-5d66-4de9-96bc-b13066173c65", nil)
	//request, err := http.Get("http://api.v3.factual.com/t/places/03c26917-5d66-4de9-96bc-b13066173c65")
	client := &http.Client{}

	if err != nil {
		fmt.Println(err)
	}

	token.SignRequestHeader(request)

	result, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", result)

	//defer request.Body.Close()
}
