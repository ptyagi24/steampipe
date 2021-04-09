package control

import (
	"encoding/json"

	"github.com/turbot/steampipe/constants"
)

func formatResults(controlPack ControlPack, format string) []byte {
	switch format {
	case constants.ArgJSON:
		return formatJSON(controlPack)
	default:
		return nil
	}
}

func formatJSON(controlPack ControlPack) []byte {
	res, _ := json.MarshalIndent(controlPack, "", "  ")
	return res
}
