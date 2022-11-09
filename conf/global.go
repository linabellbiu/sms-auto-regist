package conf

type DefaultCollectConfig struct {
	Cron         string   `yaml:"cron"`
	Host         string   `yaml:"host"`
	CompileRegex string   `yaml:"compile_regex"`
	Keywords     []string `yaml:"keywords"`
}

type GlobalConfig struct {
	CollectSourceHtml CollectSourceHtml `yaml:"collect_source_html"`
	Orc               Orc               `yaml:"orc"`
}

var Global = &GlobalConfig{}
