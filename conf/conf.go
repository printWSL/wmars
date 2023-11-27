package conf

type ConfigType struct {
	DB      DBConfig
	Ansible AnsibleConfig
}

type DBConfig struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
}

type AnsibleConfig struct {
}
