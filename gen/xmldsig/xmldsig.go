// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for http://www.w3.org/2000/09/xmldsig# with prefix 'ds'
package xmldsig

import (
	"encoding/xml"
	"fmt"
)

// XSD Elements

type Signature struct {
	Xmlns          []xml.Attr         `xml:",attr"`
	Id             *string            `xml:"Id,attr,omitempty"`
	SignedInfo     SignedInfoType     `xml:"SignedInfo"`
	SignatureValue SignatureValueType `xml:"SignatureValue"`
	KeyInfo        *KeyInfoType       `xml:"KeyInfo,omitempty"`
	Object         []*ObjectType      `xml:"Object,omitempty"`
}

// UnmarshalXML is a custom unmarshaller that allows us to capture the xmlns attributes
func (v *Signature) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		if (attr.Name.Space == "" && attr.Name.Local == "xmlns") || (attr.Name.Space == "xmlns") {
			newAttr := xml.Attr{}
			newAttr.Value = attr.Value
			newAttr.Name = xml.Name{}
			if attr.Name.Space == "" {
				newAttr.Name.Local = attr.Name.Local
			} else {
				newAttr.Name.Local = fmt.Sprintf("%s:%s", attr.Name.Space, attr.Name.Local)
			}
			v.Xmlns = append(v.Xmlns, newAttr)
		}
	}

	// Go on with unmarshalling.
	type vv Signature
	return d.DecodeElement((*vv)(v), &start)
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Signature) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = v.Xmlns
	if v.Id != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "Id"}, Value: *v.Id})
	}
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.SignedInfo, xml.StartElement{Name: xml.Name{Local: "ds:SignedInfo"}})
	e.EncodeElement(v.SignatureValue, xml.StartElement{Name: xml.Name{Local: "ds:SignatureValue"}})
	e.EncodeElement(v.KeyInfo, xml.StartElement{Name: xml.Name{Local: "ds:KeyInfo"}})
	e.EncodeElement(v.Object, xml.StartElement{Name: xml.Name{Local: "ds:Object"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD ComplexType declarations

type SignatureValueType struct {
	Id   *string `xml:"Id,attr,omitempty"`
	Data string  `xml:",chardata"`
}

type SignedInfoType struct {
	Id                     *string                    `xml:"Id,attr,omitempty"`
	CanonicalizationMethod CanonicalizationMethodType `xml:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethodType        `xml:"SignatureMethod"`
	Reference              []ReferenceType            `xml:"Reference"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SignedInfoType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = []xml.Attr{}
	if v.Id != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "Id"}, Value: *v.Id})
	}
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.CanonicalizationMethod, xml.StartElement{Name: xml.Name{Local: "ds:CanonicalizationMethod"}})
	e.EncodeElement(v.SignatureMethod, xml.StartElement{Name: xml.Name{Local: "ds:SignatureMethod"}})
	e.EncodeElement(v.Reference, xml.StartElement{Name: xml.Name{Local: "ds:Reference"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type CanonicalizationMethodType struct {
	Algorithm string  `xml:"Algorithm,attr"`
	Item      *string `xml:",any,omitempty"`
}

type SignatureMethodType struct {
	Algorithm        string                `xml:"Algorithm,attr"`
	HMACOutputLength *HMACOutputLengthType `xml:"HMACOutputLength,omitempty"`
	Item             *string               `xml:",any,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SignatureMethodType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = []xml.Attr{}

	attr = append(attr, xml.Attr{Name: xml.Name{Local: "Algorithm"}, Value: v.Algorithm})

	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.HMACOutputLength, xml.StartElement{Name: xml.Name{Local: "ds:HMACOutputLength"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ReferenceType struct {
	Id           *string          `xml:"Id,attr,omitempty"`
	URI          *string          `xml:"URI,attr,omitempty"`
	Type         *string          `xml:"Type,attr,omitempty"`
	Transforms   *TransformsType  `xml:"Transforms,omitempty"`
	DigestMethod DigestMethodType `xml:"DigestMethod"`
	DigestValue  DigestValueType  `xml:"DigestValue"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ReferenceType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = []xml.Attr{}
	if v.Id != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "Id"}, Value: *v.Id})
	}
	if v.URI != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "URI"}, Value: *v.URI})
	}
	if v.Type != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "Type"}, Value: *v.Type})
	}
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.Transforms, xml.StartElement{Name: xml.Name{Local: "ds:Transforms"}})
	e.EncodeElement(v.DigestMethod, xml.StartElement{Name: xml.Name{Local: "ds:DigestMethod"}})
	e.EncodeElement(v.DigestValue, xml.StartElement{Name: xml.Name{Local: "ds:DigestValue"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type TransformsType struct {
	Transform []TransformType `xml:"Transform"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TransformsType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Transform, xml.StartElement{Name: xml.Name{Local: "ds:Transform"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type TransformType struct {
	Algorithm string    `xml:"Algorithm,attr"`
	XPath     []*string `xml:"XPath,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v TransformType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = []xml.Attr{}

	attr = append(attr, xml.Attr{Name: xml.Name{Local: "Algorithm"}, Value: v.Algorithm})

	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.XPath, xml.StartElement{Name: xml.Name{Local: "ds:XPath"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type DigestMethodType struct {
	Algorithm string  `xml:"Algorithm,attr"`
	Item      *string `xml:",any,omitempty"`
}

type KeyInfoType struct {
	Id              *string                `xml:"Id,attr,omitempty"`
	KeyName         []*string              `xml:"KeyName,omitempty"`
	KeyValue        []*KeyValueType        `xml:"KeyValue,omitempty"`
	RetrievalMethod []*RetrievalMethodType `xml:"RetrievalMethod,omitempty"`
	X509Data        []*X509DataType        `xml:"X509Data,omitempty"`
	PGPData         []*PGPDataType         `xml:"PGPData,omitempty"`
	SPKIData        []*SPKIDataType        `xml:"SPKIData,omitempty"`
	MgmtData        []*string              `xml:"MgmtData,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v KeyInfoType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = []xml.Attr{}
	if v.Id != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "Id"}, Value: *v.Id})
	}
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.KeyName, xml.StartElement{Name: xml.Name{Local: "ds:KeyName"}})
	e.EncodeElement(v.KeyValue, xml.StartElement{Name: xml.Name{Local: "ds:KeyValue"}})
	e.EncodeElement(v.RetrievalMethod, xml.StartElement{Name: xml.Name{Local: "ds:RetrievalMethod"}})
	e.EncodeElement(v.X509Data, xml.StartElement{Name: xml.Name{Local: "ds:X509Data"}})
	e.EncodeElement(v.PGPData, xml.StartElement{Name: xml.Name{Local: "ds:PGPData"}})
	e.EncodeElement(v.SPKIData, xml.StartElement{Name: xml.Name{Local: "ds:SPKIData"}})
	e.EncodeElement(v.MgmtData, xml.StartElement{Name: xml.Name{Local: "ds:MgmtData"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type KeyValueType struct {
	DSAKeyValue *DSAKeyValueType `xml:"DSAKeyValue,omitempty"`
	RSAKeyValue *RSAKeyValueType `xml:"RSAKeyValue,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v KeyValueType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.DSAKeyValue, xml.StartElement{Name: xml.Name{Local: "ds:DSAKeyValue"}})
	e.EncodeElement(v.RSAKeyValue, xml.StartElement{Name: xml.Name{Local: "ds:RSAKeyValue"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type RetrievalMethodType struct {
	URI        string          `xml:"URI,attr"`
	Type       *string         `xml:"Type,attr,omitempty"`
	Transforms *TransformsType `xml:"Transforms,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RetrievalMethodType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = []xml.Attr{}

	attr = append(attr, xml.Attr{Name: xml.Name{Local: "URI"}, Value: v.URI})

	if v.Type != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "Type"}, Value: *v.Type})
	}
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.Transforms, xml.StartElement{Name: xml.Name{Local: "ds:Transforms"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type X509DataType struct {
	X509IssuerSerial *X509IssuerSerialType `xml:"X509IssuerSerial,omitempty"`
	X509SKI          *string               `xml:"X509SKI,omitempty"`
	X509SubjectName  *string               `xml:"X509SubjectName,omitempty"`
	X509Certificate  *string               `xml:"X509Certificate,omitempty"`
	X509CRL          *string               `xml:"X509CRL,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v X509DataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.X509IssuerSerial, xml.StartElement{Name: xml.Name{Local: "ds:X509IssuerSerial"}})
	e.EncodeElement(v.X509SKI, xml.StartElement{Name: xml.Name{Local: "ds:X509SKI"}})
	e.EncodeElement(v.X509SubjectName, xml.StartElement{Name: xml.Name{Local: "ds:X509SubjectName"}})
	e.EncodeElement(v.X509Certificate, xml.StartElement{Name: xml.Name{Local: "ds:X509Certificate"}})
	e.EncodeElement(v.X509CRL, xml.StartElement{Name: xml.Name{Local: "ds:X509CRL"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type X509IssuerSerialType struct {
	X509IssuerName   string `xml:"X509IssuerName"`
	X509SerialNumber string `xml:"X509SerialNumber"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v X509IssuerSerialType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.X509IssuerName, xml.StartElement{Name: xml.Name{Local: "ds:X509IssuerName"}})
	e.EncodeElement(v.X509SerialNumber, xml.StartElement{Name: xml.Name{Local: "ds:X509SerialNumber"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type PGPDataType struct {
	PGPKeyID     *string `xml:"PGPKeyID,omitempty"`
	PGPKeyPacket *string `xml:"PGPKeyPacket,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PGPDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.PGPKeyID, xml.StartElement{Name: xml.Name{Local: "ds:PGPKeyID"}})
	e.EncodeElement(v.PGPKeyPacket, xml.StartElement{Name: xml.Name{Local: "ds:PGPKeyPacket"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SPKIDataType struct {
	SPKISexp string  `xml:"SPKISexp"`
	Item     *string `xml:",any,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SPKIDataType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.SPKISexp, xml.StartElement{Name: xml.Name{Local: "ds:SPKISexp"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ObjectType struct {
	Id       *string `xml:"Id,attr,omitempty"`
	MimeType *string `xml:"MimeType,attr,omitempty"`
	Encoding *string `xml:"Encoding,attr,omitempty"`
	Item     any     `xml:",any"`
}

type ManifestType struct {
	Id        *string         `xml:"Id,attr,omitempty"`
	Reference []ReferenceType `xml:"Reference"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ManifestType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = []xml.Attr{}
	if v.Id != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "Id"}, Value: *v.Id})
	}
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.Reference, xml.StartElement{Name: xml.Name{Local: "ds:Reference"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SignaturePropertiesType struct {
	Id                *string                 `xml:"Id,attr,omitempty"`
	SignatureProperty []SignaturePropertyType `xml:"SignatureProperty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SignaturePropertiesType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var attr = []xml.Attr{}
	if v.Id != nil {
		attr = append(attr, xml.Attr{Name: xml.Name{Local: "Id"}, Value: *v.Id})
	}
	e.EncodeToken(xml.StartElement{
		Name: xml.Name{Local: start.Name.Local},
		Attr: attr,
	})
	e.EncodeElement(v.SignatureProperty, xml.StartElement{Name: xml.Name{Local: "ds:SignatureProperty"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SignaturePropertyType struct {
	Target string  `xml:"Target,attr"`
	Id     *string `xml:"Id,attr,omitempty"`
}

type DSAKeyValueType struct {
	G *CryptoBinary `xml:"G,omitempty"`
	Y CryptoBinary  `xml:"Y"`
	J *CryptoBinary `xml:"J,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DSAKeyValueType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.G, xml.StartElement{Name: xml.Name{Local: "ds:G"}})
	e.EncodeElement(v.Y, xml.StartElement{Name: xml.Name{Local: "ds:Y"}})
	e.EncodeElement(v.J, xml.StartElement{Name: xml.Name{Local: "ds:J"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type RSAKeyValueType struct {
	Modulus  CryptoBinary `xml:"Modulus"`
	Exponent CryptoBinary `xml:"Exponent"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v RSAKeyValueType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.Modulus, xml.StartElement{Name: xml.Name{Local: "ds:Modulus"}})
	e.EncodeElement(v.Exponent, xml.StartElement{Name: xml.Name{Local: "ds:Exponent"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD SimpleType declarations

type CryptoBinary string

type DigestValueType string

type HMACOutputLengthType int64
