package factual

import "oauth1"

//import "fmt"
import "net/http"
import "net/url"
import "io/ioutil"
import "./read"
import "errors"
import "fmt"

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

type tokenOauth struct {
	*oauth1.AccessToken
}

type factualRead struct {
	tokenOauth
	read.Read
}

func NewToken(consumerKey, consumerSecret string) tokenOauth {
	t := tokenOauth{oauth1.NewAccessToken(consumerKey, consumerSecret, "", "")}
	return t
}

func NewRead(t tokenOauth, tableName string) (*factualRead, error) {
	fR := new(factualRead)
	f, err := read.NewReader(tableName)
	if err != nil {
		return nil, err
	}
	fR.tokenOauth = t
	fR.Read = f

	return fR, nil

}

func (F *factualRead) GetReadURL() string {
	return FACTUAL_BASE_URL + F.String()
}

func (F *factualRead) Get() (string, error) {
	if F.Read == nil {
		return "", errors.New("Factual Request has no read data")
	}

	s, err := F.tokenOauth.Get(F.GetReadURL())

	if err != nil {
		return "", err
	}

	return s, nil

}

func (T tokenOauth) Get(url string) (string, error) {
	encodedURL := EncodeURLParameters(url)
	request, err := http.NewRequest("GET", encodedURL, nil)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	T.SignRequestHeader(request)
	result, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer result.Body.Close()

	fmt.Println("Throttle Allocation => ", result.Header.Get("X-Factual-Throttle-Allocation"))
	content, err := ioutil.ReadAll(result.Body)

	if err != nil {
		return "", err
	}

	return string(content), nil

}
