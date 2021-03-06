// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DataSource string

// Enum values for DataSource
const (
	DataSourceEvent                DataSource = "EVENT"
	DataSourceModel_score          DataSource = "MODEL_SCORE"
	DataSourceExternal_model_score DataSource = "EXTERNAL_MODEL_SCORE"
)

type DataType string

// Enum values for DataType
const (
	DataTypeString  DataType = "STRING"
	DataTypeInteger DataType = "INTEGER"
	DataTypeFloat   DataType = "FLOAT"
	DataTypeBoolean DataType = "BOOLEAN"
)

type DetectorVersionStatus string

// Enum values for DetectorVersionStatus
const (
	DetectorVersionStatusDraft    DetectorVersionStatus = "DRAFT"
	DetectorVersionStatusActive   DetectorVersionStatus = "ACTIVE"
	DetectorVersionStatusInactive DetectorVersionStatus = "INACTIVE"
)

type Language string

// Enum values for Language
const (
	LanguageDetectorpl Language = "DETECTORPL"
)

type ModelEndpointStatus string

// Enum values for ModelEndpointStatus
const (
	ModelEndpointStatusAssociated  ModelEndpointStatus = "ASSOCIATED"
	ModelEndpointStatusDissociated ModelEndpointStatus = "DISSOCIATED"
)

type ModelInputDataFormat string

// Enum values for ModelInputDataFormat
const (
	ModelInputDataFormatCsv  ModelInputDataFormat = "TEXT_CSV"
	ModelInputDataFormatJson ModelInputDataFormat = "APPLICATION_JSON"
)

type ModelOutputDataFormat string

// Enum values for ModelOutputDataFormat
const (
	ModelOutputDataFormatCsv       ModelOutputDataFormat = "TEXT_CSV"
	ModelOutputDataFormatJsonlines ModelOutputDataFormat = "APPLICATION_JSONLINES"
)

type ModelSource string

// Enum values for ModelSource
const (
	ModelSourceSagemaker ModelSource = "SAGEMAKER"
)

type ModelTypeEnum string

// Enum values for ModelTypeEnum
const (
	ModelTypeEnumOnline_fraud_insights ModelTypeEnum = "ONLINE_FRAUD_INSIGHTS"
)

type ModelVersionStatus string

// Enum values for ModelVersionStatus
const (
	ModelVersionStatusActive   ModelVersionStatus = "ACTIVE"
	ModelVersionStatusInactive ModelVersionStatus = "INACTIVE"
)

type RuleExecutionMode string

// Enum values for RuleExecutionMode
const (
	RuleExecutionModeAll_matched   RuleExecutionMode = "ALL_MATCHED"
	RuleExecutionModeFirst_matched RuleExecutionMode = "FIRST_MATCHED"
)

type TrainingDataSourceEnum string

// Enum values for TrainingDataSourceEnum
const (
	TrainingDataSourceEnumExternal_events TrainingDataSourceEnum = "EXTERNAL_EVENTS"
)
