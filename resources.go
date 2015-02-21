package gobilla

import (
	"encoding/json"
	"fmt"
)

/*
Canonical URI constants.
*/
const (
	ButtonURI   = "/live/website/button"
	CampaignURI = "/live/website/campaign"
)

/*
Buttons represents the button resource of Usabilla API.
*/
type Buttons struct {
	Feedback FeedbackItem
}

/*
Get function of Buttons resource returns all the buttons
taking into account the specified query params.

Accepted query params are:

- limit string
*/
func (buttons *Buttons) Get(params map[string]string) (ButtonData, error) {
	apiRequest := APIRequest{
		CanonicalURI: ButtonURI,
	}

	resp, err := apiRequest.Get(params)
	if err != nil {
		panic(err)
	}

	data := ButtonData{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

/*
FeedbackItem represents the feedback item resource of Usabilla API.
*/
type FeedbackItem struct {
}

/*
Get function of FeedbackItem resource returns all the feedback items
for a specific button, taking into account the passed query params.

Accepted query params are:

- since string (Time stamp)
*/
func (feedbackItem *FeedbackItem) Get(buttonID string, params map[string]string) (FeedbackData, error) {
	feedbackURI := fmt.Sprintf(ButtonURI+"/%s/feedback", buttonID)

	apiRequest := &APIRequest{
		CanonicalURI: feedbackURI,
	}

	resp, err := apiRequest.Get(params)
	if err != nil {
		panic(err)
	}

	data := FeedbackData{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

/*
Campaigns represents the campaign resource of Usabilla API.
*/
type Campaigns struct {
	Results CampaignResults
}

/*
Get function of Campaigns resource returns all the campaigns
taking into account the passed query params.

Accepted query params are:

- limit string
- since string (Time stamp)
*/
func (campaigns *Campaigns) Get(params map[string]string) (CampaignData, error) {
	apiRequest := APIRequest{
		CanonicalURI: CampaignURI,
	}

	resp, err := apiRequest.Get(params)
	if err != nil {
		panic(err)
	}

	data := CampaignData{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

/*
CampaignResults represents the campaign result resource of Usabilla API.
*/
type CampaignResults struct {
}

/*
Get function of CampaignResults resource returns all the campaign result items
for a specific campaign, taking into account the passed query params.

Accepted query params are:

- limit int
- since string (Time stamp)
*/
func (campaignResults *CampaignResults) Get(campaignID string, params map[string]string) (CampaignResultData, error) {
	campaignURI := fmt.Sprintf(CampaignURI+"/%s/results", campaignID)

	apiRequest := APIRequest{
		CanonicalURI: campaignURI,
	}

	resp, err := apiRequest.Get(params)
	if err != nil {
		panic(err)
	}

	data := CampaignResultData{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
