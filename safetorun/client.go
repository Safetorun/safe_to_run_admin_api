package safetorun

type Client struct {
	AuthToken string
}

func New(authToken string) Client {
	return Client{
		authToken,
	}
}
