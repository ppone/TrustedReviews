package main

import "oauth1"
import "fmt"
import "net/http"

func main() {
	token := oauth1.NewAccessToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO", "", "")
	fmt.Println(token)
	request, err := http.NewRequest("GET", "http://host.net/resource?a=b&c=d", nil)

	if err != nil {
		fmt.Println(err)
	}

	token.SignRequestQuery(request)
	fmt.Println(request)
}
