// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// lookupipv4 function returns a list of IPv4 addresses returned by DNS.
// uses default system resolver on the host running the provider.
// takes one string parameter: the host name to look up.
// returns a list of string results with IPv4 addresses.
// fails if the DNS lookup fails or if the host name is invalid.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &lookupipv4Function{}

type lookupipv4Function struct{}

func newLookupipv4Function() function.Function {
	return &lookupipv4Function{}
}

func (f *lookupipv4Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "lookupipv4"
}

func (f *lookupipv4Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns a list of IPv4 addresses returned by DNS",
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

func (f *lookupipv4Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var host string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &host))

	addrs, err := net.DefaultResolver.LookupIP(ctx, "ip4", host)
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
