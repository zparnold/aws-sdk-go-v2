// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

type Action string

// Enum values for Action
const (
	ActionOpenApp  Action = "OPEN_APP"
	ActionDeepLink Action = "DEEP_LINK"
	ActionUrl      Action = "URL"
)

func (enum Action) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Action) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AttributeType string

// Enum values for AttributeType
const (
	AttributeTypeInclusive AttributeType = "INCLUSIVE"
	AttributeTypeExclusive AttributeType = "EXCLUSIVE"
)

func (enum AttributeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AttributeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CampaignStatus string

// Enum values for CampaignStatus
const (
	CampaignStatusScheduled      CampaignStatus = "SCHEDULED"
	CampaignStatusExecuting      CampaignStatus = "EXECUTING"
	CampaignStatusPendingNextRun CampaignStatus = "PENDING_NEXT_RUN"
	CampaignStatusCompleted      CampaignStatus = "COMPLETED"
	CampaignStatusPaused         CampaignStatus = "PAUSED"
	CampaignStatusDeleted        CampaignStatus = "DELETED"
)

func (enum CampaignStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CampaignStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChannelType string

// Enum values for ChannelType
const (
	ChannelTypePush            ChannelType = "PUSH"
	ChannelTypeGcm             ChannelType = "GCM"
	ChannelTypeApns            ChannelType = "APNS"
	ChannelTypeApnsSandbox     ChannelType = "APNS_SANDBOX"
	ChannelTypeApnsVoip        ChannelType = "APNS_VOIP"
	ChannelTypeApnsVoipSandbox ChannelType = "APNS_VOIP_SANDBOX"
	ChannelTypeAdm             ChannelType = "ADM"
	ChannelTypeSms             ChannelType = "SMS"
	ChannelTypeVoice           ChannelType = "VOICE"
	ChannelTypeEmail           ChannelType = "EMAIL"
	ChannelTypeBaidu           ChannelType = "BAIDU"
	ChannelTypeCustom          ChannelType = "CUSTOM"
)

func (enum ChannelType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChannelType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeliveryStatus string

// Enum values for DeliveryStatus
const (
	DeliveryStatusSuccessful       DeliveryStatus = "SUCCESSFUL"
	DeliveryStatusThrottled        DeliveryStatus = "THROTTLED"
	DeliveryStatusTemporaryFailure DeliveryStatus = "TEMPORARY_FAILURE"
	DeliveryStatusPermanentFailure DeliveryStatus = "PERMANENT_FAILURE"
	DeliveryStatusUnknownFailure   DeliveryStatus = "UNKNOWN_FAILURE"
	DeliveryStatusOptOut           DeliveryStatus = "OPT_OUT"
	DeliveryStatusDuplicate        DeliveryStatus = "DUPLICATE"
)

func (enum DeliveryStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeliveryStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DimensionType string

// Enum values for DimensionType
const (
	DimensionTypeInclusive DimensionType = "INCLUSIVE"
	DimensionTypeExclusive DimensionType = "EXCLUSIVE"
)

func (enum DimensionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DimensionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Duration string

// Enum values for Duration
const (
	DurationHr24  Duration = "HR_24"
	DurationDay7  Duration = "DAY_7"
	DurationDay14 Duration = "DAY_14"
	DurationDay30 Duration = "DAY_30"
)

func (enum Duration) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Duration) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EndpointTypesElement string

// Enum values for EndpointTypesElement
const (
	EndpointTypesElementPush            EndpointTypesElement = "PUSH"
	EndpointTypesElementGcm             EndpointTypesElement = "GCM"
	EndpointTypesElementApns            EndpointTypesElement = "APNS"
	EndpointTypesElementApnsSandbox     EndpointTypesElement = "APNS_SANDBOX"
	EndpointTypesElementApnsVoip        EndpointTypesElement = "APNS_VOIP"
	EndpointTypesElementApnsVoipSandbox EndpointTypesElement = "APNS_VOIP_SANDBOX"
	EndpointTypesElementAdm             EndpointTypesElement = "ADM"
	EndpointTypesElementSms             EndpointTypesElement = "SMS"
	EndpointTypesElementVoice           EndpointTypesElement = "VOICE"
	EndpointTypesElementEmail           EndpointTypesElement = "EMAIL"
	EndpointTypesElementBaidu           EndpointTypesElement = "BAIDU"
	EndpointTypesElementCustom          EndpointTypesElement = "CUSTOM"
)

func (enum EndpointTypesElement) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EndpointTypesElement) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FilterType string

// Enum values for FilterType
const (
	FilterTypeSystem   FilterType = "SYSTEM"
	FilterTypeEndpoint FilterType = "ENDPOINT"
)

func (enum FilterType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FilterType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Format string

// Enum values for Format
const (
	FormatCsv  Format = "CSV"
	FormatJson Format = "JSON"
)

func (enum Format) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Format) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Frequency string

// Enum values for Frequency
const (
	FrequencyOnce    Frequency = "ONCE"
	FrequencyHourly  Frequency = "HOURLY"
	FrequencyDaily   Frequency = "DAILY"
	FrequencyWeekly  Frequency = "WEEKLY"
	FrequencyMonthly Frequency = "MONTHLY"
	FrequencyEvent   Frequency = "EVENT"
)

func (enum Frequency) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Frequency) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Include string

// Enum values for Include
const (
	IncludeAll  Include = "ALL"
	IncludeAny  Include = "ANY"
	IncludeNone Include = "NONE"
)

func (enum Include) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Include) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusCreated                    JobStatus = "CREATED"
	JobStatusPreparingForInitialization JobStatus = "PREPARING_FOR_INITIALIZATION"
	JobStatusInitializing               JobStatus = "INITIALIZING"
	JobStatusProcessing                 JobStatus = "PROCESSING"
	JobStatusPendingJob                 JobStatus = "PENDING_JOB"
	JobStatusCompleting                 JobStatus = "COMPLETING"
	JobStatusCompleted                  JobStatus = "COMPLETED"
	JobStatusFailing                    JobStatus = "FAILING"
	JobStatusFailed                     JobStatus = "FAILED"
)

