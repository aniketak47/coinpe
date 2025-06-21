package sendgrid

func Init(apiKey string) *client {
	client := &client{
		baseURL: "https://api.sendgrid.com",
		apiKey:  apiKey,
	}
	return client
}
