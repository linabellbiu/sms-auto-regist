package conf

type DefaultCollectConfig struct {
	Cron string `yaml:"cron"`
}

type GlobalConfig struct {
	CollectSourceHtml CollectSourceHtml `yaml:"collect_source_html"`
}
