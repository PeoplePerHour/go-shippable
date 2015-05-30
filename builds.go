package shippable

type Build struct {
	BaseCommitRef              *string        `json:"baseCommitRef"`
	BeforeCommitSha            *string        `json:"beforeCommitSha"`
	Branch                     *string        `json:"branch"`
	BranchCoveragePercent      float64        `json:"branchCoveragePercent"`
	BranchHead                 *string        `json:"branchHead"`
	BuildGroupNumber           *int           `json:"buildGroupNumber"`
	BuildItemStepOrder         *[]interface{} `json:"buildItemStepOrder"`
	BuildRunnerVersion         *string        `json:"buildRunnerVersion"`
	Builds                     *[]BaseBuild   `json:"builds"`
	CommitSha                  *string        `json:"commitSha"`
	CommitURL                  *string        `json:"commitUrl"`
	Committer                  Person         `json:"committer"`
	CompareURL                 *string        `json:"compareUrl"`
	CreatedByAccountID         *string        `json:"createdByAccountId"`
	CreatedDate                *string        `json:"createdDate"`
	DurationCumulative         *string        `json:"durationCumulative"`
	EmailNotifications         *[]interface{} `json:"emailNotifications"`
	EmailOnFailure             *string        `json:"emailOnFailure"`
	EmailOnSuccess             *string        `json:"emailOnSuccess"`
	ID                         *string        `json:"id"`
	ImageID                    *string        `json:"imageId"`
	ImageName                  *string        `json:"imageName"`
	IsAutoBuild                *bool          `json:"isAutoBuild"`
	IsAutoCommit               *bool          `json:"isAutoCommit"`
	IsAutoPush                 *bool          `json:"isAutoPush"`
	IsCompleted                *bool          `json:"isCompleted"`
	IsPullRequest              *bool          `json:"isPullRequest"`
	IsReRun                    *bool          `json:"isReRun"`
	Language                   *string        `json:"language"`
	LastAuthor                 *Person        `json:"lastAuthor"`
	LastCommitShortDescription *string        `json:"lastCommitShortDescription"`
	Network                    *string        `json:"network"`
	Notifications              *Notifications `json:"notifications"`
	ParallelizedTest           *bool          `json:"parallelizedTest"`
	Privileged                 *bool          `json:"privileged"`
	ProjectID                  *string        `json:"projectId"`
	PullRequestNumber          *int           `json:"pullRequestNumber"`
	RepositoryFileCount        *interface{}   `json:"repositoryFileCount"`
	RepositorySize             *interface{}   `json:"repositorySize"`
	RequiresDedicatedHost      *bool          `json:"requiresDedicatedHost"`
	RunCommand                 *string        `json:"runCommand"`
	SequenceCoveragePercent    *int           `json:"sequenceCoveragePercent"`
	Settings                   *BuildSettings `json:"settings"`
	ShouldArchive              *bool          `json:"shouldArchive"`
	ShouldDecryptSecureEnvs    *bool          `json:"shouldDecryptSecureEnvs"`
	Status                     *int           `json:"status"`
	StatusMessage              *string        `json:"statusMessage"`
	TestsFailed                *int           `json:"testsFailed"`
	TestsPassed                *int           `json:"testsPassed"`
	TestsSkipped               *int           `json:"testsSkipped"`
	TimeoutMS                  *int           `json:"timeoutMS"`
	TotalTests                 *int           `json:"totalTests"`
	TriggeredBy                *Person        `json:"triggeredBy"`
	UpdatedDate                *string        `json:"updatedDate"`
}

type BuildSteps struct {
	Build  *BuildStep `json:"build"`
	Commit *BuildStep `json:"commit"`
	Pull   *BuildStep `json:"pull"`
	Report *BuildStep `json:"report"`
	Upload *BuildStep `json:"upload"`
}

type BaseBuild struct {
	_id                     *string         `json:"_id"`
	BranchCoveragePercent   *int            `json:"branchCoveragePercent"`
	BuildNumber             *int            `json:"buildNumber"`
	CommitTag               *string         `json:"commitTag"`
	ConsoleLogBytes         *int            `json:"consoleLogBytes"`
	ConsoleLogLineCount     *int            `json:"consoleLogLineCount"`
	DeprovisionStatusDate   *string         `json:"deprovisionStatusDate"`
	Duration                *int            `json:"duration"`
	EndDate                 *string         `json:"endDate"`
	Environment             *string         `json:"environment"`
	Gemfile                 *string         `json:"gemfile"`
	ImageCommitStatusDate   *string         `json:"imageCommitStatusDate"`
	IsBuildCompleted        *bool           `json:"isBuildCompleted"`
	IsCompleted             *bool           `json:"isCompleted"`
	IsFailureAllowed        *bool           `json:"isFailureAllowed"`
	IsSubscriptionHost      *bool           `json:"isSubscriptionHost"`
	Jdk                     *string         `json:"jdk"`
	MatrixValues            *[]MatrixResult `json:"matrixValues"`
	QueuedDate              *string         `json:"queuedDate"`
	SequenceCoveragePercent *int            `json:"sequenceCoveragePercent"`
	Size                    interface{}     `json:"size"`
	StartDate               *string         `json:"startDate"`
	Status                  *int            `json:"status"`
	Steps                   *BuildSteps     `json:"steps"`
	TestsFailed             *int            `json:"testsFailed"`
	TestsPassed             *int            `json:"testsPassed"`
	TestsSkipped            *int            `json:"testsSkipped"`
	TotalTests              *int            `json:"totalTests"`
	Version                 *string         `json:"version"`
}

type BuildStep struct {
	Duration  *int      `json:"duration"`
	EndTime   *string   `json:"endTime"`
	Report    *[]Report `json:"report"`
	StartTime *string   `json:"startTime"`
}

type Report struct {
	Status *int    `json:"status"`
	Time   *string `json:"time"`
}

type BuildSettings struct {
	ImageID      *string `json:"imageId"`
	ImageOptions *struct {
		NetworkMode *string `json:"networkMode"`
		Privileged  *bool   `json:"privileged"`
	} `json:"imageOptions"`
	PullImageName         *string `json:"pullImageName"`
	RequiresDedicatedHost *bool   `json:"requiresDedicatedHost"`
	RunCommand            *string `json:"runCommand"`
}

type MatrixResult struct {
	_id   *string `json:"_id"`
	Name  *string `json:"name"`
	Value *string `json:"value"`
}
