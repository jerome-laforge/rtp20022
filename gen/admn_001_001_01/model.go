// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 with prefix 'sr'
package admn_001_001_01

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification4ADMN struct {
	FinInstnId FinancialInstitutionIdentification7ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 FinInstnId"`
}

type ClearingSystemMemberIdentification2ADMN struct {
	MmbId Min11Max11Text `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 MmbId"`
}

type DocumentTCH struct {
	XMLName       xml.Name
	AdmnSignOnReq SignOnRequestTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 AdmnSignOnReq"`
}

type FinancialInstitutionIdentification7ADMN struct {
	ClrSysMmbId ClearingSystemMemberIdentification2ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 ClrSysMmbId"`
}

type GrpHdr struct {
	MsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 CreDtTm"`
}

type GrpHdrTCH struct {
	MsgId   Max35TextTCH    `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 CreDtTm"`
}

type SignOnReq struct {
	InstrId  Max35Text                                        `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 InstrId"`
	InstgAgt BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 InstgAgt"`
	InstdAgt BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 InstdAgt"`
}

type SignOnRequest struct {
	GrpHdr    GrpHdr    `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 GrpHdr"`
	SignOnReq SignOnReq `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 SignOnReq"`
}

type SignOnRequestTCH struct {
	GrpHdr    GrpHdrTCH    `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 GrpHdr"`
	SignOnReq SignOnReqTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 SignOnReq"`
}

type SignOnReqTCH struct {
	InstrId  Max35TextTCH2                                    `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 InstrId"`
	InstgAgt BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 InstgAgt"`
	InstdAgt BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.001.001.01 InstdAgt"`
}

// XSD SimpleType declarations

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type Min11Max11Text string
