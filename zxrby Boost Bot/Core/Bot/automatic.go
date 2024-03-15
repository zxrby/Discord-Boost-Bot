package Bot

import (
	"BoostTool/Core/Discord"
	"BoostTool/Core/Utils"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Automation() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	if config.SellixAutomationSettings.Enabled {
		go EditSellixStock()
		r.POST("/sellix", handleSellixWebhook)
	} else if config.SellpassAutomationSettings.Enabled {
		r.POST("/sellpass", handleSellpassWebhook)
	} else if config.SellappAutomationSettings.Enabled {
		go EditSellappStock()
		r.POST("/sellapp", handleSellappWebhook)
	}
	// Start the server
	Utils.LogInfo("Listening and Serving HTTP Port", "Port", config.Port)
	r.Run(":" + config.Port)

}

func handleSellixWebhook(c *gin.Context) {

	var file string
	var invite string
	var payload map[string]interface{}

	_ = c.ShouldBindJSON(&payload)
	c.Status(200)

	if payload["event"] == "order:paid" {
		d := payload["data"].(map[string]interface{})
		if d["product_id"] == config.SellixAutomationSettings.ProductSettings.ThreeMonthProductID || d["product_id"] == config.SellixAutomationSettings.ProductSettings.OneMonthProductID {
			custom := d["custom_fields"].(map[string]interface{})
			invitelink := custom[config.InviteFieldName].(string)
			email := d["customer_email"].(string)
			title := d["product_title"].(string)
			orderID := d["uniqid"].(string)
			amount := d["quantity"].(float64)
			gateway := d["gateway"].(string)

			inviteParts := strings.Split(invitelink, "/")
			count := len(inviteParts)

			if count == 4 {
				invite = inviteParts[3]
			} else if count == 2 {
				invite = inviteParts[1]
			} else {
				invite = invitelink
			}

			if strings.Contains(title, "1") {
				file = "1 Month Tokens.txt"
			} else if strings.Contains(title, "3") {
				file = "3 Month Tokens.txt"
			} else {
				Utils.LogError("Failed to Auto Boost, Improper Duration Specified", "", "")
			}

			//c.String(http.StatusOK, fmt.Sprintf("We have successfully started boosting your server! If any problems occur please join our discord server: %v", config.DiscordSettings.SupportServer))
			if invite == "" || amount == 0 || file == "" {
				Utils.LogError("Missing Fields for Boosting, Check Automation Settings on Website or Config!", "", "")
				return
			}

			Utils.ClearScreen()
			Utils.PrintASCII()

			respo, err := Discord.BoostServer(invite, int(amount)*2, file)
			if err != nil {
				Utils.LogError(err.Error(), "", "")
				return
			}

			var success string
			if len(respo.SuccessTokens) != 0 {
				success = strings.Join(respo.SuccessTokens, "\n")
			} else {
				success = "No succeeded tokens."
			}

			var failed string
			if len(respo.FailedTokens) != 0 {
				failed = strings.Join(respo.FailedTokens, "\n")
			} else {
				failed = "No failed tokens."
			}

			embed := discordgo.MessageEmbed{
				Title:       "Sellix Order",
				Description: fmt.Sprintf("We have boosted a server successfully.\n**Success**: %v | **Failed**: %v\n**Elapsed Time:** %v", respo.Success, respo.Failed, respo.ElapsedTime),
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Order Information",
						Value:  fmt.Sprintf("```\nOrder ID: %v\nProduct Title: %v\nBoost Amount: %v\nServer Invite: %v\n```", orderID, title, amount*2, invite),
						Inline: false,
					},
					{
						Name:   "Customer Information",
						Value:  fmt.Sprintf("```\nEmail: %v\nPayment Method: %v\n```", email, gateway),
						Inline: false,
					},
					{
						Name:   "Succeeded Tokens",
						Value:  fmt.Sprintf("```\n%v\n```", success),
						Inline: false,
					},
					{
						Name:   "Failed Tokens",
						Value:  fmt.Sprintf("```\n%v\n```", failed),
						Inline: false,
					},
				},
				Color: int(EmbedColor),
			}

			s.ChannelMessageSendEmbed(config.DiscordSettings.LogsChannel, &embed)

			return
		}

	}

	return

}

