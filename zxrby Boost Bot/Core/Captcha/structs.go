package Captcha

type hcopPayload struct {
	APIKey   string `json:"api_key"`
	TaskType string `json:"task_type"`
	Data     struct {
		Rqdata    string `json:"rqdata"`
		Useragent string `json:"useragent"`
		Sitekey   string `json:"sitekey"`
		Proxy     string `json:"proxy"`
		Host      string `json:"host"`
	} `json:"data"`
}

type hcopResponse struct {
	Error  bool   `json:"error"`
	TaskID string `json:"task_id"`
}

type hcoptPay struct {
	APIKey string `json:"api_key"`
	TaskID string `json:"task_id"`
}

type hcopSol struct {
	Error bool `json:"error"`
	Task  struct {
		CaptchaKey string `json:"captcha_key"`
		Refunded   bool   `json:"refunded"`
		State      string `json:"state"`
	} `json:"task"`
}

type capmonsterTaskID struct {
	ErrorID int `json:"errorId"`
	TaskID  int `json:"taskId"`
}

type capmonsterGetTask struct {
	ErrorID  int    `json:"errorId"`
	Status   string `json:"status"`
	Solution struct {
		GRecaptchaResponse string `json:"gRecaptchaResponse"`
		RespKey            string `json:"respKey"`
		UserAgent          string `json:"userAgent"`
	} `json:"solution"`
}
