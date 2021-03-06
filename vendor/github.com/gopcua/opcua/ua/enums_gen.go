// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

// Code generated by cmd/service. DO NOT EDIT!

package ua

type NodeIDType uint8

func NodeIDTypeFromString(s string) NodeIDType {
	switch s {
	case "TwoByte":
		return 0
	case "FourByte":
		return 1
	case "Numeric":
		return 2
	case "String":
		return 3
	case "Guid":
		return 4
	case "ByteString":
		return 5
	default:
		return 0
	}
}

const (
	NodeIDTypeTwoByte    NodeIDType = 0
	NodeIDTypeFourByte   NodeIDType = 1
	NodeIDTypeNumeric    NodeIDType = 2
	NodeIDTypeString     NodeIDType = 3
	NodeIDTypeGUID       NodeIDType = 4
	NodeIDTypeByteString NodeIDType = 5
)

type NamingRuleType uint32

func NamingRuleTypeFromString(s string) NamingRuleType {
	switch s {
	case "Mandatory":
		return 1
	case "Optional":
		return 2
	case "Constraint":
		return 3
	default:
		return 0
	}
}

const (
	NamingRuleTypeMandatory  NamingRuleType = 1
	NamingRuleTypeOptional   NamingRuleType = 2
	NamingRuleTypeConstraint NamingRuleType = 3
)

type OpenFileMode uint32

func OpenFileModeFromString(s string) OpenFileMode {
	switch s {
	case "Read":
		return 1
	case "Write":
		return 2
	case "EraseExisting":
		return 4
	case "Append":
		return 8
	default:
		return 0
	}
}

const (
	OpenFileModeRead          OpenFileMode = 1
	OpenFileModeWrite         OpenFileMode = 2
	OpenFileModeEraseExisting OpenFileMode = 4
	OpenFileModeAppend        OpenFileMode = 8
)

type IdentityCriteriaType uint32

func IdentityCriteriaTypeFromString(s string) IdentityCriteriaType {
	switch s {
	case "UserName":
		return 1
	case "Thumbprint":
		return 2
	case "Role":
		return 3
	case "GroupId":
		return 4
	case "Anonymous":
		return 5
	case "AuthenticatedUser":
		return 6
	default:
		return 0
	}
}

const (
	IdentityCriteriaTypeUserName          IdentityCriteriaType = 1
	IdentityCriteriaTypeThumbprint        IdentityCriteriaType = 2
	IdentityCriteriaTypeRole              IdentityCriteriaType = 3
	IdentityCriteriaTypeGroupID           IdentityCriteriaType = 4
	IdentityCriteriaTypeAnonymous         IdentityCriteriaType = 5
	IdentityCriteriaTypeAuthenticatedUser IdentityCriteriaType = 6
)

type TrustListMasks uint32

func TrustListMasksFromString(s string) TrustListMasks {
	switch s {
	case "None":
		return 0
	case "TrustedCertificates":
		return 1
	case "TrustedCrls":
		return 2
	case "IssuerCertificates":
		return 4
	case "IssuerCrls":
		return 8
	case "All":
		return 15
	default:
		return 0
	}
}

const (
	TrustListMasksNone                TrustListMasks = 0
	TrustListMasksTrustedCertificates TrustListMasks = 1
	TrustListMasksTrustedCrls         TrustListMasks = 2
	TrustListMasksIssuerCertificates  TrustListMasks = 4
	TrustListMasksIssuerCrls          TrustListMasks = 8
	TrustListMasksAll                 TrustListMasks = 15
)

type PubSubState uint32

func PubSubStateFromString(s string) PubSubState {
	switch s {
	case "Disabled":
		return 0
	case "Paused":
		return 1
	case "Operational":
		return 2
	case "Error":
		return 3
	default:
		return 0
	}
}

const (
	PubSubStateDisabled    PubSubState = 0
	PubSubStatePaused      PubSubState = 1
	PubSubStateOperational PubSubState = 2
	PubSubStateError       PubSubState = 3
)

type DataSetFieldFlags uint16

func DataSetFieldFlagsFromString(s string) DataSetFieldFlags {
	switch s {
	case "None":
		return 0
	case "PromotedField":
		return 1
	default:
		return 0
	}
}

const (
	DataSetFieldFlagsNone          DataSetFieldFlags = 0
	DataSetFieldFlagsPromotedField DataSetFieldFlags = 1
)

type DataSetFieldContentMask uint32

func DataSetFieldContentMaskFromString(s string) DataSetFieldContentMask {
	switch s {
	case "None":
		return 0
	case "StatusCode":
		return 1
	case "SourceTimestamp":
		return 2
	case "ServerTimestamp":
		return 4
	case "SourcePicoSeconds":
		return 8
	case "ServerPicoSeconds":
		return 16
	case "RawData":
		return 32
	default:
		return 0
	}
}

