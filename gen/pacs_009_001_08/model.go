// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 with prefix 'c9'
package pacs_009_001_08

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType declarations

type AccountIdentification4Choice struct {
	Othr *GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Othr,omitempty"`
}

type ActiveCurrencyAndAmount struct {
	Value ActiveCurrencyAndAmountSimpleType `xml:",chardata"`
	Ccy   ActiveCurrencyCode                `xml:"Ccy,attr"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 FinInstnId"`
}

type BranchAndFinancialInstitutionIdentification6TCH struct {
	FinInstnId FinancialInstitutionIdentification18TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 FinInstnId"`
}

type CashAccount38 struct {
	Id AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Id"`
}

type ClearingSystemIdentification3Choice struct {
	Cd *ExternalCashClearingSystem1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Cd,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	MmbId Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 MmbId"`
}

type ClearingSystemMemberIdentification2TCH struct {
	MmbId Max35TextTCH3 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 MmbId"`
}

type CreditTransferTransaction36 struct {
	PmtId          PaymentIdentification7                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 PmtId"`
	PmtTpInf       PaymentTypeInformation28                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 PmtTpInf"`
	IntrBkSttlmAmt ActiveCurrencyAndAmount                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 IntrBkSttlmAmt"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstdAgt"`
	Dbtr           BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Dbtr"`
	DbtrAcct       *CashAccount38                               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 DbtrAcct,omitempty"`
	Cdtr           BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Cdtr"`
	CdtrAcct       CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtrAcct"`
	RmtInf         *RemittanceInformation2                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 RmtInf,omitempty"`
}

type CreditTransferTransaction36TCH struct {
	PmtId          PaymentIdentification7TCH                       `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 PmtId"`
	PmtTpInf       PaymentTypeInformation28TCH                     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 PmtTpInf"`
	IntrBkSttlmAmt ActiveCurrencyAndAmount                         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 IntrBkSttlmAmt"`
	InstgAgt       BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstgAgt"`
	InstdAgt       BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstdAgt"`
	Dbtr           BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Dbtr"`
	DbtrAcct       *CashAccount38                                  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 DbtrAcct,omitempty"`
	Cdtr           BranchAndFinancialInstitutionIdentification6TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Cdtr"`
	CdtrAcct       CashAccount38                                   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtrAcct"`
	RmtInf         *RemittanceInformation2TCH                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 RmtInf,omitempty"`
}

type DocumentTCH struct {
	XMLName  xml.Name
	FICdtTrf FinancialInstitutionCreditTransferV08TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 FICdtTrf"`
}

type FinancialInstitutionCreditTransferV08 struct {
	GrpHdr      GroupHeader93               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 GrpHdr"`
	CdtTrfTxInf CreditTransferTransaction36 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtTrfTxInf"`
}

type FinancialInstitutionCreditTransferV08TCH struct {
	GrpHdr      GroupHeader93TCH               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 GrpHdr"`
	CdtTrfTxInf CreditTransferTransaction36TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CdtTrfTxInf"`
}

type FinancialInstitutionIdentification18 struct {
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSysMmbId"`
}

type FinancialInstitutionIdentification18TCH struct {
	ClrSysMmbId ClearingSystemMemberIdentification2TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSysMmbId"`
}

type GenericAccountIdentification1 struct {
	Id Max34Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Id"`
}

type GroupHeader93 struct {
	MsgId             Max35Text               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 MsgId"`
	CreDtTm           rtp.ISODateTime         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CreDtTm"`
	NbOfTxs           Max1NumericText         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 NbOfTxs"`
	TtlIntrBkSttlmAmt ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 TtlIntrBkSttlmAmt"`
	IntrBkSttlmDt     rtp.ISODate             `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 IntrBkSttlmDt"`
	SttlmInf          SettlementInstruction7  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SttlmInf"`
}

type GroupHeader93TCH struct {
	MsgId             Max35TextTCH              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 MsgId"`
	CreDtTm           rtp.ISODateTime           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 CreDtTm"`
	NbOfTxs           Max1NumericText           `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 NbOfTxs"`
	TtlIntrBkSttlmAmt ActiveCurrencyAndAmount   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 TtlIntrBkSttlmAmt"`
	IntrBkSttlmDt     rtp.ISODate               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 IntrBkSttlmDt"`
	SttlmInf          SettlementInstruction7TCH `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SttlmInf"`
}

type PaymentIdentification7 struct {
	InstrId    Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstrId"`
	EndToEndId Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 EndToEndId"`
	TxId       Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 TxId"`
	ClrSysRef  *Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSysRef,omitempty"`
}

type PaymentIdentification7TCH struct {
	InstrId    Max35TextTCH2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 InstrId"`
	EndToEndId Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 EndToEndId"`
	TxId       Max35Text     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 TxId"`
	ClrSysRef  *Max35Text    `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSysRef,omitempty"`
}

type PaymentTypeInformation28 struct {
	SvcLvl ServiceLevel8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SvcLvl"`
}

type PaymentTypeInformation28TCH struct {
	SvcLvl ServiceLevel8Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SvcLvl"`
}

type RemittanceInformation2 struct {
	Ustrd *Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Ustrd,omitempty"`
}

type RemittanceInformation2TCH struct {
	Ustrd *Max140Text `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Ustrd,omitempty"`
}

type ServiceLevel8Choice struct {
	Cd *ExternalServiceLevel1Code `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 Cd,omitempty"`
}

type SettlementInstruction7 struct {
	SttlmMtd SettlementMethod1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SttlmMtd"`
	ClrSys   ClearingSystemIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSys"`
}

type SettlementInstruction7TCH struct {
	SttlmMtd SettlementMethod1Code               `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 SttlmMtd"`
	ClrSys   ClearingSystemIdentification3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08 ClrSys"`
}

// XSD SimpleType declarations

type ActiveCurrencyAndAmountSimpleType rtp.Amount

type ActiveCurrencyCode string

const ActiveCurrencyCodeUsd ActiveCurrencyCode = "USD"

type ExternalCashClearingSystem1Code string

const ExternalCashClearingSystem1CodeTch ExternalCashClearingSystem1Code = "TCH"

type ExternalServiceLevel1Code string

const ExternalServiceLevel1CodeSdva ExternalServiceLevel1Code = "SDVA"

type Max140Text string

type Max1NumericText string

type Max34Text string

type Max35Text string

type Max35TextTCH string

type Max35TextTCH2 string

type Max35TextTCH3 string

type SettlementMethod1Code string

const SettlementMethod1CodeClrg SettlementMethod1Code = "CLRG"