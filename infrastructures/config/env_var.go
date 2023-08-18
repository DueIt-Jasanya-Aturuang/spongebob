package config

var (
	AppPort      string
	AppStatus    string
	AppGracefull string

	PgHost string
	PgPort string
	PgUser string
	PgPass string
	PgName string
	PgSSL  string

	RedisHost string
	RedisPort string
	RedisDb   int
	RedisPass string

	MinIoAccessKey string
	MinIoSecretKey string
	MinIoEndpoint  string
	MinIoPort      string
	MinIoBucket    string
	MinIoSSL       string

	DefaultImage string
	AesCFB       string
	AesCBC       string
	AesCBCIV     string
)
