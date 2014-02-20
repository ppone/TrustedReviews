package main

import "oauth1"
import "fmt"
import "net/http"
import "net/url"
import "io/ioutil"
import "strings"

func main() {
	token := oauth1.NewAccessToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO", "", "")
	urlstring := "http://api.v3.factual.com/t/restaurants-us?filters={\"$and\":[{\"region\":{\"$eq\":\"FL\"}},{\"category_labels\":{\"$eq\":\"[%5C\"SOCIAL%5C\",%5C\"FOOD+AND+DINING%5C\",%5C\"RESTAURANTS%5C\"]\"}},{\"chain_name\":{\"$eq\":\"MCDONALD'S\"}}]}"
	//urlstring = strings.Replace(urlstring, " ", "+", -1)
	//urlstring = url.QueryEscape(urlstring)
	purl, err := url.Parse(urlstring)

	if err != nil {
		fmt.Println("ERROR ->", err)

	}

	fmt.Println("PURL.PATH -> ", purl.Path)
	fmt.Println("PURL.HOST -> ", purl.Host)
	fmt.Println("PURL.SCHEME -> ", purl.Scheme)
	fmt.Println("PURL -> ", purl.RawQuery)

	k := purl.Query()

	koutput := ""

	for e, f := range k {
		//fencoded := strings.Replace(f[0], " ", "+", -1)
		//fencoded := url.QueryEscape(f[0])
		koutput += e
		koutput += "="
		koutput += url.QueryEscape(f[0])
		koutput += "&"
	}

	koutput = strings.TrimRight(koutput, "&")

	fmt.Println("koutput", koutput)

	ut := purl.Scheme + "://" + purl.Host + purl.Path + "?" + koutput
	fmt.Println(ut)

	urlstring = ut

	request, err := http.NewRequest("GET", urlstring, nil)
	client := &http.Client{}

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	token.SignRequestHeader(request)
	//request.URL.RawQuery = "q=Coffee,Los Angeles&limit=1"
	//request.URL.Opaque = "q=Coffee,Los%20Angeles&limit=1"

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
