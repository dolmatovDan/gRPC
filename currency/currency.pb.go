// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.12.4
// source: currency.proto

package currency

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// all currencies
type Currencies int32

const (
	Currencies_EUR Currencies = 0
	Currencies_USD Currencies = 1
	Currencies_JPY Currencies = 2
	Currencies_BGN Currencies = 3
	Currencies_CZK Currencies = 4
	Currencies_DKK Currencies = 5
	Currencies_GBP Currencies = 6
	Currencies_HUF Currencies = 7
	Currencies_PLN Currencies = 8
	Currencies_RON Currencies = 9
	Currencies_SEK Currencies = 10
	Currencies_CHF Currencies = 11
	Currencies_ISK Currencies = 12
	Currencies_NOK Currencies = 13
	Currencies_HRK Currencies = 14
	Currencies_RUB Currencies = 15
	Currencies_TRY Currencies = 16
	Currencies_AUD Currencies = 17
	Currencies_BRL Currencies = 18
	Currencies_CAD Currencies = 19
	Currencies_CNY Currencies = 20
	Currencies_HKD Currencies = 21
	Currencies_IDR Currencies = 22
	Currencies_ILS Currencies = 23
	Currencies_INR Currencies = 24
	Currencies_KRW Currencies = 25
	Currencies_MXN Currencies = 26
	Currencies_MYR Currencies = 27
	Currencies_NZD Currencies = 28
	Currencies_PHP Currencies = 29
	Currencies_SGD Currencies = 30
	Currencies_THB Currencies = 31
	Currencies_ZAR Currencies = 32
)

// Enum value maps for Currencies.
var (
	Currencies_name = map[int32]string{
		0:  "EUR",
		1:  "USD",
		2:  "JPY",
		3:  "BGN",
		4:  "CZK",
		5:  "DKK",
		6:  "GBP",
		7:  "HUF",
		8:  "PLN",
		9:  "RON",
		10: "SEK",
		11: "CHF",
		12: "ISK",
		13: "NOK",
		14: "HRK",
		15: "RUB",
		16: "TRY",
		17: "AUD",
		18: "BRL",
		19: "CAD",
		20: "CNY",
		21: "HKD",
		22: "IDR",
		23: "ILS",
		24: "INR",
		25: "KRW",
		26: "MXN",
		27: "MYR",
		28: "NZD",
		29: "PHP",
		30: "SGD",
		31: "THB",
		32: "ZAR",
	}
	Currencies_value = map[string]int32{
		"EUR": 0,
		"USD": 1,
		"JPY": 2,
		"BGN": 3,
		"CZK": 4,
		"DKK": 5,
		"GBP": 6,
		"HUF": 7,
		"PLN": 8,
		"RON": 9,
		"SEK": 10,
		"CHF": 11,
		"ISK": 12,
		"NOK": 13,
		"HRK": 14,
		"RUB": 15,
		"TRY": 16,
		"AUD": 17,
		"BRL": 18,
		"CAD": 19,
		"CNY": 20,
		"HKD": 21,
		"IDR": 22,
		"ILS": 23,
		"INR": 24,
		"KRW": 25,
		"MXN": 26,
		"MYR": 27,
		"NZD": 28,
		"PHP": 29,
		"SGD": 30,
		"THB": 31,
		"ZAR": 32,
	}
)

func (x Currencies) Enum() *Currencies {
	p := new(Currencies)
	*p = x
	return p
}

func (x Currencies) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currencies) Descriptor() protoreflect.EnumDescriptor {
	return file_currency_proto_enumTypes[0].Descriptor()
}

func (Currencies) Type() protoreflect.EnumType {
	return &file_currency_proto_enumTypes[0]
}

func (x Currencies) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currencies.Descriptor instead.
func (Currencies) EnumDescriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

type RateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Base          Currencies             `protobuf:"varint,1,opt,name=Base,proto3,enum=Currencies" json:"Base,omitempty"`
	Destination   Currencies             `protobuf:"varint,2,opt,name=Destination,proto3,enum=Currencies" json:"Destination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	mi := &file_currency_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{0}
}

func (x *RateRequest) GetBase() Currencies {
	if x != nil {
		return x.Base
	}
	return Currencies_EUR
}

func (x *RateRequest) GetDestination() Currencies {
	if x != nil {
		return x.Destination
	}
	return Currencies_EUR
}

type RateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Rate          float32                `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	mi := &file_currency_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_currency_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_currency_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_currency_proto protoreflect.FileDescriptor

