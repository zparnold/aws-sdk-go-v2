// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccountRoleStatus string

// Enum values for AccountRoleStatus
const (
	AccountRoleStatusReady           AccountRoleStatus = "READY"
	AccountRoleStatusCreating        AccountRoleStatus = "CREATING"
	AccountRoleStatusPendingdeletion AccountRoleStatus = "PENDING_DELETION"
	AccountRoleStatusDeleting        AccountRoleStatus = "DELETING"
	AccountRoleStatusDeleted         AccountRoleStatus = "DELETED"
)

type CustomerPolicyScopeIdType string

// Enum values for CustomerPolicyScopeIdType
const (
	CustomerPolicyScopeIdTypeAccount  CustomerPolicyScopeIdType = "ACCOUNT"
	CustomerPolicyScopeIdTypeOrg_unit CustomerPolicyScopeIdType = "ORG_UNIT"
)

type DependentServiceName string

// Enum values for DependentServiceName
const (
	DependentServiceNameAwsconfig              DependentServiceName = "AWSCONFIG"
	DependentServiceNameAwswaf                 DependentServiceName = "AWSWAF"
	DependentServiceNameAwsshieldadvanced      DependentServiceName = "AWSSHIELD_ADVANCED"
	DependentServiceNameAwsvirtualprivatecloud DependentServiceName = "AWSVPC"
)

type PolicyComplianceStatusType string

// Enum values for PolicyComplianceStatusType
const (
	PolicyComplianceStatusTypeCompliant    PolicyComplianceStatusType = "COMPLIANT"
	PolicyComplianceStatusTypeNoncompliant PolicyComplianceStatusType = "NON_COMPLIANT"
)

type RemediationActionType string

// Enum values for RemediationActionType
const (
	RemediationActionTypeRemove RemediationActionType = "REMOVE"
	RemediationActionTypeModify RemediationActionType = "MODIFY"
)

type SecurityServiceType string

// Enum values for SecurityServiceType
const (
	SecurityServiceTypeWaf                           SecurityServiceType = "WAF"
	SecurityServiceTypeWafv2                         SecurityServiceType = "WAFV2"
	SecurityServiceTypeShield_advanced               SecurityServiceType = "SHIELD_ADVANCED"
	SecurityServiceTypeSecurity_groups_common        SecurityServiceType = "SECURITY_GROUPS_COMMON"
	SecurityServiceTypeSecurity_groups_content_audit SecurityServiceType = "SECURITY_GROUPS_CONTENT_AUDIT"
	SecurityServiceTypeSecurity_groups_usage_audit   SecurityServiceType = "SECURITY_GROUPS_USAGE_AUDIT"
)

type ViolationReason string

// Enum values for ViolationReason
const (
	ViolationReasonWebaclmissingrulegroup                  ViolationReason = "WEB_ACL_MISSING_RULE_GROUP"
	ViolationReasonResourcemissingwebacl                   ViolationReason = "RESOURCE_MISSING_WEB_ACL"
	ViolationReasonResourceincorrectwebacl                 ViolationReason = "RESOURCE_INCORRECT_WEB_ACL"
	ViolationReasonResourcemissingshieldprotection         ViolationReason = "RESOURCE_MISSING_SHIELD_PROTECTION"
	ViolationReasonResourcemissingwebaclorshieldprotection ViolationReason = "RESOURCE_MISSING_WEB_ACL_OR_SHIELD_PROTECTION"
	ViolationReasonResourcemissingsecuritygroup            ViolationReason = "RESOURCE_MISSING_SECURITY_GROUP"
	ViolationReasonResourceviolatesauditsecuritygroup      ViolationReason = "RESOURCE_VIOLATES_AUDIT_SECURITY_GROUP"
	ViolationReasonSecuritygroupunused                     ViolationReason = "SECURITY_GROUP_UNUSED"
	ViolationReasonSecuritygroupredundant                  ViolationReason = "SECURITY_GROUP_REDUNDANT"
)