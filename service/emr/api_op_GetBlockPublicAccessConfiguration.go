// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the Amazon EMR block public access configuration for your AWS account in
// the current Region. For more information see Configure Block Public Access for
// Amazon EMR
// (https://docs.aws.amazon.com/emr/latest/ManagementGuide/configure-block-public-access.html)
// in the Amazon EMR Management Guide.
func (c *Client) GetBlockPublicAccessConfiguration(ctx context.Context, params *GetBlockPublicAccessConfigurationInput, optFns ...func(*Options)) (*GetBlockPublicAccessConfigurationOutput, error) {
	stack := middleware.NewStack("GetBlockPublicAccessConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetBlockPublicAccessConfigurationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBlockPublicAccessConfiguration(options.Region), middleware.Before)

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
			OperationName: "GetBlockPublicAccessConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetBlockPublicAccessConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBlockPublicAccessConfigurationInput struct {
}

type GetBlockPublicAccessConfigurationOutput struct {
	// Properties that describe the AWS principal that created the
	// BlockPublicAccessConfiguration using the PutBlockPublicAccessConfiguration
	// action as well as the date and time that the configuration was created. Each
	// time a configuration for block public access is updated, Amazon EMR updates this
	// metadata.
	BlockPublicAccessConfigurationMetadata *types.BlockPublicAccessConfigurationMetadata
	// A configuration for Amazon EMR block public access. The configuration applies to
	// all clusters created in your account for the current Region. The configuration
	// specifies whether block public access is enabled. If block public access is
	// enabled, security groups associated with the cluster cannot have rules that
	// allow inbound traffic from 0.0.0.0/0 or ::/0 on a port, unless the port is
	// specified as an exception using PermittedPublicSecurityGroupRuleRanges in the
	// BlockPublicAccessConfiguration. By default, Port 22 (SSH) is an exception, and
	// public access is allowed on this port. You can change this by updating the block
	// public access configuration to remove the exception. For accounts that created
	// clusters in a Region before November 25, 2019, block public access is disabled
	// by default in that Region. To use this feature, you must manually enable and
	// configure it. For accounts that did not create an EMR cluster in a Region before
	// this date, block public access is enabled by default in that Region.
	BlockPublicAccessConfiguration *types.BlockPublicAccessConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetBlockPublicAccessConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetBlockPublicAccessConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetBlockPublicAccessConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBlockPublicAccessConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "GetBlockPublicAccessConfiguration",
	}
}
