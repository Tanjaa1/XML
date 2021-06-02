package Enum

type NotificationType int
const (
	FOLLOW_REQUEST NotificationType = iota
	TAG
	LIKE
	DISLIKE
	COMMENT
)
