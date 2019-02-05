package user

import "time"

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserProfile struct {
	Oauth            string      `json:"oauth" bson:"oauth"`
	Email            string      `json:"email" bson:"email"`
	Image            string      `json:"image" bson:"image"`
	Username         string      `json:"username" bson:"username"`
	Name             string      `json:"name" bson:"name"`
	Partner          string      `json:"partner" bson:"partner"`
	Location         string      `json:"location" bson:"location"`
	Private          bool        `json:"private" bson:"private"`
	IsMember         bool        `json:"isMember" bson:"isMember"`
	PartnerFacebook  string      `json:"partnerFacebook" bson:"partnerFacebook"`
	PartnerTwitter   string      `json:"partnerTwitter" bson:"partnerTwitter"`
	PartnerInstagram string      `json:"partnerInstagram" bson:"partnerInstagram"`
	MemberFacebook   string      `json:"memberFacebook" bson:"memberFacebook"`
	MemberTwitter    string      `json:"memberTwitter" bson:"memberTwitter"`
	MemberInstagram  string      `json:"memberInstagram" bson:"memberInstagram"`
	SocialMedia      SocialMedia `json:"socialMedia" bson:"socialMedia"`
}

type Membership struct {
	IsMember       bool      `json:"isMember" bson:"isMember"`
	MemberShipDate time.Time `json:"membershipDate" bson:"membershipDate"`
	ExpirationDate time.Time `json:"expirationDate" bson:"expirationDate"`
}

type SocialMedia struct {
	Facebook  SocialMediaProfile `json:"facebook" bson:"facebook"`
	Twitter   SocialMediaProfile `json:"twitter" bson:"twitter"`
	Instagram SocialMediaProfile `json:"snapchat" bson:"instagram"`
}

type SocialMediaProfile struct {
	Username string `json:"username" bson:"username"`
	Private  bool   `json:"private" bson:"private"`
}
