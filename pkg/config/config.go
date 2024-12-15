package config

const ()

type Schema struct {
	Environment   string `env:"environment"`
	HttpPort      int    `env:"http_port"`
	GrpcPort      int    `env:"grpc_port"`
	AuthSecret    string `env:"auth_secret"`
	DatabaseURI   string `env:"database_uri"`
	RedisURI      string `env:"redis_uri"`
	RedisPassword string `env:"redis_password"`
	RedisDB       int    `env:"redis_db"`
}

var (
	cfg Schema
)

func LoadConfig() *Schema {
	// TODO parsing yaml files and flags
	return &cfg
}

func GetConfig() *Schema {
	return &cfg
}
