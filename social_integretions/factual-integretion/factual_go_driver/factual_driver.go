package main

import "oauth1"
import "fmt"
import "net/http"
import "net/url"
import "io/ioutil"
import "./read"
import "errors"

//import "strings"

const VERSION = "v3"
const FACTUAL_BASE_URL = "http://api.v3.factual.com"

func EncodeURLParameters(urlstring string) string {
	encodedURLValue, err := url.Parse(urlstring)
	if err != nil {
		panic("Could not parse the url")
	}
	encodedURLValue.RawQuery = encodedURLValue.Query().Encode()

	return encodedURLValue.String()
}

type factualRequest struct {
	token *oauth1.AccessToken
	read.Read
}

func NewRequest(consumerKey, consumerSecret string) *factualRequest {
	f := new(factualRequest)
	f.token = oauth1.NewAccessToken(consumerKey, consumerSecret, "", "")
	return f
}

func (F *factualRequest) NewRead(tableName string) (*factualRequest, error) {
	f, err := read.NewReader(tableName)
	if err != nil {
		return nil, err
	}
	F.Read = f
	return F, nil

}

func (F *factualRequest) GetReadURL() string {
	return FACTUAL_BASE_URL + F.String()
}

func (F *factualRequest) GetRead() (string, error) {
	if F.Read == nil {
		return "", errors.New("Factual Request has no read data")
	}
	encodedURL := EncodeURLParameters(F.GetReadURL())
	request, err := http.NewRequest("GET", encodedURL, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	F.token.SignRequestHeader(request)
	result, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer result.Body.Close()

	content, err := ioutil.ReadAll(result.Body)

	if err != nil {
		return "", err
	}

	return string(content), nil

}

func main() {

	fR := NewRequest("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO")

	_, err := fR.NewRead("restaurants-us")

	if err != nil {
		fmt.Println(err)
		return
	}

	fR.AddQuery("pizza", "sushi")
	fR.AddLimit(5)
	fR.AddSelect("factual_id")

	fmt.Println("frData", fR.GetReadURL())

	s, err := fR.GetRead()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)

	//token := oauth1.NewAccessToken("Pbu7jRdBErgLW07g9c25JtGcwwt1KmpoxRTfFL3x", "vC4AgocPBhxe0GFkTsetoiuEAJEgqz6MCbAnXEoO", "", "")
	/*
		r, err := read.NewRead("restaurants-us")
		if err != nil {
			fmt.Println(err)
			return
		}

		r, err = r.AddQuery("Coffee", "\"Los Angeles\"")

		if err != nil {
			fmt.Println(err)
			return
		}

		r = r.AddLimit(5)
		//r, err = r.AddSelect("cuisine")

		if err != nil {
			fmt.Println(err)
			return
		}*/

	/*fmt.Println("TESTING RAW QUERY STRING", FACTUAL_BASE_URL+r.String())
	urlstring := EncodeURLParameters(FACTUAL_BASE_URL + r.String())
	durlstring := EncodeURLParameters("http://api.v3.factual.com/t/restaurants-us?q=Coffee,\"Los Angeles\"&limit=5&select=factual_id")
	fmt.Println("TESTING QUERY STRING", durlstring, urlstring)
	request, err := http.NewRequest("GET", durlstring, nil)
	client := &http.Client{}

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	token.SignRequestHeader(request)

	result, err := client.Do(request)

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	//_, err := ioutil.ReadAll(result.Body)
	content, err := ioutil.ReadAll(result.Body)

	defer result.Body.Close()

	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	//fmt.Println(request)
	fmt.Println("result => ", result)

	fmt.Println("request => ", request, "REQUEST_URL => ", request.URL)
	fmt.Printf("%s", content)
	*/

}