func handleSellappWebhook(c *gin.Context) {
	c.Status(200)
	client := http.Client{Timeout: time.Second * 30}
	headers := http.Header{
		"content-type":  {"application/json"},
		"accept":        {"text/json"},
		"authorization": {"Bearer " + config.SellappAutomationSettings.APIKey},
	}

	var amount int
	var file string
	var invite string
	var invitelink string

	threemonthid, _ := strconv.Atoi(config.SellappAutomationSettings.ProductSettings.ThreeMonthProductID)
	onemonthid, _ := strconv.Atoi(config.SellappAutomationSettings.ProductSettings.OneMonthProductID)

	// Parse the Sellapp webhook payload
	var payload SellappOrderCompleted
	_ = c.ShouldBindJSON(&payload)

	if payload.Event == "order.completed" {
		var SellappProductIn SellappOrderInfo
		uniqID := payload.Data.ID
		req, _ := http.NewRequest("GET", fmt.Sprintf("https://sell.app/api/v1/invoices/%v", uniqID), nil)
		req.Header = headers
		resp1, _ := client.Do(req)

		defer resp1.Body.Close()

		body, _ := io.ReadAll(resp1.Body)
		_ = json.Unmarshal(body, &SellappProductIn)

		for _, productinfo := range SellappProductIn.Data.Products {
			if productinfo.ID == threemonthid || productinfo.ID == onemonthid {
				for _, variantsinfo := range productinfo.Variants {
					amount = variantsinfo.Quantity * 2
					for _, additionalinfo := range variantsinfo.AdditionalInformation {
						if additionalinfo.Label == config.InviteFieldName {
							invitelink = additionalinfo.Value
							break
						}

					}
				}

				title := productinfo.Title
				method := SellappProductIn.Data.Payment.Gateway.Type
				email := SellappProductIn.Data.Payment.Gateway.Data.CustomerEmail

				if strings.Contains(title, "1") {
					file = "1 Month Tokens.txt"
				} else if strings.Contains(title, "3") {
					file = "3 Month Tokens.txt"
				} else {
					Utils.LogError("Failed to Auto Boost, Improper Duration Specified", "", "")
				}

				inviteParts := strings.Split(invitelink, "/")
				count := len(inviteParts)

				if count == 4 {
					invite = inviteParts[3]
				} else if count == 2 {
					invite = inviteParts[1]
				} else {
					invite = invitelink
				}

				//c.String(http.StatusOK, fmt.Sprintf("We have successfully started boosting your server! If any problems occur please join our discord server: %v", config.DiscordSettings.SupportServer))

				if invite == "" || amount == 0 || file == "" {
					Utils.LogError("Missing Fields for Boosting, Check Automation Settings on Website or Config!", "", "")
					return
				}

				Utils.ClearScreen()
				Utils.PrintASCII()

				respo, err := Discord.BoostServer(invite, amount, file)
				if err != nil {
					Utils.LogError(err.Error(), "", "")
					return
				}

				var success string
				if len(respo.SuccessTokens) != 0 {
					success = strings.Join(respo.SuccessTokens, "\n")
				} else {
					success = "No succeeded tokens."
				}

				var failed string
				if len(respo.FailedTokens) != 0 {
					failed = strings.Join(respo.FailedTokens, "\n")
				} else {
					failed = "No failed tokens."
				}

				//descrip := fmt.Sprintf("**Order ID:** %v\n**Email:** %v\n**Product Title:** %v\n**Amount of Boosts:** %v\n**Duration of Tokens:** %v\n**Server Invite:** [%v](https://discord.gg/%v)", uniqID, email, title, amount, durationbeforeparse, invite, invite)
				embed := discordgo.MessageEmbed{
					Title:       "Sellapp Order",
					Description: fmt.Sprintf("We have boosted a server successfully.\n**Success**: %v | **Failed**: %v\n**Elapsed Time:** %v", respo.Success, respo.Failed, respo.ElapsedTime),
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "Order Information",
							Value:  fmt.Sprintf("```\nOrder ID: %v\nProduct Title: %v\nBoost Amount: %v\nServer Invite: %v\n```", uniqID, title, amount, invite),
							Inline: false,
						},
						{
							Name:   "Customer Information",
							Value:  fmt.Sprintf("```\nEmail: %v\nPayment Method: %v\n```", email, method),
							Inline: false,
						},
						{
							Name:   "Succeeded Tokens",
							Value:  fmt.Sprintf("```\n%v\n```", success),
							Inline: false,
						},
						{
							Name:   "Failed Tokens",
							Value:  fmt.Sprintf("```\n%v\n```", failed),
							Inline: false,
						},
					},
					Color: int(EmbedColor),
				}

				s.ChannelMessageSendEmbed(config.DiscordSettings.LogsChannel, &embed)

			}
		}
	}

}

