package shippable

// Notifications are used to pingback external services about the outcome of a build.
type Notifications struct {
	IRC IRCNotification `json:"irc"`
}

// IRCNotification is used to report the outcome of a build on IRC.
type IRCNotification struct {
	Channels             []interface{} `json:"channels"`
	NotifyOnPullRequests bool          `json:"notifyOnPullRequests"`
}
