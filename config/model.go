package config

type Config struct {
	Database DB  `yaml:"database"`
	Redis    Rdb `yaml:"redis"`
}

type DB struct {
	Name    string `yaml:"name"`
	Migrate bool   `yaml:"migrate"`
}
type Rdb struct {
	Name string `yaml:"name"`
}