func hashIP(ip net.IP) []byte {
	// Convert the IP address to a byte slice
	ipBytes := []byte(ip)

	// Create a SHA-256 hash
	hash := sha256.Sum256(ipBytes)

	// Return the hash as a byte slice
	return hash[:]
}

func handleSellpassWebhook(c *gin.Context) {
	c.Status(200)

	var t []string
	var amountbeforeparse string
	var durationbeforeparse string
	var amount int
	var invitelink string
	var file string
	var invite string
	var shopid int
	var title string

	var data struct {
		Data []DataItem `json:"data"`
	}
	headers := http.Header{
		"authority":       {"api.sellpass.io"},
		"accept":          {"application/json, text/plain, */*"},
		"accept-language": {"en-US,en;q=0.9"},
		"authorization":   {"Bearer " + config.SellpassAutomationSettings.APIKey},
	}
	client := http.Client{Timeout: time.Second * 60}
	req, _ := http.NewRequest("GET", "https://dev.sellpass.io/self/shops", nil)
	req.Header = headers
	resp, _ := client.Do(req)
	bodytext, _ := io.ReadAll(resp.Body)

	_ = json.Unmarshal(bodytext, &data)
	for _, item := range data.Data {
		shopid = item.ShopData.ID
	}

	// Parse the Sellpass webhook payload
	var payload map[string]interface{}
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		return
	}
	// Check if the "data" key exists in the payload
	invoiceID, ok := payload["InvoiceId"].(string)
	if !ok {
		Utils.LogError("Received Sellpass webhook with no or invalid 'InvoiceId' field", "", "")
		return
	}
	url := fmt.Sprintf("https://dev.sellpass.io/self/%v/invoices/%v", shopid, invoiceID)
	req2, _ := http.NewRequest("GET", url, nil)
	req2.Header = headers
	var Order SellpassWebData
	resp2, _ := client.Do(req2)
	bodytext2, _ := io.ReadAll(resp2.Body)
	_ = json.Unmarshal(bodytext2, &Order)

	for _, parts := range Order.Data.PartInvoices {
		title = parts.Product.Title
	}
	t = strings.Split(title, "|")
	amountbeforeparse = t[0]
	durationbeforeparse = t[1]

	if strings.Contains(amountbeforeparse, "30") {
		amount = 30
	} else if strings.Contains(amountbeforeparse, "14") {
		amount = 14
	} else if strings.Contains(amountbeforeparse, "12") {
		amount = 12
	} else if strings.Contains(amountbeforeparse, "10") {
		amount = 10
	} else if strings.Contains(amountbeforeparse, "8") {
		amount = 8
	} else if strings.Contains(amountbeforeparse, "6") {
		amount = 6
	} else if strings.Contains(amountbeforeparse, "4") {
		amount = 4
	} else if strings.Contains(amountbeforeparse, "2") {
		amount = 2
	} else {
		Utils.LogError("Failed to Auto Boost, Improper Amount of Boosts Specified", "", "")
	}

	if strings.Contains(durationbeforeparse, "1") {
		file = "1 Month Tokens.txt"
	} else if strings.Contains(durationbeforeparse, "3") {
		file = "3 Month Tokens.txt"
	} else {
		Utils.LogError("Failed to Auto Boost, Improper Duration Specified", "", "")
	}

	fmt.Sprintf("%v", file)
	for _, field := range Order.Data.PartInvoices {
		for _, fieldValue := range field.CustomFields {
			if fieldValue.CustomField.Name == config.InviteFieldName {
				invitelink = fieldValue.ValueString
			}
		}
	}

	email := Order.Data.CustomerInfo.CustomerForShop.Customer.Email
	ip := Order.Data.CustomerInfo.CurrentIP.IP
	ipp := net.ParseIP(ip)
	hashedIP := hashIP(ipp)

	inviteParts := strings.Split(invitelink, "/")
	count := len(inviteParts)

	if count == 4 {
		invite = inviteParts[3]
	} else if count == 2 {
		invite = inviteParts[1]
	} else {
		invite = invitelink
	}

	c.String(http.StatusOK, fmt.Sprintf("We have successfully started boosting your server! If any problems occur please join our discord server: %v", config.DiscordSettings.SupportServer))

	if invite == "" || amount == 0 || file == "" {
		Utils.LogError("Missing Fields for Boosting, Check Automation Settings on Website or Config!", "", "")
		return
	}

	Utils.ClearScreen()
	Utils.PrintASCII()

	respo, err := Discord.BoostServer(invite, amount, file)
	if err != nil {
		Utils.LogError(err.Error(), "", "")
		return
	}

	var success string
	if len(respo.SuccessTokens) != 0 {
		success = strings.Join(respo.SuccessTokens, "\n")
	} else {
		success = "No succeeded tokens."
	}

	var failed string
	if len(respo.FailedTokens) != 0 {
		failed = strings.Join(respo.FailedTokens, "\n")
	} else {
		failed = "No failed tokens."
	}

	embed := discordgo.MessageEmbed{
		Title:       "Sellpass Order",
		Description: fmt.Sprintf("We have boosted a server successfully.\n**Success**: %v | **Failed**: %v\n**Elapsed Time:** %v", respo.Success, respo.Failed, respo.ElapsedTime),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Order Information",
				Value:  fmt.Sprintf("```\nOrder ID: %v\nProduct Title: %v\nBoost Amount: %v\nServer Invite: %v\n```", invoiceID, title, amount, invite),
				Inline: false,
			},
			{
				Name:   "Customer Information",
				Value:  fmt.Sprintf("```\nEmail: %v\nIP Address: %x\n```", email, hashedIP),
				Inline: false,
			},
			{
				Name:   "Succeeded Tokens",
				Value:  fmt.Sprintf("```\n%v\n```", success),
				Inline: false,
			},
			{
				Name:   "Failed Tokens",
				Value:  fmt.Sprintf("```\n%v\n```", failed),
				Inline: false,
			},
		},
		Color: int(EmbedColor),
	}

	s.ChannelMessageSendEmbed(config.DiscordSettings.LogsChannel, &embed)

}

