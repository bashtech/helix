package helix

// GetChatBadgeParams ...
type GetChatBadgeParams struct {
	BroadcasterID string `query:"broadcaster_id"`
}

// GetChatBadgeResponse ...
type GetChatBadgeResponse struct {
	ResponseCommon
	Data ManyChatBadge
}

// ManyChatBadge ...
type ManyChatBadge struct {
	Badges []ChatBadge `json:"data"`
}

// ChatBadge ...
type ChatBadge struct {
	SetID    string         `json:"set_id"`
	Versions []BadgeVersion `json:"versions"`
}

// BadgeVersion ...
type BadgeVersion struct {
	ID         string `json:"id"`
	ImageUrl1x string `json:"image_url_1x"`
	ImageUrl2x string `json:"image_url_2x"`
	ImageUrl4x string `json:"image_url_4x"`
}

// GetChatEmoteParams ...
type GetChatEmoteParams struct {
	BroadcasterID string `query:"broadcaster_id"`
}

// GetChatSetEmoteParams ...
type GetChatSetEmoteParams struct {
	EmoteSetID string `query:"emote_set_id"`
}

// GetChatEmoteResponse ...
type GetChatEmoteResponse struct {
	ResponseCommon
	Data ManyChatEmote
}

// ManyChatEmote ...
type ManyChatEmote struct {
	Emotes []ChatEmote `json:"data"`
}

// ChatEmote ...
type ChatEmote struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	Versions   []EmoteImages  `json:"images"`
	Tier       string         `json:"tier"`
	EmoteType  string         `json:"emote_type"`
	EmoteSetID string         `json:"emote_set_id"`
}

// BadgeVersion ...
type EmoteImages struct {
	Url1x string `json:"url_1x"`
	Url2x string `json:"url_2x"`
	Url4x string `json:"url_4x"`
}


// GetChannelChatBadges ...
func (c *Client) GetChannelChatBadges(params *GetChatBadgeParams) (*GetChatBadgeResponse, error) {
	resp, err := c.get("/chat/badges", &ManyChatBadge{}, params)
	if err != nil {
		return nil, err
	}

	channels := &GetChatBadgeResponse{}
	resp.HydrateResponseCommon(&channels.ResponseCommon)
	channels.Data.Badges = resp.Data.(*ManyChatBadge).Badges

	return channels, nil
}

// GetGlobalChatBadges ...
func (c *Client) GetGlobalChatBadges() (*GetChatBadgeResponse, error) {
	resp, err := c.get("/chat/badges/global", &ManyChatBadge{}, nil)
	if err != nil {
		return nil, err
	}

	channels := &GetChatBadgeResponse{}
	resp.HydrateResponseCommon(&channels.ResponseCommon)
	channels.Data.Badges = resp.Data.(*ManyChatBadge).Badges

	return channels, nil
}

// GetGlobalChatEmotes ...
func (c *Client) GetGlobalChatEmotes() (*GetChatEmoteResponse, error) {
	resp, err := c.get("/chat/emotes/global", &ManyChatEmote{}, nil)
	if err != nil {
		return nil, err
	}

	channels := &GetChatEmoteResponse{}
	resp.HydrateResponseCommon(&channels.ResponseCommon)
	channels.Data.Emotes = resp.Data.(*ManyChatEmote).Emotes

	return channels, nil
}

// GetChannelChatEmotes ...
func (c *Client) GetChannelChatEmotes(params *GetChatEmoteParams) (*GetChatEmoteResponse, error) {
	resp, err := c.get("/chat/emotes", &ManyChatEmote{}, params)
	if err != nil {
		return nil, err
	}

	channels := &GetChatEmoteResponse{}
	resp.HydrateResponseCommon(&channels.ResponseCommon)
	channels.Data.Emotes = resp.Data.(*ManyChatEmote).Emotes

	return channels, nil
}

// GetSetChatEmotes ...
func (c *Client) GetSetChatEmotes(params *GetChatSetEmoteParams) (*GetChatEmoteResponse, error) {
	resp, err := c.get("/chat/emotes/set", &ManyChatEmote{}, params)
	if err != nil {
		return nil, err
	}

	channels := &GetChatEmoteResponse{}
	resp.HydrateResponseCommon(&channels.ResponseCommon)
	channels.Data.Emotes = resp.Data.(*ManyChatEmote).Emotes

	return channels, nil
}