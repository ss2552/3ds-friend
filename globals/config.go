package globals

type config struct {
	PostgresURI              string
	PostgresMaxConnections   int64
	AESKey                   string
	AuthenticationServerPort uint16
	SecureServerHost         string
	SecureServerPort         uint16
	HealthCheckPort          uint16 `envconf:"optional"`
	EnableBella              bool   `envconf:"optional"`
}

var Config *config = &config{}
