// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// urlpath function returns the path portion of the url.
// takes one string parameter: valid url (e.g. "https://example.com/path?query=1#fragment").
// returns a string result with path portion (e.g., "/path").
// fails if the url is invalid (e.g. missing scheme) or cannot be parsed.

package main

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &urlpathFunction{}

type urlpathFunction struct{}

func newUrlpathFunction() function.Function {
	return &urlpathFunction{}
}

func (f *urlpathFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "urlpath"
}

func (f *urlpathFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns the path portion from the URL",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "url",
				Description: "Valid URL to extract path from (e.g. https://example.com/path?query=1#fragment)",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *urlpathFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
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

	path := u.Path

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &path))
}
