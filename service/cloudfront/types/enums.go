// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CachePolicyCookieBehavior string

// Enum values for CachePolicyCookieBehavior
const (
	CachePolicyCookieBehaviorNone      CachePolicyCookieBehavior = "none"
	CachePolicyCookieBehaviorWhitelist CachePolicyCookieBehavior = "whitelist"
	CachePolicyCookieBehaviorAllexcept CachePolicyCookieBehavior = "allExcept"
	CachePolicyCookieBehaviorAll       CachePolicyCookieBehavior = "all"
)

type CachePolicyHeaderBehavior string

// Enum values for CachePolicyHeaderBehavior
const (
	CachePolicyHeaderBehaviorNone      CachePolicyHeaderBehavior = "none"
	CachePolicyHeaderBehaviorWhitelist CachePolicyHeaderBehavior = "whitelist"
)

type CachePolicyQueryStringBehavior string

// Enum values for CachePolicyQueryStringBehavior
const (
	CachePolicyQueryStringBehaviorNone      CachePolicyQueryStringBehavior = "none"
	CachePolicyQueryStringBehaviorWhitelist CachePolicyQueryStringBehavior = "whitelist"
	CachePolicyQueryStringBehaviorAllexcept CachePolicyQueryStringBehavior = "allExcept"
	CachePolicyQueryStringBehaviorAll       CachePolicyQueryStringBehavior = "all"
)

type CachePolicyType string

// Enum values for CachePolicyType
const (
	CachePolicyTypeManaged CachePolicyType = "managed"
	CachePolicyTypeCustom  CachePolicyType = "custom"
)

type CertificateSource string

// Enum values for CertificateSource
const (
	CertificateSourceCloudfront CertificateSource = "cloudfront"
	CertificateSourceIam        CertificateSource = "iam"
	CertificateSourceAcm        CertificateSource = "acm"
)

type EventType string

// Enum values for EventType
const (
	EventTypeViewerRequest  EventType = "viewer-request"
	EventTypeViewerResponse EventType = "viewer-response"
	EventTypeOriginRequest  EventType = "origin-request"
	EventTypeOriginResponse EventType = "origin-response"
)

type Format string

// Enum values for Format
const (
	FormatUrlencoded Format = "URLEncoded"
)

type GeoRestrictionType string

// Enum values for GeoRestrictionType
const (
	GeoRestrictionTypeBlacklist GeoRestrictionType = "blacklist"
	GeoRestrictionTypeWhitelist GeoRestrictionType = "whitelist"
	GeoRestrictionTypeNone      GeoRestrictionType = "none"
)

type HttpVersion string

// Enum values for HttpVersion
const (
	HttpVersionHttp11 HttpVersion = "http1.1"
	HttpVersionHttp2  HttpVersion = "http2"
)

type ICPRecordalStatus string

// Enum values for ICPRecordalStatus
const (
	ICPRecordalStatusApproved  ICPRecordalStatus = "APPROVED"
	ICPRecordalStatusSuspended ICPRecordalStatus = "SUSPENDED"
	ICPRecordalStatusPending   ICPRecordalStatus = "PENDING"
)

type ItemSelection string

// Enum values for ItemSelection
const (
	ItemSelectionNone      ItemSelection = "none"
	ItemSelectionWhitelist ItemSelection = "whitelist"
	ItemSelectionAll       ItemSelection = "all"
)

type Method string

// Enum values for Method
const (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodPatch   Method = "PATCH"
	MethodOptions Method = "OPTIONS"
	MethodDelete  Method = "DELETE"
)

type MinimumProtocolVersion string

// Enum values for MinimumProtocolVersion
const (
	MinimumProtocolVersionSslv3       MinimumProtocolVersion = "SSLv3"
	MinimumProtocolVersionTlsv1       MinimumProtocolVersion = "TLSv1"
	MinimumProtocolVersionTlsv1_2016  MinimumProtocolVersion = "TLSv1_2016"
	MinimumProtocolVersionTlsv11_2016 MinimumProtocolVersion = "TLSv1.1_2016"
	MinimumProtocolVersionTlsv12_2018 MinimumProtocolVersion = "TLSv1.2_2018"
	MinimumProtocolVersionTlsv12_2019 MinimumProtocolVersion = "TLSv1.2_2019"
)