const (
	DataSetFieldContentMaskNone              DataSetFieldContentMask = 0
	DataSetFieldContentMaskStatusCode        DataSetFieldContentMask = 1
	DataSetFieldContentMaskSourceTimestamp   DataSetFieldContentMask = 2
	DataSetFieldContentMaskServerTimestamp   DataSetFieldContentMask = 4
	DataSetFieldContentMaskSourcePicoSeconds DataSetFieldContentMask = 8
	DataSetFieldContentMaskServerPicoSeconds DataSetFieldContentMask = 16
	DataSetFieldContentMaskRawData           DataSetFieldContentMask = 32
)

type OverrideValueHandling uint32

func OverrideValueHandlingFromString(s string) OverrideValueHandling {
	switch s {
	case "Disabled":
		return 0
	case "LastUsableValue":
		return 1
	case "OverrideValue":
		return 2
	default:
		return 0
	}
}

const (
	OverrideValueHandlingDisabled        OverrideValueHandling = 0
	OverrideValueHandlingLastUsableValue OverrideValueHandling = 1
	OverrideValueHandlingOverrideValue   OverrideValueHandling = 2
)

type DataSetOrderingType uint32

func DataSetOrderingTypeFromString(s string) DataSetOrderingType {
	switch s {
	case "Undefined":
		return 0
	case "AscendingWriterId":
		return 1
	case "AscendingWriterIdSingle":
		return 2
	default:
		return 0
	}
}

const (
	DataSetOrderingTypeUndefined               DataSetOrderingType = 0
	DataSetOrderingTypeAscendingWriterID       DataSetOrderingType = 1
	DataSetOrderingTypeAscendingWriterIDSingle DataSetOrderingType = 2
)

type UADPNetworkMessageContentMask uint32

func UADPNetworkMessageContentMaskFromString(s string) UADPNetworkMessageContentMask {
	switch s {
	case "None":
		return 0
	case "PublisherId":
		return 1
	case "GroupHeader":
		return 2
	case "WriterGroupId":
		return 4
	case "GroupVersion":
		return 8
	case "NetworkMessageNumber":
		return 16
	case "SequenceNumber":
		return 32
	case "PayloadHeader":
		return 64
	case "Timestamp":
		return 128
	case "PicoSeconds":
		return 256
	case "DataSetClassId":
		return 512
	case "PromotedFields":
		return 1024
	default:
		return 0
	}
}

const (
	UADPNetworkMessageContentMaskNone                 UADPNetworkMessageContentMask = 0
	UADPNetworkMessageContentMaskPublisherID          UADPNetworkMessageContentMask = 1
	UADPNetworkMessageContentMaskGroupHeader          UADPNetworkMessageContentMask = 2
	UADPNetworkMessageContentMaskWriterGroupID        UADPNetworkMessageContentMask = 4
	UADPNetworkMessageContentMaskGroupVersion         UADPNetworkMessageContentMask = 8
	UADPNetworkMessageContentMaskNetworkMessageNumber UADPNetworkMessageContentMask = 16
	UADPNetworkMessageContentMaskSequenceNumber       UADPNetworkMessageContentMask = 32
	UADPNetworkMessageContentMaskPayloadHeader        UADPNetworkMessageContentMask = 64
	UADPNetworkMessageContentMaskTimestamp            UADPNetworkMessageContentMask = 128
	UADPNetworkMessageContentMaskPicoSeconds          UADPNetworkMessageContentMask = 256
	UADPNetworkMessageContentMaskDataSetClassID       UADPNetworkMessageContentMask = 512
	UADPNetworkMessageContentMaskPromotedFields       UADPNetworkMessageContentMask = 1024
)

type UADPDataSetMessageContentMask uint32

func UADPDataSetMessageContentMaskFromString(s string) UADPDataSetMessageContentMask {
	switch s {
	case "None":
		return 0
	case "Timestamp":
		return 1
	case "PicoSeconds":
		return 2
	case "Status":
		return 4
	case "MajorVersion":
		return 8
	case "MinorVersion":
		return 16
	case "SequenceNumber":
		return 32
	default:
		return 0
	}
}

const (
	UADPDataSetMessageContentMaskNone           UADPDataSetMessageContentMask = 0
	UADPDataSetMessageContentMaskTimestamp      UADPDataSetMessageContentMask = 1
	UADPDataSetMessageContentMaskPicoSeconds    UADPDataSetMessageContentMask = 2
	UADPDataSetMessageContentMaskStatus         UADPDataSetMessageContentMask = 4
	UADPDataSetMessageContentMaskMajorVersion   UADPDataSetMessageContentMask = 8
	UADPDataSetMessageContentMaskMinorVersion   UADPDataSetMessageContentMask = 16
	UADPDataSetMessageContentMaskSequenceNumber UADPDataSetMessageContentMask = 32
)

type JSONNetworkMessageContentMask uint32

func JSONNetworkMessageContentMaskFromString(s string) JSONNetworkMessageContentMask {
	switch s {
	case "None":
		return 0
	case "NetworkMessageHeader":
		return 1
	case "DataSetMessageHeader":
		return 2
	case "SingleDataSetMessage":
		return 4
	case "PublisherId":
		return 8
	case "DataSetClassId":
		return 16
	case "ReplyTo":
		return 32
	default:
		return 0
	}
}

