package conf

type Orc struct {
	Source string `yaml:"source"`
	Image  string `yaml:"image""`
	Cmd    string `yaml:"cmd"`
}
