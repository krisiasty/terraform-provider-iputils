// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// isvalid function returns true if provided address is valid.
// takes one string parameters: the IP address (ipv4 or ipv6).
// returns a bool result - true if address is valid, or false otherwise.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &isvalidFunction{}

type isvalidFunction struct{}

func newIsvalidFunction() function.Function {
	return &isvalidFunction{}
}

func (f *isvalidFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "isvalid"
}

func (f *isvalidFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns true if provided address is valid",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "address",
				Description: "IP address (ipv4 or ipv6)",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *isvalidFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var address string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &address))

	// Parse arguments
	ip := net.ParseIP(address)

	// Return false if the address is not valid
	if ip == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, false))
		return
	}

	// Assume the address is valid
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, true))
}
