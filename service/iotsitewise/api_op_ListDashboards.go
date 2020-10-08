// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a paginated list of dashboards for an AWS IoT SiteWise Monitor
// project.
func (c *Client) ListDashboards(ctx context.Context, params *ListDashboardsInput, optFns ...func(*Options)) (*ListDashboardsOutput, error) {
	if params == nil {
		params = &ListDashboardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDashboards", params, optFns, addOperationListDashboardsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDashboardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDashboardsInput struct {

	// The ID of the project.
	//
	// This member is required.
	ProjectId *string

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string
}

type ListDashboardsOutput struct {

	// A list that summarizes each dashboard in the project.
	//
	// This member is required.
	DashboardSummaries []*types.DashboardSummary

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDashboardsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDashboards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDashboards{}, middleware.After)
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
	addOpListDashboardsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDashboards(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDashboards(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "ListDashboards",
	}
}
