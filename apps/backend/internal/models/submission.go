package models

type Submission struct {
	No            	string `json:"no"`
	Id            	string `json:"id"`
	VoteUrl       	string `json:"voteUrl"`
	Title         	string `json:"title"`
	TitleUrl      	string `json:"titleUrl"`
	Site          	string `json:"site"`
	SiteUrl       	string `json:"siteUrl"`
	Score         	string `json:"score"`
	ByUser        	string `json:"byUser"`
	ByUrl         	string `json:"byUrl"`
	Age           	string `json:"age"`
	AgeUrl        	string `json:"ageUrl"`
	HideUrl       	string `json:"hideUrl"`
	PastUrl       	string `json:"pastUrl"`
	FavoriteUrl   	string `json:"favoriteUrl"`
	Comments   	  	string `json:"comments"`
	CommentsUrl   	string `json:"commentsUrl"`
	FormAction   	string `json:"formAction"`
	FormParentVal   string `json:"formParentVal"`
	FormGotoVal   	string `json:"formGotoVal"`
	FormHmacVal   	string `json:"formHmacVal"`
}