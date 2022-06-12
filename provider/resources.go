// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nexus

import (
	"fmt"
	"path/filepath"

	"github.com/SimCubeLtd/pulumi-nexus/provider/pkg/version"
	nexus "github.com/datadrivers/terraform-provider-nexus"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
)

const (
	mainPkg = "nexus"
	mainMod = "index"
)

func Provider() tfbridge.ProviderInfo {
	p := shimv2.NewProvider(nexus.Provider())

	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "nexus",
		DisplayName:       "Nexus",
		Publisher:         "SimCubeLtd",
		LogoURL:           "https://digital.ai/sites/default/files/pictures/styles/maxwidth_300/public/pt_logos/NexusRepo_Icon.png",
		PluginDownloadURL: "https://github.com/SimCubeLtd/pulumi-nexus/releases/download/v${VERSION}",
		Description:       "A Pulumi package for creating and managing nexus resources.",
		Keywords:          []string{"pulumi", "nexus", "category/cloud"},
		License:           "MIT",
		Homepage:          "https://github.com/SimCubeLtd/pulumi-nexus",
		Repository:        "https://github.com/SimCubeLtd/pulumi-nexus",
		GitHubOrg:         "datadrivers",
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources:         map[string]*tfbridge.ResourceInfo{
			// "nexus_item": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Item"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"password": {
			// 			Secret: tfbridge.True(),
			// 		},
			// 	},
			// },
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// "nexus_item": {
			// 	Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetItem"),
			// },
			// "nexus_vault": {
			// 	Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetVault"),
			// },
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@simcubeltd/pulumi-nexus",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "simcubeltd_pulumi_nexus",
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/SimCubeLtd/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