const (
	JSONNetworkMessageContentMaskNone                 JSONNetworkMessageContentMask = 0
	JSONNetworkMessageContentMaskNetworkMessageHeader JSONNetworkMessageContentMask = 1
	JSONNetworkMessageContentMaskDataSetMessageHeader JSONNetworkMessageContentMask = 2
	JSONNetworkMessageContentMaskSingleDataSetMessage JSONNetworkMessageContentMask = 4
	JSONNetworkMessageContentMaskPublisherID          JSONNetworkMessageContentMask = 8
	JSONNetworkMessageContentMaskDataSetClassID       JSONNetworkMessageContentMask = 16
	JSONNetworkMessageContentMaskReplyTo              JSONNetworkMessageContentMask = 32
)

type JSONDataSetMessageContentMask uint32

func JSONDataSetMessageContentMaskFromString(s string) JSONDataSetMessageContentMask {
	switch s {
	case "None":
		return 0
	case "DataSetWriterId":
		return 1
	case "MetaDataVersion":
		return 2
	case "SequenceNumber":
		return 4
	case "Timestamp":
		return 8
	case "Status":
		return 16
	default:
		return 0
	}
}

const (
	JSONDataSetMessageContentMaskNone            JSONDataSetMessageContentMask = 0
	JSONDataSetMessageContentMaskDataSetWriterID JSONDataSetMessageContentMask = 1
	JSONDataSetMessageContentMaskMetaDataVersion JSONDataSetMessageContentMask = 2
	JSONDataSetMessageContentMaskSequenceNumber  JSONDataSetMessageContentMask = 4
	JSONDataSetMessageContentMaskTimestamp       JSONDataSetMessageContentMask = 8
	JSONDataSetMessageContentMaskStatus          JSONDataSetMessageContentMask = 16
)

type BrokerTransportQoS uint32

func BrokerTransportQoSFromString(s string) BrokerTransportQoS {
	switch s {
	case "NotSpecified":
		return 0
	case "BestEffort":
		return 1
	case "AtLeastOnce":
		return 2
	case "AtMostOnce":
		return 3
	case "ExactlyOnce":
		return 4
	default:
		return 0
	}
}

const (
	BrokerTransportQoSNotSpecified BrokerTransportQoS = 0
	BrokerTransportQoSBestEffort   BrokerTransportQoS = 1
	BrokerTransportQoSAtLeastOnce  BrokerTransportQoS = 2
	BrokerTransportQoSAtMostOnce   BrokerTransportQoS = 3
	BrokerTransportQoSExactlyOnce  BrokerTransportQoS = 4
)

type DiagnosticsLevel uint32

func DiagnosticsLevelFromString(s string) DiagnosticsLevel {
	switch s {
	case "Basic":
		return 0
	case "Advanced":
		return 1
	case "Info":
		return 2
	case "Log":
		return 3
	case "Debug":
		return 4
	default:
		return 0
	}
}

const (
	DiagnosticsLevelBasic    DiagnosticsLevel = 0
	DiagnosticsLevelAdvanced DiagnosticsLevel = 1
	DiagnosticsLevelInfo     DiagnosticsLevel = 2
	DiagnosticsLevelLog      DiagnosticsLevel = 3
	DiagnosticsLevelDebug    DiagnosticsLevel = 4
)

type PubSubDiagnosticsCounterClassification uint32

func PubSubDiagnosticsCounterClassificationFromString(s string) PubSubDiagnosticsCounterClassification {
	switch s {
	case "Information":
		return 0
	case "Error":
		return 1
	default:
		return 0
	}
}

const (
	PubSubDiagnosticsCounterClassificationInformation PubSubDiagnosticsCounterClassification = 0
	PubSubDiagnosticsCounterClassificationError       PubSubDiagnosticsCounterClassification = 1
)

type IDType uint32

func IDTypeFromString(s string) IDType {
	switch s {
	case "Numeric":
		return 0
	case "String":
		return 1
	case "Guid":
		return 2
	case "Opaque":
		return 3
	default:
		return 0
	}
}

const (
	IDTypeNumeric IDType = 0
	IDTypeString  IDType = 1
	IDTypeGUID    IDType = 2
	IDTypeOpaque  IDType = 3
)

type NodeClass uint32

func NodeClassFromString(s string) NodeClass {
	switch s {
	case "Unspecified":
		return 0
	case "Object":
		return 1
	case "Variable":
		return 2
	case "Method":
		return 4
	case "ObjectType":
		return 8
	case "VariableType":
		return 16
	case "ReferenceType":
		return 32
	case "DataType":
		return 64
	case "View":
		return 128
	default:
		return 0
	}
}

const (
	NodeClassUnspecified   NodeClass = 0
	NodeClassObject        NodeClass = 1
	NodeClassVariable      NodeClass = 2
	NodeClassMethod        NodeClass = 4
	NodeClassObjectType    NodeClass = 8
	NodeClassVariableType  NodeClass = 16
	NodeClassReferenceType NodeClass = 32
	NodeClassDataType      NodeClass = 64
	NodeClassView          NodeClass = 128
)

type PermissionType uint32

