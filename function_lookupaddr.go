// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// lookupaddr function returns a list of IP addresses returned by DNS.
// can return both IPv4 and IPv6 addresses.
// uses default system resolver on the host running the provider.
// takes one string parameter: the host name to look up.
// returns a list of string results with IP addresses.
// fails if the DNS lookup fails or if the host name is invalid.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &lookupaddrFunction{}

type lookupaddrFunction struct{}

func newlookupaddrFunction() function.Function {
	return &lookupaddrFunction{}
}

func (f *lookupaddrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "lookupaddr"
}

func (f *lookupaddrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns a list of IP addresses returned by DNS (both IPv4 and IPv6)",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "host",
				Description: "hostname to look up",
			},
		},
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
	}
}

func (f *lookupaddrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var host string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &host))

	addrs, err := net.DefaultResolver.LookupIP(ctx, "ip", host)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}
	result := make([]string, len(addrs))
	for i, ip := range addrs {
		result[i] = ip.String()
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &result))
}
