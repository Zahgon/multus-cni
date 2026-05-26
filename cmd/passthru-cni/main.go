// Package: passthru-cni
package main

import (
	"github.com/containernetworking/cni/pkg/skel"
	cniTypes "github.com/containernetworking/cni/pkg/types"
	cniVersion "github.com/containernetworking/cni/pkg/version"
)

// NetConf is a CNI configuration structure
type NetConf struct {
	cniTypes.NetConf
}

func main() {
	skel.PluginMain(
		cmdAdd,
		nil,
		cmdDel,
		cniVersion.PluginSupports("0.3.0", "0.3.1", "0.4.0", "1.0.0", "1.1.0"),
		"Passthrough CNI Plugin v1.0",
	)
}

func cmdAdd(args *skel.CmdArgs) error { _ = "STUB: not implemented"; return nil }

// Create an empty but valid CNI result

func cmdDel(_ *skel.CmdArgs) error {
	_ = "STUB: not implemented"
	// Nothing to do for DEL command, just return nil
	return nil
}

func loadNetConf(bytes []byte) (*NetConf, error) { _ = "STUB: not implemented"; return nil, nil }