func PermissionTypeFromString(s string) PermissionType {
	switch s {
	case "None":
		return 0
	case "Browse":
		return 1
	case "ReadRolePermissions":
		return 2
	case "WriteAttribute":
		return 4
	case "WriteRolePermissions":
		return 8
	case "WriteHistorizing":
		return 16
	case "Read":
		return 32
	case "Write":
		return 64
	case "ReadHistory":
		return 128
	case "InsertHistory":
		return 256
	case "ModifyHistory":
		return 512
	case "DeleteHistory":
		return 1024
	case "ReceiveEvents":
		return 2048
	case "Call":
		return 4096
	case "AddReference":
		return 8192
	case "RemoveReference":
		return 16384
	case "DeleteNode":
		return 32768
	case "AddNode":
		return 65536
	default:
		return 0
	}
}

const (
	PermissionTypeNone                 PermissionType = 0
	PermissionTypeBrowse               PermissionType = 1
	PermissionTypeReadRolePermissions  PermissionType = 2
	PermissionTypeWriteAttribute       PermissionType = 4
	PermissionTypeWriteRolePermissions PermissionType = 8
	PermissionTypeWriteHistorizing     PermissionType = 16
	PermissionTypeRead                 PermissionType = 32
	PermissionTypeWrite                PermissionType = 64
	PermissionTypeReadHistory          PermissionType = 128
	PermissionTypeInsertHistory        PermissionType = 256
	PermissionTypeModifyHistory        PermissionType = 512
	PermissionTypeDeleteHistory        PermissionType = 1024
	PermissionTypeReceiveEvents        PermissionType = 2048
	PermissionTypeCall                 PermissionType = 4096
	PermissionTypeAddReference         PermissionType = 8192
	PermissionTypeRemoveReference      PermissionType = 16384
	PermissionTypeDeleteNode           PermissionType = 32768
	PermissionTypeAddNode              PermissionType = 65536
)

type AccessLevelType uint8

func AccessLevelTypeFromString(s string) AccessLevelType {
	switch s {
	case "None":
		return 0
	case "CurrentRead":
		return 1
	case "CurrentWrite":
		return 2
	case "HistoryRead":
		return 4
	case "HistoryWrite":
		return 8
	case "SemanticChange":
		return 16
	case "StatusWrite":
		return 32
	case "TimestampWrite":
		return 64
	default:
		return 0
	}
}

const (
	AccessLevelTypeNone           AccessLevelType = 0
	AccessLevelTypeCurrentRead    AccessLevelType = 1
	AccessLevelTypeCurrentWrite   AccessLevelType = 2
	AccessLevelTypeHistoryRead    AccessLevelType = 4
	AccessLevelTypeHistoryWrite   AccessLevelType = 8
	AccessLevelTypeSemanticChange AccessLevelType = 16
	AccessLevelTypeStatusWrite    AccessLevelType = 32
	AccessLevelTypeTimestampWrite AccessLevelType = 64
)

type AccessLevelExType uint32

func AccessLevelExTypeFromString(s string) AccessLevelExType {
	switch s {
	case "None":
		return 0
	case "CurrentRead":
		return 1
	case "CurrentWrite":
		return 2
	case "HistoryRead":
		return 4
	case "HistoryWrite":
		return 8
	case "SemanticChange":
		return 16
	case "StatusWrite":
		return 32
	case "TimestampWrite":
		return 64
	case "NonatomicRead":
		return 256
	case "NonatomicWrite":
		return 512
	case "WriteFullArrayOnly":
		return 1024
	default:
		return 0
	}
}

const (
	AccessLevelExTypeNone               AccessLevelExType = 0
	AccessLevelExTypeCurrentRead        AccessLevelExType = 1
	AccessLevelExTypeCurrentWrite       AccessLevelExType = 2
	AccessLevelExTypeHistoryRead        AccessLevelExType = 4
	AccessLevelExTypeHistoryWrite       AccessLevelExType = 8
	AccessLevelExTypeSemanticChange     AccessLevelExType = 16
	AccessLevelExTypeStatusWrite        AccessLevelExType = 32
	AccessLevelExTypeTimestampWrite     AccessLevelExType = 64
	AccessLevelExTypeNonatomicRead      AccessLevelExType = 256
	AccessLevelExTypeNonatomicWrite     AccessLevelExType = 512
	AccessLevelExTypeWriteFullArrayOnly AccessLevelExType = 1024
)

type EventNotifierType uint8

func EventNotifierTypeFromString(s string) EventNotifierType {
	switch s {
	case "None":
		return 0
	case "SubscribeToEvents":
		return 1
	case "HistoryRead":
		return 4
	case "HistoryWrite":
		return 8
	default:
		return 0
	}
}

const (
	EventNotifierTypeNone              EventNotifierType = 0
	EventNotifierTypeSubscribeToEvents EventNotifierType = 1
	EventNotifierTypeHistoryRead       EventNotifierType = 4
	EventNotifierTypeHistoryWrite      EventNotifierType = 8
)

type StructureType uint32

func StructureTypeFromString(s string) StructureType {
	switch s {
	case "Structure":
		return 0
	case "StructureWithOptionalFields":
		return 1
	case "Union":
		return 2
	default:
		return 0
	}
}

