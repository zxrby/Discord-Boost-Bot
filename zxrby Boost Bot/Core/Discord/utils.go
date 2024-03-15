package Discord

import (
	"BoostTool/Core/Utils"
	"time"
)

type BoostResponse struct {
	Success       int
	Failed        int
	SuccessTokens []string
	FailedTokens  []string
	ElapsedTime   time.Duration
}

type CaptchaStruct struct {
	Success         bool `json:"success"`
	CaptchaServices struct {
		Hcoptcha struct {
			State   int  `json:"state"`
			Working bool `json:"working"`
		} `json:"hcoptcha"`
		Capsolver struct {
			State   int  `json:"state"`
			Working bool `json:"working"`
		} `json:"capsolver"`
	} `json:"captchaServices"`
}

type FingerprintResponse struct {
	Fingerprint string    `json:"fingerprint"`
	Assignments [][]int64 `json:"assignments"`
}

func ReloadFiles(path string) *Utils.Cycle {

	Tokens, _ := Utils.NewFromFile(path)
	return Tokens
}
