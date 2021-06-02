package Enum

type Offence int
const (
	RACISM Offence = iota
	SEXUAL_CONTENT
	HATE_SPEECH
	SPAM
	SUICIDE
	DRUG_USE
	OTHER
)