const (
	StructureTypeStructure                   StructureType = 0
	StructureTypeStructureWithOptionalFields StructureType = 1
	StructureTypeUnion                       StructureType = 2
)

type ApplicationType uint32

func ApplicationTypeFromString(s string) ApplicationType {
	switch s {
	case "Server":
		return 0
	case "Client":
		return 1
	case "ClientAndServer":
		return 2
	case "DiscoveryServer":
		return 3
	default:
		return 0
	}
}

const (
	ApplicationTypeServer          ApplicationType = 0
	ApplicationTypeClient          ApplicationType = 1
	ApplicationTypeClientAndServer ApplicationType = 2
	ApplicationTypeDiscoveryServer ApplicationType = 3
)

type MessageSecurityMode uint32

func MessageSecurityModeFromString(s string) MessageSecurityMode {
	switch s {
	case "Invalid":
		return 0
	case "None":
		return 1
	case "Sign":
		return 2
	case "SignAndEncrypt":
		return 3
	default:
		return 0
	}
}

const (
	MessageSecurityModeInvalid        MessageSecurityMode = 0
	MessageSecurityModeNone           MessageSecurityMode = 1
	MessageSecurityModeSign           MessageSecurityMode = 2
	MessageSecurityModeSignAndEncrypt MessageSecurityMode = 3
)

type UserTokenType uint32

func UserTokenTypeFromString(s string) UserTokenType {
	switch s {
	case "Anonymous":
		return 0
	case "UserName":
		return 1
	case "Certificate":
		return 2
	case "IssuedToken":
		return 3
	default:
		return 0
	}
}

const (
	UserTokenTypeAnonymous   UserTokenType = 0
	UserTokenTypeUserName    UserTokenType = 1
	UserTokenTypeCertificate UserTokenType = 2
	UserTokenTypeIssuedToken UserTokenType = 3
)

type SecurityTokenRequestType uint32

func SecurityTokenRequestTypeFromString(s string) SecurityTokenRequestType {
	switch s {
	case "Issue":
		return 0
	case "Renew":
		return 1
	default:
		return 0
	}
}

const (
	SecurityTokenRequestTypeIssue SecurityTokenRequestType = 0
	SecurityTokenRequestTypeRenew SecurityTokenRequestType = 1
)

type NodeAttributesMask uint32

func NodeAttributesMaskFromString(s string) NodeAttributesMask {
	switch s {
	case "None":
		return 0
	case "AccessLevel":
		return 1
	case "ArrayDimensions":
		return 2
	case "BrowseName":
		return 4
	case "ContainsNoLoops":
		return 8
	case "DataType":
		return 16
	case "Description":
		return 32
	case "DisplayName":
		return 64
	case "EventNotifier":
		return 128
	case "Executable":
		return 256
	case "Historizing":
		return 512
	case "InverseName":
		return 1024
	case "IsAbstract":
		return 2048
	case "MinimumSamplingInterval":
		return 4096
	case "NodeClass":
		return 8192
	case "NodeId":
		return 16384
	case "Symmetric":
		return 32768
	case "UserAccessLevel":
		return 65536
	case "UserExecutable":
		return 131072
	case "UserWriteMask":
		return 262144
	case "ValueRank":
		return 524288
	case "WriteMask":
		return 1048576
	case "Value":
		return 2097152
	case "DataTypeDefinition":
		return 4194304
	case "RolePermissions":
		return 8388608
	case "AccessRestrictions":
		return 16777216
	case "All":
		return 33554431
	case "BaseNode":
		return 26501220
	case "Object":
		return 26501348
	case "ObjectType":
		return 26503268
	case "Variable":
		return 26571383
	case "VariableType":
		return 28600438
	case "Method":
		return 26632548
	case "ReferenceType":
		return 26537060
	case "View":
		return 26501356
	default:
		return 0
	}
}

const (
	NodeAttributesMaskNone                    NodeAttributesMask = 0
	NodeAttributesMaskAccessLevel             NodeAttributesMask = 1
	NodeAttributesMaskArrayDimensions         NodeAttributesMask = 2
	NodeAttributesMaskBrowseName              NodeAttributesMask = 4
	NodeAttributesMaskContainsNoLoops         NodeAttributesMask = 8
	NodeAttributesMaskDataType                NodeAttributesMask = 16
	NodeAttributesMaskDescription             NodeAttributesMask = 32
	NodeAttributesMaskDisplayName             NodeAttributesMask = 64
	NodeAttributesMaskEventNotifier           NodeAttributesMask = 128
	NodeAttributesMaskExecutable              NodeAttributesMask = 256
	NodeAttributesMaskHistorizing             NodeAttributesMask = 512
	NodeAttributesMaskInverseName             NodeAttributesMask = 1024
	NodeAttributesMaskIsAbstract              NodeAttributesMask = 2048
	NodeAttributesMaskMinimumSamplingInterval NodeAttributesMask = 4096
	NodeAttributesMaskNodeClass               NodeAttributesMask = 8192
	NodeAttributesMaskNodeID                  NodeAttributesMask = 16384
	NodeAttributesMaskSymmetric               NodeAttributesMask = 32768
	NodeAttributesMaskUserAccessLevel         NodeAttributesMask = 65536
	NodeAttributesMaskUserExecutable          NodeAttributesMask = 131072
	NodeAttributesMaskUserWriteMask           NodeAttributesMask = 262144
	NodeAttributesMaskValueRank               NodeAttributesMask = 524288
	NodeAttributesMaskWriteMask               NodeAttributesMask = 1048576
	NodeAttributesMaskValue                   NodeAttributesMask = 2097152
	NodeAttributesMaskDataTypeDefinition      NodeAttributesMask = 4194304
	NodeAttributesMaskRolePermissions         NodeAttributesMask = 8388608
	NodeAttributesMaskAccessRestrictions      NodeAttributesMask = 16777216
	NodeAttributesMaskAll                     NodeAttributesMask = 33554431
	NodeAttributesMaskBaseNode                NodeAttributesMask = 26501220
	NodeAttributesMaskObject                  NodeAttributesMask = 26501348
	NodeAttributesMaskObjectType              NodeAttributesMask = 26503268
	NodeAttributesMaskVariable                NodeAttributesMask = 26571383
	NodeAttributesMaskVariableType            NodeAttributesMask = 28600438
	NodeAttributesMaskMethod                  NodeAttributesMask = 26632548
	NodeAttributesMaskReferenceType           NodeAttributesMask = 26537060
	NodeAttributesMaskView                    NodeAttributesMask = 26501356
)