func (enum JobStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MessageType string

// Enum values for MessageType
const (
	MessageTypeTransactional MessageType = "TRANSACTIONAL"
	MessageTypePromotional   MessageType = "PROMOTIONAL"
)

func (enum MessageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MessageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Mode string

// Enum values for Mode
const (
	ModeDelivery Mode = "DELIVERY"
	ModeFilter   Mode = "FILTER"
)

func (enum Mode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Mode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Operator string

// Enum values for Operator
const (
	OperatorAll Operator = "ALL"
	OperatorAny Operator = "ANY"
)

func (enum Operator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Operator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RecencyType string

// Enum values for RecencyType
const (
	RecencyTypeActive   RecencyType = "ACTIVE"
	RecencyTypeInactive RecencyType = "INACTIVE"
)

func (enum RecencyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RecencyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SegmentType string

// Enum values for SegmentType
const (
	SegmentTypeDimensional SegmentType = "DIMENSIONAL"
	SegmentTypeImport      SegmentType = "IMPORT"
)

func (enum SegmentType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SegmentType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeAll  SourceType = "ALL"
	SourceTypeAny  SourceType = "ANY"
	SourceTypeNone SourceType = "NONE"
)

func (enum SourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type State string

// Enum values for State
const (
	StateDraft     State = "DRAFT"
	StateActive    State = "ACTIVE"
	StateCompleted State = "COMPLETED"
	StateCancelled State = "CANCELLED"
	StateClosed    State = "CLOSED"
)

func (enum State) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum State) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TemplateType string

// Enum values for TemplateType
const (
	TemplateTypeEmail TemplateType = "EMAIL"
	TemplateTypeSms   TemplateType = "SMS"
	TemplateTypeVoice TemplateType = "VOICE"
	TemplateTypePush  TemplateType = "PUSH"
)

func (enum TemplateType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TemplateType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Type string

// Enum values for Type
const (
	TypeAll  Type = "ALL"
	TypeAny  Type = "ANY"
	TypeNone Type = "NONE"
)

func (enum Type) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Type) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}