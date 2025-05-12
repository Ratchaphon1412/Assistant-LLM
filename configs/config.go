package configs

type Config struct {
	// Server Configuration
	ServerPort string `env:"SERVER_PORT" envDefault:"8000"`
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
}
