// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// rangehost returns n-th host in specified IP address range.
// index (n) can be negative, which means counting from the end of the range.
// takes two parameters:
//    string parameters: two IP addresses separated by hyphen "-" (ipv4 or ipv6).
//    int64 parameter: index of the host to return (0 is the first host, -1 is the last host).
// returns IP address as a string.
// fails if the range is invalid or if the index is out of bounds.

package main

import (
	"context"
	"math/big"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"go4.org/netipx"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &rangehostFunction{}

type rangehostFunction struct{}

func newRangehostFunction() function.Function {
	return &rangehostFunction{}
}

func (f *rangehostFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "rangehost"
}

func (f *rangehostFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns the size of the IP range",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "range",
				Description: "IP address range (ipv4 or ipv6)",
			},
			function.Int64Parameter{
				Name:        "index",
				Description: "index of the host to return (0 is the first host, -1 is the last host)",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *rangehostFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var iprange string
	var index int64
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &iprange, &index))

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
	addrInt := new(big.Int)
	if index >= 0 {
		indexInt := new(big.Int).SetInt64(index)
		addrInt.Add(fromInt, indexInt)
		if addrInt.Cmp(toInt) > 0 {
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Index out of bounds"))
			return
		}
	} else {
		indexInt := new(big.Int).SetInt64(index + 1)
		addrInt.Add(toInt, indexInt)
		if addrInt.Cmp(fromInt) < 0 {
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Index out of bounds"))
			return
		}
	}

	// Convert calculated address stored as big.Int back to netip.Addr
	var addrBytes [16]byte
	addrInt.FillBytes(addrBytes[:])
	addr := netip.AddrFrom16(addrBytes).Unmap()
	if !addr.IsValid() {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Error calculating IP address"))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, addr.String()))
}
