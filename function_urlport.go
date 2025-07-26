// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// urlport function returns the port part of the url.
// takes one string parameter: valid url (e.g. "https://example.com/path?query=1#fragment").
// returns a string result with the port number (e.g., "443").
// if the port is not specified in the URL, it will be inferred for most common schemes (http, https, ftp, ssh).
// fails if the url is invalid (e.g. missing scheme) or cannot be parsed.

package main

import (
	"context"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &urlportFunction{}

type urlportFunction struct{}

func newUrlportFunction() function.Function {
	return &urlportFunction{}
}

func (f *urlportFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "urlport"
}

func (f *urlportFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Returns the port part of the URL",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "url",
				Description: "Valid URL to extract port number from (e.g. https://example.com:8443/path?query=1#fragment)",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *urlportFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var urlstring string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &urlstring))

	// Parse arguments
	u, err := url.Parse(urlstring)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid URL"))
		return
	}

	// Fail if the scheme is not specified
	if u.Scheme == "" {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid URL - missing scheme"))
		return
	}

	port := u.Port()

	// If port is not specified, infer it based on the scheme
	if port == "" {
		switch u.Scheme {
		case "http":
			port = "80"
		case "https":
			port = "443"
		case "ftp":
			port = "21"
		case "ssh":
			port = "22"
		}
	}

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, &port))
}
