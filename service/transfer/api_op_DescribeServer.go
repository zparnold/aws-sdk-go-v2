// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a file transfer protocol-enabled server that you specify by passing
// the ServerId parameter.  <p>The response contains a description of a server's
// properties. When you set <code>EndpointType</code> to VPC, the response will
// contain the <code>EndpointDetails</code>.</p>
func (c *Client) DescribeServer(ctx context.Context, params *DescribeServerInput, optFns ...func(*Options)) (*DescribeServerOutput, error) {
	stack := middleware.NewStack("DescribeServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeServerMiddlewares(stack)
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
	addOpDescribeServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeServer(options.Region), middleware.Before)
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
			OperationName: "DescribeServer",
			Err:           err,
		}
	}
	out := result.(*DescribeServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeServerInput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server.
	//
	// This member is required.
	ServerId *string
}

type DescribeServerOutput struct {

	// An array containing the properties of a file transfer protocol-enabled server
	// with the ServerID you specified.
	//
	// This member is required.
	Server *types.DescribedServer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DescribeServer",
	}
}
