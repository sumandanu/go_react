package models

type News struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	TitleUrl      string `json:"titleUrl"`
	Site          string `json:"site"`
	SiteUrl       string `json:"siteUrl"`
	Score         string `json:"score"`
	ByUser        string `json:"byUser"`
	ByUrl         string `json:"byUrl"`
	Age           string `json:"age"`
	AgeUrl        string `json:"ageUrl"`
	HideUrl       string `json:"hideUrl"`
	PastUrl       string `json:"pastUrl"`
	CommentsUrl   string `json:"commentsUrl"`
  }