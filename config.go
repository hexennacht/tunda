package tunda

type configuration struct {
	RedisAddress      string `envconfig:"REDIS_ADDRESS" required:"true" default:"localhost:6379"`
	RedisDB           int    `envconfig:"REDIS_DB" default:"0"`
	RedisUsername     string `envconfig:"REDIS_USERNAME" default:""`
	RedisPassword     string `envconfig:"REDIS_PASSWORD" default:""`
	RedisWriteTimeout int    `envconfig:"REDIS_WRITE_TIMEOUT" default:"5"`
	RedisTimeout      int    `envconfig:"REDIS_TIMEOUT" default:"5"`
	RedisKeepAlive    int    `envconfig:"REDIS_KEEP_ALIVE" default:"5"`

	DefaultSpawnWorker int `envconfig:"DEFAULT_SPAWN_WORKER" deafault:"20"`
}
