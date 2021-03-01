package main

import (
	"fmt"
	"os"
	"time"

	usabilla "github.com/cxomni/api-go"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: test <key> <secret>")
		return
	}

	key := os.Args[1]
	secret := os.Args[2]

	// Pass the key and secret which should be defined as ENV vars
	usabilla := usabilla.New(key, secret, nil)

	resource := usabilla.Campaigns()

	// Get the first ten buttons
	// params := map[string]string{"limit": "90"}
	//buttons, err := usabilla.InpageWidgets().Get(params)

	/*if err != nil {
		fmt.Printf("ERROR: %#v", err)
	} else {
		fmt.Printf("InPage: %#v\n", buttons)
	}*/

	campaigns, err := resource.Get(map[string]string{
		"limit": "100",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if campaigns.Count == 0 {
		fmt.Println("no campaigns")
		return
	}

	for _, campaign := range campaigns.Items {
		fmt.Printf("-----\nID: %s, %s\n", campaign.ID, campaign.Name)

		// params := map[string]string{"since": "1582538940000"}
		params := map[string]string{"since": fmt.Sprintf("%d", time.Now().Unix()*1000-90*86400000)}
		feedback, err := resource.Results().Get(campaign.ID, params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("feedback.length: ", len(feedback.Items))
		if len(feedback.Items) > 0 {
			fmt.Printf("Feedback for button %s:\n", campaign.ButtonID)
			for _, feedback := range feedback.Items {
				/*if feedback.ID == "5e53a45bb96ea1031d0114cd" {

									fmt.Printf("%#v\n", feedback)
				        }*/
				fmt.Printf("%s: %#v\n", feedback.ID, feedback)
			}
		}
	}
}