func EditSellixStock() {
	client := http.Client{Timeout: time.Second * 30}
	headers := http.Header{
		"authority":         {"dev.sellix.io"},
		"accept":            {"application/json, text/plain, */*"},
		"accept-language":   {"en-US,en;q=0.9"},
		"authorization":     {"Bearer " + config.SellixAutomationSettings.APIKey},
		"x-sellix-merchant": {config.SellixAutomationSettings.ShopName},
	}
	for {
		if config.SellixAutomationSettings.ProductSettings.ThreeMonthProductID != "" && config.SellixAutomationSettings.ProductSettings.ThreeMonthProductPrice != 0 {
			var SellixProductInfo SellixProductInfo
			var Serials []string
			Amount := Utils.Get3MTokens()
			for i := 0; i < Amount; i++ {

				Serials = append(Serials, "Your Boosts Will be Delivered Shortly")
			}

			req, _ := http.NewRequest("GET", "https://dev.sellix.io/v1/products/"+config.SellixAutomationSettings.ProductSettings.ThreeMonthProductID, nil)
			req.Header = headers
			resp, _ := client.Do(req)
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			_ = json.Unmarshal(body, &SellixProductInfo)

			payload := map[string]interface{}{
				"title":        SellixProductInfo.Data.Product.Title,
				"price":        config.SellixAutomationSettings.ProductSettings.ThreeMonthProductPrice,
				"description":  SellixProductInfo.Data.Product.Description,
				"type":         "SERIALS",
				"min_quantity": 1,
				"max_quantity": Amount,
				"serials":      Serials,
			}
			jsonPayload, _ := json.Marshal(payload)
			req2, _ := http.NewRequest("PUT", "https://dev.sellix.io/v1/products/"+config.SellixAutomationSettings.ProductSettings.ThreeMonthProductID, bytes.NewReader(jsonPayload))
			req2.Header = headers
			resp2, _ := client.Do(req2)
			defer resp2.Body.Close()

			time.Sleep(time.Second * 1)
		}
		if config.SellixAutomationSettings.ProductSettings.OneMonthProductID != "" && config.SellixAutomationSettings.ProductSettings.OneMonthProductPrice != 0 {
			var SellixProductInfo SellixProductInfo
			var Serials []string
			Amount := Utils.Get1mTokens()
			for i := 0; i < Amount; i++ {

				Serials = append(Serials, "Your Boosts Will be Delivered Shortly")
			}

			req, _ := http.NewRequest("GET", "https://dev.sellix.io/v1/products/"+config.SellixAutomationSettings.ProductSettings.OneMonthProductID, nil)
			req.Header = headers
			resp, _ := client.Do(req)
			defer resp.Body.Close()
			body, _ := io.ReadAll(resp.Body)
			_ = json.Unmarshal(body, &SellixProductInfo)

			payload := map[string]interface{}{
				"title":        SellixProductInfo.Data.Product.Title,
				"price":        config.SellixAutomationSettings.ProductSettings.OneMonthProductPrice,
				"description":  SellixProductInfo.Data.Product.Description,
				"type":         "SERIALS",
				"min_quantity": 1,
				"max_quantity": Amount,
				"serials":      Serials,
			}
			jsonPayload, _ := json.Marshal(payload)
			req2, _ := http.NewRequest("PUT", "https://dev.sellix.io/v1/products/"+config.SellixAutomationSettings.ProductSettings.OneMonthProductID, bytes.NewReader(jsonPayload))
			req2.Header = headers
			resp2, _ := client.Do(req2)
			defer resp2.Body.Close()

			time.Sleep(time.Second * 1)
		}
	}

}

