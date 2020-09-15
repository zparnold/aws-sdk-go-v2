// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Add a parameter to the system.
func (c *Client) PutParameter(ctx context.Context, params *PutParameterInput, optFns ...func(*Options)) (*PutParameterOutput, error) {
	stack := middleware.NewStack("PutParameter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutParameterMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutParameterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutParameter(options.Region), middleware.Before)

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
			OperationName: "PutParameter",
			Err:           err,
		}
	}
	out := result.(*PutParameterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutParameterInput struct {
	// The type of parameter that you want to add to the system. SecureString is not
	// currently supported for AWS CloudFormation templates or in the China Regions.
	// Items in a StringList must be separated by a comma (,). You can't use other
	// punctuation or special character to escape items in the list. If you have a
	// parameter value that requires a comma, then use the String data type. Specifying
	// a parameter type is not required when updating a parameter. You must specify a
	// parameter type when creating a parameter.
	Type types.ParameterType
	// Optional metadata that you assign to a resource. Tags enable you to categorize a
	// resource in different ways, such as by purpose, owner, or environment. For
	// example, you might want to tag a Systems Manager parameter to identify the type
	// of resource to which it applies, the environment, or the type of configuration
	// data referenced by the parameter. In this case, you could specify the following
	// key name/value pairs:  <ul> <li> <p> <code>Key=Resource,Value=S3bucket</code>
	// </p> </li> <li> <p> <code>Key=OS,Value=Windows</code> </p> </li> <li> <p>
	// <code>Key=ParameterType,Value=LicenseKey</code> </p> </li> </ul> <note> <p>To
	// add tags to an existing Systems Manager parameter, use the
	// <a>AddTagsToResource</a> action.</p> </note>
	Tags []*types.Tag
	// One or more policies to apply to a parameter. This action takes a JSON array.
	// Parameter Store supports the following policy types:  <p>Expiration: This policy
	// deletes the parameter after it expires. When you create the policy,  you specify
	// the expiration date. You can update the expiration date and time by updating the
	// policy. Updating the parameter does not affect the expiration date and time.
	// When the expiration time is reached, Parameter Store deletes the parameter.
	// ExpirationNotification: This policy triggers an event in Amazon CloudWatch
	// Events that notifies you about the expiration. By using this policy, you can
	// receive notification before or after the expiration time is reached, in units of
	// days or hours. NoChangeNotification: This policy triggers a CloudWatch event if
	// a parameter has not been modified for a specified period of time. This policy
	// type is useful when, for example, a secret needs to be changed within a period
	// of time, but it has not been changed.  <p>All existing policies are preserved
	// until you send new policies or an empty policy. For more  information about
	// parameter policies, see Assigning parameter policies
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-policies.html).
	Policies *string
	// Overwrite an existing parameter. If not specified, will default to "false".
	Overwrite *bool
	// A regular expression used to validate the parameter value. For example, for
	// String types with values restricted to numbers, you can specify the following:
	// AllowedPattern=^\d+$
	AllowedPattern *string
	// The parameter value that you want to add to the system. Standard parameters have
	// a value limit of 4 KB. Advanced parameters have a value limit of 8 KB.
	Value *string
	// The KMS Key ID that you want to use to encrypt a parameter. Either the default
	// AWS Key Management Service (AWS KMS) key automatically assigned to your AWS
	// account or a custom key. Required for parameters that use the SecureString data
	// type. If you don't specify a key ID, the system uses the default key associated
	// with your AWS account.
	//
	//     * To use your default AWS KMS key, choose the
	// SecureString data type, and do not specify the Key ID when you create the
	// parameter. The system automatically populates Key ID with your default KMS
	// key.
	//
	//     * To use a custom KMS key, choose the SecureString data type with the
	// Key ID parameter.
	KeyId *string
	// The data type for a String parameter. Supported data types include plain text
	// and Amazon Machine Image IDs. The following data type values are supported.
	//
	//
	// * text
	//
	//     * aws:ec2:image
	//
	// When you create a String parameter and specify
	// aws:ec2:image, Systems Manager validates the parameter value is in the required
	// format, such as ami-12345abcdeEXAMPLE, and that the specified AMI is available
	// in your AWS account. For more information, see Native parameter support for
	// Amazon Machine Image IDs
	// (http://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	// in the AWS Systems Manager User Guide.
	DataType *string
	// The fully qualified name of the parameter that you want to add to the system.
	// The fully qualified name includes the complete hierarchy of the parameter path
	// and name. For parameters in a hierarchy, you must include a leading forward
	// slash character (/) when you create or reference a parameter. For example:
	// /Dev/DBServer/MySQL/db-string13 Naming Constraints:
	//
	//     * Parameter names are
	// case sensitive.
	//
	//     * A parameter name must be unique within an AWS Region
	//
	//
	// * A parameter name can't be prefixed with "aws" or "ssm" (case-insensitive).
	//
	//
	// * Parameter names can include only the following symbols and letters:
	// a-zA-Z0-9_.-/
	//
	//     * A parameter name can't include spaces.
	//
	//     * Parameter
	// hierarchies are limited to a maximum depth of fifteen levels.
	//
	// For additional
	// information about valid values for parameter names, see About requirements and
	// constraints for parameter names
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html)
	// in the AWS Systems Manager User Guide. The maximum length constraint listed
	// below includes capacity for additional system attributes that are not part of
	// the name. The maximum length for a parameter name, including the full length of
	// the parameter ARN, is 1011 characters. For example, the length of the following
	// parameter name is 65 characters, not 20 characters:
	// arn:aws:ssm:us-east-2:111122223333:parameter/ExampleParameterName
	Name *string
	// The parameter tier to assign to a parameter. Parameter Store offers a standard
	// tier and an advanced tier for parameters. Standard parameters have a content
	// size limit of 4 KB and can't be configured to use parameter policies. You can
	// create a maximum of 10,000 standard parameters for each Region in an AWS
	// account. Standard parameters are offered at no additional cost. Advanced
	// parameters have a content size limit of 8 KB and can be configured to use
	// parameter policies. You can create a maximum of 100,000 advanced parameters for
	// each Region in an AWS account. Advanced parameters incur a charge. For more
	// information, see Standard and advanced parameter tiers
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html)
	// in the AWS Systems Manager User Guide. You can change a standard parameter to an
	// advanced parameter any time. But you can't revert an advanced parameter to a
	// standard parameter. Reverting an advanced parameter to a standard parameter
	// would result in data loss because the system would truncate the size of the
	// parameter from 8 KB to 4 KB. Reverting would also remove any policies attached
	// to the parameter. Lastly, advanced parameters use a different form of encryption
	// than standard parameters. If you no longer need an advanced parameter, or if you
	// no longer want to incur charges for an advanced parameter, you must delete it
	// and recreate it as a new standard parameter. Using the Default Tier
	// Configuration In PutParameter requests, you can specify the tier to create the
	// parameter in. Whenever you specify a tier in the request, Parameter Store
	// creates or updates the parameter according to that request. However, if you do
	// not specify a tier in a request, Parameter Store assigns the tier based on the
	// current Parameter Store default tier configuration. The default tier when you
	// begin using Parameter Store is the standard-parameter tier. If you use the
	// advanced-parameter tier, you can specify one of the following as the default:
	//
	//
	// * Advanced: With this option, Parameter Store evaluates all requests as advanced
	// parameters.
	//
	//     * Intelligent-Tiering: With this option, Parameter Store
	// evaluates each request to determine if the parameter is standard or advanced. If
	// the request doesn't include any options that require an advanced parameter, the
	// parameter is created in the standard-parameter tier. If one or more options
	// requiring an advanced parameter are included in the request, Parameter Store
	// create a parameter in the advanced-parameter tier. This approach helps control
	// your parameter-related costs by always creating standard parameters unless an
	// advanced parameter is necessary.
	//
	// Options that require an advanced parameter
	// include the following:
	//
	//     * The content size of the parameter is more than 4
	// KB.
	//
	//     * The parameter uses a parameter policy.
	//
	//     * More than 10,000
	// parameters already exist in your AWS account in the current Region.
	//
	// For more
	// information about configuring the default tier option, see Specifying a default
	// parameter tier
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/ps-default-tier.html)
	// in the AWS Systems Manager User Guide.
	Tier types.ParameterTier
	// Information about the parameter that you want to add to the system. Optional but
	// recommended. Do not enter personally identifiable information in this field.
	Description *string
}

type PutParameterOutput struct {
	// The tier assigned to the parameter.
	Tier types.ParameterTier
	// The new version number of a parameter. If you edit a parameter value, Parameter
	// Store automatically creates a new version and assigns this new version a unique
	// ID. You can reference a parameter version ID in API actions or in Systems
	// Manager documents (SSM documents). By default, if you don't specify a specific
	// version, the system returns the latest parameter value when a parameter is
	// called.
	Version *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutParameterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutParameter{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutParameter{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutParameter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "PutParameter",
	}
}
