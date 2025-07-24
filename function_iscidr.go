// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// iscidr function returns true if provided address in CIDR format is valid.
// takes one string parameters: the IP address in CIDR format (ipv4 or ipv6).
// returns a bool result - true if address is valid, or false otherwise.

package main

import (
	"context"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &iscidrFunction{}

type iscidrFunction struct{}

func newIscidrFunction() function.Function {
	return &iscidrFunction{}
}

func (f *iscidrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "iscidr"
}

func (f *iscidrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns true if provided address in CIDR format is valid",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "cidr",
				Description: "IP address in CIDR format (ipv4 or ipv6)",
			},
		},
		Return: function.BoolReturn{},
	}
}

func (f *iscidrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var cidr string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &cidr))

	// Parse arguments
	prefix, err := netip.ParsePrefix(cidr)
	if err != nil || !prefix.IsValid() {
		resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, false))
		return
	}

	// Assume the address is valid cidr
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, true))
}
