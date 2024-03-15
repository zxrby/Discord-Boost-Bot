package Discord

import (
	"BoostTool/Core/Utils"
	"errors"
	"sync"
	"time"
)

var config, _ = Utils.LoadConfig()
var mutex sync.Mutex

func BoostServer(invite string, amount int, file string) (*BoostResponse, error) {
	Tokens := ReloadFiles(file)

	if len(Tokens.List)*2 < amount {
		Utils.LogError("Not Enough Tokens to Boost", "", "")
		return nil, errors.New("Not Enough Tokens to Boost")
	}

	var wg sync.WaitGroup
	conc := make(chan struct{}, 1000)

	success := 0
	failed := 0

	var successTokens []string
	var failedTokens []string
	var tokensUsed []string
	times := time.Now()

	for i := 0; i < int(amount)/2; i++ {
		wg.Add(1)
		go func(token string) {
			defer wg.Done()
			conc <- struct{}{}

			tokensUsed = append(tokensUsed, token)

			c := New(token)
			err := c.GetRequiredElements()
			if err != nil {
				failed++
				failedTokens = append(failedTokens, token+" Error: "+err.Error())
				Utils.AppendTextToFile(token+"\n", "failed.txt")
				return
			}
			err = c.IsValidInvite(invite)
			if err != nil {
				failed++
				failedTokens = append(failedTokens, token+" Error: "+err.Error())
				Utils.AppendTextToFile(token+"\n", "failed.txt")
				return
			}

			err = c.JoinServer(invite)
			if err != nil {
				failed++
				failedTokens = append(failedTokens, token+" Error: "+err.Error())
				Utils.AppendTextToFile(token+"\n", "failed.txt")
				return
			}

			go Utils.RemoveToken(token, file)
			go Utils.AppendTextToFile(token+"\n", "used.txt")

			err = c.GetSubscriptionSlots()
			if err != nil {
				failed++
				failedTokens = append(failedTokens, token+" Error: "+err.Error())
				Utils.AppendTextToFile(token+"\n", "failed.txt")
				return
			}

			err = c.BoostServer()
			if err != nil {
				failed++
				failedTokens = append(failedTokens, token+" Error: "+err.Error())
				Utils.AppendTextToFile(token+"\n", "failed.txt")
				return
			}
			successTokens = append(successTokens, token)
			success++

			c.CustomizeTokens()

			<-conc
		}(Tokens.Next())

	}
	wg.Wait()
	ElapsedTime := time.Since(times) - 2

	_ = ReloadFiles(file)

	return &BoostResponse{
		Success:       success,
		Failed:        failed,
		SuccessTokens: successTokens,
		FailedTokens:  failedTokens,
		ElapsedTime:   ElapsedTime,
	}, nil
}
