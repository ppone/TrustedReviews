package filters

func Blank(keyword string, b bool) string {
	if b == true {
		return keyword + ":" + "{\"$blank\":true}"
	} else {
		return keyword + ":" + "{\"$blank\":false}"
	}
}

func BeginsWith(keyword string, value string) string {

	return "{" + "\"" + keyword + ":" + "{\"bw\":\"" + value + "\"}}"

}

func BeginsWithAny(keyword string, values ...string) string {

	return "{" + "\"" + keyword + ":" + "{\"$blank\":true}" + "}"

}
