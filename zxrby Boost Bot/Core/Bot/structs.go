package Bot

import "time"

type DataItem struct {
	ID       int `json:"id"`
	ShopData struct {
		ID int `json:"id"`
	} `json:"shop"`
}

type SellpassWebData struct {
	Data struct {
		ID           string `json:"id"`
		CustomerInfo struct {
			CustomerForShop struct {
				ID             int       `json:"id"`
				CreatedAt      time.Time `json:"createdAt"`
				TotalSpent     int       `json:"totalSpent"`
				TotalPurchases int       `json:"totalPurchases"`
				IsBlocked      bool      `json:"isBlocked"`
				Customer       struct {
					Email string `json:"email"`
				} `json:"customer"`
				CustomerID    int `json:"customerId"`
				Visits        int `json:"visits"`
				AverageReview int `json:"averageReview"`
				ShopID        int `json:"shopId"`
			} `json:"customerForShop"`
			CurrentIP struct {
				ID             int       `json:"id"`
				DateTime       time.Time `json:"dateTime"`
				IP             string    `json:"ip"`
				Country        string    `json:"country"`
				City           string    `json:"city"`
				ConnectionType int       `json:"connectionType"`
				RiskScore      int       `json:"riskScore"`
				IsoCode        string    `json:"isoCode"`
				Isp            string    `json:"isp"`
			} `json:"currentIp"`
			Useragent string `json:"useragent"`
			InvoiceID string `json:"invoiceId"`
		} `json:"customerInfo"`
		CustomerInfoID int `json:"customerInfoId"`
		PartInvoices   []struct {
			ID      int `json:"id"`
			Product struct {
				ID                int       `json:"id"`
				UniquePath        string    `json:"uniquePath"`
				Title             string    `json:"title"`
				Description       string    `json:"description"`
				ShortDescription  string    `json:"shortDescription"`
				Unlisted          bool      `json:"unlisted"`
				Private           bool      `json:"private"`
				OnHold            bool      `json:"onHold"`
				IsInStock         bool      `json:"isInStock"`
				IsInstantDelivery bool      `json:"isInstantDelivery"`
				CreatedAt         time.Time `json:"createdAt"`
				UpdatedAt         time.Time `json:"updatedAt"`
				ListingID         int       `json:"listingId"`
				IsDeleted         bool      `json:"isDeleted"`
				IsInternal        bool      `json:"isInternal"`
				ShopID            int       `json:"shopId"`
			} `json:"product"`
			Quantity            int      `json:"quantity"`
			DeliveryStatus      int      `json:"deliveryStatus"`
			DeliveredGoods      []string `json:"deliveredGoods"`
			TotalDeliveredGoods []string `json:"totalDeliveredGoods"`
			CustomFields        []struct {
				ID          int `json:"id"`
				CustomField struct {
					ID          int    `json:"id"`
					Type        int    `json:"type"`
					Name        string `json:"name"`
					Required    bool   `json:"required"`
					ValueString string `json:"valueString"`
				} `json:"customField"`
				ValueString string `json:"valueString"`
			} `json:"customFields"`
			Replacements []any  `json:"replacements"`
			InvoiceID    string `json:"invoiceId"`
			RawPrice     int    `json:"rawPrice"`
			RawPriceUSD  int    `json:"rawPriceUSD"`
			EndPrice     int    `json:"endPrice"`
			EndPriceUSD  int    `json:"endPriceUSD"`
		} `json:"partInvoices"`
		Status      int `json:"status"`
		RawPrice    int `json:"rawPrice"`
		RawPriceUSD int `json:"rawPriceUSD"`
		EndPrice    int `json:"endPrice"`
		EndPriceUSD int `json:"endPriceUSD"`
		Gateway     struct {
			GatewayName int `json:"gatewayName"`
		} `json:"gateway"`
		Currency string `json:"currency"`
		Tickets  []any  `json:"tickets"`
		Timeline []struct {
			ID     int       `json:"id"`
			Time   time.Time `json:"time"`
			Status int       `json:"status"`
		} `json:"timeline"`
		ForPayPalEmail struct {
			HostedURL string `json:"hostedUrl"`
			ID        int    `json:"id"`
			InvoiceID string `json:"invoiceId"`
		} `json:"forPayPalEmail"`
		ShopID            int       `json:"shopId"`
		ExpiresAt         time.Time `json:"expiresAt"`
		ManuallyCompleted bool      `json:"manuallyCompleted"`
		HideFromStats     bool      `json:"hideFromStats"`
		IsBalanceTopUp    bool      `json:"isBalanceTopUp"`
	} `json:"data"`
}

