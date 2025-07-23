// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

// tocidr function converts an IP address and subnet mask to CIDR notation.
// takes two string parameters: the IP address and the subnet mask.
// returns a string result with IP address and subnet mask in CIDR format (e.g., "192.168.128.1/24").

package main

import (
	"context"
	"net"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &tocidrFunction{}

type tocidrFunction struct{}

func newTocidrFunction() function.Function {
	return &tocidrFunction{}
}

func (f *tocidrFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "tocidr"
}

func (f *tocidrFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Description: "Convert IP address and subnet mask to CIDR notation",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "address",
				Description: "IP address (ipv4)",
			},
			function.StringParameter{
				Name:        "mask",
				Description: "Subnet mask (e.g. 255.255.255.0)",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *tocidrFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Read Terraform argument data into the variables
	var address, netmask string
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &address, &netmask))

	// Parse arguments
	ip := net.ParseIP(address)
	mask := net.ParseIP(netmask)
	if ip == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid IP address"))
		return
	}

	if mask == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Invalid netmask"))
		return
	}

	if ip.To4() == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("IPv6 address provided but only IPv4 addresses are supported"))
		return
	}

	if mask.To4() == nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Only IPv4 netmasks are supported"))
		return
	}

	// Convert mask to net.IPMask
	ones, bits := net.IPMask(mask.To4()).Size()

	if ones == 0 || bits == 0 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError("Non-canonical subnet mask"))
		return
	}

	// Construct CIDR notation
	cidr := address + "/" + strconv.Itoa(ones)

	// Set the result
	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, cidr))
}
