package parser

type situation string

const (
	situationPending           situation = ""
	situationLabel             situation = "label"
	situationJump              situation = "jump"
	situationCall              situation = "call"
	situationFakeLabel         situation = "fakelabel"
	situationFakeJump          situation = "fakejump"
	situationScreen            situation = "screen"
	situationFromScreenToOther situation = "fromscreentoother"
)