type AttributeWriteMask uint32

func AttributeWriteMaskFromString(s string) AttributeWriteMask {
	switch s {
	case "None":
		return 0
	case "AccessLevel":
		return 1
	case "ArrayDimensions":
		return 2
	case "BrowseName":
		return 4
	case "ContainsNoLoops":
		return 8
	case "DataType":
		return 16
	case "Description":
		return 32
	case "DisplayName":
		return 64
	case "EventNotifier":
		return 128
	case "Executable":
		return 256
	case "Historizing":
		return 512
	case "InverseName":
		return 1024
	case "IsAbstract":
		return 2048
	case "MinimumSamplingInterval":
		return 4096
	case "NodeClass":
		return 8192
	case "NodeId":
		return 16384
	case "Symmetric":
		return 32768
	case "UserAccessLevel":
		return 65536
	case "UserExecutable":
		return 131072
	case "UserWriteMask":
		return 262144
	case "ValueRank":
		return 524288
	case "WriteMask":
		return 1048576
	case "ValueForVariableType":
		return 2097152
	case "DataTypeDefinition":
		return 4194304
	case "RolePermissions":
		return 8388608
	case "AccessRestrictions":
		return 16777216
	case "AccessLevelEx":
		return 33554432
	default:
		return 0
	}
}

const (
	AttributeWriteMaskNone                    AttributeWriteMask = 0
	AttributeWriteMaskAccessLevel             AttributeWriteMask = 1
	AttributeWriteMaskArrayDimensions         AttributeWriteMask = 2
	AttributeWriteMaskBrowseName              AttributeWriteMask = 4
	AttributeWriteMaskContainsNoLoops         AttributeWriteMask = 8
	AttributeWriteMaskDataType                AttributeWriteMask = 16
	AttributeWriteMaskDescription             AttributeWriteMask = 32
	AttributeWriteMaskDisplayName             AttributeWriteMask = 64
	AttributeWriteMaskEventNotifier           AttributeWriteMask = 128
	AttributeWriteMaskExecutable              AttributeWriteMask = 256
	AttributeWriteMaskHistorizing             AttributeWriteMask = 512
	AttributeWriteMaskInverseName             AttributeWriteMask = 1024
	AttributeWriteMaskIsAbstract              AttributeWriteMask = 2048
	AttributeWriteMaskMinimumSamplingInterval AttributeWriteMask = 4096
	AttributeWriteMaskNodeClass               AttributeWriteMask = 8192
	AttributeWriteMaskNodeID                  AttributeWriteMask = 16384
	AttributeWriteMaskSymmetric               AttributeWriteMask = 32768
	AttributeWriteMaskUserAccessLevel         AttributeWriteMask = 65536
	AttributeWriteMaskUserExecutable          AttributeWriteMask = 131072
	AttributeWriteMaskUserWriteMask           AttributeWriteMask = 262144
	AttributeWriteMaskValueRank               AttributeWriteMask = 524288
	AttributeWriteMaskWriteMask               AttributeWriteMask = 1048576
	AttributeWriteMaskValueForVariableType    AttributeWriteMask = 2097152
	AttributeWriteMaskDataTypeDefinition      AttributeWriteMask = 4194304
	AttributeWriteMaskRolePermissions         AttributeWriteMask = 8388608
	AttributeWriteMaskAccessRestrictions      AttributeWriteMask = 16777216
	AttributeWriteMaskAccessLevelEx           AttributeWriteMask = 33554432
)

type BrowseDirection uint32

func BrowseDirectionFromString(s string) BrowseDirection {
	switch s {
	case "Forward":
		return 0
	case "Inverse":
		return 1
	case "Both":
		return 2
	case "Invalid":
		return 3
	default:
		return 0
	}
}

