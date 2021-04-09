package control

import (
	"github.com/turbot/steampipe/utils"
	"sync"
	"time"
)

type ControlReportingOptions struct {
	OutputDirectory string
	OutputFormats   []string
	WithColor       bool
	WithProgress    bool
}

type ControlType struct {
	ControlID   string `json:"control_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ControlResult struct {
	Status   string `json:"status"`
	Reason   string `json:"reason"`
	Resource string `json:"resource"`
}

type ControlRun struct {
	Type    ControlType     `json:"type"`
	Results []ControlResult `json:"results"`
}

type ControlPack struct {
	Timestamp   time.Time    `json:"timestamp"`
	ControlRuns []ControlRun `json:"control_runs"`
}

const (
	ControlAlarm   = "alarm"
	ControlError   = "error"
	ControlInfo    = "info"
	ControlOK      = "ok"
	ControlSkipped = "skipped"
)

func getPluralisedControlsText(count int) string {
	if count == 1 {
		return "control"
	}
	return "controls"
}

func RunControl(reportingOptions ControlReportingOptions) {
	spinner := utils.ShowSpinner("Running controls")
	// TODO how do I actually get this? Simulate it coming from a channel for now
	controlChan := make(chan ControlPack)
	go runControls(controlChan)
	controlPack := <-controlChan

	utils.StopSpinner(spinner)

	// Generate output formats in parallel
	var wg sync.WaitGroup

	// Wait for each formatter to complete
	for _, outputFormat := range reportingOptions.OutputFormats {
		wg.Add(1)
		go outputResults(controlPack, outputFormat, reportingOptions.OutputDirectory, &wg)
	}

	// Wait for CLI output to complete
	wg.Add(1)
	go displayControlResults(controlPack, reportingOptions, &wg)

	wg.Wait()
}

func runControls(stream chan ControlPack) {
	controlPack := ControlPack{
		Timestamp: time.Now(),
		ControlRuns: []ControlRun{
			{Type: ControlType{
				ControlID:   "aws.cis.v130.1.1",
				Title:       "Maintain current contact details",
				Description: "Amazon S3 provides a variety of no, or low, cost encryption options to protect data at rest.",
			}, Results: []ControlResult{
				{
					Status: ControlInfo,
				},
			}}, {Type: ControlType{
				ControlID:   "aws.cis.v130.2.1.1",
				Title:       "Ensure all S3 buckets employ encryption-at-rest",
				Description: "Ensure contact email and telephone details for AWS accounts are current and map to more than one individual in your organization. An AWS account supports a number of contact details, and AWS will use these to contact the account owner if activity judged to be in breach of Acceptable Use Policy or indicative of likely security compromise is observed by the AWS Abuse team. Contact details should not be for a single individual, as circumstances may arise where that individual is unavailable. Email contact details should point to a mail alias which forwards email to multiple individuals within the organization; where feasible, phone contact details should point to a PABX hunt group or other call-forwarding system.",
			}, Results: []ControlResult{
				{
					Status: ControlOK,
				},
			}}, {Type: ControlType{
				ControlID:   "aws.cis.v130.2.1.2",
				Title:       "Ensure S3 Bucket Policy allows HTTPS requests",
				Description: "At the Amazon S3 bucket level, you can configure permissions through a bucket policy making the objects accessible only through HTTPS.",
			}, Results: []ControlResult{
				{
					Status: ControlAlarm,
				},
			}}, {Type: ControlType{
				ControlID:   "aws.cis.v130.2.2.1",
				Title:       "Ensure EBS volume encryption is enabled",
				Description: "Elastic Compute Cloud (EC2) supports encryption at rest when using the Elastic Block Store (EBS) service. While disabled by default, forcing encryption at EBS volume creation is supported.",
			}, Results: []ControlResult{
				{
					Status: ControlOK,
				},
			}},
		},
	}
	stream <- controlPack
}
