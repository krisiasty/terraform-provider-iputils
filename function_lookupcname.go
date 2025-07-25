// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// lookupcname function returns canonical name (CNAME) record for specified host.
// uses default system resolver on the host running the provider.
// takes one string parameter: the host name to look up.
// returns a string results with host canonical name.
// fails if the DNS lookup fails or if the host name is invalid.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &lookupcnameFunction{}

type lookupcnameFunction struct{}

func newLookupcnameFunction() function.Function {
	return &lookupcnameFunction{}
}

func (f *lookupcnameFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "lookupcname"
}

func (f *lookupcnameFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns canonical name (CNAME) record for specified host",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "host",
				Description: "hostname to look up",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *lookupcnameFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var host string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &host))

	cname, err := net.DefaultResolver.LookupCNAME(ctx, host)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &cname))
}
