package shippable

type Notifications struct {
	Irc IrcNotification `json:"irc"`
}

type IrcNotification struct {
	Channels             []interface{} `json:"channels"`
	NotifyOnPullRequests bool          `json:"notifyOnPullRequests"`
}
