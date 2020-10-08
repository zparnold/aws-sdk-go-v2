// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Simulate how a set of IAM policies attached to an IAM entity works with a list
// of API operations and AWS resources to determine the policies' effective
// permissions. The entity can be an IAM user, group, or role. If you specify a
// user, then the simulation also includes all of the policies that are attached to
// groups that the user belongs to. You can optionally include a list of one or
// more additional policies specified as strings to include in the simulation. If
// you want to simulate only policies specified as strings, use
// SimulateCustomPolicy () instead. You can also optionally include one
// resource-based policy to be evaluated with each of the resources included in the
// simulation. The simulation does not perform the API operations; it only checks
// the authorization to determine if the simulated policies allow or deny the
// operations. Note: This API discloses information about the permissions granted
// to other users. If you do not want users to see other user's permissions, then
// consider allowing them to use SimulateCustomPolicy () instead. Context keys are
// variables maintained by AWS and its services that provide details about the
// context of an API query request. You can use the Condition element of an IAM
// policy to evaluate context keys. To get the list of context keys that the
// policies require for correct simulation, use GetContextKeysForPrincipalPolicy
// (). If the output is long, you can use the MaxItems and Marker parameters to
// paginate the results.
func (c *Client) SimulatePrincipalPolicy(ctx context.Context, params *SimulatePrincipalPolicyInput, optFns ...func(*Options)) (*SimulatePrincipalPolicyOutput, error) {
	if params == nil {
		params = &SimulatePrincipalPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SimulatePrincipalPolicy", params, optFns, addOperationSimulatePrincipalPolicyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*SimulatePrincipalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SimulatePrincipalPolicyInput struct {

	// A list of names of API operations to evaluate in the simulation. Each operation
	// is evaluated for each resource. Each operation must include the service
	// identifier, such as iam:CreateUser.
	//
	// This member is required.
	ActionNames []*string

	// The Amazon Resource Name (ARN) of a user, group, or role whose policies you want
	// to include in the simulation. If you specify a user, group, or role, the
	// simulation includes all policies that are associated with that entity. If you
	// specify a user, the simulation also includes all policies that are attached to
	// any groups the user belongs to. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	//
	// This member is required.
	PolicySourceArn *string

	// The ARN of the IAM user that you want to specify as the simulated caller of the
	// API operations. If you do not specify a CallerArn, it defaults to the ARN of the
	// user that you specify in PolicySourceArn, if you specified a user. If you
	// include both a PolicySourceArn (for example,
	// arn:aws:iam::123456789012:user/David) and a CallerArn (for example,
	// arn:aws:iam::123456789012:user/Bob), the result is that you simulate calling the
	// API operations as Bob, as if Bob had David's policies. You can specify only the
	// ARN of an IAM user. You cannot specify the ARN of an assumed role, federated
	// user, or a service principal. CallerArn is required if you include a
	// ResourcePolicy and the PolicySourceArn is not the ARN for an IAM user. This is
	// required so that the resource-based policy's Principal element has a value to
	// use in evaluating the policy. For more information about ARNs, see Amazon
	// Resource Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	CallerArn *string

	// A list of context keys and corresponding values for the simulation to use.
	// Whenever a context key is evaluated in one of the simulated IAM permissions
	// policies, the corresponding value is supplied.
	ContextEntries []*types.ContextEntry

	// Use this parameter only when paginating results and only after you receive a
	// response indicating that the results are truncated. Set it to the value of the
	// Marker element in the response that you received to indicate where the next call
	// should start.
	Marker *string

	// Use this only when paginating results to indicate the maximum number of items
	// you want in the response. If additional items exist beyond the maximum you
	// specify, the IsTruncated response element is true. If you do not include this
	// parameter, the number of items defaults to 100. Note that IAM might return fewer
	// results, even when there are more results available. In that case, the
	// IsTruncated response element returns true, and Marker contains a value to
	// include in the subsequent call that tells the service where to continue from.
	MaxItems *int32

	// The IAM permissions boundary policy to simulate. The permissions boundary sets
	// the maximum permissions that the entity can have. You can input only one
	// permissions boundary when you pass a policy to this operation. An IAM entity can
	// only have one permissions boundary in effect at a time. For example, if a
	// permissions boundary is attached to an entity and you pass in a different
	// permissions boundary policy using this parameter, then the new permissions
	// boundary policy is used for the simulation. For more information about
	// permissions boundaries, see Permissions Boundaries for IAM Entities
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
	// in the IAM User Guide. The policy input is specified as a string containing the
	// complete, valid JSON text of a permissions boundary policy. The regex pattern
	// (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of
	// characters consisting of the following:
	//
	//     * Any printable ASCII character
	// ranging from the space character (\u0020) through the end of the ASCII character
	// range
	//
	//     * The printable characters in the Basic Latin and Latin-1 Supplement
	// character set (through \u00FF)
	//
	//     * The special characters tab (\u0009), line
	// feed (\u000A), and carriage return (\u000D)
	PermissionsBoundaryPolicyInputList []*string

	// An optional list of additional policy documents to include in the simulation.
	// Each document is specified as a string containing the complete, valid JSON text
	// of an IAM policy. The regex pattern (http://wikipedia.org/wiki/regex) used to
	// validate this parameter is a string of characters consisting of the following:
	//
	//
	// * Any printable ASCII character ranging from the space character (\u0020)
	// through the end of the ASCII character range
	//
	//     * The printable characters in
	// the Basic Latin and Latin-1 Supplement character set (through \u00FF)
	//
	//     * The
	// special characters tab (\u0009), line feed (\u000A), and carriage return
	// (\u000D)
	PolicyInputList []*string

	// A list of ARNs of AWS resources to include in the simulation. If this parameter
	// is not provided, then the value defaults to * (all resources). Each API in the
	// ActionNames parameter is evaluated for each resource in this list. The
	// simulation determines the access result (allowed or denied) of each combination
	// and reports it in the response. The simulation does not automatically retrieve
	// policies for the specified resources. If you want to include a resource policy
	// in the simulation, then you must include the policy as a string in the
	// ResourcePolicy parameter. For more information about ARNs, see Amazon Resource
	// Names (ARNs) and AWS Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference.
	ResourceArns []*string

	// Specifies the type of simulation to run. Different API operations that support
	// resource-based policies require different combinations of resources. By
	// specifying the type of simulation to run, you enable the policy simulator to
	// enforce the presence of the required resources to ensure reliable simulation
	// results. If your simulation does not match one of the following scenarios, then
	// you can omit this parameter. The following list shows each of the supported
	// scenario values and the resources that you must define to run the simulation.
	// Each of the EC2 scenarios requires that you specify instance, image, and
	// security group resources. If your scenario includes an EBS volume, then you must
	// specify that volume as a resource. If the EC2 scenario includes VPC, then you
	// must supply the network interface resource. If it includes an IP subnet, then
	// you must specify the subnet resource. For more information on the EC2 scenario
	// options, see Supported Platforms
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html)
	// in the Amazon EC2 User Guide.
	//
	//     * EC2-Classic-InstanceStore instance, image,
	// security group
	//
	//     * EC2-Classic-EBS instance, image, security group, volume
	//
	//
	// * EC2-VPC-InstanceStore instance, image, security group, network interface
	//
	//
	// * EC2-VPC-InstanceStore-Subnet instance, image, security group, network
	// interface, subnet
	//
	//     * EC2-VPC-EBS instance, image, security group, network
	// interface, volume
	//
	//     * EC2-VPC-EBS-Subnet instance, image, security group,
	// network interface, subnet, volume
	ResourceHandlingOption *string

	// An AWS account ID that specifies the owner of any simulated resource that does
	// not identify its owner in the resource ARN. Examples of resource ARNs include an
	// S3 bucket or object. If ResourceOwner is specified, it is also used as the
	// account owner of any ResourcePolicy included in the simulation. If the
	// ResourceOwner parameter is not specified, then the owner of the resources and
	// the resource policy defaults to the account of the identity provided in
	// CallerArn. This parameter is required only if you specify a resource-based
	// policy and account that owns the resource is different from the account that
	// owns the simulated calling user CallerArn.
	ResourceOwner *string

	// A resource-based policy to include in the simulation provided as a string. Each
	// resource in the simulation is treated as if it had this policy attached. You can
	// include only one resource-based policy in a simulation. The regex pattern
	// (http://wikipedia.org/wiki/regex) used to validate this parameter is a string of
	// characters consisting of the following:
	//
	//     * Any printable ASCII character
	// ranging from the space character (\u0020) through the end of the ASCII character
	// range
	//
	//     * The printable characters in the Basic Latin and Latin-1 Supplement
	// character set (through \u00FF)
	//
	//     * The special characters tab (\u0009), line
	// feed (\u000A), and carriage return (\u000D)
	ResourcePolicy *string
}

// Contains the response to a successful SimulatePrincipalPolicy () or
// SimulateCustomPolicy () request.
type SimulatePrincipalPolicyOutput struct {

	// The results of the simulation.
	EvaluationResults []*types.EvaluationResult

	// A flag that indicates whether there are more items to return. If your results
	// were truncated, you can make a subsequent pagination request using the Marker
	// request parameter to retrieve more items. Note that IAM might return fewer than
	// the MaxItems number of results even when there are more results available. We
	// recommend that you check IsTruncated after every call to ensure that you receive
	// all your results.
	IsTruncated *bool

	// When IsTruncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent pagination request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSimulatePrincipalPolicyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSimulatePrincipalPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSimulatePrincipalPolicy{}, middleware.After)
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
	addOpSimulatePrincipalPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSimulatePrincipalPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSimulatePrincipalPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "SimulatePrincipalPolicy",
	}
}