type SellixProductInfo struct {
	Status int `json:"status"`
	Data   struct {
		Product struct {
			ID                      int      `json:"id"`
			Uniqid                  string   `json:"uniqid"`
			Slug                    string   `json:"slug"`
			ShopID                  int      `json:"shop_id"`
			Type                    string   `json:"type"`
			Subtype                 any      `json:"subtype"`
			Title                   string   `json:"title"`
			Currency                string   `json:"currency"`
			PayWhatYouWant          int      `json:"pay_what_you_want"`
			Price                   int      `json:"price"`
			PriceDisplay            int      `json:"price_display"`
			PriceDiscount           int      `json:"price_discount"`
			AffiliateRevenuePercent int      `json:"affiliate_revenue_percent"`
			PriceVariants           any      `json:"price_variants"`
			Description             string   `json:"description"`
			LicensingEnabled        int      `json:"licensing_enabled"`
			LicensePeriod           any      `json:"license_period"`
			ImageAttachment         string   `json:"image_attachment"`
			FileAttachment          string   `json:"file_attachment"`
			YoutubeLink             any      `json:"youtube_link"`
			VolumeDiscounts         []any    `json:"volume_discounts"`
			RecurringInterval       any      `json:"recurring_interval"`
			RecurringIntervalCount  any      `json:"recurring_interval_count"`
			TrialPeriod             any      `json:"trial_period"`
			PaypalProductID         any      `json:"paypal_product_id"`
			PaypalPlanID            any      `json:"paypal_plan_id"`
			StripePriceID           any      `json:"stripe_price_id"`
			DiscordIntegration      int      `json:"discord_integration"`
			DiscordOptional         int      `json:"discord_optional"`
			DiscordSetRole          int      `json:"discord_set_role"`
			DiscordServerID         string   `json:"discord_server_id"`
			DiscordRoleID           any      `json:"discord_role_id"`
			DiscordRemoveRole       int      `json:"discord_remove_role"`
			QuantityMin             int      `json:"quantity_min"`
			QuantityMax             int      `json:"quantity_max"`
			QuantityWarning         int      `json:"quantity_warning"`
			Gateways                []string `json:"gateways"`
			CustomFields            []struct {
				Type     string `json:"type"`
				Name     string `json:"name"`
				Default  string `json:"default"`
				Required bool   `json:"required"`
			} `json:"custom_fields"`
			CryptoConfirmationsNeeded int    `json:"crypto_confirmations_needed"`
			MaxRiskLevel              int    `json:"max_risk_level"`
			BlockVpnProxies           bool   `json:"block_vpn_proxies"`
			DeliveryText              string `json:"delivery_text"`
			DeliveryTime              any    `json:"delivery_time"`
			ServiceText               string `json:"service_text"`
			StockDelimiter            string `json:"stock_delimiter"`
			Stock                     int    `json:"stock"`
			DynamicWebhook            string `json:"dynamic_webhook"`
			Bestseller                int    `json:"bestseller"`
			SortPriority              int    `json:"sort_priority"`
			Unlisted                  bool   `json:"unlisted"`
			OnHold                    int    `json:"on_hold"`
			TermsOfService            string `json:"terms_of_service"`
			Warranty                  int    `json:"warranty"`
			WarrantyText              string `json:"warranty_text"`
			WatermarkEnabled          int    `json:"watermark_enabled"`
			WatermarkText             string `json:"watermark_text"`
			RedirectLink              any    `json:"redirect_link"`
			LabelSingular             any    `json:"label_singular"`
			LabelPlural               any    `json:"label_plural"`
			Private                   bool   `json:"private"`
			CreatedAt                 int    `json:"created_at"`
			UpdatedAt                 int    `json:"updated_at"`
			UpdatedBy                 int    `json:"updated_by"`
			MarketplaceCategoryID     int    `json:"marketplace_category_id"`
			Name                      string `json:"name"`
			ImageName                 any    `json:"image_name"`
			ImageStorage              any    `json:"image_storage"`
			CloudflareImageID         any    `json:"cloudflare_image_id"`
			Feedback                  struct {
				Total    int   `json:"total"`
				Positive int   `json:"positive"`
				Neutral  int   `json:"neutral"`
				Negative int   `json:"negative"`
				Numbers  []any `json:"numbers"`
				List     []any `json:"list"`
			} `json:"feedback"`
			Categories          []any `json:"categories"`
			PaymentGatewaysFees []any `json:"payment_gateways_fees"`
			PriceConversions    struct {
				Cad    float64 `json:"CAD"`
				Hkd    float64 `json:"HKD"`
				Isk    float64 `json:"ISK"`
				Php    int     `json:"PHP"`
				Dkk    float64 `json:"DKK"`
				Huf    float64 `json:"HUF"`
				Czk    float64 `json:"CZK"`
				Gbp    float64 `json:"GBP"`
				Ron    float64 `json:"RON"`
				Sek    float64 `json:"SEK"`
				Idr    float64 `json:"IDR"`
				Inr    float64 `json:"INR"`
				Brl    float64 `json:"BRL"`
				Rub    float64 `json:"RUB"`
				Hrk    float64 `json:"HRK"`
				Jpy    float64 `json:"JPY"`
				Thb    float64 `json:"THB"`
				Chf    float64 `json:"CHF"`
				Eur    float64 `json:"EUR"`
				Myr    float64 `json:"MYR"`
				Bgn    float64 `json:"BGN"`
				Try    float64 `json:"TRY"`
				Cny    float64 `json:"CNY"`
				Nok    float64 `json:"NOK"`
				Nzd    float64 `json:"NZD"`
				Zar    float64 `json:"ZAR"`
				Usd    int     `json:"USD"`
				Mxn    float64 `json:"MXN"`
				Sgd    float64 `json:"SGD"`
				Aud    float64 `json:"AUD"`
				Ils    float64 `json:"ILS"`
				Krw    float64 `json:"KRW"`
				Pln    float64 `json:"PLN"`
				Crypto struct {
					Btc   string `json:"BTC"`
					Bnb   string `json:"BNB"`
					Eth   string `json:"ETH"`
					Ltc   string `json:"LTC"`
					Bch   string `json:"BCH"`
					Nano  string `json:"NANO"`
					Xmr   string `json:"XMR"`
					Sol   string `json:"SOL"`
					Xrp   string `json:"XRP"`
					Cro   string `json:"CRO"`
					Usdc  string `json:"USDC"`
					Usdt  string `json:"USDT"`
					Trx   string `json:"TRX"`
					Ccd   string `json:"CCD"`
					Matic string `json:"MATIC"`
					Ape   string `json:"APE"`
					Pepe  string `json:"PEPE"`
					Dai   string `json:"DAI"`
					Weth  string `json:"WETH"`
					Shib  string `json:"SHIB"`
				} `json:"crypto"`
			} `json:"price_conversions"`
			Serials       []any  `json:"serials"`
			Webhooks      []any  `json:"webhooks"`
			Theme         string `json:"theme"`
			DarkMode      int    `json:"dark_mode"`
			VatPercentage string `json:"vat_percentage"`
			TaxDetails    struct {
				VatPercentage              string `json:"vat_percentage"`
				TaxConfiguration           string `json:"tax_configuration"`
				TaxConfigurationData       []any  `json:"tax_configuration_data"`
				DisplayTaxOnStorefront     int    `json:"display_tax_on_storefront"`
				DisplayTaxCustomFields     int    `json:"display_tax_custom_fields"`
				ValidationOnlyForCompanies int    `json:"validation_only_for_companies"`
				ValidateVatNumber          int    `json:"validate_vat_number"`
				PricesTaxInclusive         int    `json:"prices_tax_inclusive"`
			} `json:"tax_details"`
			AverageScore any   `json:"average_score"`
			SoldCount    int   `json:"sold_count"`
			Addons       []any `json:"addons"`
		} `json:"product"`
	} `json:"data"`
	Error   any    `json:"error"`
	Message any    `json:"message"`
	Env     string `json:"env"`
	Log     struct {
		Cache           string `json:"cache"`
		OneStart        int64  `json:"1.start"`
		TwoShop         int    `json:"2.shop"`
		ThreeErrors     int    `json:"3.errors"`
		FourInfocard    int    `json:"4.infocard"`
		SixAveragescore int    `json:"6.averagescore"`
		SevenSoldCount  int    `json:"7.sold_count"`
		EightRegex      int    `json:"8.regex"`
		NineView        int    `json:"9.view"`
	} `json:"log"`
}

