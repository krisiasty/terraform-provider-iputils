// Copyright (c) Krzysztof Ciepłucha
// SPDX-License-Identifier: MIT

package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

var (
	// version stromg will be set by the build system
	version string = "dev"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "Enable debug mode")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/krisiasty/iputils",
	}

	err := providerserver.Serve(context.Background(), newIputilsProvider(version), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
