// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a health check. Amazon Route 53 does not prevent you from deleting a
// health check even if the health check is associated with one or more resource
// record sets. If you delete a health check and you don't update the associated
// resource record sets, the future status of the health check can't be predicted
// and may change. This will affect the routing of DNS queries for your DNS
// failover configuration. For more information, see Replacing and Deleting Health
// Checks
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-deleting.html#health-checks-deleting.html)
// in the Amazon Route 53 Developer Guide.  <p>If you're using AWS Cloud Map and
// you configured Cloud Map to create a Route 53 health check when you register an
// instance, you can't use the Route 53 <code>DeleteHealthCheck</code> command to
// delete the health check. The health check is deleted automatically when you
// deregister the instance; there can be a delay of several hours before the health
// check is deleted from Route 53. </p>
func (c *Client) DeleteHealthCheck(ctx context.Context, params *DeleteHealthCheckInput, optFns ...func(*Options)) (*DeleteHealthCheckOutput, error) {
	if params == nil {
		params = &DeleteHealthCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteHealthCheck", params, optFns, addOperationDeleteHealthCheckMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This action deletes a health check.
type DeleteHealthCheckInput struct {

	// The ID of the health check that you want to delete.
	//
	// This member is required.
	HealthCheckId *string
}

// An empty element.
type DeleteHealthCheckOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteHealthCheckMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteHealthCheck{}, middleware.After)
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
	addOpDeleteHealthCheckValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteHealthCheck(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteHealthCheck(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "DeleteHealthCheck",
	}
}