var cvalue = config.InviteFieldName

type SellixProductPaid struct {
	Event string `json:"event"`
	Data  struct {
		ID                         int    `json:"id"`
		Uniqid                     string `json:"uniqid"`
		RecurringBillingID         any    `json:"recurring_billing_id"`
		Type                       string `json:"type"`
		Subtype                    any    `json:"subtype"`
		Total                      int    `json:"total"`
		TotalDisplay               int    `json:"total_display"`
		ProductVariants            any    `json:"product_variants"`
		ExchangeRate               int    `json:"exchange_rate"`
		CryptoExchangeRate         int    `json:"crypto_exchange_rate"`
		CryptoExchangeRateFix      string `json:"crypto_exchange_rate_fix"`
		Currency                   string `json:"currency"`
		ShopID                     int    `json:"shop_id"`
		ShopImageName              any    `json:"shop_image_name"`
		ShopImageStorage           any    `json:"shop_image_storage"`
		ShopCloudflareImageID      any    `json:"shop_cloudflare_image_id"`
		Name                       string `json:"name"`
		CustomerEmail              string `json:"customer_email"`
		AffiliateRevenueCustomerID any    `json:"affiliate_revenue_customer_id"`
		PaypalEmailDelivery        bool   `json:"paypal_email_delivery"`
		ProductID                  string `json:"product_id"`
		ProductTitle               string `json:"product_title"`
		ProductType                string `json:"product_type"`
		SubscriptionID             any    `json:"subscription_id"`
		SubscriptionTime           any    `json:"subscription_time"`
		Gateway                    string `json:"gateway"`
		Blockchain                 any    `json:"blockchain"`
		PaypalApm                  any    `json:"paypal_apm"`
		StripeApm                  any    `json:"stripe_apm"`
		PaypalEmail                any    `json:"paypal_email"`
		PaypalOrderID              string `json:"paypal_order_id"`
		PaypalPayerEmail           any    `json:"paypal_payer_email"`
		PaypalFee                  int    `json:"paypal_fee"`
		PaypalSubscriptionID       any    `json:"paypal_subscription_id"`
		PaypalSubscriptionLink     any    `json:"paypal_subscription_link"`
		LexOrderID                 any    `json:"lex_order_id"`
		LexPaymentMethod           any    `json:"lex_payment_method"`
		PaydashPaymentID           any    `json:"paydash_paymentID"`
		VirtualPaymentsID          any    `json:"virtual_payments_id"`
		StripeClientSecret         any    `json:"stripe_client_secret"`
		StripePriceID              any    `json:"stripe_price_id"`
		SkrillEmail                any    `json:"skrill_email"`
		SkrillSid                  any    `json:"skrill_sid"`
		SkrillLink                 any    `json:"skrill_link"`
		PerfectmoneyID             any    `json:"perfectmoney_id"`
		BinanceInvoiceID           any    `json:"binance_invoice_id"`
		BinanceQrcode              any    `json:"binance_qrcode"`
		BinanceCheckoutURL         any    `json:"binance_checkout_url"`
		CryptoAddress              any    `json:"crypto_address"`
		CryptoAmount               int    `json:"crypto_amount"`
		CryptoReceived             int    `json:"crypto_received"`
		CryptoURI                  any    `json:"crypto_uri"`
		CryptoConfirmationsNeeded  int    `json:"crypto_confirmations_needed"`
		CryptoScheduledPayout      bool   `json:"crypto_scheduled_payout"`
		CryptoPayout               int    `json:"crypto_payout"`
		FeeBilled                  bool   `json:"fee_billed"`
		BillInfo                   any    `json:"bill_info"`
		CashappQrcode              any    `json:"cashapp_qrcode"`
		CashappNote                any    `json:"cashapp_note"`
		CashappCashtag             any    `json:"cashapp_cashtag"`
		Country                    string `json:"country"`
		Location                   string `json:"location"`
		IP                         string `json:"ip"`
		IsVpnOrProxy               bool   `json:"is_vpn_or_proxy"`
		UserAgent                  string `json:"user_agent"`
		Quantity                   int    `json:"quantity"`
		CouponID                   any    `json:"coupon_id"`
		CustomFields               map[string]string
		DeveloperInvoice           bool   `json:"developer_invoice"`
		DeveloperTitle             any    `json:"developer_title"`
		DeveloperWebhook           any    `json:"developer_webhook"`
		DeveloperReturnURL         any    `json:"developer_return_url"`
		Status                     string `json:"status"`
		StatusDetails              string `json:"status_details"`
		VoidDetails                any    `json:"void_details"`
		Discount                   int    `json:"discount"`
		FeePercentage              int    `json:"fee_percentage"`
		FeeBreakdown               string `json:"fee_breakdown"`
		DiscountBreakdown          struct {
			Log struct {
				Coupon struct {
					Total         int `json:"total"`
					Coupon        int `json:"coupon"`
					TotalDisplay  int `json:"total_display"`
					CouponDisplay int `json:"coupon_display"`
				} `json:"coupon"`
				BundleDiscount []any `json:"bundle_discount"`
				VolumeDiscount struct {
					Total                 int `json:"total"`
					TotalDisplay          int `json:"total_display"`
					VolumeDiscount        int `json:"volume_discount"`
					VolumeDiscountDisplay int `json:"volume_discount_display"`
				} `json:"volume_discount"`
			} `json:"log"`
			Tax struct {
				Percentage string `json:"percentage"`
			} `json:"tax"`
			Addons []any `json:"addons"`
			Coupon []any `json:"coupon"`
			TaxLog struct {
				Vat                 string `json:"vat"`
				Type                string `json:"type"`
				VatTotal            int    `json:"vat_total"`
				TotalPreVat         int    `json:"total_pre_vat"`
				TotalWithVat        int    `json:"total_with_vat"`
				VatPercentage       string `json:"vat_percentage"`
				VatTotalDisplay     int    `json:"vat_total_display"`
				TotalPreVatDisplay  int    `json:"total_pre_vat_display"`
				TotalWithVatDisplay int    `json:"total_with_vat_display"`
			} `json:"tax_log"`
			Products   []any `json:"products"`
			Currencies struct {
				Default string `json:"default"`
				Display string `json:"display"`
			} `json:"currencies"`
			GatewayFee      []any `json:"gateway_fee"`
			PriceDiscount   []any `json:"price_discount"`
			BundleDiscounts []any `json:"bundle_discounts"`
			VolumeDiscounts struct {
				Six522Ac1E96167 struct {
					Type          any `json:"type"`
					Amount        int `json:"amount"`
					Percentage    any `json:"percentage"`
					AmountDisplay int `json:"amount_display"`
				} `json:"6522ac1e96167"`
			} `json:"volume_discounts"`
		} `json:"discount_breakdown"`
		DayValue      int    `json:"day_value"`
		Day           string `json:"day"`
		Month         string `json:"month"`
		Year          int    `json:"year"`
		ProductAddons any    `json:"product_addons"`
		BundleConfig  any    `json:"bundle_config"`
		CreatedAt     int    `json:"created_at"`
		UpdatedAt     int    `json:"updated_at"`
		UpdatedBy     int    `json:"updated_by"`
		IPInfo        struct {
			Success            bool    `json:"success"`
			Message            string  `json:"message"`
			FraudScore         int     `json:"fraud_score"`
			CountryCode        string  `json:"country_code"`
			Region             string  `json:"region"`
			City               string  `json:"city"`
			Isp                string  `json:"ISP"`
			Asn                int     `json:"ASN"`
			OperatingSystem    string  `json:"operating_system"`
			Browser            string  `json:"browser"`
			Organization       string  `json:"organization"`
			IsCrawler          bool    `json:"is_crawler"`
			Timezone           string  `json:"timezone"`
			Mobile             bool    `json:"mobile"`
			Host               string  `json:"host"`
			Proxy              int     `json:"proxy"`
			Vpn                bool    `json:"vpn"`
			Tor                bool    `json:"tor"`
			ActiveVpn          bool    `json:"active_vpn"`
			ActiveTor          bool    `json:"active_tor"`
			DeviceBrand        string  `json:"device_brand"`
			DeviceModel        string  `json:"device_model"`
			RecentAbuse        bool    `json:"recent_abuse"`
			BotStatus          bool    `json:"bot_status"`
			ConnectionType     string  `json:"connection_type"`
			AbuseVelocity      string  `json:"abuse_velocity"`
			ZipCode            string  `json:"zip_code"`
			Latitude           float64 `json:"latitude"`
			Longitude          float64 `json:"longitude"`
			RequestID          string  `json:"request_id"`
			TransactionDetails struct {
				ValidBillingAddress       any      `json:"valid_billing_address"`
				ValidShippingAddress      any      `json:"valid_shipping_address"`
				ValidBillingEmail         bool     `json:"valid_billing_email"`
				ValidShippingEmail        any      `json:"valid_shipping_email"`
				RiskyBillingPhone         any      `json:"risky_billing_phone"`
				RiskyShippingPhone        any      `json:"risky_shipping_phone"`
				BillingPhoneCountry       any      `json:"billing_phone_country"`
				BillingPhoneCountryCode   any      `json:"billing_phone_country_code"`
				ShippingPhoneCountry      any      `json:"shipping_phone_country"`
				ShippingPhoneCountryCode  any      `json:"shipping_phone_country_code"`
				BillingPhoneCarrier       any      `json:"billing_phone_carrier"`
				ShippingPhoneCarrier      any      `json:"shipping_phone_carrier"`
				BillingPhoneLineType      string   `json:"billing_phone_line_type"`
				ShippingPhoneLineType     string   `json:"shipping_phone_line_type"`
				FraudulentBehavior        bool     `json:"fraudulent_behavior"`
				BinCountry                any      `json:"bin_country"`
				BinType                   string   `json:"bin_type"`
				RiskyUsername             any      `json:"risky_username"`
				ValidBillingPhone         any      `json:"valid_billing_phone"`
				ValidShippingPhone        any      `json:"valid_shipping_phone"`
				LeakedBillingEmail        bool     `json:"leaked_billing_email"`
				LeakedShippingEmail       any      `json:"leaked_shipping_email"`
				LeakedUserData            bool     `json:"leaked_user_data"`
				IsPrepaidCard             any      `json:"is_prepaid_card"`
				PhoneNameIdentityMatch    string   `json:"phone_name_identity_match"`
				PhoneEmailIdentityMatch   string   `json:"phone_email_identity_match"`
				PhoneAddressIdentityMatch string   `json:"phone_address_identity_match"`
				EmailNameIdentityMatch    string   `json:"email_name_identity_match"`
				NameAddressIdentityMatch  string   `json:"name_address_identity_match"`
				AddressEmailIdentityMatch string   `json:"address_email_identity_match"`
				RiskScore                 int      `json:"risk_score"`
				BinBankName               any      `json:"bin_bank_name"`
				RiskFactors               []string `json:"risk_factors"`
			} `json:"transaction_details"`
			Asn0 int    `json:"asn"`
			Isp0 string `json:"isp"`
		} `json:"ip_info"`
		Serials       []string `json:"serials"`
		LockedSerials []any    `json:"locked_serials"`
		Webhooks      []struct {
			Uniqid       string `json:"uniqid"`
			URL          string `json:"url"`
			Event        string `json:"event"`
			Retries      int    `json:"retries"`
			ResponseCode int    `json:"response_code"`
			CreatedAt    int    `json:"created_at"`
		} `json:"webhooks"`
		PaypalDispute    any   `json:"paypal_dispute"`
		ProductDownloads []any `json:"product_downloads"`
		PaymentLinkID    any   `json:"payment_link_id"`
		License          bool  `json:"license"`
		StatusHistory    []struct {
			ID        int    `json:"id"`
			InvoiceID string `json:"invoice_id"`
			Status    string `json:"status"`
			Details   string `json:"details"`
			CreatedAt int    `json:"created_at"`
		} `json:"status_history"`
		AmlWallets         []any  `json:"aml_wallets"`
		CryptoTransactions []any  `json:"crypto_transactions"`
		PaypalClientID     string `json:"paypal_client_id"`
		Product            struct {
			Uniqid                  string `json:"uniqid"`
			Title                   string `json:"title"`
			RedirectLink            any    `json:"redirect_link"`
			Description             string `json:"description"`
			PriceDisplay            int    `json:"price_display"`
			Currency                string `json:"currency"`
			ImageName               any    `json:"image_name"`
			ImageStorage            any    `json:"image_storage"`
			PayWhatYouWant          int    `json:"pay_what_you_want"`
			AffiliateRevenuePercent int    `json:"affiliate_revenue_percent"`
			CloudflareImageID       any    `json:"cloudflare_image_id"`
			LabelSingular           any    `json:"label_singular"`
			LabelPlural             any    `json:"label_plural"`
			Feedback                struct {
				Total    int   `json:"total"`
				Positive int   `json:"positive"`
				Neutral  int   `json:"neutral"`
				Negative int   `json:"negative"`
				List     []any `json:"list"`
			} `json:"feedback"`
			AverageScore              any   `json:"average_score"`
			ID                        int   `json:"id"`
			ShopID                    int   `json:"shop_id"`
			Price                     int   `json:"price"`
			QuantityMin               int   `json:"quantity_min"`
			QuantityMax               int   `json:"quantity_max"`
			QuantityWarning           int   `json:"quantity_warning"`
			Gateways                  []any `json:"gateways"`
			CryptoConfirmationsNeeded int   `json:"crypto_confirmations_needed"`
			MaxRiskLevel              int   `json:"max_risk_level"`
			BlockVpnProxies           bool  `json:"block_vpn_proxies"`
			Private                   bool  `json:"private"`
			Stock                     int   `json:"stock"`
			Unlisted                  bool  `json:"unlisted"`
			SortPriority              int   `json:"sort_priority"`
			CreatedAt                 int   `json:"created_at"`
			UpdatedAt                 int   `json:"updated_at"`
			UpdatedBy                 int   `json:"updated_by"`
		} `json:"product"`
		TotalConversions struct {
			Cad    float64 `json:"CAD"`
			Hkd    float64 `json:"HKD"`
			Isk    float64 `json:"ISK"`
			Php    int     `json:"PHP"`
			Dkk    float64 `json:"DKK"`
			Huf    float64 `json:"HUF"`
			Czk    float64 `json:"CZK"`
			Gbp    float64 `json:"GBP"`
			Ron    float64 `json:"RON"`
			Sek    float64 `json:"SEK"`
			Idr    float64 `json:"IDR"`
			Inr    float64 `json:"INR"`
			Brl    float64 `json:"BRL"`
			Rub    float64 `json:"RUB"`
			Hrk    float64 `json:"HRK"`
			Jpy    float64 `json:"JPY"`
			Thb    float64 `json:"THB"`
			Chf    float64 `json:"CHF"`
			Eur    float64 `json:"EUR"`
			Myr    float64 `json:"MYR"`
			Bgn    float64 `json:"BGN"`
			Try    float64 `json:"TRY"`
			Cny    float64 `json:"CNY"`
			Nok    float64 `json:"NOK"`
			Nzd    float64 `json:"NZD"`
			Zar    float64 `json:"ZAR"`
			Usd    string  `json:"USD"`
			Mxn    float64 `json:"MXN"`
			Sgd    float64 `json:"SGD"`
			Aud    float64 `json:"AUD"`
			Ils    float64 `json:"ILS"`
			Krw    float64 `json:"KRW"`
			Pln    float64 `json:"PLN"`
			Crypto struct {
				Btc   string `json:"BTC"`
				Bnb   string `json:"BNB"`
				Eth   string `json:"ETH"`
				Ltc   string `json:"LTC"`
				Bch   string `json:"BCH"`
				Nano  string `json:"NANO"`
				Xmr   string `json:"XMR"`
				Sol   string `json:"SOL"`
				Xrp   string `json:"XRP"`
				Cro   string `json:"CRO"`
				Usdc  string `json:"USDC"`
				Usdt  string `json:"USDT"`
				Trx   string `json:"TRX"`
				Ccd   string `json:"CCD"`
				Matic string `json:"MATIC"`
				Ape   string `json:"APE"`
				Pepe  string `json:"PEPE"`
				Dai   string `json:"DAI"`
				Weth  string `json:"WETH"`
				Shib  string `json:"SHIB"`
			} `json:"crypto"`
		} `json:"total_conversions"`
		Theme      string `json:"theme"`
		DarkMode   int    `json:"dark_mode"`
		CryptoMode any    `json:"crypto_mode"`
		Products   []struct {
			Uniqid                  string `json:"uniqid"`
			Title                   string `json:"title"`
			RedirectLink            any    `json:"redirect_link"`
			Description             string `json:"description"`
			PriceDisplay            string `json:"price_display"`
			Currency                string `json:"currency"`
			ImageName               any    `json:"image_name"`
			ImageStorage            any    `json:"image_storage"`
			PayWhatYouWant          int    `json:"pay_what_you_want"`
			AffiliateRevenuePercent int    `json:"affiliate_revenue_percent"`
			CloudflareImageID       any    `json:"cloudflare_image_id"`
			LabelSingular           any    `json:"label_singular"`
			LabelPlural             any    `json:"label_plural"`
			Feedback                struct {
				Total    int   `json:"total"`
				Positive int   `json:"positive"`
				Neutral  int   `json:"neutral"`
				Negative int   `json:"negative"`
				List     []any `json:"list"`
			} `json:"feedback"`
			AverageScore any `json:"average_score"`
		} `json:"products"`
		GatewaysAvailable            []string `json:"gateways_available"`
		ShopPaymentGatewaysFees      []any    `json:"shop_payment_gateways_fees"`
		ShopPaypalCreditCard         bool     `json:"shop_paypal_credit_card"`
		ShopForcePaypalEmailDelivery bool     `json:"shop_force_paypal_email_delivery"`
		ShopWalletconnectID          any      `json:"shop_walletconnect_id"`
		VoidTimes                    []struct {
			Gateways []string `json:"gateways"`
			Conf     struct {
				Void       int `json:"void"`
				WaitPeriod any `json:"wait_period"`
			} `json:"conf,omitempty"`
			Conf0 struct {
				Void                    int `json:"void"`
				WaitPeriod              int `json:"wait_period"`
				Partial                 int `json:"partial"`
				WaitingForConfirmations int `json:"waiting_for_confirmations"`
			} `json:"conf,omitempty"`
			Conf1 struct {
				Void                    int `json:"void"`
				WaitPeriod              int `json:"wait_period"`
				Partial                 int `json:"partial"`
				WaitingForConfirmations int `json:"waiting_for_confirmations"`
			} `json:"conf,omitempty"`
			Conf2 struct {
				Void                    int `json:"void"`
				WaitPeriod              int `json:"wait_period"`
				Partial                 int `json:"partial"`
				WaitingForConfirmations int `json:"waiting_for_confirmations"`
			} `json:"conf,omitempty"`
		} `json:"void_times"`
	} `json:"data"`
}

