package mqtt

type ConnectionOptions struct {
	Username string
	Password string
	Host     string
	Port     string
	ClientID string
}

type Options struct {
	Connection ConnectionOptions
	Topic	   string
}
