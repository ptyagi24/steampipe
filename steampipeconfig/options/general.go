package options

import "github.com/turbot/steampipe/constants"

// General
type General struct {
	LogLevel    *string `hcl:"log_level"`
	UpdateCheck *string `hcl:"update_check"`
}

// ConfigMap :: create a config map to pass to viper
func (c General) ConfigMap() map[string]interface{} {
	// only add keys which are non null
	res := map[string]interface{}{}
	if c.LogLevel != nil {
		res[constants.ArgLogLevel] = c.LogLevel
	}
	if c.UpdateCheck != nil {
		res[constants.ArgUpdateCheck] = c.UpdateCheck
	}

	return res
}
