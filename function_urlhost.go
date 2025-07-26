// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// urlhost function returns host portion of the url.
// takes one string parameter: valid url (e.g. "https://example.com/path?query=1#fragment").
// returns a string result with host portion (e.g., "example.com").
// fails if the url is invalid (e.g. missing scheme) or cannot be parsed.

package main

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &urlhostFunction{}

type urlhostFunction struct{}

func newUrlhostFunction() function.Function {
	return &urlhostFunction{}
}

func (f *urlhostFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "urlhost"
}

func (f *urlhostFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns host portion from the URL",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "url",
				Description: "Valid URL to extract host from (e.g. https://example.com/path?query=1#fragment)",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *urlhostFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var urlstring string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &urlstring))

	// Parse arguments
	u, err := url.Parse(urlstring)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid URL - cannot parse"))
		return
	}

	// If scheme is not specified, assume "http" and try to parse again
	if u.Scheme == "" {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid URL - missing scheme"))
		return
	}

	host := u.Hostname()

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &host))
}
