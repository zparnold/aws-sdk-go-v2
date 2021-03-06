// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the backend environments for an Amplify app.
func (c *Client) ListBackendEnvironments(ctx context.Context, params *ListBackendEnvironmentsInput, optFns ...func(*Options)) (*ListBackendEnvironmentsOutput, error) {
	stack := middleware.NewStack("ListBackendEnvironments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListBackendEnvironmentsMiddlewares(stack)
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
	addOpListBackendEnvironmentsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBackendEnvironments(options.Region), middleware.Before)
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
			OperationName: "ListBackendEnvironments",
			Err:           err,
		}
	}
	out := result.(*ListBackendEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the list backend environments request.
type ListBackendEnvironmentsInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name of the backend environment
	EnvironmentName *string

	// The maximum number of records to list in a single response.
	MaxResults *int32

	// A pagination token. Set to null to start listing backend environments from the
	// start. If a non-null pagination token is returned in a result, pass its value in
	// here to list more backend environments.
	NextToken *string
}

// The result structure for the list backend environments result.
type ListBackendEnvironmentsOutput struct {

	// The list of backend environments for an Amplify app.
	//
	// This member is required.
	BackendEnvironments []*types.BackendEnvironment

	// A pagination token. If a non-null pagination token is returned in a result, pass
	// its value in another request to retrieve more entries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListBackendEnvironmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListBackendEnvironments{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListBackendEnvironments{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBackendEnvironments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "ListBackendEnvironments",
	}
}
