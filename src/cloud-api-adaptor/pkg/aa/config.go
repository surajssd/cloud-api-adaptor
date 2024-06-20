package aa

import (
	"fmt"
	"strings"

	toml "github.com/pelletier/go-toml/v2"
)

const (
	DefaultAaConfigPath = "/run/peerpod/aa.toml"
)

type AAConfig struct {
	TokenCfg struct {
		CocoAs struct {
			URL string `toml:"url"`
		} `toml:"coco_as"`
		Kbs struct {
			URL string `toml:"url"`
		} `toml:"kbs"`
	} `toml:"token_configs"`
}

func parseAAKBCParams(aaKBCParams string) (string, error) {
	parts := strings.SplitN(aaKBCParams, "::", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid aa-kbs-params input: %s", aaKBCParams)
	}
	_, url := parts[0], parts[1]
	return url, nil
}

func CreateConfigFile(aaKBCParams string) (string, error) {
	url, err := parseAAKBCParams(aaKBCParams)
	if err != nil {
		return "", err
	}

	config := AAConfig{}
	config.TokenCfg.CocoAs.URL = ""
	config.TokenCfg.Kbs.URL = url

	bytes, err := toml.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
