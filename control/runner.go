package control

type ControlType struct {
	ControlID   string `json:"control_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ControlResult struct {
	Error    error
	Status   string `json:"status"`
	Reason   string `json:"reason"`
	Resource string `json:"resource"`
}

type ControlRun struct {
	Type    ControlType
	Results []ControlResult
}

const (
	ControlAlarm   = "alarm"
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

func RunControl(output string, outputDir string) {
	//fmt.Println("format", format)
	//fmt.Println("output-dir", outputDir)

	controlPack := []ControlRun{
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
	}

	displayControlsTable(controlPack)
	displayControlStatusesTable(controlPack)

	//totalAlarm := 1
	//totalOK := 2
	//
	//fmt.Println("Control results")
	//fmt.Println("")
	//fmt.Println(text.FgRed.Sprintf("%d %s in alarm", totalAlarm, getPluralisedControlsText(totalAlarm)))
	//fmt.Println(text.FgGreen.Sprintf("%d %s in OK", totalOK, getPluralisedControlsText(totalOK)))

	//fmt.Println(controls)

	//formattedOutput := FormatControl(controls, format)
	//fmt.Println(formattedOutput)
}
