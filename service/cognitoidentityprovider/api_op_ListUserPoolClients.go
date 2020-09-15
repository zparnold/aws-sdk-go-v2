// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the clients that have been created for the specified user pool.
func (c *Client) ListUserPoolClients(ctx context.Context, params *ListUserPoolClientsInput, optFns ...func(*Options)) (*ListUserPoolClientsOutput, error) {
	stack := middleware.NewStack("ListUserPoolClients", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListUserPoolClientsMiddlewares(stack)
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
	addOpListUserPoolClientsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListUserPoolClients(options.Region), middleware.Before)

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
			OperationName: "ListUserPoolClients",
			Err:           err,
		}
	}
	out := result.(*ListUserPoolClientsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to list the user pool clients.
type ListUserPoolClientsInput struct {
	// The user pool ID for the user pool where you want to list user pool clients.
	UserPoolId *string
	// The maximum number of results you want the request to return when listing the
	// user pool clients.
	MaxResults *int32
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
}

// Represents the response from the server that lists user pool clients.
type ListUserPoolClientsOutput struct {
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
	// The user pool clients in the response that lists user pool clients.
	UserPoolClients []*types.UserPoolClientDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListUserPoolClientsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListUserPoolClients{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListUserPoolClients{}, middleware.After)
}

func newServiceMetadataMiddleware_opListUserPoolClients(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ListUserPoolClients",
	}
}