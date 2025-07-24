// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// prefixsize function calculates the size of the IP prefix.
// takes one string parameters: IP prefix in CIDR notation (ipv4 or ipv6).
// returns an int64 result with the size of the prefix.
// fails if prefix is invalid or the size is too large to fit in int64.

package main

import (
	"context"
	"math/big"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &prefixsizeFunction{}

type prefixsizeFunction struct{}

func newPrefixsizeFunction() function.Function {
	return &prefixsizeFunction{}
}

func (f *prefixsizeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "prefixsize"
}

func (f *prefixsizeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns the size of the IP prefix",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "prefix",
				Description: "IP prefix in CIDR notation (ipv4 or ipv6)",
			},
		},
		Return: function.Int64Return{},
	}
}

func (f *prefixsizeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var prefix string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &prefix))

	// Parse arguments
	p, err := netip.ParsePrefix(prefix)
	if err != nil || !p.IsValid() {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid IP prefix"))
		return
	}

	if p.Addr() != p.Masked().Addr() {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Prefix is not in canonical form"))
		return
	}

	// Calculate the size of the prefix
	var bits int
	if p.Addr().Is4() {
		bits = 32
	} else {
		bits = 128
	}
	ones := p.Bits()
	size := new(big.Int).Lsh(big.NewInt(1), uint(bits-ones))

	if !size.IsInt64() {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("IP prefix size is too large to fit in int64"))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, size.Int64()))
}
