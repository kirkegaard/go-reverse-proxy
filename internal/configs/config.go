package configs

type resource struct {
	Name        string
	Endpoint    string
	Destination string
}

type configuration struct {
	Server struct {
		Host string
		Port string
	}
	Resources []resource
}

func NewConfiguration() (*configuration, error) {
	Config := &configuration{}
	Config.Server.Host = "localhost"
	Config.Server.Port = "6969"
	Config.Resources = []resource{
		{Name: "api", Endpoint: "/api", Destination: "http://localhost:8080/api"},
	}

	return Config, nil
}
