package models

type ScrapeResult[T any] struct {
	Data     	T   	`json:"data"`
	NextLink  	string	`json:"nextLink"`
}