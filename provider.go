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
		newTocidrFunction,      // tocidr(address, mask string) string
		newFromcidrFunction,    // fromcidr(cidr string) object { address: string, netmask: string }
		newCidrhostFunction,    // cidrhost(cidr string) string
		newCidrmaskFunction,    // cidrmask(cidr string) string
		newCidrbitsFunction,    // cidrbits(cidr string) int32
		newMaskcidrFunction,    // maskcidr(mask string) string
		newMaskbitsFunction,    // maskbits(mask string) int32
		newIsvalidFunction,     // isvalid(address string) bool
		newIscidrFunction,      // cidrvalid(address string) bool
		newIsrangeFunction,     // isrange(range string) bool
		newRangesizeFunction,   // rangesize(range string) int64
		newRangehostFunction,   // rangehost(range string, index int64) string
		newPrefixsizeFunction,  // prefixsize(prefix string) int64
		newLookupipv4Function,  // lookupipv4(host string) list(string)
		newLookupipv6Function,  // lookupipv6(host string) list(string)
		newlookupaddrFunction,  // lookupaddr(host string) list(string)
		newLookuphostFunction,  // lookuphost(address string) list(string)
		newLookupcnameFunction, // lookupcname(host string) string
	}
}
