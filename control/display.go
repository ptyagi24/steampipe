package control

import (
	"bytes"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/turbot/steampipe/constants"
	"github.com/turbot/steampipe/display"
)

func getControlStatusText(status string) string {
	switch status {
	case ControlAlarm:
		return text.FgRed.Sprint(status)
	case ControlOK:
		return text.FgGreen.Sprint(status)
	case ControlInfo:
		return text.FgHiBlue.Sprint(status)
	case ControlSkipped:
		fallthrough
	default:
		return status
	}
}

func displayControlsTable(controlPack []ControlRun) {
	// the buffer to put the output data in
	outbuf := bytes.NewBufferString("")

	// the table
	t := table.NewWriter()
	t.SetOutputMirror(outbuf)
	t.SetStyle(table.StyleDefault)
	t.Style().Format.Header = text.FormatDefault

	headers := table.Row{
		"Status",
		"ID",
		"Name",
		"Description",
	}
	colConfigs := []table.ColumnConfig{
		{
			Name:     "Status",
			Number:   1,
			WidthMax: constants.MaxColumnWidth,
		},
		{
			Name:     "ID",
			Number:   2,
			WidthMax: constants.MaxColumnWidth,
		},
		{
			Name:     "Name",
			Number:   3,
			WidthMax: constants.MaxColumnWidth,
		},
		{
			Name:     "Description",
			Number:   4,
			WidthMax: constants.MaxColumnWidth,
		},
	}

	t.SetColumnConfigs(colConfigs)
	t.AppendHeader(headers)

	for _, control := range controlPack {
		for _, result := range control.Results {
			row := table.Row{
				getControlStatusText(result.Status),
				control.Type.ControlID,
				control.Type.Title,
				control.Type.Description,
			}
			t.AppendRow(row)
		}
	}

	t.Render()
	fmt.Println("Control details")
	fmt.Println("")
	display.ShowPaged(outbuf.String())
	fmt.Println("")
}

func displayControlStatusesTable(controlPack []ControlRun) {
	// the buffer to put the output data in
	outbuf := bytes.NewBufferString("")

	// the table
	t := table.NewWriter()
	t.SetOutputMirror(outbuf)
	t.SetStyle(table.StyleDefault)
	t.Style().Format.Header = text.FormatDefault

	headers := table.Row{
		"Status",
		"Total",
	}
	colConfigs := []table.ColumnConfig{
		{
			Name:     "Status",
			Number:   1,
			WidthMax: constants.MaxColumnWidth,
		},
		{
			Name:     "Total",
			Number:   2,
			WidthMax: constants.MaxColumnWidth,
		},
	}

	t.SetColumnConfigs(colConfigs)
	t.AppendHeader(headers)

	alarmTotal := 0
	okTotal := 0
	infoTotal := 0
	skippedTotal := 0
	totalControls := 0

	for _, control := range controlPack {
		for _, result := range control.Results {
			totalControls += 1
			switch result.Status {
			case ControlAlarm:
				alarmTotal += 1
			case ControlInfo:
				infoTotal += 1
			case ControlOK:
				okTotal += 1
			case ControlSkipped:
				skippedTotal += 1
			}
		}
	}

	alarmRow := table.Row{
		getControlStatusText(ControlAlarm),
		alarmTotal,
	}
	okRow := table.Row{
		getControlStatusText(ControlOK),
		okTotal,
	}
	infoRow := table.Row{
		getControlStatusText(ControlInfo),
		infoTotal,
	}
	skippedRow := table.Row{
		getControlStatusText(ControlSkipped),
		skippedTotal,
	}
	t.AppendRow(alarmRow)
	t.AppendRow(okRow)
	t.AppendRow(infoRow)
	t.AppendRow(skippedRow)

	t.Render()
	fmt.Println(fmt.Sprintf("%d %s", totalControls, getPluralisedControlsText(totalControls)))
	fmt.Println("")
	display.ShowPaged(outbuf.String())
	fmt.Println("")
}
