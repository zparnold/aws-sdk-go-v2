// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds health-based detection to the Shield Advanced protection for a resource.
// Shield Advanced health-based detection uses the health of your AWS resource to
// improve responsiveness and accuracy in attack detection and mitigation. You
// define the health check in Route 53 and then associate it with your Shield
// Advanced protection. For more information, see Shield Advanced Health-Based
// Detection
// (https://docs.aws.amazon.com/waf/latest/developerguide/ddos-overview.html#ddos-advanced-health-check-option)
// in the AWS WAF and AWS Shield Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) AssociateHealthCheck(ctx context.Context, params *AssociateHealthCheckInput, optFns ...func(*Options)) (*AssociateHealthCheckOutput, error) {
	stack := middleware.NewStack("AssociateHealthCheck", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateHealthCheckMiddlewares(stack)
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
	addOpAssociateHealthCheckValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateHealthCheck(options.Region), middleware.Before)

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
			OperationName: "AssociateHealthCheck",
			Err:           err,
		}
	}
	out := result.(*AssociateHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateHealthCheckInput struct {
	// The Amazon Resource Name (ARN) of the health check to associate with the
	// protection.
	HealthCheckArn *string
	// The unique identifier (ID) for the Protection () object to add the health check
	// association to.
	ProtectionId *string
}

type AssociateHealthCheckOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateHealthCheckMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateHealthCheck{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateHealthCheck{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateHealthCheck(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "AssociateHealthCheck",
	}
}