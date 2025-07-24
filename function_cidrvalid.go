// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// cidrvalid function returns true if provided address in CIDR format is valid.
// takes one string parameters: the IP address in CIDR format (ipv4 or ipv6).
// returns a bool result - true if address is valid, or false otherwise.

package main

import (
	"context"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &cidrvalidFunction{}

type cidrvalidFunction struct{}

func newCidrvalidFunction() function.Function {
	return &cidrvalidFunction{}
}

func (f *cidrvalidFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "cidrvalid"
}

func (f *cidrvalidFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns true if provided address in CIDR format is valid",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "address",
				Description: "IP address in CIDR format (ipv4 or ipv6)",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *cidrvalidFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var address string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &address))

	// Parse arguments
	prefix, err := netip.ParsePrefix(address)
	if err != nil || !prefix.IsValid() {
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, false))
		return
	}

	// Assume the address is valid
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, true))
}
