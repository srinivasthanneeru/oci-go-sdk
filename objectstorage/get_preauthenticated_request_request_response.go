// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// GetPreauthenticatedRequestRequest wrapper for the GetPreauthenticatedRequest operation
type GetPreauthenticatedRequestRequest struct {

	// The top-level namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The unique identifier for the pre-authenticated request (PAR). This can be used to manage the PAR
	// such as GET or DELETE the PAR
	ParID *string `mandatory:"true" contributesTo:"path" name:"parId"`

	// The client request ID for tracing.
	OpcClientRequestID *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`
}

func (request GetPreauthenticatedRequestRequest) String() string {
	return common.PointerString(request)
}

// GetPreauthenticatedRequestResponse wrapper for the GetPreauthenticatedRequest operation
type GetPreauthenticatedRequestResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PreauthenticatedRequestSummary instance
	PreauthenticatedRequestSummary `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestID *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, please provide this request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPreauthenticatedRequestResponse) String() string {
	return common.PointerString(response)
}