func EditSellappStock() {
	client := http.Client{Timeout: time.Second * 30}
	headers := http.Header{
		"content-type":  {"application/json"},
		"accept":        {"text/json"},
		"authorization": {"Bearer " + config.SellappAutomationSettings.APIKey},
	}
	for {
		if config.SellappAutomationSettings.ProductSettings.ThreeMonthProductID != "" {
			var Serials []string
			var payload map[string]interface{}
			Amount := Utils.Get3MTokens()
			for i := 0; i < Amount; i++ {
				Serials = append(Serials, "Your Boosts Will be Delivered Shortly")
			}

			if len(Serials) == 0 {
				payload = map[string]interface{}{
					"visibility": "ON_HOLD",
					"deliverable": map[string]interface{}{
						"types": []string{"TEXT"},
						"data": map[string]interface{}{
							"serials":         []string{"1", "2"},
							"removeDuplicate": false,
						},
					},
				}

			} else {
				payload = map[string]interface{}{
					"visibility": "PUBLIC",
					"deliverable": map[string]interface{}{
						"types": []string{"TEXT"},
						"data": map[string]interface{}{
							"serials":         Serials,
							"removeDuplicate": false,
						},
					},
				}
			}

			jsonPayload, _ := json.Marshal(payload)
			req, _ := http.NewRequest("PATCH", "https://sell.app/api/v1/listings/"+config.SellappAutomationSettings.ProductSettings.ThreeMonthProductID, bytes.NewReader(jsonPayload))
			req.Header = headers
			resp1, _ := client.Do(req)
			defer resp1.Body.Close()
			time.Sleep(time.Second * 1)

		}
		if config.SellappAutomationSettings.ProductSettings.OneMonthProductID != "" {
			var Serials []string
			var payload map[string]interface{}
			Amount := Utils.Get1mTokens()
			for i := 0; i < Amount; i++ {
				Serials = append(Serials, "Your Boosts Will be Delivered Shortly")
			}

			if len(Serials) == 0 {
				payload = map[string]interface{}{
					"visibility": "ON_HOLD",
					"deliverable": map[string]interface{}{
						"types": []string{"TEXT"},
						"data": map[string]interface{}{
							"serials":         []string{"1", "2"},
							"removeDuplicate": false,
						},
					},
				}

			} else {
				payload = map[string]interface{}{
					"visibility": "PUBLIC",
					"deliverable": map[string]interface{}{
						"types": []string{"TEXT"},
						"data": map[string]interface{}{
							"serials":         Serials,
							"removeDuplicate": false,
						},
					},
				}
			}

			jsonPayload, _ := json.Marshal(payload)
			req, _ := http.NewRequest("PATCH", "https://sell.app/api/v1/listings/"+config.SellappAutomationSettings.ProductSettings.OneMonthProductID, bytes.NewReader(jsonPayload))
			req.Header = headers
			resp1, _ := client.Do(req)
			defer resp1.Body.Close()
			time.Sleep(time.Second * 1)

		}
	}
}

