package shippable

/*
func GetAccounts() []string                {}
func GetAccount(id string) *Account        {}
func DeleteAccount(id string) bool         {}
func (a *Account) GetIdentities() []string {}
*/

type Account struct {
	AccountAuthorizationLastSyncEndDate   string     `json:"accountAuthorizationLastSyncEndDate"`
	AccountAuthorizationLastSyncStartDate string     `json:"accountAuthorizationLastSyncStartDate"`
	BraintreeCustomerID                   string     `json:"braintreeCustomerId"`
	Created                               string     `json:"created"`
	ID                                    string     `json:"id"`
	Identities                            []Identity `json:"identities"`
	LastAccountSyncEndDate                string     `json:"lastAccountSyncEndDate"`
	LastAccountSyncStartDate              string     `json:"lastAccountSyncStartDate"`
	LastUsedIdentityID                    string     `json:"lastUsedIdentityId"`
	SystemRoles                           []string   `json:"systemRoles"`
}

type Identity struct {
	AvatarID                  string      `json:"avatarId"`
	AvatarURL                 string      `json:"avatarUrl"`
	DisplayName               string      `json:"displayName"`
	Email                     string      `json:"email"`
	Emails                    []Email     `json:"emails"`
	EnforceScopes             []string    `json:"enforceScopes"`
	ID                        string      `json:"id"`
	MigratedProviderID        bool        `json:"migratedProviderId"`
	Provider                  string      `json:"provider"`
	ProviderBlog              string      `json:"providerBlog"`
	ProviderCompany           string      `json:"providerCompany"`
	ProviderFollowerCount     int         `json:"providerFollowerCount"`
	ProviderID                string      `json:"providerId"`
	ProviderLocation          string      `json:"providerLocation"`
	ProviderOwnedPrivateRepos interface{} `json:"providerOwnedPrivateRepos"`
	ProviderPrivateGists      interface{} `json:"providerPrivateGists"`
	ProviderPublicGistCount   int         `json:"providerPublicGistCount"`
	ProviderPublicRepoCount   int         `json:"providerPublicRepoCount"`
	ProviderTotalPrivateRepos interface{} `json:"providerTotalPrivateRepos"`
	ProviderType              string      `json:"providerType"`
	Scopes                    []string    `json:"scopes"`
	UserName                  string      `json:"userName"`
}

type Email struct {
	Email    string `json:"email"`
	Primary  bool   `json:"primary"`
	Verified bool   `json:"verified"`
}
