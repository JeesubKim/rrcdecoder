package rrcdecoder

//https://www.etsi.org/deliver/etsi_ts/136300_136399/136331/15.20.00_60/ts_136331v152000p.pdf
//315 page
type BCCH_BCH_Message struct {
	message BCCH_BCH_MessageType
}

type BCCH_BCH_MessageType MasterInformationBlock

type BCCH_BCH_Message_MBMS struct {
	message BCCH_BCH_MessageTYPE_MBMS_r14
}

type BCCH_BCH_MessageTYPE_MBMS_r14 MasterInformationBlock_MBMS_r14
