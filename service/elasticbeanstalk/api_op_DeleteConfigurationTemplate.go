// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified configuration template. When you launch an environment
// using a configuration template, the environment gets a copy of the template. You
// can delete or modify the environment's copy of the template without affecting
// the running environment.
func (c *Client) DeleteConfigurationTemplate(ctx context.Context, params *DeleteConfigurationTemplateInput, optFns ...func(*Options)) (*DeleteConfigurationTemplateOutput, error) {
	stack := middleware.NewStack("DeleteConfigurationTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteConfigurationTemplateMiddlewares(stack)
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
	addOpDeleteConfigurationTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationTemplate(options.Region), middleware.Before)
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
			OperationName: "DeleteConfigurationTemplate",
			Err:           err,
		}
	}
	out := result.(*DeleteConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to delete a configuration template.
type DeleteConfigurationTemplateInput struct {

	// The name of the application to delete the configuration template from.
	//
	// This member is required.
	ApplicationName *string

	// The name of the configuration template to delete.
	//
	// This member is required.
	TemplateName *string
}

type DeleteConfigurationTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteConfigurationTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteConfigurationTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteConfigurationTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteConfigurationTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DeleteConfigurationTemplate",
	}
}