const (
	BrowseDirectionForward BrowseDirection = 0
	BrowseDirectionInverse BrowseDirection = 1
	BrowseDirectionBoth    BrowseDirection = 2
	BrowseDirectionInvalid BrowseDirection = 3
)

type BrowseResultMask uint32

func BrowseResultMaskFromString(s string) BrowseResultMask {
	switch s {
	case "None":
		return 0
	case "ReferenceTypeId":
		return 1
	case "IsForward":
		return 2
	case "NodeClass":
		return 4
	case "BrowseName":
		return 8
	case "DisplayName":
		return 16
	case "TypeDefinition":
		return 32
	case "All":
		return 63
	case "ReferenceTypeInfo":
		return 3
	case "TargetInfo":
		return 60
	default:
		return 0
	}
}

const (
	BrowseResultMaskNone              BrowseResultMask = 0
	BrowseResultMaskReferenceTypeID   BrowseResultMask = 1
	BrowseResultMaskIsForward         BrowseResultMask = 2
	BrowseResultMaskNodeClass         BrowseResultMask = 4
	BrowseResultMaskBrowseName        BrowseResultMask = 8
	BrowseResultMaskDisplayName       BrowseResultMask = 16
	BrowseResultMaskTypeDefinition    BrowseResultMask = 32
	BrowseResultMaskAll               BrowseResultMask = 63
	BrowseResultMaskReferenceTypeInfo BrowseResultMask = 3
	BrowseResultMaskTargetInfo        BrowseResultMask = 60
)

type FilterOperator uint32

func FilterOperatorFromString(s string) FilterOperator {
	switch s {
	case "Equals":
		return 0
	case "IsNull":
		return 1
	case "GreaterThan":
		return 2
	case "LessThan":
		return 3
	case "GreaterThanOrEqual":
		return 4
	case "LessThanOrEqual":
		return 5
	case "Like":
		return 6
	case "Not":
		return 7
	case "Between":
		return 8
	case "InList":
		return 9
	case "And":
		return 10
	case "Or":
		return 11
	case "Cast":
		return 12
	case "InView":
		return 13
	case "OfType":
		return 14
	case "RelatedTo":
		return 15
	case "BitwiseAnd":
		return 16
	case "BitwiseOr":
		return 17
	default:
		return 0
	}
}

const (
	FilterOperatorEquals             FilterOperator = 0
	FilterOperatorIsNull             FilterOperator = 1
	FilterOperatorGreaterThan        FilterOperator = 2
	FilterOperatorLessThan           FilterOperator = 3
	FilterOperatorGreaterThanOrEqual FilterOperator = 4
	FilterOperatorLessThanOrEqual    FilterOperator = 5
	FilterOperatorLike               FilterOperator = 6
	FilterOperatorNot                FilterOperator = 7
	FilterOperatorBetween            FilterOperator = 8
	FilterOperatorInList             FilterOperator = 9
	FilterOperatorAnd                FilterOperator = 10
	FilterOperatorOr                 FilterOperator = 11
	FilterOperatorCast               FilterOperator = 12
	FilterOperatorInView             FilterOperator = 13
	FilterOperatorOfType             FilterOperator = 14
	FilterOperatorRelatedTo          FilterOperator = 15
	FilterOperatorBitwiseAnd         FilterOperator = 16
	FilterOperatorBitwiseOr          FilterOperator = 17
)

type TimestampsToReturn uint32

func TimestampsToReturnFromString(s string) TimestampsToReturn {
	switch s {
	case "Source":
		return 0
	case "Server":
		return 1
	case "Both":
		return 2
	case "Neither":
		return 3
	case "Invalid":
		return 4
	default:
		return 0
	}
}

const (
	TimestampsToReturnSource  TimestampsToReturn = 0
	TimestampsToReturnServer  TimestampsToReturn = 1
	TimestampsToReturnBoth    TimestampsToReturn = 2
	TimestampsToReturnNeither TimestampsToReturn = 3
	TimestampsToReturnInvalid TimestampsToReturn = 4
)

type HistoryUpdateType uint32

func HistoryUpdateTypeFromString(s string) HistoryUpdateType {
	switch s {
	case "Insert":
		return 1
	case "Replace":
		return 2
	case "Update":
		return 3
	case "Delete":
		return 4
	default:
		return 0
	}
}

const (
	HistoryUpdateTypeInsert  HistoryUpdateType = 1
	HistoryUpdateTypeReplace HistoryUpdateType = 2
	HistoryUpdateTypeUpdate  HistoryUpdateType = 3
	HistoryUpdateTypeDelete  HistoryUpdateType = 4
)

type PerformUpdateType uint32

func PerformUpdateTypeFromString(s string) PerformUpdateType {
	switch s {
	case "Insert":
		return 1
	case "Replace":
		return 2
	case "Update":
		return 3
	case "Remove":
		return 4
	default:
		return 0
	}
}

const (
	PerformUpdateTypeInsert  PerformUpdateType = 1
	PerformUpdateTypeReplace PerformUpdateType = 2
	PerformUpdateTypeUpdate  PerformUpdateType = 3
	PerformUpdateTypeRemove  PerformUpdateType = 4
)

type MonitoringMode uint32

