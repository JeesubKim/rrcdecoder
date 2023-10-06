package rrcdecoder

type DL_CCCH_Message struct {
	message DL_CCCH_MessageType
}

type DL_CCCH_MessageType interface{}

type RRCConnectionReestablishment struct {
}
type RRCConnectionReestablishmentReject struct {
}
type RRCConnectionReject struct {
}
type RRCConnectionSetup struct {
}

type RRCEarlyDataComplete_r15 struct {
}
