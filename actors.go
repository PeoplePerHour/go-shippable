package shippable

// Person holds info of a Shippable API user and/or Project contributor.
type Person struct {
	AvatarURL   *string `json:"avatarUrl"`
	DisplayName *string `json:"displayName"`
	Email       *string `json:"email"`
	Login       *string `json:"login"`
}
