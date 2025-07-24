// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// cidrbits function returns number of bits in the mask from the address in CIDR notation.
// takes one string parameters: the IP address and the subnet mask in CIDR notation (e.g. "192.168.128.1/24").
// returns an int32 result with number of bits (e.g., 24).
// supports both ipv4 and ipv6.
// fails if the address is not in CIDR notation.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &cidrbitsFunction{}

type cidrbitsFunction struct{}

func newCidrbitsFunction() function.Function {
	return &cidrbitsFunction{}
}

func (f *cidrbitsFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "cidrbits"
}

func (f *cidrbitsFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Return number of bits in the mask from the address in CIDR notation",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "cidr",
				Description: "IP address and subnet mask in CIDR notation (ipv4 or ipv6)",
			},
		},
		Return: function.Int32Return{},
	}
}

func (f *cidrbitsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var cidr string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &cidr))

	// Parse arguments
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid CIDR address"))
		return
	}

	// Set the result
	ones, _ := network.Mask.Size()
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, ones))
}
