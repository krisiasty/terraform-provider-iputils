// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// cidrhost function returns host portion from the address in CIDR notation.
// takes one string parameters: the IP address and the subnet mask in CIDR notation (e.g. "192.168.128.1/24").
// returns a string result with IP address (e.g., "192.168.128.1").

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &cidrhostFunction{}

type cidrhostFunction struct{}

func newCidrhostFunction() function.Function {
	return &cidrhostFunction{}
}

func (f *cidrhostFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "cidrhost"
}

func (f *cidrhostFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Return host portion from the address in CIDR notation.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "cidr",
				Description: "IP address and subnet mask in CIDR notation (ipv4)",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *cidrhostFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var cidr string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &cidr))

	// Parse arguments
	addr, _, err := net.ParseCIDR(cidr)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid CIDR address"))
		return
	}

	if addr.To4() == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Only IPv4 addresses are supported"))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, addr.String()))
}
