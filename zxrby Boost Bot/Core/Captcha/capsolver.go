package Captcha

import (
	"BoostTool/Core/Utils"
	capsolver_go "github.com/capsolver/capsolver-go"
)

func Captcha(apikey, host, auth, webkey, rqdata, token1 string) string {
	capSolver := capsolver_go.CapSolver{apikey}
	s, err := capSolver.Solve(
		map[string]interface{}{
			"type":        "HCaptchaTurboTask",
			"websiteURL":  "https://discord.com/",
			"websiteKey":  webkey,
			"isInvisible": false,
			"enterprisePayload": map[string]interface{}{
				"rqdata": rqdata,
			},
			"proxy":     "http:" + host + ":" + auth,
			"userAgent": "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9017 Chrome/108.0.5359.215 Electron/22.3.12 Safari/537.36",
		})

	if err != nil {
		Utils.LogError("Error Captcha failed", "Error", err.Error())
		return ""
	}

	Utils.LogInfo("Captcha Task", "ID", s.TaskId)

	return s.Solution.GRecaptchaResponse
}
