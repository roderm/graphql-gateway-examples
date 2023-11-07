package data

import "github.com/AlekSi/pointer"

type User interface {
	IsEntity()
}

type AWSUser struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	AwsID    *string `json:"awsID"`
}

func (AWSUser) IsUser()   {}
func (AWSUser) IsEntity() {}

type GitHubUser struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	GithubID *string `json:"githubID"`
}

func (GitHubUser) IsUser()   {}
func (GitHubUser) IsEntity() {}

var User1 = &GitHubUser{
	ID:       "1",
	Username: "tony",
	GithubID: pointer.ToString("gh_abc"),
}
var User2 = &AWSUser{
	ID:       "2",
	Username: "jeff",
	AwsID:    pointer.ToString("aw_abc"),
}
