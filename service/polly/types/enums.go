// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Engine string

// Enum values for Engine
const (
	EngineStandard Engine = "standard"
	EngineNeural   Engine = "neural"
)

type Gender string

// Enum values for Gender
const (
	GenderFemale Gender = "Female"
	GenderMale   Gender = "Male"
)

type LanguageCode string

// Enum values for LanguageCode
const (
	LanguageCodeArb     LanguageCode = "arb"
	LanguageCodeCmnCn   LanguageCode = "cmn-CN"
	LanguageCodeCyGb    LanguageCode = "cy-GB"
	LanguageCodeDaDk    LanguageCode = "da-DK"
	LanguageCodeDeDe    LanguageCode = "de-DE"
	LanguageCodeEnAu    LanguageCode = "en-AU"
	LanguageCodeEnGb    LanguageCode = "en-GB"
	LanguageCodeEnGbWls LanguageCode = "en-GB-WLS"
	LanguageCodeEnIn    LanguageCode = "en-IN"
	LanguageCodeEnUs    LanguageCode = "en-US"
	LanguageCodeEsEs    LanguageCode = "es-ES"
	LanguageCodeEsMx    LanguageCode = "es-MX"
	LanguageCodeEsUs    LanguageCode = "es-US"
	LanguageCodeFrCa    LanguageCode = "fr-CA"
	LanguageCodeFrFr    LanguageCode = "fr-FR"
	LanguageCodeIsIs    LanguageCode = "is-IS"
	LanguageCodeItIt    LanguageCode = "it-IT"
	LanguageCodeJaJp    LanguageCode = "ja-JP"
	LanguageCodeHiIn    LanguageCode = "hi-IN"
	LanguageCodeKoKr    LanguageCode = "ko-KR"
	LanguageCodeNbNo    LanguageCode = "nb-NO"
	LanguageCodeNlNl    LanguageCode = "nl-NL"
	LanguageCodePlPl    LanguageCode = "pl-PL"
	LanguageCodePtBr    LanguageCode = "pt-BR"
	LanguageCodePtPt    LanguageCode = "pt-PT"
	LanguageCodeRoRo    LanguageCode = "ro-RO"
	LanguageCodeRuRu    LanguageCode = "ru-RU"
	LanguageCodeSvSe    LanguageCode = "sv-SE"
	LanguageCodeTrTr    LanguageCode = "tr-TR"
)

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatJson       OutputFormat = "json"
	OutputFormatMp3        OutputFormat = "mp3"
	OutputFormatOgg_vorbis OutputFormat = "ogg_vorbis"
	OutputFormatPcm        OutputFormat = "pcm"
)

type SpeechMarkType string

// Enum values for SpeechMarkType
const (
	SpeechMarkTypeSentence SpeechMarkType = "sentence"
	SpeechMarkTypeSsml     SpeechMarkType = "ssml"
	SpeechMarkTypeViseme   SpeechMarkType = "viseme"
	SpeechMarkTypeWord     SpeechMarkType = "word"
)

type TaskStatus string

// Enum values for TaskStatus
const (
	TaskStatusScheduled   TaskStatus = "scheduled"
	TaskStatusIn_progress TaskStatus = "inProgress"
	TaskStatusCompleted   TaskStatus = "completed"
	TaskStatusFailed      TaskStatus = "failed"
)

type TextType string

// Enum values for TextType
const (
	TextTypeSsml TextType = "ssml"
	TextTypeText TextType = "text"
)

type VoiceId string

// Enum values for VoiceId
const (
	VoiceIdAditi     VoiceId = "Aditi"
	VoiceIdAmy       VoiceId = "Amy"
	VoiceIdAstrid    VoiceId = "Astrid"
	VoiceIdBianca    VoiceId = "Bianca"
	VoiceIdBrian     VoiceId = "Brian"
	VoiceIdCamila    VoiceId = "Camila"
	VoiceIdCarla     VoiceId = "Carla"
	VoiceIdCarmen    VoiceId = "Carmen"
	VoiceIdCeline    VoiceId = "Celine"
	VoiceIdChantal   VoiceId = "Chantal"
	VoiceIdConchita  VoiceId = "Conchita"
	VoiceIdCristiano VoiceId = "Cristiano"
	VoiceIdDora      VoiceId = "Dora"
	VoiceIdEmma      VoiceId = "Emma"
	VoiceIdEnrique   VoiceId = "Enrique"
	VoiceIdEwa       VoiceId = "Ewa"
	VoiceIdFiliz     VoiceId = "Filiz"
	VoiceIdGeraint   VoiceId = "Geraint"
	VoiceIdGiorgio   VoiceId = "Giorgio"
	VoiceIdGwyneth   VoiceId = "Gwyneth"
	VoiceIdHans      VoiceId = "Hans"
	VoiceIdInes      VoiceId = "Ines"
	VoiceIdIvy       VoiceId = "Ivy"
	VoiceIdJacek     VoiceId = "Jacek"
	VoiceIdJan       VoiceId = "Jan"
	VoiceIdJoanna    VoiceId = "Joanna"
	VoiceIdJoey      VoiceId = "Joey"
	VoiceIdJustin    VoiceId = "Justin"
	VoiceIdKarl      VoiceId = "Karl"
	VoiceIdKendra    VoiceId = "Kendra"
	VoiceIdKevin     VoiceId = "Kevin"
	VoiceIdKimberly  VoiceId = "Kimberly"
	VoiceIdLea       VoiceId = "Lea"
	VoiceIdLiv       VoiceId = "Liv"
	VoiceIdLotte     VoiceId = "Lotte"
	VoiceIdLucia     VoiceId = "Lucia"
	VoiceIdLupe      VoiceId = "Lupe"
	VoiceIdMads      VoiceId = "Mads"
	VoiceIdMaja      VoiceId = "Maja"
	VoiceIdMarlene   VoiceId = "Marlene"
	VoiceIdMathieu   VoiceId = "Mathieu"
	VoiceIdMatthew   VoiceId = "Matthew"
	VoiceIdMaxim     VoiceId = "Maxim"
	VoiceIdMia       VoiceId = "Mia"
	VoiceIdMiguel    VoiceId = "Miguel"
	VoiceIdMizuki    VoiceId = "Mizuki"
	VoiceIdNaja      VoiceId = "Naja"
	VoiceIdNicole    VoiceId = "Nicole"
	VoiceIdPenelope  VoiceId = "Penelope"
	VoiceIdRaveena   VoiceId = "Raveena"
	VoiceIdRicardo   VoiceId = "Ricardo"
	VoiceIdRuben     VoiceId = "Ruben"
	VoiceIdRussell   VoiceId = "Russell"
	VoiceIdSalli     VoiceId = "Salli"
	VoiceIdSeoyeon   VoiceId = "Seoyeon"
	VoiceIdTakumi    VoiceId = "Takumi"
	VoiceIdTatyana   VoiceId = "Tatyana"
	VoiceIdVicki     VoiceId = "Vicki"
	VoiceIdVitoria   VoiceId = "Vitoria"
	VoiceIdZeina     VoiceId = "Zeina"
	VoiceIdZhiyu     VoiceId = "Zhiyu"
)
