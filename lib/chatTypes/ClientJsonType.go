package chatTypes

// ClientJsonType is used for json view of Client struct
type clientJsonType struct {
	Address `json:"address"`
	Name    string `json:"name"`
	Message `json:"message"`
}

// Address
type Address struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// Message
type Message struct {
	Timestamp   string `json:"timestamp"`
	MessageText string `json:"messageText"`
}
