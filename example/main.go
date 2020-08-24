package main

import (
	"fmt"
	"os"

	usabilla "github.com/cxomni/api-go"
)

func buttons(usabilla *usabilla.Usabilla) {
	b := usabilla.Buttons()

	buttons, err := b.Get(nil)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	for _, button := range buttons.Items {
		resp, err := b.Feedback().Get(button.ID, nil)
		if err != nil {
			fmt.Errorf("%s", err)
		}
		count := 0
		fmt.Printf("START PRINTING FEEDBACK FOR BUTTON: %s\n", button.ID)
		for _, item := range resp.Items {
			fmt.Printf("FEEDBACK %s\n", item.ID)
			count++
		}
		fmt.Printf("RECEIVED %d FEEDBACK ITEMS\n", count)
	}
	fmt.Printf("RECEIVED FEEDBACK FROM %d BUTTONS\n", buttons.Count)
}

func buttonsIterator(usabilla *usabilla.Usabilla) {
	b := usabilla.Buttons()

	buttons, err := b.Get(nil)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	for _, button := range buttons.Items {
		count := 0
		fmt.Printf("START PRINTING FEEDBACK FOR BUTTON: %s\n", button.ID)
		for item := range b.Feedback().Iterate(button.ID, nil) {
			fmt.Printf("FEEDBACK %s\n", item.ID)
			count++
		}
		fmt.Printf("RECEIVED %d FEEDBACK ITEMS\n", count)
	}
	fmt.Printf("RECEIVED FEEDBACK FROM %d BUTTONS\n", buttons.Count)
}

func appCampaigns(usabilla *usabilla.Usabilla) {
	resource := usabilla.AppCampaigns()

	campaigns, err := resource.Get(nil)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	for _, campaign := range campaigns.Items {
		count := 0
		fmt.Printf("START PRINTING RESULTS FOR CAMPAIGN: %s\n", campaign.ID)
		for result := range resource.Results().Iterate(campaign.ID, nil) {
			fmt.Printf("Result %s\n", result.ID)
			count++
		}
		fmt.Printf("Received %d results\n", count)
	}

	fmt.Printf("FOUND %d CAMPAIGNS IN TOTAL\n", campaigns.Count)
}

func main() {
	key := os.Getenv("USABILLA_API_KEY")
	secret := os.Getenv("USABILLA_API_SECRET")

	// You can pass a custom http.Client
	// We pass nil to use the http.DefaultClient
	usabilla := usabilla.New(key, secret, nil)

	// Uses a simple GET to get all feedback items for all buttons.
	buttons(usabilla)

	// Uses a channel of feedback items, and once all items have been
	// consumed and the response HasMore then it fires a new request
	// for all feedback items for all buttons.
	buttonsIterator(usabilla)

	// Display App campaigns and their results
	appCampaigns(usabilla)
}
