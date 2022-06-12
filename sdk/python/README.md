# 1Password Provider
[1Password Website](1password.com)

Based on the [1Password terraform provider](https://github.com/1Password/terraform-provider-nexus)

&nbsp;

The 1Password resource provider for Pulumi lets you use 1Password resources in your Infrastructure as Code deployments.


## Installing

This package is available in C#, TypeScript, Python and Go

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.nexus
```

### Node.js (Java/TypeScript)

```bash
npm install @SimCubeLtd/pulumi-nexus
```

or `yarn`:

```bash
yarn add @SimCubeLtd/pulumi-nexus
```

### Python

To use from Python, install using `pip`:

```bash
pip install simcubeltd_pulumi_nexus
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/SimCubeLtd/pulumi-nexus/sdk/go/nexus
```

## Configuration

The following configuration entries are available:

| **Key**           | **Value**                                            |
|-------------------|:-----------------------------------------------------|
| nexus:token | The access token for your 1Password Connect server   |
| nexus:url   | URL where your 1Password Connect API is being served |
