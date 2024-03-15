package Captcha

import (
	"BoostTool/Core/Utils"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)



// Define Hcoptcha function
func Hcoptcha(apikey, host, auth, webkey, rqdata, token1 string) string {
	client := http.Client{Timeout: 30 * time.Second}
	proxy := auth + "@" + host

	payload := map[string]interface{}{
		"task_type": "hcaptchaEnterprise",
		"api_key":   apikey,
		"data": map[string]string{
			"sitekey": webkey,
			"url":     "https://discord.com",
			"proxy":   proxy,
			"rqdata":  rqdata,
		},
	}
	jsonPayload, _ := json.Marshal(payload)

	req1, _ := http.NewRequest("POST", "https://api.hcoptcha.online/api/createTask", bytes.NewBuffer(jsonPayload))
	req1.Header.Set("Content-Type", "application/json")

	resp1, err := client.Do(req1)
	if err != nil {
		Utils.LogError("Error Captcha failed", "Error", err.Error())
		return ""
	}
	bodytext1, _ := io.ReadAll(resp1.Body)

	defer resp1.Body.Close()
	var hcopresp hcopResponse

	err = json.Unmarshal(bodytext1, &hcopresp)
	if err != nil {
		Utils.LogError("Error Captcha failed", "Error", err.Error())
		return ""
	}

	if hcopresp.TaskID != "" {
		Utils.LogInfo("Captcha Task", "Task ID", hcopresp.TaskID)
		for i := 0; i < 10; i++ {
			json2 := hcoptPay{
				APIKey: apikey,
				TaskID: hcopresp.TaskID,
			}
			p2, _ := json.Marshal(json2)

			req2, _ := http.NewRequest("POST", "https://api.hcoptcha.online/api/getTaskData", bytes.NewReader(p2))
			req2.Header.Set("Content-Type", "application/json")

			resp2, err := client.Do(req2)
			if err != nil {
				Utils.LogError("Error Captcha failed", "Error", err.Error())
				return ""
			}

			defer resp2.Body.Close()
			var hcoptresp hcopSol
			bodytext2, _ := io.ReadAll(resp2.Body)

			err = json.Unmarshal(bodytext2, &hcoptresp)
			if err != nil {
				Utils.LogError("Error Captcha failed", "Error", err.Error())
				return ""
			}

			if hcoptresp.Task.State != "completed" && hcoptresp.Task.State != "processing" {
				Utils.LogError("Couldn't Solved Captcha, Retrying", "Token", token1)
				return ""
			} else if hcoptresp.Task.State == "completed" {
				return hcoptresp.Task.CaptchaKey
			} else {
				time.Sleep(time.Second * 2)
				continue
			}
		}
	} else {
		Utils.LogError("Couldn't Get Captcha Task ID, Check API Key or Contact Support", "Error", string(bodytext1))
	}

	return ""
}
