// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListProtocolsListsInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether the lists to retrieve are default lists owned by AWS Firewall
	// Manager.
	DefaultLists *bool `type:"boolean"`

	// The maximum number of objects that you want AWS Firewall Manager to return
	// for this request. If more objects are available, in the response, AWS Firewall
	// Manager provides a NextToken value that you can use in a subsequent call
	// to get the next batch of objects.
	//
	// If you don't specify this, AWS Firewall Manager returns all available objects.
	//
	// MaxResults is a required field
	MaxResults *int64 `min:"1" type:"integer" required:"true"`

	// If you specify a value for MaxResults in your list request, and you have
	// more objects than the maximum, AWS Firewall Manager returns this token in
	// the response. For all but the first request, you provide the token returned
	// by the prior request in the request parameters, to retrieve the next batch
	// of objects.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListProtocolsListsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProtocolsListsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProtocolsListsInput"}

	if s.MaxResults == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaxResults"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListProtocolsListsOutput struct {
	_ struct{} `type:"structure"`

	// If you specify a value for MaxResults in your list request, and you have
	// more objects than the maximum, AWS Firewall Manager returns this token in
	// the response. You can use this token in subsequent requests to retrieve the
	// next batch of objects.
	NextToken *string `min:"1" type:"string"`

	// An array of ProtocolsListDataSummary objects.
	ProtocolsLists []ProtocolsListDataSummary `type:"list"`
}

// String returns the string representation
func (s ListProtocolsListsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListProtocolsLists = "ListProtocolsLists"

// ListProtocolsListsRequest returns a request value for making API operation for
// Firewall Management Service.
//
// Returns an array of ProtocolsListDataSummary objects.
//
//    // Example sending a request using ListProtocolsListsRequest.
//    req := client.ListProtocolsListsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/ListProtocolsLists
func (c *Client) ListProtocolsListsRequest(input *ListProtocolsListsInput) ListProtocolsListsRequest {
	op := &aws.Operation{
		Name:       opListProtocolsLists,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListProtocolsListsInput{}
	}

	req := c.newRequest(op, input, &ListProtocolsListsOutput{})

	return ListProtocolsListsRequest{Request: req, Input: input, Copy: c.ListProtocolsListsRequest}
}

// ListProtocolsListsRequest is the request type for the
// ListProtocolsLists API operation.
type ListProtocolsListsRequest struct {
	*aws.Request
	Input *ListProtocolsListsInput
	Copy  func(*ListProtocolsListsInput) ListProtocolsListsRequest
}

// Send marshals and sends the ListProtocolsLists API request.
func (r ListProtocolsListsRequest) Send(ctx context.Context) (*ListProtocolsListsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProtocolsListsResponse{
		ListProtocolsListsOutput: r.Request.Data.(*ListProtocolsListsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListProtocolsListsResponse is the response type for the
// ListProtocolsLists API operation.
type ListProtocolsListsResponse struct {
	*ListProtocolsListsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProtocolsLists request.
func (r *ListProtocolsListsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}