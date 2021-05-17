package ptcpayclient

//***********************************************
//* Copyright (c) 2021 Ulbora Labs LLC
//* Copyright (c) 2021 Ken Williamson
//***********************************************

const (
	userAgent = "Ulbora btyPayClient"
)

// go mod init github.com/Ulbora/BTCPayClient

//Headers Headers
type Headers struct {
	headers map[string]string
	//mu      sync.Mutex
}

//Set Set
func (h *Headers) Set(key string, value string) {
	//h.mu.Lock()
	//defer h.mu.Unlock()
	if h.headers == nil {
		h.headers = make(map[string]string)
	}
	h.headers[key] = value
}

//Client Client
type Client interface {
	//GetClientID() string
	//Token(req *TokenRequest) *TokenResponse
	//PairClient(code string) *TokenResponse
	//GetPairingCodeRequest(code string) string
	//GetRates(currencyPairs []string, storeID string) *RateResponse
	//CreateInvoice(inv *InvoiceReq) *InvoiceResponse
	//GetInvoice(invoiceID string) *InvoiceResponse
	//GetInvoices(args *InvoiceArgs) *InvoiceListResponse
}

// go mod init github.com/Ulbora/BTCPayClientGreenfieldV1
