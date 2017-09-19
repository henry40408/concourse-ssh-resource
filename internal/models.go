package internal

type Source struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	PrivateKey string `json:"private_key"`
}

type Version struct{}

type Params struct {
	Script string `json:"script"`
}

type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
