package config

import "github.com/spf13/viper"

type (
	Container struct {
		App             *App
		Token           *Token
		Redis           *Redis
		DB              *DB
		HTTP            *HTTP
		MongoDB         *MongoDB
		MaintenanceMode string
		JWTSecretKey    string
	}

	App struct {
		Name         string
		Env          string
		JWTSecretKey string
	}

	Token struct {
		Duration string
	}

	Redis struct {
		Addr     string
		Password string
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}

	MongoDB struct {
		ConnectionURI string
	}
)

func New() (*Container, error) {
	viper.SetConfigFile("../app.yml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	app := &App{
		Name:         viper.GetString("APP_NAME"),
		Env:          viper.GetString("APP_ENV"),
		JWTSecretKey: viper.GetString("JWT_SECRET_KEY"),
	}

	token := &Token{
		Duration: viper.GetString("TOKEN_DURATION"),
	}

	redis := &Redis{
		Addr:     viper.GetString("REDIS_ADDR"),
		Password: viper.GetString("REDIS_PASSWORD"),
	}

	db := &DB{
		Connection: viper.GetString("DB_CONNECTION"),
		Host:       viper.GetString("DB_HOST"),
		Port:       viper.GetString("DB_PORT"),
		User:       viper.GetString("DB_USER"),
		Password:   viper.GetString("DB_PASSWORD"),
		Name:       viper.GetString("DB_NAME"),
	}

	http := &HTTP{
		Env:            viper.GetString("APP_ENV"),
		URL:            viper.GetString("HTTP_URL"),
		Port:           viper.GetString("HTTP_PORT"),
		AllowedOrigins: viper.GetString("HTTP_ALLOWED_ORIGINS"),
	}

	mongoDB := &MongoDB{
		ConnectionURI: viper.GetString("MONGODB_CONNECTION_URI"),
	}

	maintenanceMode := viper.GetString("MAINTENANCE_MODE")
	return &Container{
		App:             app,
		Token:           token,
		Redis:           redis,
		DB:              db,
		HTTP:            http,
		MongoDB:         mongoDB,
		MaintenanceMode: maintenanceMode,
		JWTSecretKey:    app.JWTSecretKey,
	}, nil
}
