// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified security groups or all of your security groups. A
// security group is for use with instances either in the EC2-Classic platform or
// in a specific VPC. For more information, see Amazon EC2 Security Groups
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html)
// in the Amazon Elastic Compute Cloud User Guide and Security Groups for Your VPC
// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_SecurityGroups.html)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) DescribeSecurityGroups(ctx context.Context, params *DescribeSecurityGroupsInput, optFns ...func(*Options)) (*DescribeSecurityGroupsOutput, error) {
	stack := middleware.NewStack("DescribeSecurityGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeSecurityGroupsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSecurityGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "DescribeSecurityGroups",
			Err:           err,
		}
	}
	out := result.(*DescribeSecurityGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSecurityGroupsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters. If using multiple filters for rules, the results include security
	// groups for which any combination of rules - not necessarily a single rule -
	// match all filters.
	//
	//     * description - The description of the security group.
	//
	//
	// * egress.ip-permission.cidr - An IPv4 CIDR block for an outbound security group
	// rule.
	//
	//     * egress.ip-permission.from-port - For an outbound rule, the start of
	// port range for the TCP and UDP protocols, or an ICMP type number.
	//
	//     *
	// egress.ip-permission.group-id - The ID of a security group that has been
	// referenced in an outbound security group rule.
	//
	//     *
	// egress.ip-permission.group-name - The name of a security group that has been
	// referenced in an outbound security group rule.
	//
	//     *
	// egress.ip-permission.ipv6-cidr - An IPv6 CIDR block for an outbound security
	// group rule.
	//
	//     * egress.ip-permission.prefix-list-id - The ID of a prefix list
	// to which a security group rule allows outbound access.
	//
	//     *
	// egress.ip-permission.protocol - The IP protocol for an outbound security group
	// rule (tcp | udp | icmp or a protocol number).
	//
	//     *
	// egress.ip-permission.to-port - For an outbound rule, the end of port range for
	// the TCP and UDP protocols, or an ICMP code.
	//
	//     * egress.ip-permission.user-id
	// - The ID of an AWS account that has been referenced in an outbound security
	// group rule.
	//
	//     * group-id - The ID of the security group.
	//
	//     * group-name -
	// The name of the security group.
	//
	//     * ip-permission.cidr - An IPv4 CIDR block
	// for an inbound security group rule.
	//
	//     * ip-permission.from-port - For an
	// inbound rule, the start of port range for the TCP and UDP protocols, or an ICMP
	// type number.
	//
	//     * ip-permission.group-id - The ID of a security group that has
	// been referenced in an inbound security group rule.
	//
	//     *
	// ip-permission.group-name - The name of a security group that has been referenced
	// in an inbound security group rule.
	//
	//     * ip-permission.ipv6-cidr - An IPv6 CIDR
	// block for an inbound security group rule.
	//
	//     * ip-permission.prefix-list-id -
	// The ID of a prefix list from which a security group rule allows inbound
	// access.
	//
	//     * ip-permission.protocol - The IP protocol for an inbound security
	// group rule (tcp | udp | icmp or a protocol number).
	//
	//     * ip-permission.to-port
	// - For an inbound rule, the end of port range for the TCP and UDP protocols, or
	// an ICMP code.
	//
	//     * ip-permission.user-id - The ID of an AWS account that has
	// been referenced in an inbound security group rule.
	//
	//     * owner-id - The AWS
	// account ID of the owner of the security group.
	//
	//     * tag: - The key/value
	// combination of a tag assigned to the resource. Use the tag key in the filter
	// name and the tag value as the filter value. For example, to find all resources
	// that have a tag with the key Owner and the value TeamA, specify tag:Owner for
	// the filter name and TeamA for the filter value.
	//
	//     * tag-key - The key of a
	// tag assigned to the resource. Use this filter to find all resources assigned a
	// tag with a specific key, regardless of the tag value.
	//
	//     * vpc-id - The ID of
	// the VPC specified when the security group was created.
	Filters []*types.Filter

	// The IDs of the security groups. Required for security groups in a nondefault
	// VPC. Default: Describes all your security groups.
	GroupIds []*string

	// [EC2-Classic and default VPC only] The names of the security groups. You can
	// specify either the security group name or the security group ID. For security
	// groups in a nondefault VPC, use the group-name filter to describe security
	// groups by name. Default: Describes all your security groups.
	GroupNames []*string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value. This
	// value can be between 5 and 1000. If this parameter is not specified, then all
	// results are returned.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string
}

type DescribeSecurityGroupsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the security groups.
	SecurityGroups []*types.SecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeSecurityGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeSecurityGroups{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSecurityGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSecurityGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSecurityGroups",
	}
}
