// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 with prefix 're'
package admn_006_001_01

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification4ADMN struct {
	FinInstnId FinancialInstitutionIdentification7ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 FinInstnId"`
}

type ClearingSystemMemberIdentification2ADMN struct {
	MmbId Min11Max11Text `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 MmbId"`
}

type DocumentTCH struct {
	XMLName      xml.Name
	AdmnEchoResp EchoResponseTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 AdmnEchoResp"`
}

type EchoResp struct {
	InstgAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 InstdAgt"`
	OrgnlInstrId Max35Text                                        `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 OrgnlInstrId"`
	FnctnCd      EchoCode                                         `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 FnctnCd"`
	TxSts        TransactionIndividualStatus3CodeEcho             `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 TxSts"`
}

type EchoResponse struct {
	GrpHdr       GrpHdr   `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 GrpHdr"`
	EchoResponse EchoResp `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 EchoResponse"`
}

type EchoResponseTCH struct {
	GrpHdr       GrpHdrTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 GrpHdr"`
	EchoResponse EchoResp  `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 EchoResponse"`
}

type FinancialInstitutionIdentification7ADMN struct {
	ClrSysMmbId ClearingSystemMemberIdentification2ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 ClrSysMmbId"`
}

type GrpHdr struct {
	MsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 CreDtTm"`
}

type GrpHdrTCH struct {
	MsgId   Max35TextTCH    `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 CreDtTm"`
}

// XSD SimpleType declarations

type EchoCode string

const EchoCode731 EchoCode = "731"

type Max35Text string

type Max35TextTCH string

type Min11Max11Text string

type TransactionIndividualStatus3CodeEcho string

const TransactionIndividualStatus3CodeEchoActc TransactionIndividualStatus3CodeEcho = "ACTC"
