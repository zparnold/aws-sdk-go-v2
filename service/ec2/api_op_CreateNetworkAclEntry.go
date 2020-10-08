// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an entry (a rule) in a network ACL with the specified rule number. Each
// network ACL has a set of numbered ingress rules and a separate set of numbered
// egress rules. When determining whether a packet should be allowed in or out of a
// subnet associated with the ACL, we process the entries in the ACL according to
// the rule numbers, in ascending order. Each network ACL has a set of ingress
// rules and a separate set of egress rules. We recommend that you leave room
// between the rule numbers (for example, 100, 110, 120, ...), and not number them
// one right after the other (for example, 101, 102, 103, ...). This makes it
// easier to add a rule between existing ones without having to renumber the rules.
// After you add an entry, you can't modify it; you must either replace it, or
// create an entry and delete the old one. For more information about network ACLs,
// see Network ACLs
// (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html) in the Amazon
// Virtual Private Cloud User Guide.
func (c *Client) CreateNetworkAclEntry(ctx context.Context, params *CreateNetworkAclEntryInput, optFns ...func(*Options)) (*CreateNetworkAclEntryOutput, error) {
	if params == nil {
		params = &CreateNetworkAclEntryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNetworkAclEntry", params, optFns, addOperationCreateNetworkAclEntryMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNetworkAclEntryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNetworkAclEntryInput struct {

	// Indicates whether this is an egress rule (rule is applied to traffic leaving the
	// subnet).
	//
	// This member is required.
	Egress *bool

	// The ID of the network ACL.
	//
	// This member is required.
	NetworkAclId *string

	// The protocol number. A value of "-1" means all protocols. If you specify "-1" or
	// a protocol number other than "6" (TCP), "17" (UDP), or "1" (ICMP), traffic on
	// all ports is allowed, regardless of any ports or ICMP types or codes that you
	// specify. If you specify protocol "58" (ICMPv6) and specify an IPv4 CIDR block,
	// traffic for all ICMP types and codes allowed, regardless of any that you
	// specify. If you specify protocol "58" (ICMPv6) and specify an IPv6 CIDR block,
	// you must specify an ICMP type and code.
	//
	// This member is required.
	Protocol *string

	// Indicates whether to allow or deny the traffic that matches the rule.
	//
	// This member is required.
	RuleAction types.RuleAction

	// The rule number for the entry (for example, 100). ACL entries are processed in
	// ascending order by rule number. Constraints: Positive integer from 1 to 32766.
	// The range 32767 to 65535 is reserved for internal use.
	//
	// This member is required.
	RuleNumber *int32

	// The IPv4 network range to allow or deny, in CIDR notation (for example
	// 172.16.0.0/24). We modify the specified CIDR block to its canonical form; for
	// example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18.
	CidrBlock *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// ICMP protocol: The ICMP or ICMPv6 type and code. Required if specifying protocol
	// 1 (ICMP) or protocol 58 (ICMPv6) with an IPv6 CIDR block.
	IcmpTypeCode *types.IcmpTypeCode

	// The IPv6 network range to allow or deny, in CIDR notation (for example
	// 2001:db8:1234:1a00::/64).
	Ipv6CidrBlock *string

	// TCP or UDP protocols: The range of ports the rule applies to. Required if
	// specifying protocol 6 (TCP) or 17 (UDP).
	PortRange *types.PortRange
}

type CreateNetworkAclEntryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateNetworkAclEntryMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateNetworkAclEntry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateNetworkAclEntry{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateNetworkAclEntryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetworkAclEntry(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateNetworkAclEntry(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateNetworkAclEntry",
	}
}
