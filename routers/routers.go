package routers

// Available modula como a requisição deve ser instanciada
type Available struct {
	Url    string
	Method string
}

var UrlAvailable = Available{
	Url:    "https://brapi.dev/api/available",
	Method: "GET",
}

// Currency modula como a requisição deve ser instanciada
type Currency struct {
	UrlUsdToBrl string
	UrlEurToBrl string
	Method      string
}

var UrlCurrency = Currency{
	UrlUsdToBrl: "https://brapi.dev/api/v2/currency?currency=USD-BRL",
	UrlEurToBrl: "https://brapi.dev/api/v2/currency?currency=EUR-BRL",
	Method:      "GET",
}
