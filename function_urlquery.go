// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// urlquery function returns query options from the url.
// takes one string parameter: valid url (e.g. "https://example.com/path?opt1=val1&opt2=val2").
// returns a map of query options from the URL; each value is a list of strings
// fails if the url is invalid (e.g. missing scheme) or cannot be parsed.

package main

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &urlqueryFunction{}

type urlqueryFunction struct{}

func newUrlqueryFunction() function.Function {
	return &urlqueryFunction{}
}

func (f *urlqueryFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "urlquery"
}

func (f *urlqueryFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns a map of query options from the URL; each value is a list of strings",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "url",
				Description: "Valid URL to extract path from (e.g. https://example.com/path?query=1#fragment)",
			},
		},
		Return: function.MapReturn{
			ElementType: types.ListType{ElemType: types.StringType},
		},
	}
}

func (f *urlqueryFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var urlstring string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &urlstring))

	// Parse arguments
	u, err := url.Parse(urlstring)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid URL - cannot parse"))
		return
	}

	// Fail if the scheme is not specified
	if u.Scheme == "" {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid URL - missing scheme"))
		return
	}

	query := u.Query()

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &query))
}