type SellpassPInfo struct {
	Data struct {
		ID              int    `json:"id"`
		Path            string `json:"path"`
		SearchWordsMeta string `json:"searchWordsMeta"`
		Position        int    `json:"position"`
		MinPrice        int    `json:"minPrice"`
		Main            bool   `json:"main"`
		Seo             struct {
			MetaTitle       string `json:"metaTitle"`
			MetaDescription string `json:"metaDescription"`
		} `json:"seo"`
		Visibility int `json:"visibility"`
		Type       int `json:"type"`
		Product    struct {
			ID               int    `json:"id"`
			UniquePath       string `json:"uniquePath"`
			Title            string `json:"title"`
			Description      string `json:"description"`
			ShortDescription string `json:"shortDescription"`
			OnHold           bool   `json:"onHold"`
			Terms            bool   `json:"terms"`
			Variants         []struct {
				ID           int    `json:"id"`
				Title        string `json:"title"`
				PriceDetails struct {
					Amount   int    `json:"amount"`
					Currency string `json:"currency"`
				} `json:"priceDetails"`
				ProductType int `json:"productType"`
				AsDynamic   struct {
					ExternalURL string `json:"externalUrl"`
					MinAmount   int    `json:"minAmount"`
					IsInternal  bool   `json:"isInternal"`
					MaxAmount   int    `json:"maxAmount"`
				} `json:"asDynamic"`
				Gateways []struct {
					Gateway int `json:"gateway"`
				} `json:"gateways"`
				CustomFields []struct {
					ID          int    `json:"id"`
					Type        int    `json:"type"`
					Name        string `json:"name"`
					Required    bool   `json:"required"`
					ValueString string `json:"valueString"`
				} `json:"customFields"`
			} `json:"variants"`
			CreatedAt time.Time `json:"createdAt"`
		} `json:"product"`
	} `json:"data"`
}