type OriginProtocolPolicy string

// Enum values for OriginProtocolPolicy
const (
	OriginProtocolPolicyHttpOnly    OriginProtocolPolicy = "http-only"
	OriginProtocolPolicyMatchViewer OriginProtocolPolicy = "match-viewer"
	OriginProtocolPolicyHttpsOnly   OriginProtocolPolicy = "https-only"
)

type OriginRequestPolicyCookieBehavior string

// Enum values for OriginRequestPolicyCookieBehavior
const (
	OriginRequestPolicyCookieBehaviorNone      OriginRequestPolicyCookieBehavior = "none"
	OriginRequestPolicyCookieBehaviorWhitelist OriginRequestPolicyCookieBehavior = "whitelist"
	OriginRequestPolicyCookieBehaviorAll       OriginRequestPolicyCookieBehavior = "all"
)

type OriginRequestPolicyHeaderBehavior string

// Enum values for OriginRequestPolicyHeaderBehavior
const (
	OriginRequestPolicyHeaderBehaviorNone                            OriginRequestPolicyHeaderBehavior = "none"
	OriginRequestPolicyHeaderBehaviorWhitelist                       OriginRequestPolicyHeaderBehavior = "whitelist"
	OriginRequestPolicyHeaderBehaviorAllviewer                       OriginRequestPolicyHeaderBehavior = "allViewer"
	OriginRequestPolicyHeaderBehaviorAllviewerandwhitelistcloudfront OriginRequestPolicyHeaderBehavior = "allViewerAndWhitelistCloudFront"
)

type OriginRequestPolicyQueryStringBehavior string

// Enum values for OriginRequestPolicyQueryStringBehavior
const (
	OriginRequestPolicyQueryStringBehaviorNone      OriginRequestPolicyQueryStringBehavior = "none"
	OriginRequestPolicyQueryStringBehaviorWhitelist OriginRequestPolicyQueryStringBehavior = "whitelist"
	OriginRequestPolicyQueryStringBehaviorAll       OriginRequestPolicyQueryStringBehavior = "all"
)

type OriginRequestPolicyType string

// Enum values for OriginRequestPolicyType
const (
	OriginRequestPolicyTypeManaged OriginRequestPolicyType = "managed"
	OriginRequestPolicyTypeCustom  OriginRequestPolicyType = "custom"
)

type PriceClass string

// Enum values for PriceClass
const (
	PriceClassPriceclass_100 PriceClass = "PriceClass_100"
	PriceClassPriceclass_200 PriceClass = "PriceClass_200"
	PriceClassPriceclass_all PriceClass = "PriceClass_All"
)

type SslProtocol string

// Enum values for SslProtocol
const (
	SslProtocolSslv3  SslProtocol = "SSLv3"
	SslProtocolTlsv1  SslProtocol = "TLSv1"
	SslProtocolTlsv11 SslProtocol = "TLSv1.1"
	SslProtocolTlsv12 SslProtocol = "TLSv1.2"
)

type SSLSupportMethod string

// Enum values for SSLSupportMethod
const (
	SSLSupportMethodSniOnly SSLSupportMethod = "sni-only"
	SSLSupportMethodVip     SSLSupportMethod = "vip"
)

type ViewerProtocolPolicy string

// Enum values for ViewerProtocolPolicy
const (
	ViewerProtocolPolicyAllowAll        ViewerProtocolPolicy = "allow-all"
	ViewerProtocolPolicyHttpsOnly       ViewerProtocolPolicy = "https-only"
	ViewerProtocolPolicyRedirectToHttps ViewerProtocolPolicy = "redirect-to-https"
)
