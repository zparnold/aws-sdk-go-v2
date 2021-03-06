// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a contact with a given address book.
func (c *Client) AssociateContactWithAddressBook(ctx context.Context, params *AssociateContactWithAddressBookInput, optFns ...func(*Options)) (*AssociateContactWithAddressBookOutput, error) {
	stack := middleware.NewStack("AssociateContactWithAddressBook", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateContactWithAddressBookMiddlewares(stack)
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
	addOpAssociateContactWithAddressBookValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateContactWithAddressBook(options.Region), middleware.Before)
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
			OperationName: "AssociateContactWithAddressBook",
			Err:           err,
		}
	}
	out := result.(*AssociateContactWithAddressBookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateContactWithAddressBookInput struct {

	// The ARN of the address book with which to associate the contact.
	//
	// This member is required.
	AddressBookArn *string

	// The ARN of the contact to associate with an address book.
	//
	// This member is required.
	ContactArn *string
}

type AssociateContactWithAddressBookOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateContactWithAddressBookMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateContactWithAddressBook{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateContactWithAddressBook{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateContactWithAddressBook(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "AssociateContactWithAddressBook",
	}
}