const file_currency_proto_rawDesc = "" +
	"\n" +
	"\x0ecurrency.proto\"]\n" +
	"\vRateRequest\x12\x1f\n" +
	"\x04Base\x18\x01 \x01(\x0e2\v.CurrenciesR\x04Base\x12-\n" +
	"\vDestination\x18\x02 \x01(\x0e2\v.CurrenciesR\vDestination\"\"\n" +
	"\fRateResponse\x12\x12\n" +
	"\x04rate\x18\x01 \x01(\x02R\x04rate*\xb5\x02\n" +
	"\n" +
	"Currencies\x12\a\n" +
	"\x03EUR\x10\x00\x12\a\n" +
	"\x03USD\x10\x01\x12\a\n" +
	"\x03JPY\x10\x02\x12\a\n" +
	"\x03BGN\x10\x03\x12\a\n" +
	"\x03CZK\x10\x04\x12\a\n" +
	"\x03DKK\x10\x05\x12\a\n" +
	"\x03GBP\x10\x06\x12\a\n" +
	"\x03HUF\x10\a\x12\a\n" +
	"\x03PLN\x10\b\x12\a\n" +
	"\x03RON\x10\t\x12\a\n" +
	"\x03SEK\x10\n" +
	"\x12\a\n" +
	"\x03CHF\x10\v\x12\a\n" +
	"\x03ISK\x10\f\x12\a\n" +
	"\x03NOK\x10\r\x12\a\n" +
	"\x03HRK\x10\x0e\x12\a\n" +
	"\x03RUB\x10\x0f\x12\a\n" +
	"\x03TRY\x10\x10\x12\a\n" +
	"\x03AUD\x10\x11\x12\a\n" +
	"\x03BRL\x10\x12\x12\a\n" +
	"\x03CAD\x10\x13\x12\a\n" +
	"\x03CNY\x10\x14\x12\a\n" +
	"\x03HKD\x10\x15\x12\a\n" +
	"\x03IDR\x10\x16\x12\a\n" +
	"\x03ILS\x10\x17\x12\a\n" +
	"\x03INR\x10\x18\x12\a\n" +
	"\x03KRW\x10\x19\x12\a\n" +
	"\x03MXN\x10\x1a\x12\a\n" +
	"\x03MYR\x10\x1b\x12\a\n" +
	"\x03NZD\x10\x1c\x12\a\n" +
	"\x03PHP\x10\x1d\x12\a\n" +
	"\x03SGD\x10\x1e\x12\a\n" +
	"\x03THB\x10\x1f\x12\a\n" +
	"\x03ZAR\x10 22\n" +
	"\bCurrency\x12&\n" +
	"\aGetRate\x12\f.RateRequest\x1a\r.RateResponseB\vZ\t/currencyb\x06proto3"

var (
	file_currency_proto_rawDescOnce sync.Once
	file_currency_proto_rawDescData []byte
)

func file_currency_proto_rawDescGZIP() []byte {
	file_currency_proto_rawDescOnce.Do(func() {
		file_currency_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_currency_proto_rawDesc), len(file_currency_proto_rawDesc)))
	})
	return file_currency_proto_rawDescData
}

var file_currency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_currency_proto_goTypes = []any{
	(Currencies)(0),      // 0: Currencies
	(*RateRequest)(nil),  // 1: RateRequest
	(*RateResponse)(nil), // 2: RateResponse
}
var file_currency_proto_depIdxs = []int32{
	0, // 0: RateRequest.Base:type_name -> Currencies
	0, // 1: RateRequest.Destination:type_name -> Currencies
	1, // 2: Currency.GetRate:input_type -> RateRequest
	2, // 3: Currency.GetRate:output_type -> RateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_currency_proto_init() }
func file_currency_proto_init() {
	if File_currency_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_currency_proto_rawDesc), len(file_currency_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_currency_proto_goTypes,
		DependencyIndexes: file_currency_proto_depIdxs,
		EnumInfos:         file_currency_proto_enumTypes,
		MessageInfos:      file_currency_proto_msgTypes,
	}.Build()
	File_currency_proto = out.File
	file_currency_proto_goTypes = nil
	file_currency_proto_depIdxs = nil
}
