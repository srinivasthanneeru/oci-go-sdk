// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity


type CreateApiKeyDetails struct {

    // The public key.  Must be an RSA key in PEM format.
    Key string `json:"key,omitempty"`
}