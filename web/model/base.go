package model

type App struct {
	*Server `yaml:"server"`
	*Db     `yaml:"db"`
	*Log    `yaml:"log"`
}

type Server struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type Db struct {
	Mongodb Mongodb `yaml:"mongodb"`
}

type Mongodb struct {
	Db   int    `yaml:"db"`
	Addr string `yaml:"addr"`
}

type Log struct {
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
	Path       string `yaml:"path"`
}
