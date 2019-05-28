// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/UpdateTypedLinkFacetRequest
type UpdateTypedLinkFacetInput struct {
	_ struct{} `type:"structure"`

	// Attributes update structure.
	//
	// AttributeUpdates is a required field
	AttributeUpdates []TypedLinkFacetAttributeUpdate `type:"list" required:"true"`

	// The order of identity attributes for the facet, from most significant to
	// least significant. The ability to filter typed links considers the order
	// that the attributes are defined on the typed link facet. When providing ranges
	// to a typed link selection, any inexact ranges must be specified at the end.
	// Any attributes that do not have a range specified are presumed to match the
	// entire range. Filters are interpreted in the order of the attributes on the
	// typed link facet, not the order in which they are supplied to any API calls.
	// For more information about identity attributes, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
	//
	// IdentityAttributeOrder is a required field
	IdentityAttributeOrder []string `type:"list" required:"true"`

	// The unique name of the typed link facet.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns.
	//
	// SchemaArn is a required field
	SchemaArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTypedLinkFacetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTypedLinkFacetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTypedLinkFacetInput"}

	if s.AttributeUpdates == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeUpdates"))
	}

	if s.IdentityAttributeOrder == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityAttributeOrder"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.SchemaArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("SchemaArn"))
	}
	if s.AttributeUpdates != nil {
		for i, v := range s.AttributeUpdates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AttributeUpdates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTypedLinkFacetInput) MarshalFields(e protocol.FieldEncoder) error {

	if len(s.AttributeUpdates) > 0 {
		v := s.AttributeUpdates

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "AttributeUpdates", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if len(s.IdentityAttributeOrder) > 0 {
		v := s.IdentityAttributeOrder

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "IdentityAttributeOrder", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SchemaArn != nil {
		v := *s.SchemaArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "x-amz-data-partition", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/UpdateTypedLinkFacetResponse
type UpdateTypedLinkFacetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateTypedLinkFacetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTypedLinkFacetOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateTypedLinkFacet = "UpdateTypedLinkFacet"

// UpdateTypedLinkFacetRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Updates a TypedLinkFacet. For more information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
//
//    // Example sending a request using UpdateTypedLinkFacetRequest.
//    req := client.UpdateTypedLinkFacetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/UpdateTypedLinkFacet
func (c *Client) UpdateTypedLinkFacetRequest(input *UpdateTypedLinkFacetInput) UpdateTypedLinkFacetRequest {
	op := &aws.Operation{
		Name:       opUpdateTypedLinkFacet,
		HTTPMethod: "PUT",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/typedlink/facet",
	}

	if input == nil {
		input = &UpdateTypedLinkFacetInput{}
	}

	req := c.newRequest(op, input, &UpdateTypedLinkFacetOutput{})
	return UpdateTypedLinkFacetRequest{Request: req, Input: input, Copy: c.UpdateTypedLinkFacetRequest}
}

// UpdateTypedLinkFacetRequest is the request type for the
// UpdateTypedLinkFacet API operation.
type UpdateTypedLinkFacetRequest struct {
	*aws.Request
	Input *UpdateTypedLinkFacetInput
	Copy  func(*UpdateTypedLinkFacetInput) UpdateTypedLinkFacetRequest
}

// Send marshals and sends the UpdateTypedLinkFacet API request.
func (r UpdateTypedLinkFacetRequest) Send(ctx context.Context) (*UpdateTypedLinkFacetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTypedLinkFacetResponse{
		UpdateTypedLinkFacetOutput: r.Request.Data.(*UpdateTypedLinkFacetOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTypedLinkFacetResponse is the response type for the
// UpdateTypedLinkFacet API operation.
type UpdateTypedLinkFacetResponse struct {
	*UpdateTypedLinkFacetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTypedLinkFacet request.
func (r *UpdateTypedLinkFacetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}