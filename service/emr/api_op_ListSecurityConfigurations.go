// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the security configurations visible to this account, providing their
// creation dates and times, and their names. This call returns a maximum of 50
// clusters per call, but returns a marker to track the paging of the cluster list
// across multiple ListSecurityConfigurations calls.
func (c *Client) ListSecurityConfigurations(ctx context.Context, params *ListSecurityConfigurationsInput, optFns ...func(*Options)) (*ListSecurityConfigurationsOutput, error) {
	stack := middleware.NewStack("ListSecurityConfigurations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListSecurityConfigurationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityConfigurations(options.Region), middleware.Before)
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
			OperationName: "ListSecurityConfigurations",
			Err:           err,
		}
	}
	out := result.(*ListSecurityConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityConfigurationsInput struct {

	// The pagination token that indicates the set of results to retrieve.
	Marker *string
}

type ListSecurityConfigurationsOutput struct {

	// A pagination token that indicates the next set of results to retrieve. Include
	// the marker in the next ListSecurityConfiguration call to retrieve the next page
	// of results, if required.
	Marker *string

	// The creation date and time, and name, of each security configuration.
	SecurityConfigurations []*types.SecurityConfigurationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListSecurityConfigurationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListSecurityConfigurations{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSecurityConfigurations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSecurityConfigurations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListSecurityConfigurations",
	}
}
