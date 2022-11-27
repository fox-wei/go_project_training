package conf

type DBConf struct {
	Driver string

	Host string //address

	Port string

	User string

	Passwrod string

	DbName string //*database name

	Charset string
}
