// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// An application in AppConfig is a logical unit of code that provides capabilities
// for your customers. For example, an application can be a microservice that runs
// on Amazon EC2 instances, a mobile application installed by your users, a
// serverless application using Amazon API Gateway and AWS Lambda, or any system
// you run on behalf of others.
func (c *Client) CreateApplication(ctx context.Context, params *CreateApplicationInput, optFns ...func(*Options)) (*CreateApplicationOutput, error) {
	stack := middleware.NewStack("CreateApplication", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateApplicationMiddlewares(stack)
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
	addOpCreateApplicationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateApplication(options.Region), middleware.Before)
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
			OperationName: "CreateApplication",
			Err:           err,
		}
	}
	out := result.(*CreateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateApplicationInput struct {

	// A name for the application.
	//
	// This member is required.
	Name *string

	// A description of the application.
	Description *string

	// Metadata to assign to the application. Tags help organize and categorize your
	// AppConfig resources. Each tag consists of a key and an optional value, both of
	// which you define.
	Tags map[string]*string
}

type CreateApplicationOutput struct {

	// The description of the application.
	Description *string

	// The application ID.
	Id *string

	// The application name.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateApplicationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateApplication{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateApplication{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateApplication(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "CreateApplication",
	}
}
