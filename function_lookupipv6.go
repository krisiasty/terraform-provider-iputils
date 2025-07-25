// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// lookupipv6 function returns a list of IPv6 addresses returned by DNS.
// uses default system resolver on the host running the provider.
// takes one string parameter: the host name to look up.
// returns a list of string results with IPv6 addresses.
// fails if the DNS lookup fails or if the host name is invalid.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &lookupipv6Function{}

type lookupipv6Function struct{}

func newLookupipv6Function() function.Function {
	return &lookupipv6Function{}
}

func (f *lookupipv6Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "lookupipv6"
}

func (f *lookupipv6Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns a list of IPv6 addresses returned by DNS",
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

func (f *lookupipv6Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var host string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &host))

	addrs, err := net.DefaultResolver.LookupIP(ctx, "ip6", host)
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
