package main

import "oauth1"
import "fmt"

func main() {
	token := oauth1.NewAccessToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO", "", "")
	fmt.Println(token)
}
