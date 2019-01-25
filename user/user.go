package user

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserProfile struct {
	OauthID            string      `json:"-"`
	Email              string      `json:"-"`
	Username           string      `json:"-"`
	Name               string      `json:"name"`
	Partner            string      `json:"partner"`
	City               string      `json:"city"`
	State              string      `json:"state"`
	Private            bool        `json:"private"`
	UserSocialMedia    SocialMedia `json:"userSocialMedia"`
	PartnerSocialMedia SocialMedia `json:"partnerSocialMedia"`
}

type SocialMedia struct {
	Facebook SocialMediaProfile `json:"facebook"`
	Twitter  SocialMediaProfile `json:"twitter"`
	Snapchat SocialMediaProfile `json:"snapchat"`
}

type SocialMediaProfile struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Private  bool   `json:"private"`
}
