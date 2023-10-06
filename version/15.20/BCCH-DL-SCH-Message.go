package rrcdecoder

type BCCH_DL_SCH_Message struct {
	message BCCH_DL_SCHE_MessageType
}
type BCCH_DL_SCH_Message_BR struct {
	message BCCH_DL_SCH_Message_BR_r13
}

type BCCH_DL_SCHE_MessageType interface{}

type SystemInformation struct{}
type SystemInformationBlockType1 struct{}

type BCCH_DL_SCH_Message_BR_r13 interface{}

type BCCH_DL_SCH_Message_MBMS_r14 interface{}

type SystemInformation_MBMS_r14 struct {
}

type SystemInformationBlockType1_MBMS_r14 struct {
}
