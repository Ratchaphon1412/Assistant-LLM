package configs

type Config struct {
	// Server Configuration
	ServerPort string `env:"SERVER_PORT" envDefault:"8000"`
	CLIENT_URL string `env:"CLIENT_URL"`
	// Database Configuration
	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     string `env:"DB_PORT"`
	DB_USER     string `env:"DB_USER"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_NAME     string `env:"DB_NAME"`
	DB_TIMEZONE string `env:"DB_TIMEZONE"`
	DB_SSL_MODE string `env:"DB_SSL_MODE"`

	// Google OAuth Configuration
	GOOGLE_CLIENT_ID     string `env:"GOOGLE_CLIENT_ID"`
	GOOGLE_CLIENT_SECRET string `env:"GOOGLE_CLIENT_SECRET"`
	GOOGLE_REDIRECT_URL  string `env:"GOOGLE_REDIRECT_URL"`

	// JWT Configuration
	JWT_SECRET      string `env:"JWT_SECRET"`
	JWT_COOKIE_NAME string `env:"JWT_COOKIE_NAME"`
	JWT_SECURE      bool   `env:"JWT_SECURE"`
	JWT_HTTP_ONLY   bool   `env:"JWT_HTTP_ONLY"`

	// Temporal Configuration
	TEMPORAL_HOST          string `env:"TEMPORAL_HOST"`
	TEMPORAL_PORT          string `env:"TEMPORAL_PORT"`
	TEMPORAL_NAMESPACE     string `env:"TEMPORAL_NAMESPACE"`
	TEMPORAL_TASK_QUEUE    string `env:"TEMPORAL_TASK_QUEUE"`
	TEMPORAL_WORKFLOW_NAME string `env:"TEMPORAL_WORKFLOW_NAME"`

	// Redis Configuration
	REDIS_ADDR string `env:"REDIS_ADDR"`

	REDIS_USERNAME  string `env:"REDIS_USERNAME"`
	REDIS_PASSWORD  string `env:"REDIS_PASSWORD"`
	REDIS_DATABASE  int    `env:"REDIS_DATABASE"`
	REDIS_POOLFIFO  bool   `env:"REDIS_POOLFIFO"`
	REDIS_POOL_SIZE int    `env:"REDIS_POOL_SIZE"`

	// OpenWeather
	OPEN_WEATHER_MAP_API_KEY string `env:"OPEN_WEATHER_MAP_API_KEY"`
	OPEN_WEATHER_MAP_URL     string `env:"OPEN_WEATHER_MAP_URL"`
}
