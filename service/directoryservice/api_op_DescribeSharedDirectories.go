// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the shared directories in your account.
func (c *Client) DescribeSharedDirectories(ctx context.Context, params *DescribeSharedDirectoriesInput, optFns ...func(*Options)) (*DescribeSharedDirectoriesOutput, error) {
	if params == nil {
		params = &DescribeSharedDirectoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSharedDirectories", params, optFns, addOperationDescribeSharedDirectoriesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSharedDirectoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSharedDirectoriesInput struct {

	// Returns the identifier of the directory in the directory owner account.
	//
	// This member is required.
	OwnerDirectoryId *string

	// The number of shared directories to return in the response object.
	Limit *int32

	// The DescribeSharedDirectoriesResult.NextToken value from a previous call to
	// DescribeSharedDirectories (). Pass null if this is the first call.
	NextToken *string

	// A list of identifiers of all shared directories in your account.
	SharedDirectoryIds []*string
}

type DescribeSharedDirectoriesOutput struct {

	// If not null, token that indicates that more results are available. Pass this
	// value for the NextToken parameter in a subsequent call to
	// DescribeSharedDirectories () to retrieve the next set of items.
	NextToken *string

	// A list of all shared directories in your account.
	SharedDirectories []*types.SharedDirectory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeSharedDirectoriesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSharedDirectories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSharedDirectories{}, middleware.After)
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
	addOpDescribeSharedDirectoriesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSharedDirectories(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeSharedDirectories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeSharedDirectories",
	}
}
