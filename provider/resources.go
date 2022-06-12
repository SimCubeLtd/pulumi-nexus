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
	nexus "github.com/SimCubeLtd/terraform-provider-nexus/provider"
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
		Description:       "A Pulumi package for creating and managing SonaType Nexus resources.",
		Keywords:          []string{"pulumi", "nexus", "category/cloud"},
		License:           "MIT",
		Homepage:          "https://github.com/SimCubeLtd/pulumi-nexus",
		Repository:        "https://github.com/SimCubeLtd/pulumi-nexus",
		GitHubOrg:         "datadrivers",
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"nexus_anonymous":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusAnonymous")},
			"nexus_blobstore":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusBlobstore")},
			"nexus_blobstore_azure":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusBlobstoreAzure")},
			"nexus_blobstore_file":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusBlobstoreFile")},
			"nexus_blobstore_group":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusBlobstoreGroup")},
			"nexus_blobstore_s3":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusBlobstoreS3")},
			"nexus_content_selector":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusContentSelector")},
			"nexus_privilege":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusPrivilege")},
			"nexus_repository":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepository")},
			"nexus_repository_apt_hosted":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryAptHosted")},
			"nexus_repository_apt_proxy":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryAptProxy")},
			"nexus_repository_docker_group":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryDockerGroup")},
			"nexus_repository_docker_hosted":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryDockerHosted")},
			"nexus_repository_docker_proxy":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryDockerProxy")},
			"nexus_repository_maven_hosted":   {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryMavenHosted")},
			"nexus_repository_yum_group":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryYumGroup")},
			"nexus_repository_yum_hosted":     {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryYumHosted")},
			"nexus_repository_yum_proxy":      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRepositoryYumProxy")},
			"nexus_role":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRole")},
			"nexus_routing_rule":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusRoutingRule")},
			"nexus_script":                    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusScript")},
			"nexus_security_anonymous":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecurityAnonymous")},
			"nexus_security_content_selector": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecurityContentSelector")},
			"nexus_security_ldap":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecurityLdap")},
			"nexus_security_ldap_order":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecurityLdapOrder")},
			"nexus_security_realms":           {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecurityRealms")},
			"nexus_security_role":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecurityRole")},
			"nexus_security_saml":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecuritySaml")},
			"nexus_security_user":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecurityUser")},
			"nexus_security_user_token":       {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusSecurityUserToken")},
			"nexus_user":                      {Tok: tfbridge.MakeResource(mainPkg, mainMod, "NexusUser")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"nexus_anonymous":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusAnonymous")},
			"nexus_blobstore":                 {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusBlobstore")},
			"nexus_blobstore_azure":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusBlobstoreAzure")},
			"nexus_blobstore_file":            {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusBlobstoreFile")},
			"nexus_blobstore_group":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusBlobstoreGroup")},
			"nexus_blobstore_s3":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusBlobstoreS3")},
			"nexus_privileges":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusPrivileges")},
			"nexus_repository":                {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepository")},
			"nexus_repository_apt_hosted":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryAptHosted")},
			"nexus_repository_apt_proxy":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryAptProxy")},
			"nexus_repository_docker_group":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryDockerGroup")},
			"nexus_repository_docker_hosted":  {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryDockerHosted")},
			"nexus_repository_docker_proxy":   {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryDockerProxy")},
			"nexus_repository_list":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryList")},
			"nexus_repository_yum_group":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryYumGroup")},
			"nexus_repository_yum_hosted":     {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryYumHosted")},
			"nexus_repository_yum_proxy":      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRepositoryYumProxy")},
			"nexus_routing_rule":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusRoutingRule")},
			"nexus_security_anonymous":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusSecurityAnonymous")},
			"nexus_security_content_selector": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusSecurityContentSelector")},
			"nexus_security_ldap":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusSecurityLdap")},
			"nexus_security_realms":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusSecurityRealms")},
			"nexus_security_role":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusSecurityRole")},
			"nexus_security_saml":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusSecuritySaml")},
			"nexus_security_user":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusSecurityUser")},
			"nexus_security_user_token":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusSecurityUserToken")},
			"nexus_user":                      {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "GetNexusUser")},
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
