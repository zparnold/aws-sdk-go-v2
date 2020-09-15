// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the CA certificates registered for your AWS account. The results are
// paginated with a default page size of 25. You can use the returned marker to
// retrieve additional results.
func (c *Client) ListCACertificates(ctx context.Context, params *ListCACertificatesInput, optFns ...func(*Options)) (*ListCACertificatesOutput, error) {
	stack := middleware.NewStack("ListCACertificates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListCACertificatesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListCACertificates(options.Region), middleware.Before)

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
			OperationName: "ListCACertificates",
			Err:           err,
		}
	}
	out := result.(*ListCACertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for the ListCACertificates operation.
type ListCACertificatesInput struct {
	// Determines the order of the results.
	AscendingOrder *bool
	// The marker for the next set of results.
	Marker *string
	// The result page size.
	PageSize *int32
}

// The output from the ListCACertificates operation.
type ListCACertificatesOutput struct {
	// The CA certificates registered in your AWS account.
	Certificates []*types.CACertificate
	// The current position within the list of CA certificates.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListCACertificatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListCACertificates{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListCACertificates{}, middleware.After)
}

func newServiceMetadataMiddleware_opListCACertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListCACertificates",
	}
}
