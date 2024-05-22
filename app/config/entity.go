package config

type Conf struct {
	App struct {
		Name       string `env:"APP_NAME"`
		Name_api   string `env:"APP_NAME_API"`
		Port       string `env:"APP_PORT"`
		Mode       string `env:"APP_MODE"`
		Url        string `env:"APP_URL"`
		Secret_key string `env:"APP_SECRET"`
	}
	MonggoDb_local struct {
		User string `env:"MONGODB_USER_LOCAL"`
		Port string `env:"MONGODB_PORT_LOCAL"`
	}
	MonggoDb_staging struct {
		User string `env:"MONGODB_USER_STAGING"`
		Pass string `env:"MONGODB_PASSWORD_STAGING"`
		Host string `env:"MONGODB_HOST_STAGING"`
	}
	MonggoDb_prod struct {
		User string `env:"MONGODB_USER_PROD"`
		Pass string `env:"MONGODB_PASSWORD_PROD"`
		Host string `env:"MONGODB_HOST_PROD"`
	}
	Db struct {
		Host string `env:"DB_HOST_LOCAL"`
		Name string `env:"DB_NAME_LOCAL"`
		User string `env:"DB_USER_LOCAL"`
		Pass string `env:"DB_PASSWORD_LOCAL"`
		Port string `env:"DB_PORT_LOCAL"`
	}
	Db_develop struct {
		Host string `env:"DB_HOST_DEVELOP"`
		Name string `env:"DB_NAME_DEVELOP"`
		User string `env:"DB_USER_DEVELOP"`
		Pass string `env:"DB_PASSWORD_DEVELOP"`
		Port string `env:"DB_PORT_DEVELOP"`
	}
	Db_staging struct {
		Host string `env:"DB_HOST_STAGING"`
		Name string `env:"DB_NAME_STAGING"`
		User string `env:"DB_USER_STAGING"`
		Pass string `env:"DB_PASSWORD_STAGING"`
		Port string `env:"DB_PORT_STAGING"`
	}
	Db_prod struct {
		Host string `env:"DB_HOST_PROD"`
		Name string `env:"DB_NAME_PROD"`
		User string `env:"DB_USER_PROD"`
		Pass string `env:"DB_PASSWORD_PROD"`
		Port string `env:"DB_PORT_PROD"`
	}
	Rabbitmq_staging struct {
		Host string `env:"RABBIT_MQ_HOST_STAGING"`
		User string `env:"RABBIT_MQ_USER_STAGING"`
		Pass string `env:"RABBIT_MQ_PASSWORD_STAGING"`
		Port string `env:"RABBIT_MQ_PORT_STAGING"`
	}
	Rabbitmq_local struct {
		Host string `env:"RABBIT_MQ_HOST_LOCAL"`
		User string `env:"RABBIT_MQ_USER_LOCAL"`
		Pass string `env:"RABBIT_MQ_PASSWORD_LOCAL"`
		Port string `env:"RABBIT_MQ_PORT_LOCAL"`
	}
	Rabbitmq_prod struct {
		Host string `env:"RABBIT_MQ_HOST_PROD"`
		User string `env:"RABBIT_MQ_USER_PROD"`
		Pass string `env:"RABBIT_MQ_PASSWORD_PROD"`
		Port string `env:"RABBIT_MQ_PORT_PROD"`
	}
	Redis_staging struct {
		RedisAddress  string `env:"REDIS_ADDRESS_STG"`
		RedisPassword string `env:"REDIS_PASSWORD_STG"`
	}
	Redis_prod struct {
		RedisAddress  string `env:"REDIS_ADDRESS_PROD"`
		RedisPassword string `env:"REDIS_PASSWORD_PROD"`
	}
	Firebase_staging struct {
		Type                        string `env:"FIREBASE_TYPE"`
		Project_id                  string `env:"FIREBASE_PROJECT_ID"`
		Private_key_id              string `env:"FIREBASE_PRIVATE_KEY_ID"`
		Private_key                 string `env:"FIREBASE_PRIVATE_KEY"`
		Client_email                string `env:"FIREBASE_CLIENT_EMAIL"`
		Client_id                   string `env:"FIREBASE_CLIENT_ID"`
		Auth_uri                    string `env:"FIREBASE_AUTH_URL"`
		Token_uri                   string `env:"FIREBASE_TOKEN_URL"`
		Auth_provider_x509_cert_url string `env:"FIREBASE_AUTH_PROVIDER_X509_CERT_URL"`
		Client_x509_cert_url        string `env:"FIREBASE_CLIENT_X509_CERT_URL"`
	}
	Firebase_prod struct {
		Type                        string `env:"FIREBASE_TYPE"`
		Project_id                  string `env:"FIREBASE_PROJECT_ID"`
		Private_key_id              string `env:"FIREBASE_PRIVATE_KEY_ID"`
		Private_key                 string `env:"FIREBASE_PRIVATE_KEY"`
		Client_email                string `env:"FIREBASE_CLIENT_EMAIL"`
		Client_id                   string `env:"FIREBASE_CLIENT_ID"`
		Auth_uri                    string `env:"FIREBASE_AUTH_URL"`
		Token_uri                   string `env:"FIREBASE_TOKEN_URL"`
		Auth_provider_x509_cert_url string `env:"FIREBASE_AUTH_PROVIDER_X509_CERT_URL"`
		Client_x509_cert_url        string `env:"FIREBASE_CLIENT_X509_CERT_URL"`
	}
	AuthServer_local struct {
		Url string `env:"AUTH_SERVER_LOCAL"`
	}
	AuthServer_develop struct {
		Url string `env:"AUTH_SERVER_DEVELOP"`
	}
	AuthServer_staging struct {
		Url string `env:"AUTH_SERVER_STAGING"`
	}
	AuthServer_prod struct {
		Url string `env:"AUTH_SERVER_PRODUCTION"`
	}
	Sentry struct {
		Dsn                    string  `env:"SENTRY_DSN"`
		TracesSampleRate       float64 `env:"SENTRY_TRACE_SAMPLE_RATE"`
		ProfileTraceSampleRate float64 `env:"SENTRY_PROFILE_TRACE_SAMPLE_RATE"`
	}
}
