// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// iputils is a Terraform provider that provides utility functions for IP address manipulation
// there are no resources or data sources implemented in this provider
// no configuration options are supported

package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure IPUtilsProvider satisfies various provider interfaces.
var _ provider.Provider = &iputilsProvider{}
var _ provider.ProviderWithFunctions = &iputilsProvider{}

type iputilsProvider struct {
	version string
}

func newIputilsProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &iputilsProvider{
			version: version,
		}
	}
}

func (p *iputilsProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "iputils"
	resp.Version = p.version
}

func (p *iputilsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *iputilsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "provider does not support any configuration options", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *iputilsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *iputilsProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

// Functions defines the functions implemented in the provider.
func (p *iputilsProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		newTocidrFunction, // tocidr(address, mask string) string
	}
}
