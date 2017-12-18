package chatTypes

// ClientJSONType is used for json view of Client struct
type clientJSONType struct {
	Address `json:"address"`
	Name    string `json:"name"`
	Message `json:"message"`
}

// Address nested struct for ClientJsonType
type Address struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

// Message nested struct for ClientJsonType
type Message struct {
	Timestamp   string `json:"timestamp"`
	MessageText string `json:"messageText"`
}
