// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a logger definition version.
func (c *Client) GetLoggerDefinitionVersion(ctx context.Context, params *GetLoggerDefinitionVersionInput, optFns ...func(*Options)) (*GetLoggerDefinitionVersionOutput, error) {
	if params == nil {
		params = &GetLoggerDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLoggerDefinitionVersion", params, optFns, addOperationGetLoggerDefinitionVersionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLoggerDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoggerDefinitionVersionInput struct {

	// The ID of the logger definition.
	//
	// This member is required.
	LoggerDefinitionId *string

	// The ID of the logger definition version. This value maps to the ''Version''
	// property of the corresponding ''VersionInformation'' object, which is returned
	// by ''ListLoggerDefinitionVersions'' requests. If the version is the last one
	// that was associated with a logger definition, the value also maps to the
	// ''LatestVersion'' property of the corresponding ''DefinitionInformation''
	// object.
	//
	// This member is required.
	LoggerDefinitionVersionId *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
}

type GetLoggerDefinitionVersionOutput struct {

	// The ARN of the logger definition version.
	Arn *string

	// The time, in milliseconds since the epoch, when the logger definition version
	// was created.
	CreationTimestamp *string

	// Information about the logger definition version.
	Definition *types.LoggerDefinitionVersion

	// The ID of the logger definition version.
	Id *string

	// The version of the logger definition version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLoggerDefinitionVersionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLoggerDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLoggerDefinitionVersion{}, middleware.After)
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
	addOpGetLoggerDefinitionVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoggerDefinitionVersion(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetLoggerDefinitionVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetLoggerDefinitionVersion",
	}
}
