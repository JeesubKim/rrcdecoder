package rrcdecoder

type DL_DCCH_Message struct {
	message DL_DCCH_MessageType
}

type DL_DCCH_MessageType interface{}

type CSFBParametersResponseCDMA2000 struct{}
type DLInformationTransfer struct{}
type HandoverFromEUTRAPreparationRequest struct{}
type MobilityFromEUTRACommand struct{}
type RRCConnectionReconfiguration struct{}
type RRCConnectionRelease struct{}
type SecurityModeCommand struct{}
type UECapabilityEnquiry struct{}
type CounterCheck struct{}
type UEInformationRequest_r9 struct{}
type LoggedMeasurementConfiguration_r10 struct{}
type RNReconfiguration_r10 struct{}
type RRCConnectionResume_r13 struct{}
