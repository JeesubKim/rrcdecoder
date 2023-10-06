package rrcdecoder

type UL_CCCH_Message struct {
	message UL_CCCH_MessageType
}

type UL_CCCH_MessageType interface{}

type RRCConnectionReestablishmentRequest struct{}
type RRCConnectionRequest struct{}

type RRCConnectionResumeRequest_r13 struct{}

type RRCEarlyDataRequest_r15 struct{}
