// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all valid regions for Amazon Lightsail. Use the include
// availability zones parameter to also return the Availability Zones in a region.
func (c *Client) GetRegions(ctx context.Context, params *GetRegionsInput, optFns ...func(*Options)) (*GetRegionsOutput, error) {
	stack := middleware.NewStack("GetRegions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRegionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRegions(options.Region), middleware.Before)
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
			OperationName: "GetRegions",
			Err:           err,
		}
	}
	out := result.(*GetRegionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRegionsInput struct {

	// A Boolean value indicating whether to also include Availability Zones in your
	// get regions request. Availability Zones are indicated with a letter: e.g.,
	// us-east-2a.
	IncludeAvailabilityZones *bool

	// >A Boolean value indicating whether to also include Availability Zones for
	// databases in your get regions request. Availability Zones are indicated with a
	// letter (e.g., us-east-2a).
	IncludeRelationalDatabaseAvailabilityZones *bool
}

type GetRegionsOutput struct {

	// An array of key-value pairs containing information about your get regions
	// request.
	Regions []*types.Region

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRegionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRegions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRegions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRegions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRegions",
	}
}