//func EditSellpassStock() {
//	var shopid int
//	var data struct {
//		Data []DataItem `json:"data"`
//	}
//
//	client := http.Client{Timeout: time.Second * 60}
//	headers := http.Header{
//		"authority":       {"api.sellpass.io"},
//		"accept":          {"application/json"},
//		"content-type":    {"application/json"},
//		"accept-language": {"en-US,en;q=0.9"},
//		"authorization":   {"Bearer " + config.SellpassAutomationSettings.APIKey},
//	}
//	req, _ := http.NewRequest("GET", "https://dev.sellpass.io/self/shops", nil)
//	req.Header = headers
//	resp, _ := client.Do(req)
//	bodytext, _ := io.ReadAll(resp.Body)
//	_ = json.Unmarshal(bodytext, &data)
//	for _, item := range data.Data {
//		shopid = item.ShopData.ID
//	}
//
//	for {
//		if config.SellpassAutomationSettings.ProductSettings.ThreeMonthProductID != "" {
//			var ProductInfo SellpassPInfo
//			var Serials []string
//			var vid int
//			var vtitle string
//			var gateway int
//			var price int
//			var currency string
//			req, _ := http.NewRequest("GET", fmt.Sprintf("https://dev.sellpass.io/self/%v/v2/products/%v", shopid, config.SellpassAutomationSettings.ProductSettings.ThreeMonthProductID), nil)
//			req.Header = headers
//			resp, _ := client.Do(req)
//			body, _ := io.ReadAll(resp.Body)
//			fmt.Println(string(body))
//			_ = json.Unmarshal(body, &ProductInfo)
//
//			Amount := Utils.Get3MTokens()
//			for i := 0; i < Amount; i++ {
//				Serials = append(Serials, "Your Boosts Will be Delivered Shortly\n")
//			}
//
//			for _, variants1 := range ProductInfo.Data.Product.Variants {
//				vid = variants1.ID
//				vtitle = variants1.Title
//				price = variants1.PriceDetails.Amount
//				currency = variants1.PriceDetails.Currency
//				for _, gate := range variants1.Gateways {
//					gateway = gate.Gateway
//				}
//			}
//
//			fmt.Println(gateway)
//			fmt.Println(vtitle)
//			fmt.Println(vid)