type SellpassProductUpdate struct {
	Title            string `json:"title"`
	ShortDescription string `json:"shortDescription"`
	Description      string `json:"description"`
	Variants         []struct {
		ID               int    `json:"id"`
		Title            string `json:"title"`
		Description      string `json:"description"`
		ShortDescription string `json:"shortDescription"`
		PriceDetails     struct {
			Amount   float64 `json:"amount"`
			Currency string  `json:"currency"`
		} `json:"priceDetails"`
		Gateways []struct {
			Gateway int `json:"gateway"`
			Rules   struct {
				BlockVpn bool `json:"blockVpn"`
			} `json:"rules"`
			Price struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"price"`
		} `json:"gateways"`
		ProductType int `json:"productType"`
		AsDynamic   struct {
			Stock       int    `json:"stock"`
			ExternalURL string `json:"externalUrl"`
			MinAmount   int    `json:"minAmount"`
			MaxAmount   int    `json:"maxAmount"`
			IsInternal  bool   `json:"isInternal"`
		} `json:"asDynamic"`
		AsSerials struct {
			Delimiter        string `json:"delimiter"`
			Serials          string `json:"serials"`
			MinAmount        int    `json:"minAmount"`
			MaxAmount        int    `json:"maxAmount"`
			RemoveDuplicates bool   `json:"removeDuplicates"`
		} `json:"asSerials"`
		AsService struct {
			Stock     int    `json:"stock"`
			Text      string `json:"text"`
			MinAmount int    `json:"minAmount"`
			MaxAmount int    `json:"maxAmount"`
		} `json:"asService"`
		CustomerNote string `json:"customerNote"`
		RedirectURL  string `json:"redirectUrl"`
		CustomFields []struct {
			ID          int    `json:"id"`
			Type        int    `json:"type"`
			Name        string `json:"name"`
			Required    bool   `json:"required"`
			ValueString string `json:"valueString"`
			Placeholder string `json:"placeholder"`
			Regex       string `json:"regex"`
			ValueInt    int    `json:"valueInt"`
			ValueBool   bool   `json:"valueBool"`
		} `json:"customFields"`
		Warranty struct {
			Text            string `json:"text"`
			DurationSeconds int    `json:"durationSeconds"`
		} `json:"warranty"`
		DiscordSocialConnectSettings struct {
			Enabled                    bool `json:"enabled"`
			Required                   bool `json:"required"`
			BeforePurchaseRequireRoles struct {
				GuildID string   `json:"guildId"`
				RoleIds []string `json:"roleIds"`
			} `json:"beforePurchaseRequireRoles"`
			BeforePurchaseServer struct {
				GuildID string   `json:"guildId"`
				RoleIds []string `json:"roleIds"`
			} `json:"beforePurchaseServer"`
			AfterPurchaseServer struct {
				GuildID string   `json:"guildId"`
				RoleIds []string `json:"roleIds"`
			} `json:"afterPurchaseServer"`
		} `json:"discordSocialConnectSettings"`
	} `json:"variants"`
	Path string `json:"path"`
	Seo  struct {
		MetaTitle       string `json:"metaTitle"`
		MetaDescription string `json:"metaDescription"`
	} `json:"seo"`
	Unlisted   bool `json:"unlisted"`
	Private    bool `json:"private"`
	OnHold     bool `json:"onHold"`
	IsInternal bool `json:"isInternal"`
}

type SellappOrderCompleted struct {
	Event string `json:"event,omitempty"`
	Data  struct {
		ID int `json:"id,omitempty"`
	}
}

type SellappOrderInfo struct {
	Data struct {
		ID      int `json:"id,omitempty"`
		Payment struct {
			Fee struct {
				Base     string `json:"base,omitempty"`
				Currency string `json:"currency,omitempty"`
				Units    int    `json:"units,omitempty"`
				Vat      int    `json:"vat,omitempty"`
				Total    struct {
					Exclusive string `json:"exclusive,omitempty"`
					Inclusive string `json:"inclusive,omitempty"`
				} `json:"total,omitempty"`
			} `json:"fee,omitempty"`
			Gateway struct {
				Data struct {
					Total struct {
						Base     string `json:"base,omitempty"`
						Currency string `json:"currency,omitempty"`
						Units    int    `json:"units,omitempty"`
						Vat      int    `json:"vat,omitempty"`
						Total    struct {
							Exclusive string `json:"exclusive,omitempty"`
							Inclusive string `json:"inclusive,omitempty"`
						} `json:"total,omitempty"`
					} `json:"total,omitempty"`
					CustomerEmail string `json:"customer_email,omitempty"`
					TransactionID string `json:"transaction_id,omitempty"`
				} `json:"data,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"gateway,omitempty"`
			Subtotal struct {
				Base     string `json:"base,omitempty"`
				Currency string `json:"currency,omitempty"`
				Units    int    `json:"units,omitempty"`
				Vat      int    `json:"vat,omitempty"`
				Total    struct {
					Exclusive string `json:"exclusive,omitempty"`
					Inclusive string `json:"inclusive,omitempty"`
				} `json:"total,omitempty"`
			} `json:"subtotal,omitempty"`
			ExpiresAt time.Time `json:"expires_at,omitempty"`
			FullPrice struct {
				Base     string `json:"base,omitempty"`
				Currency string `json:"currency,omitempty"`
				Units    int    `json:"units,omitempty"`
				Vat      int    `json:"vat,omitempty"`
				Total    struct {
					Exclusive string `json:"exclusive,omitempty"`
					Inclusive string `json:"inclusive,omitempty"`
				} `json:"total,omitempty"`
			} `json:"full_price,omitempty"`
			OriginalAmount struct {
				Base     string `json:"base,omitempty"`
				Currency string `json:"currency,omitempty"`
				Units    int    `json:"units,omitempty"`
				Vat      int    `json:"vat,omitempty"`
				Total    struct {
					Exclusive string `json:"exclusive,omitempty"`
					Inclusive string `json:"inclusive,omitempty"`
				} `json:"total,omitempty"`
			} `json:"original_amount,omitempty"`
		} `json:"payment,omitempty"`
		Status struct {
			History []struct {
				SetAt     time.Time `json:"setAt,omitempty"`
				Status    string    `json:"status,omitempty"`
				UpdatedAt time.Time `json:"updatedAt,omitempty"`
			} `json:"history,omitempty"`
			Status struct {
				SetAt     time.Time `json:"setAt,omitempty"`
				Status    string    `json:"status,omitempty"`
				UpdatedAt time.Time `json:"updatedAt,omitempty"`
			} `json:"status,omitempty"`
		} `json:"status,omitempty"`
		Webhooks            []any     `json:"webhooks,omitempty"`
		Feedback            string    `json:"feedback,omitempty"`
		CreatedAt           time.Time `json:"created_at,omitempty"`
		UpdatedAt           time.Time `json:"updated_at,omitempty"`
		StoreID             int       `json:"store_id,omitempty"`
		CouponID            any       `json:"coupon_id,omitempty"`
		SubscriptionID      any       `json:"subscription_id,omitempty"`
		CustomerInformation struct {
			ID           int    `json:"id,omitempty"`
			Email        string `json:"email,omitempty"`
			Country      string `json:"country,omitempty"`
			Location     string `json:"location,omitempty"`
			IP           string `json:"ip,omitempty"`
			Proxied      bool   `json:"proxied,omitempty"`
			BrowserAgent string `json:"browser_agent,omitempty"`
			Vat          struct {
				Amount  int    `json:"amount,omitempty"`
				Country string `json:"country,omitempty"`
			} `json:"vat,omitempty"`
		} `json:"customer_information,omitempty"`
		Products []struct {
			ID       int    `json:"id,omitempty"`
			Title    string `json:"title,omitempty"`
			URL      string `json:"url,omitempty"`
			Variants []struct {
				ID                    int    `json:"id,omitempty"`
				Title                 string `json:"title,omitempty"`
				Quantity              int    `json:"quantity,omitempty"`
				AdditionalInformation []struct {
					Key   string `json:"key,omitempty"`
					Label string `json:"label,omitempty"`
					Value string `json:"value,omitempty"`
				} `json:"additional_information,omitempty"`
			} `json:"variants,omitempty"`
		} `json:"products,omitempty"`
	} `json:"data,omitempty"`
}
