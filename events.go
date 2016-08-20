package lyfeed

type ItemCreated struct {
	GUID   string
	FeedId int
	Raw    string
}

type CommentCreated struct {
	UserId int
	Text   string
}

type FeedCreated struct {
	UsedId int
	URL    string
	Name   string
}
