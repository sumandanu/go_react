package models

type Comment struct {
	No            string `json:"no"`
	Id            string `json:"id"`
	VoteUrl       string `json:"voteUrl"`
	ByUser        string `json:"byUser"`
	ByUrl         string `json:"byUrl"`
	Age           string `json:"age"`
	AgeUrl        string `json:"ageUrl"`
	ParentUrl     string `json:"parentUrl"`
	ContextUrl    string `json:"contextUrl"`
	OnStory       string `json:"onStory"`
	OnStoryUrl    string `json:"onStoryUrl"`
	Comment       string `json:"comment"`
	NextUrl       string `json:"nextUrl"`
}