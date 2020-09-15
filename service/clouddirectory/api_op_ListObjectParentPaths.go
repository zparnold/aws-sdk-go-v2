// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves all available parent paths for any object type such as node, leaf
// node, policy node, and index node objects. For more information about objects,
// see Directory Structure
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directorystructure.html).
// Use this API to evaluate all parents for an object. The call returns all objects
// from the root of the directory up to the requested object. The API returns the
// number of paths based on user-defined MaxResults, in case there are multiple
// paths to the parent. The order of the paths and nodes returned is consistent
// among multiple API calls unless the objects are deleted or moved. Paths not
// leading to the directory root are ignored from the target object.
func (c *Client) ListObjectParentPaths(ctx context.Context, params *ListObjectParentPathsInput, optFns ...func(*Options)) (*ListObjectParentPathsOutput, error) {
	stack := middleware.NewStack("ListObjectParentPaths", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListObjectParentPathsMiddlewares(stack)
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
	addOpListObjectParentPathsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectParentPaths(options.Region), middleware.Before)

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
			OperationName: "ListObjectParentPaths",
			Err:           err,
		}
	}
	out := result.(*ListObjectParentPathsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectParentPathsInput struct {
	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32
	// The ARN of the directory to which the parent path applies.
	DirectoryArn *string
	// The reference that identifies the object whose parent paths are listed.
	ObjectReference *types.ObjectReference
	// The pagination token.
	NextToken *string
}

type ListObjectParentPathsOutput struct {
	// The pagination token.
	NextToken *string
	// Returns the path to the ObjectIdentifiers that are associated with the
	// directory.
	PathToObjectIdentifiersList []*types.PathToObjectIdentifiers

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListObjectParentPathsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListObjectParentPaths{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListObjectParentPaths{}, middleware.After)
}

func newServiceMetadataMiddleware_opListObjectParentPaths(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListObjectParentPaths",
	}
}