func MonitoringModeFromString(s string) MonitoringMode {
	switch s {
	case "Disabled":
		return 0
	case "Sampling":
		return 1
	case "Reporting":
		return 2
	default:
		return 0
	}
}

const (
	MonitoringModeDisabled  MonitoringMode = 0
	MonitoringModeSampling  MonitoringMode = 1
	MonitoringModeReporting MonitoringMode = 2
)

type DataChangeTrigger uint32

func DataChangeTriggerFromString(s string) DataChangeTrigger {
	switch s {
	case "Status":
		return 0
	case "StatusValue":
		return 1
	case "StatusValueTimestamp":
		return 2
	default:
		return 0
	}
}

const (
	DataChangeTriggerStatus               DataChangeTrigger = 0
	DataChangeTriggerStatusValue          DataChangeTrigger = 1
	DataChangeTriggerStatusValueTimestamp DataChangeTrigger = 2
)

type DeadbandType uint32

func DeadbandTypeFromString(s string) DeadbandType {
	switch s {
	case "None":
		return 0
	case "Absolute":
		return 1
	case "Percent":
		return 2
	default:
		return 0
	}
}

const (
	DeadbandTypeNone     DeadbandType = 0
	DeadbandTypeAbsolute DeadbandType = 1
	DeadbandTypePercent  DeadbandType = 2
)

type RedundancySupport uint32

func RedundancySupportFromString(s string) RedundancySupport {
	switch s {
	case "None":
		return 0
	case "Cold":
		return 1
	case "Warm":
		return 2
	case "Hot":
		return 3
	case "Transparent":
		return 4
	case "HotAndMirrored":
		return 5
	default:
		return 0
	}
}

const (
	RedundancySupportNone           RedundancySupport = 0
	RedundancySupportCold           RedundancySupport = 1
	RedundancySupportWarm           RedundancySupport = 2
	RedundancySupportHot            RedundancySupport = 3
	RedundancySupportTransparent    RedundancySupport = 4
	RedundancySupportHotAndMirrored RedundancySupport = 5
)

type ServerState uint32

func ServerStateFromString(s string) ServerState {
	switch s {
	case "Running":
		return 0
	case "Failed":
		return 1
	case "NoConfiguration":
		return 2
	case "Suspended":
		return 3
	case "Shutdown":
		return 4
	case "Test":
		return 5
	case "CommunicationFault":
		return 6
	case "Unknown":
		return 7
	default:
		return 0
	}
}

const (
	ServerStateRunning            ServerState = 0
	ServerStateFailed             ServerState = 1
	ServerStateNoConfiguration    ServerState = 2
	ServerStateSuspended          ServerState = 3
	ServerStateShutdown           ServerState = 4
	ServerStateTest               ServerState = 5
	ServerStateCommunicationFault ServerState = 6
	ServerStateUnknown            ServerState = 7
)

type ModelChangeStructureVerbMask uint32

func ModelChangeStructureVerbMaskFromString(s string) ModelChangeStructureVerbMask {
	switch s {
	case "NodeAdded":
		return 1
	case "NodeDeleted":
		return 2
	case "ReferenceAdded":
		return 4
	case "ReferenceDeleted":
		return 8
	case "DataTypeChanged":
		return 16
	default:
		return 0
	}
}

const (
	ModelChangeStructureVerbMaskNodeAdded        ModelChangeStructureVerbMask = 1
	ModelChangeStructureVerbMaskNodeDeleted      ModelChangeStructureVerbMask = 2
	ModelChangeStructureVerbMaskReferenceAdded   ModelChangeStructureVerbMask = 4
	ModelChangeStructureVerbMaskReferenceDeleted ModelChangeStructureVerbMask = 8
	ModelChangeStructureVerbMaskDataTypeChanged  ModelChangeStructureVerbMask = 16
)

type AxisScaleEnumeration uint32

func AxisScaleEnumerationFromString(s string) AxisScaleEnumeration {
	switch s {
	case "Linear":
		return 0
	case "Log":
		return 1
	case "Ln":
		return 2
	default:
		return 0
	}
}

const (
	AxisScaleEnumerationLinear AxisScaleEnumeration = 0
	AxisScaleEnumerationLog    AxisScaleEnumeration = 1
	AxisScaleEnumerationLn     AxisScaleEnumeration = 2
)

type ExceptionDeviationFormat uint32

func ExceptionDeviationFormatFromString(s string) ExceptionDeviationFormat {
	switch s {
	case "AbsoluteValue":
		return 0
	case "PercentOfValue":
		return 1
	case "PercentOfRange":
		return 2
	case "PercentOfEURange":
		return 3
	case "Unknown":
		return 4
	default:
		return 0
	}
}

const (
	ExceptionDeviationFormatAbsoluteValue    ExceptionDeviationFormat = 0
	ExceptionDeviationFormatPercentOfValue   ExceptionDeviationFormat = 1
	ExceptionDeviationFormatPercentOfRange   ExceptionDeviationFormat = 2
	ExceptionDeviationFormatPercentOfEURange ExceptionDeviationFormat = 3
	ExceptionDeviationFormatUnknown          ExceptionDeviationFormat = 4
)
