package user

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserProfile struct {
	Oauth              string      `json:"oauth" bson:"oauth"`
	Email              string      `json:"-"`
	Username           string      `json:"username" bson:"username"`
	Name               string      `json:"name" bson:"name"`
	Partner            string      `json:"partner" bson:"partner"`
	City               string      `json:"city" bson:"city"`
	State              string      `json:"state" bson:"state"`
	Private            bool        `json:"private" bson:"private"`
	UserSocialMedia    SocialMedia `json:"userSocialMedia" bson:"userSocialMedia"`
	PartnerSocialMedia SocialMedia `json:"partnerSocialMedia" bson:"partnerSocialMedia"`
}

type SocialMedia struct {
	Facebook SocialMediaProfile `json:"facebook" bson:"facebook"`
	Twitter  SocialMediaProfile `json:"twitter" bson:"twitter"`
	Snapchat SocialMediaProfile `json:"snapchat" bson:"snapchat"`
}

type SocialMediaProfile struct {
	Username string `json:"username" bson:"username"`
	Private  bool   `json:"private" bson:"private"`
}
