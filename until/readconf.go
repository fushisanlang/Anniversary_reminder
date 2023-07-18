package until

import (
	"runtime"

	"github.com/larspensjo/config"
)

var TOPIC = make(map[string]string)

func ReadConf(parameter string) string {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cfg, _ := config.ReadDefault("conf/conf.ini")

	if cfg.HasSection("Anniversary_reminder") {
		section, err := cfg.SectionOptions("Anniversary_reminder")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("Anniversary_reminder", v)
				if err == nil {
					TOPIC[v] = options
				}
			}
		}
	}
	return TOPIC[parameter]
}
