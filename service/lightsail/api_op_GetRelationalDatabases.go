// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about all of your databases in Amazon Lightsail.
func (c *Client) GetRelationalDatabases(ctx context.Context, params *GetRelationalDatabasesInput, optFns ...func(*Options)) (*GetRelationalDatabasesOutput, error) {
	if params == nil {
		params = &GetRelationalDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabases", params, optFns, addOperationGetRelationalDatabasesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabasesInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetRelationalDatabases request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string
}

type GetRelationalDatabasesOutput struct {

	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetRelationalDatabases request and specify the
	// next page token using the pageToken parameter.
	NextPageToken *string

	// An object describing the result of your get relational databases request.
	RelationalDatabases []*types.RelationalDatabase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRelationalDatabasesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabases{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabases(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRelationalDatabases(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabases",
	}
}
