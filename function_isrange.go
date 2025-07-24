// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// isrange function returns true if provided IP address range is valid.
// takes one string parameters: two IP addresses separated by hyphen "-" (ipv4 or ipv6).
// returns a bool result - true if range is valid, or false otherwise.

package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"go4.org/netipx"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &israngeFunction{}

type israngeFunction struct{}

func newIsrangeFunction() function.Function {
	return &israngeFunction{}
}

func (f *israngeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "isrange"
}

func (f *israngeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns true if provided IP address range is valid",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "range",
				Description: "IP address range (ipv4 or ipv6)",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *israngeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var iprange string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &iprange))

	// Parse arguments
	r, err := netipx.ParseIPRange(iprange)
	if err != nil || !r.IsValid() {
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, false))
		return
	}

	// Assume the range is valid
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, true))
}