//payload := SellpassProductUpdate{
//	Title:       ProductInfo.Data.Product.Title,
//	Description: ProductInfo.Data.Product.Description,
//	Variants: []struct {
//		ID               int    `json:"id"`
//		Title            string `json:"title"`
//		Description      string `json:"description"`
//		ShortDescription string `json:"shortDescription"`
//		PriceDetails     struct {
//			Amount   float64 `json:"amount"`
//			Currency string  `json:"currency"`
//		} `json:"priceDetails"`
//		Gateways []struct {
//			Gateway int `json:"gateway"`
//			Rules   struct {
//				BlockVpn bool `json:"blockVpn"`
//			} `json:"rules"`
//			Price struct {
//				Amount   float64 `json:"amount"`
//				Currency string  `json:"currency"`
//			} `json:"price"`
//		} `json:"gateways"`
//		ProductType int `json:"productType"`
//		AsDynamic   struct {
//			Stock       int    `json:"stock"`
//			ExternalURL string `json:"externalUrl"`
//			MinAmount   int    `json:"minAmount"`
//			MaxAmount   int    `json:"maxAmount"`
//			IsInternal  bool   `json:"isInternal"`
//		} `json:"asDynamic"`
//		AsSerials struct {
//			Delimiter        string `json:"delimiter"`
//			Serials          string `json:"serials"`
//			MinAmount        int    `json:"minAmount"`
//			MaxAmount        int    `json:"maxAmount"`
//			RemoveDuplicates bool   `json:"removeDuplicates"`
//		} `json:"asSerials"`
//		AsService struct {
//			Stock     int    `json:"stock"`
//			Text      string `json:"text"`
//			MinAmount int    `json:"minAmount"`
//			MaxAmount int    `json:"maxAmount"`
//		} `json:"asService"`
//		CustomerNote string `json:"customerNote"`
//		RedirectURL  string `json:"redirectUrl"`
//		CustomFields []struct {
//			ID          int    `json:"id"`
//			Type        int    `json:"type"`
//			Name        string `json:"name"`
//			Required    bool   `json:"required"`
//			ValueString string `json:"valueString"`
//			Placeholder string `json:"placeholder"`
//			Regex       string `json:"regex"`
//			ValueInt    int    `json:"valueInt"`
//			ValueBool   bool   `json:"valueBool"`
//		} `json:"customFields"`
//
//		Warranty struct {
//			Text            string `json:"text"`
//			DurationSeconds int    `json:"durationSeconds"`
//		} `json:"warranty"`
//		DiscordSocialConnectSettings struct {
//			Enabled                    bool `json:"enabled"`
//			Required                   bool `json:"required"`
//			BeforePurchaseRequireRoles struct {
//				GuildID string   `json:"guildId"`
//				RoleIds []string `json:"roleIds"`
//			} `json:"beforePurchaseRequireRoles"`
//			BeforePurchaseServer struct {
//				GuildID string   `json:"guildId"`
//				RoleIds []string `json:"roleIds"`
//			} `json:"beforePurchaseServer"`
//			AfterPurchaseServer struct {
//				GuildID string   `json:"guildId"`
//				RoleIds []string `json:"roleIds"`
//			} `json:"afterPurchaseServer"`
//		} `json:"discordSocialConnectSettings"`
//	}{
//		{
//			ID:          vid,
//			Title:       vtitle,
//			ProductType: 0,
//			Gateways: map[string]{}
//			AsSerials: struct {
//				Delimiter        string `json:"delimiter"`
//				Serials          string `json:"serials"`
//				MinAmount        int    `json:"minAmount"`
//				MaxAmount        int    `json:"maxAmount"`
//				RemoveDuplicates bool   `json:"removeDuplicates"`
//			}(struct {
//				Delimiter        string
//				Serials          string
//				MinAmount        int
//				MaxAmount        int
//				RemoveDuplicates bool
//			}{Delimiter: "\n", Serials: fmt.Sprintf("%v", Serials), MinAmount: 1, MaxAmount: 999, RemoveDuplicates: false}),
//		},
//	},
//	Path: ProductInfo.Data.Path,
//	Seo: struct {
//		MetaTitle       string `json:"metaTitle"`
//		MetaDescription string `json:"metaDescription"`
//	}(struct {
//		MetaTitle       string
//		MetaDescription string
//	}{MetaTitle: ProductInfo.Data.Seo.MetaTitle, MetaDescription: ProductInfo.Data.Seo.MetaDescription}),
//	Unlisted: false,
//	Private:  false,
//	OnHold:   false,
//}
//
//jsonp, _ := json.Marshal(payload)
//fmt.Println(string(jsonp))

//			payload := fmt.Sprintf(`{
//				"title": "%v",
//				"description": "%v",
//				"variants": [{
//					"id": %v,
//					"title": "%v",
//					"priceDetails": {
//						"amount": %v,
//						currency": "%v"
//					},
//					"gateways" : [{
//							"gateway": %v
//					}],
//					"productType": 0,
//					"asSerials": {
//						"delimiter" : "/n",
//						"serials": "Your Boosts Will be Delivered Shortly",
//						"minAmount" : 1,
//						"maxAmount" : 999,
//						"removeDuplicates": false
//					}
//				}],
//				"path": "%v",
//				"seo": {
//					"metaTitle": "%v",
//					"metaDescription": "%v"
//				},
//				"unlisted": false,
//				"private": false,
//				"onHold": false
//			}`, ProductInfo.Data.Product.Title, ProductInfo.Data.Product.Description, vid, vtitle, price, currency, gateway, ProductInfo.Data.Path, ProductInfo.Data.Seo.MetaTitle, ProductInfo.Data.Seo.MetaDescription)
//			reader := strings.NewReader(payload)
//			var data1 map[string]interface{}
//			decoder := json.NewDecoder(reader)
//			pp := decoder.Decode(&data1)
//
//			fmt.Printf("%v", pp)
//			req2, _ := http.NewRequest("PUT", fmt.Sprintf("https://dev.sellpass.io/self/%v/v2/products/%v", shopid, config.SellpassAutomationSettings.ProductSettings.ThreeMonthProductID), nil)
//			req2.Header = headers
//			resp2, _ := client.Do(req2)
//			body2, _ := io.ReadAll(resp2.Body)
//			fmt.Println("This is body2" + string(body2))
//			time.Sleep(time.Second * 10)
//		}
//	}
//
//}
