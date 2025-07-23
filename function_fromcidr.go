// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// fromcidr function converts an IP address and subnet mask from CIDR notation.
// takes one string parameters: the IP address and the subnet mask in CIDR notation.
// returns an object with IP address and subnet mask as string values.
// fails if the address is not in CIDR notation or if it is not an IPv4 address.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &fromcidrFunction{}

type fromcidrFunction struct{}

func newFromcidrFunction() function.Function {
	return &fromcidrFunction{}
}

func (f *fromcidrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "fromcidr"
}

func (f *fromcidrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns IP address and subnet mask from CIDR notation",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "cidr",
				Description: "IP address and subnet mask in CIDR notation (ipv4)",
			},
		},
		Return: function.ObjectReturn{
			AttributeTypes: map[string]attr.Type{
				"address": types.StringType,
				"netmask": types.StringType,
			},
		},
	}
}

func (f *fromcidrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var cidr string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &cidr))

	// Parse arguments
	addr, network, err := net.ParseCIDR(cidr)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid CIDR address"))
		return
	}

	if addr.To4() == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Only IPv4 addresses are supported"))
		return
	}

	// Construct the result object
	result, diag := types.ObjectValue(
		map[string]attr.Type{
			"address": types.StringType,
			"netmask": types.StringType,
		},
		map[string]attr.Value{
			"address": types.StringValue(addr.String()),
			"netmask": types.StringValue(net.IP(network.Mask).String()),
		},
	)
	if diag.HasError() {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.FuncErrorFromDiags(ctx, diag))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
