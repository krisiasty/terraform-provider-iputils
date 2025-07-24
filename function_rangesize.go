// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// rangesize function calculates the size of the IP range.
// takes one string parameters: two IP addresses separated by hyphen "-" (ipv4 or ipv6).
// returns an int64 result with the size of the range.

package main

import (
	"context"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"go4.org/netipx"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &rangesizeFunction{}

type rangesizeFunction struct{}

func newRangesizeFunction() function.Function {
	return &rangesizeFunction{}
}

func (f *rangesizeFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "rangesize"
}

func (f *rangesizeFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns the size of the IP range",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "range",
				Description: "IP address range (ipv4 or ipv6)",
			},
		},
		Return: function.Int64Return{},
	}
}

func (f *rangesizeFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var iprange string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &iprange))

	// Parse arguments
	r, err := netipx.ParseIPRange(iprange)
	if err != nil || !r.IsValid() {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid IP range"))
		return
	}

	// Calculate the size of the range using big.Int arithmetic
	from := r.From().As16()
	to := r.To().As16()
	fromInt := new(big.Int).SetBytes(from[:])
	toInt := new(big.Int).SetBytes(to[:])
	size := new(big.Int).Sub(toInt, fromInt)
	size.Add(size, big.NewInt(1))

	if !size.IsInt64() {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Range size is too large to fit in int64"))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, size.Int64()))
}
