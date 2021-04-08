package control

import (
	"encoding/json"
	"fmt"
	"github.com/turbot/steampipe/constants"
)

func FormatControl(controls []ControlRun, format string) []byte {
	if format == constants.ArgJSON {
		return formatJSON(controls)
	}
	return nil

}

func formatJSON(controls []ControlRun) []byte {
	res, _ := json.MarshalIndent(controls, "", "  ")
	fmt.Println(string(res))
	return res
}
