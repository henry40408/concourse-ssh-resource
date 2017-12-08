package models

// Source holds information to connect remote machine via SSH
type Source struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	PrivateKey string `json:"private_key"`
}

// Version is breadcrumb for Concourse CI to choose
// whether to run the pipeline or not. Response from `out` command needs
// it included to comply with Concourse specification
type Version struct {
	Timestamp string `json:"time"`
}

// Params holds script so user can run multiple scripts on the same machine
// in Concourse CI pipeline
type Params struct {
	Interpreter string `json:"interpreter"`
	Script      string `json:"script"`
	Placeholders []Placeholder `json:"placeholders"`
}

type Placeholder struct {
	Name string `json:"name"`
	Value string `json:"value"`
	File string `json:"file"`
}

// Metadata holds metadata from `in` and `out` command.
// Response from `out` command needs it included
// to comply with Concourse specification
type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
