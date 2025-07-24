// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// maskbits function returns number of bits in the mask provided in dotted decimal format.
// takes one string parameters: subnet mask in dotted decimal format (e.g. "255.255.255.0").
// returns an int32 result with number of bits (e.g., 24).
// fails if the subnet mask is invalid or not in dotted decimal format.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &maskbitsFunction{}

type maskbitsFunction struct{}

func newMaskbitsFunction() function.Function {
	return &maskbitsFunction{}
}

func (f *maskbitsFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "maskbits"
}

func (f *maskbitsFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns number of bits in the mask provided in dotted decimal format",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "mask",
				Description: "Subnet mask (e.g. 255.255.255.0)",
			},
		},
		Return: function.Int32Return{},
	}
}

func (f *maskbitsFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var netmask string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &netmask))

	// Parse arguments
	mask := net.ParseIP(netmask)

	if mask == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid netmask"))
		return
	}

	if mask.To4() == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Only IPv4 netmasks are supported"))
		return
	}

	// Convert mask to net.IPMask
	ones, bits := net.IPMask(mask.To4()).Size()
	if ones == 0 && bits == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Non-canonical subnet mask"))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, ones))
}
