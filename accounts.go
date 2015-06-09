package shippable

// GetAccounts returns a string list of your Shippable account IDs.
func (a *AccountService) GetAccounts() (accounts *[]string, resp *Response, err error) {
	req, _ := a.client.NewRequest("GET", "accounts", accounts)
	accounts = new([]string)
	resp, err = a.client.Do(req, accounts)
	return
}

// GetAccount returns information of a specific Shippable account ID.
func (a *AccountService) GetAccount(id string) (account *Account, resp *Response, err error) {
	req, err := a.client.NewRequest("GET", "accounts/"+id, nil)
	if err != nil {
		return nil, nil, err
	}
	account = new(Account)
	resp, err = a.client.Do(req, account)
	return
}

// DeleteAccount deletes the specified Shippable account ID.
func (a *AccountService) DeleteAccount(id string) (ok bool, resp *Response, err error) {
	req, err := a.client.NewRequest("DELETE", "accounts/"+id, nil)
	if err != nil {
		return false, nil, err
	}
	resp, err = a.client.Do(req, nil)
	if resp.StatusCode < 300 {
		ok = true
	}
	return
}

// GetAccountIdentities returns a string list of identity ids associated with this account.
// Your account can have multiple identities. There will always be at least one identity, and
// that is the identity of your linked github or bitbucket account. Another identity your account
// could take is the identity of an organization it belongs to.
func (a *AccountService) GetAccountIdentities(id string) (identities *[]string, resp *Response, err error) {
	req, err := a.client.NewRequest("GET", "accounts/"+id+"/identities", nil)
	if err != nil {
		return nil, nil, err
	}
	identities = new([]string)
	resp, err = a.client.Do(req, identities)
	return
}

// Account is the Shippable API user entity.
type Account struct {
	AccountAuthorizationLastSyncEndDate   *string     `json:"accountAuthorizationLastSyncEndDate"`
	AccountAuthorizationLastSyncStartDate *string     `json:"accountAuthorizationLastSyncStartDate"`
	BraintreeCustomerID                   *string     `json:"braintreeCustomerId"`
	Created                               *string     `json:"created"`
	ID                                    *string     `json:"id"`
	Identities                            *[]Identity `json:"identities"`
	LastAccountSyncEndDate                *string     `json:"lastAccountSyncEndDate"`
	LastAccountSyncStartDate              *string     `json:"lastAccountSyncStartDate"`
	LastUsedIdentityID                    *string     `json:"lastUsedIdentityId"`
	SystemRoles                           *[]string   `json:"systemRoles"`
}

// Identity of a Shippable API user with Providers like Github, BitBucket etc.
type Identity struct {
	AvatarID                  *string   `json:"avatarId"`
	AvatarURL                 *string   `json:"avatarUrl"`
	DisplayName               *string   `json:"displayName"`
	Email                     *string   `json:"email"`
	Emails                    *[]Email  `json:"emails"`
	EnforceScopes             *[]string `json:"enforceScopes"`
	ID                        *string   `json:"id"`
	MigratedProviderID        bool      `json:"migratedProviderId"`
	Provider                  *string   `json:"provider"`
	ProviderBlog              *string   `json:"providerBlog"`
	ProviderCompany           *string   `json:"providerCompany"`
	ProviderFollowerCount     int       `json:"providerFollowerCount"`
	ProviderID                *string   `json:"providerId"`
	ProviderLocation          *string   `json:"providerLocation"`
	ProviderOwnedPrivateRepos int       `json:"providerOwnedPrivateRepos"`
	ProviderPrivateGists      int       `json:"providerPrivateGists"`
	ProviderPublicGistCount   int       `json:"providerPublicGistCount"`
	ProviderPublicRepoCount   int       `json:"providerPublicRepoCount"`
	ProviderTotalPrivateRepos int       `json:"providerTotalPrivateRepos"`
	ProviderType              *string   `json:"providerType"`
	Scopes                    *[]string `json:"scopes"`
	Username                  *string   `json:"userName"`
}

// Email of a Shippable API user.
type Email struct {
	Email    *string `json:"email"`
	Primary  bool    `json:"primary"`
	Verified bool    `json:"verified"`
}
