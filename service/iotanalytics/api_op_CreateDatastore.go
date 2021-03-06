// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a data store, which is a repository for messages.
func (c *Client) CreateDatastore(ctx context.Context, params *CreateDatastoreInput, optFns ...func(*Options)) (*CreateDatastoreOutput, error) {
	stack := middleware.NewStack("CreateDatastore", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDatastoreMiddlewares(stack)
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
	addOpCreateDatastoreValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatastore(options.Region), middleware.Before)
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
			OperationName: "CreateDatastore",
			Err:           err,
		}
	}
	out := result.(*CreateDatastoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatastoreInput struct {

	// The name of the data store.
	//
	// This member is required.
	DatastoreName *string

	// Where data store data is stored. You may choose one of "serviceManagedS3" or
	// "customerManagedS3" storage. If not specified, the default is
	// "serviceManagedS3". This cannot be changed after the data store is created.
	DatastoreStorage *types.DatastoreStorage

	// How long, in days, message data is kept for the data store. When
	// "customerManagedS3" storage is selected, this parameter is ignored.
	RetentionPeriod *types.RetentionPeriod

	// Metadata which can be used to manage the data store.
	Tags []*types.Tag
}

type CreateDatastoreOutput struct {

	// The ARN of the data store.
	DatastoreArn *string

	// The name of the data store.
	DatastoreName *string

	// How long, in days, message data is kept for the data store.
	RetentionPeriod *types.RetentionPeriod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDatastoreMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDatastore{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDatastore{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDatastore(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "CreateDatastore",
	}
}
