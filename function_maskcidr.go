// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// maskcidr function converts subnet mask in dotted decimal format to CIDR format (with leading slash).
// takes one string parameters: subnet mask in dotted decimal format (e.g. "255.255.255.0").
// returns a string result with CIDR mask (e.g., "/24").
// fails if the subnet mask is invalid or not in dotted decimal format.

package main

import (
	"context"
	"net"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &maskcidrFunction{}

type maskcidrFunction struct{}

func newMaskcidrFunction() function.Function {
	return &maskcidrFunction{}
}

func (f *maskcidrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "maskcidr"
}

func (f *maskcidrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Converts subnet mask in dotted decimal format to CIDR format (with leading slash)",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "mask",
				Description: "Subnet mask (e.g. 255.255.255.0)",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *maskcidrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
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
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, "/"+strconv.Itoa(ones)))
}
