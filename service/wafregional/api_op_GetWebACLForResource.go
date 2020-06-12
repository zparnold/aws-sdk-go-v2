// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetWebACLForResourceInput struct {
	_ struct{} `type:"structure"`

	// The ARN (Amazon Resource Name) of the resource for which to get the web ACL,
	// either an application load balancer or Amazon API Gateway stage.
	//
	// The ARN should be in one of the following formats:
	//
	//    * For an Application Load Balancer: arn:aws:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	//    * For an Amazon API Gateway stage: arn:aws:apigateway:region::/restapis/api-id/stages/stage-name
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetWebACLForResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetWebACLForResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetWebACLForResourceInput"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetWebACLForResourceOutput struct {
	_ struct{} `type:"structure"`

	// Information about the web ACL that you specified in the GetWebACLForResource
	// request. If there is no associated resource, a null WebACLSummary is returned.
	WebACLSummary *WebACLSummary `type:"structure"`
}

// String returns the string representation
func (s GetWebACLForResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetWebACLForResource = "GetWebACLForResource"

// GetWebACLForResourceRequest returns a request value for making API operation for
// AWS WAF Regional.
//
//
// This is AWS WAF Classic Regional documentation. For more information, see
// AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the AWS
// WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// With the latest version, AWS WAF has a single set of endpoints for regional
// and global use.
//
// Returns the web ACL for the specified resource, either an application load
// balancer or Amazon API Gateway stage.
//
//    // Example sending a request using GetWebACLForResourceRequest.
//    req := client.GetWebACLForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/GetWebACLForResource
func (c *Client) GetWebACLForResourceRequest(input *GetWebACLForResourceInput) GetWebACLForResourceRequest {
	op := &aws.Operation{
		Name:       opGetWebACLForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetWebACLForResourceInput{}
	}

	req := c.newRequest(op, input, &GetWebACLForResourceOutput{})

	return GetWebACLForResourceRequest{Request: req, Input: input, Copy: c.GetWebACLForResourceRequest}
}

// GetWebACLForResourceRequest is the request type for the
// GetWebACLForResource API operation.
type GetWebACLForResourceRequest struct {
	*aws.Request
	Input *GetWebACLForResourceInput
	Copy  func(*GetWebACLForResourceInput) GetWebACLForResourceRequest
}

// Send marshals and sends the GetWebACLForResource API request.
func (r GetWebACLForResourceRequest) Send(ctx context.Context) (*GetWebACLForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetWebACLForResourceResponse{
		GetWebACLForResourceOutput: r.Request.Data.(*GetWebACLForResourceOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetWebACLForResourceResponse is the response type for the
// GetWebACLForResource API operation.
type GetWebACLForResourceResponse struct {
	*GetWebACLForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetWebACLForResource request.
func (r *GetWebACLForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}