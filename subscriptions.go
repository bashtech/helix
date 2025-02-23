package helix

// Subscription ...
type Subscription struct {
	BroadcasterID    string `json:"broadcaster_id"`
	BroadcasterLogin string `json:"broadcaster_login"`
	BroadcasterName  string `json:"broadcaster_name"`
	IsGift           bool   `json:"is_gift"`
	GifterID         string `json:"gifter_id"`
	GifterLogin      string `json:"gifter_login"`
	GifterName       string `json:"gifter_name"`
	Tier             string `json:"tier"`
	PlanName         string `json:"plan_name"`
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	UserLogin        string `json:"user_login"`
}

// UserSubscription ...
type UserSubscription struct {
	BroadcasterID    string `json:"broadcaster_id"`
	BroadcasterLogin string `json:"broadcaster_login"`
	BroadcasterName  string `json:"broadcaster_name"`
	IsGift           bool   `json:"is_gift"`
	GifterLogin      string `json:"gifter_login"`
	GifterName       string `json:"gifter_name"`
	Tier             string `json:"tier"`
}

// ManySubscriptions ...
type ManySubscriptions struct {
	Subscriptions []Subscription `json:"data"`
	Pagination    Pagination     `json:"pagination"`
	Total         int            `json:"total"`
}

// ManyUserSubscription ...
type ManyUserSubscriptions struct {
	UserSubscriptions []UserSubscription `json:"data"`
}

// SubscriptionsResponse ...
type SubscriptionsResponse struct {
	ResponseCommon
	Data ManySubscriptions
}

// UserSubscriptionResponse ...
type UserSubscriptionResponse struct {
	ResponseCommon
	Data ManyUserSubscriptions
}

// SubscriptionsParams ...
type SubscriptionsParams struct {
	BroadcasterID string   `query:"broadcaster_id"` // Limit 1
	UserID        []string `query:"user_id"`        // Limit 100
	After         string   `query:"after"`
	Before        string   `query:"before"`
	First         int      `query:"first,20"` // Limit 100
}

// UserSubscriptionsParams ...
type UserSubscriptionsParams struct {
	BroadcasterID string `query:"broadcaster_id"`
	UserID        string `query:"user_id"`
}

// GetSubscriptions gets subscriptions about one Twitch broadcaster.
// Broadcasters can only request their own subscriptions.
//
// Required scope: channel:read:subscriptions
func (c *Client) GetSubscriptions(params *SubscriptionsParams) (*SubscriptionsResponse, error) {
	resp, err := c.get("/subscriptions", &ManySubscriptions{}, params)
	if err != nil {
		return nil, err
	}

	subscriptions := &SubscriptionsResponse{}
	resp.HydrateResponseCommon(&subscriptions.ResponseCommon)
	subscriptions.Data.Subscriptions = resp.Data.(*ManySubscriptions).Subscriptions
	subscriptions.Data.Pagination = resp.Data.(*ManySubscriptions).Pagination

	return subscriptions, nil
}

// CheckUserSubsription Check if a specific user is subscribed to a specific channel
//
// Required scope: user:read:subscriptions
func (c *Client) CheckUserSubsription(params *UserSubscriptionsParams) (*UserSubscriptionResponse, error) {
	resp, err := c.get("/subscriptions/user", &ManyUserSubscriptions{}, params)
	if err != nil {
		return nil, err
	}

	subscriptions := &UserSubscriptionResponse{}
	resp.HydrateResponseCommon(&subscriptions.ResponseCommon)
	subscriptions.Data.UserSubscriptions = resp.Data.(*ManyUserSubscriptions).UserSubscriptions

	return subscriptions, nil
}
