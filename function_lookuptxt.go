// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// lookuptxt function returns a list of TXT records for given domain name.
// uses default system resolver on the host running the provider.
// takes one string parameter: the domain name to look up.
// returns a list of string results with TXT records.
// fails if the DNS lookup fails or if the domain name is invalid.

package main

import (
	"context"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &lookuptxtFunction{}

type lookuptxtFunction struct{}

func newLookuptxtFunction() function.Function {
	return &lookuptxtFunction{}
}

func (f *lookuptxtFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "lookuptxt"
}

func (f *lookuptxtFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns a list of TXT records for given domain name",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "name",
				Description: "domain name to look up",
			},
		},
		Return: function.ListReturn{
			ElementType: types.StringType,
		},
	}
}

func (f *lookuptxtFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var name string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &name))

	result, err := net.DefaultResolver.LookupTXT(ctx, name)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &result))
}
