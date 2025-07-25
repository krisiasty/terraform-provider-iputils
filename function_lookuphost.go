// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// lookuphost function returns a list of hostnames associated with specified IP address returned by DNS.
// (effectively reverse DNS lookup).
// uses default system resolver on the host running the provider.
// takes one string parameter: the IPv4 or IPv6 address to look up.
// returns a list of string results with DNS host names.
// fails if the DNS lookup fails or if the address is invalid.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &lookuphostFunction{}

type lookuphostFunction struct{}

func newLookuphostFunction() function.Function {
	return &lookuphostFunction{}
}

func (f *lookuphostFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "lookuphost"
}

func (f *lookuphostFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns a list of hostnames associated with specified IP address returned by DNS (reverse DNS lookup)",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "address",
				Description: "IP address to look up",
			},
		},
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
	}
}

func (f *lookuphostFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var address string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &address))

	hosts, err := net.DefaultResolver.LookupAddr(ctx, address)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Cannot resolve address: "+err.Error()))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &hosts))
}
