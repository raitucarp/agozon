package agozon

type endpoint struct {
    locale   string
    url      string
}

// create list of endpoint
type endpoints []endpoint

// Select endpoint from list
func (e *endpoints) Select(locale string) string {
    // set default endpoint from default locale
    defaultEndpoint := defaultLocale.url
    // search through endpoints
    for _, endpoint := range *e {
        // if endpoint locale match locale,
        // return endpoint url
        if endpoint.locale == locale {
            return endpoint.url
        }
    }
    // if not match, use default endpoint
    return defaultEndpoint
}

// default locale is US
var defaultLocale = endpoint{"US", "http://webservices.amazon.com/onca/xml"}
var locales = endpoints{
    {"BR", "http://webservices.amazon.com.br/onca/xml"},
    {"CA", "http://webservices.amazon.ca/onca/xml"},
    {"CN", "http://webservices.amazon.cn/onca/xml"},
    {"DE", "http://webservices.amazon.de/onca/xml"},
    {"ES", "http://webservices.amazon.es/onca/xml"},
    {"FR", "http://webservices.amazon.fr/onca/xml"},
    {"IN", "http://webservices.amazon.in/onca/xml"},
    {"IT", "http://webservices.amazon.it/onca/xml"},
    {"JP", "http://webservices.amazon.co.jp/onca/xml"},
    {"UK", "http://webservices.amazon.co.uk/onca/xml"},
    defaultLocale,
}