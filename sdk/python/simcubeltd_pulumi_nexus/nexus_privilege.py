# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NexusPrivilegeArgs', 'NexusPrivilege']

@pulumi.input_type
class NexusPrivilegeArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 type: pulumi.Input[str],
                 content_selector: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 script_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NexusPrivilege resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: Actions for the privilege (browse, read, edit, add, delete, all and run)
        :param pulumi.Input[str] type: The type of the privilege. Possible values: `application`, `repository-admin`, `repository-content-selector`,
               `repository-view`, `script`, `wildcard`
        :param pulumi.Input[str] content_selector: The content selector for the privilege
        :param pulumi.Input[str] description: A description of the privilege
        :param pulumi.Input[str] domain: The domain of the privilege
        :param pulumi.Input[str] format: The format of the privilege. Possible values: `apt`, `bower`, `conan`, `docker`, `gitlfs`, `go`, `helm`, `maven2`,
               `npm`, `nuget`, `p2`, `pypi`, `raw`, `rubygems`, `yum`
        :param pulumi.Input[str] name: The name of the privilege
        :param pulumi.Input[str] pattern: The wildcard privilege pattern
        :param pulumi.Input[str] repository: The repository of the privilege
        :param pulumi.Input[str] script_name: The script name related to the privilege
        """
        pulumi.set(__self__, "actions", actions)
        pulumi.set(__self__, "type", type)
        if content_selector is not None:
            pulumi.set(__self__, "content_selector", content_selector)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if script_name is not None:
            pulumi.set(__self__, "script_name", script_name)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Actions for the privilege (browse, read, edit, add, delete, all and run)
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the privilege. Possible values: `application`, `repository-admin`, `repository-content-selector`,
        `repository-view`, `script`, `wildcard`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="contentSelector")
    def content_selector(self) -> Optional[pulumi.Input[str]]:
        """
        The content selector for the privilege
        """
        return pulumi.get(self, "content_selector")

    @content_selector.setter
    def content_selector(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_selector", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the privilege
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain of the privilege
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        The format of the privilege. Possible values: `apt`, `bower`, `conan`, `docker`, `gitlfs`, `go`, `helm`, `maven2`,
        `npm`, `nuget`, `p2`, `pypi`, `raw`, `rubygems`, `yum`
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the privilege
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The wildcard privilege pattern
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The repository of the privilege
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="scriptName")
    def script_name(self) -> Optional[pulumi.Input[str]]:
        """
        The script name related to the privilege
        """
        return pulumi.get(self, "script_name")

    @script_name.setter
    def script_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "script_name", value)


@pulumi.input_type
class _NexusPrivilegeState:
    def __init__(__self__, *,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 content_selector: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 script_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NexusPrivilege resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: Actions for the privilege (browse, read, edit, add, delete, all and run)
        :param pulumi.Input[str] content_selector: The content selector for the privilege
        :param pulumi.Input[str] description: A description of the privilege
        :param pulumi.Input[str] domain: The domain of the privilege
        :param pulumi.Input[str] format: The format of the privilege. Possible values: `apt`, `bower`, `conan`, `docker`, `gitlfs`, `go`, `helm`, `maven2`,
               `npm`, `nuget`, `p2`, `pypi`, `raw`, `rubygems`, `yum`
        :param pulumi.Input[str] name: The name of the privilege
        :param pulumi.Input[str] pattern: The wildcard privilege pattern
        :param pulumi.Input[str] repository: The repository of the privilege
        :param pulumi.Input[str] script_name: The script name related to the privilege
        :param pulumi.Input[str] type: The type of the privilege. Possible values: `application`, `repository-admin`, `repository-content-selector`,
               `repository-view`, `script`, `wildcard`
        """
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if content_selector is not None:
            pulumi.set(__self__, "content_selector", content_selector)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if repository is not None:
            pulumi.set(__self__, "repository", repository)
        if script_name is not None:
            pulumi.set(__self__, "script_name", script_name)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Actions for the privilege (browse, read, edit, add, delete, all and run)
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter(name="contentSelector")
    def content_selector(self) -> Optional[pulumi.Input[str]]:
        """
        The content selector for the privilege
        """
        return pulumi.get(self, "content_selector")

    @content_selector.setter
    def content_selector(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_selector", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the privilege
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain of the privilege
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        The format of the privilege. Possible values: `apt`, `bower`, `conan`, `docker`, `gitlfs`, `go`, `helm`, `maven2`,
        `npm`, `nuget`, `p2`, `pypi`, `raw`, `rubygems`, `yum`
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the privilege
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The wildcard privilege pattern
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def repository(self) -> Optional[pulumi.Input[str]]:
        """
        The repository of the privilege
        """
        return pulumi.get(self, "repository")

    @repository.setter
    def repository(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository", value)

    @property
    @pulumi.getter(name="scriptName")
    def script_name(self) -> Optional[pulumi.Input[str]]:
        """
        The script name related to the privilege
        """
        return pulumi.get(self, "script_name")

    @script_name.setter
    def script_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "script_name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the privilege. Possible values: `application`, `repository-admin`, `repository-content-selector`,
        `repository-view`, `script`, `wildcard`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class NexusPrivilege(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 content_selector: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 script_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NexusPrivilege resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: Actions for the privilege (browse, read, edit, add, delete, all and run)
        :param pulumi.Input[str] content_selector: The content selector for the privilege
        :param pulumi.Input[str] description: A description of the privilege
        :param pulumi.Input[str] domain: The domain of the privilege
        :param pulumi.Input[str] format: The format of the privilege. Possible values: `apt`, `bower`, `conan`, `docker`, `gitlfs`, `go`, `helm`, `maven2`,
               `npm`, `nuget`, `p2`, `pypi`, `raw`, `rubygems`, `yum`
        :param pulumi.Input[str] name: The name of the privilege
        :param pulumi.Input[str] pattern: The wildcard privilege pattern
        :param pulumi.Input[str] repository: The repository of the privilege
        :param pulumi.Input[str] script_name: The script name related to the privilege
        :param pulumi.Input[str] type: The type of the privilege. Possible values: `application`, `repository-admin`, `repository-content-selector`,
               `repository-view`, `script`, `wildcard`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NexusPrivilegeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NexusPrivilege resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NexusPrivilegeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NexusPrivilegeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 content_selector: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 repository: Optional[pulumi.Input[str]] = None,
                 script_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NexusPrivilegeArgs.__new__(NexusPrivilegeArgs)

            if actions is None and not opts.urn:
                raise TypeError("Missing required property 'actions'")
            __props__.__dict__["actions"] = actions
            __props__.__dict__["content_selector"] = content_selector
            __props__.__dict__["description"] = description
            __props__.__dict__["domain"] = domain
            __props__.__dict__["format"] = format
            __props__.__dict__["name"] = name
            __props__.__dict__["pattern"] = pattern
            __props__.__dict__["repository"] = repository
            __props__.__dict__["script_name"] = script_name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(NexusPrivilege, __self__).__init__(
            'nexus:index/nexusPrivilege:NexusPrivilege',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            actions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            content_selector: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            pattern: Optional[pulumi.Input[str]] = None,
            repository: Optional[pulumi.Input[str]] = None,
            script_name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'NexusPrivilege':
        """
        Get an existing NexusPrivilege resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] actions: Actions for the privilege (browse, read, edit, add, delete, all and run)
        :param pulumi.Input[str] content_selector: The content selector for the privilege
        :param pulumi.Input[str] description: A description of the privilege
        :param pulumi.Input[str] domain: The domain of the privilege
        :param pulumi.Input[str] format: The format of the privilege. Possible values: `apt`, `bower`, `conan`, `docker`, `gitlfs`, `go`, `helm`, `maven2`,
               `npm`, `nuget`, `p2`, `pypi`, `raw`, `rubygems`, `yum`
        :param pulumi.Input[str] name: The name of the privilege
        :param pulumi.Input[str] pattern: The wildcard privilege pattern
        :param pulumi.Input[str] repository: The repository of the privilege
        :param pulumi.Input[str] script_name: The script name related to the privilege
        :param pulumi.Input[str] type: The type of the privilege. Possible values: `application`, `repository-admin`, `repository-content-selector`,
               `repository-view`, `script`, `wildcard`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NexusPrivilegeState.__new__(_NexusPrivilegeState)

        __props__.__dict__["actions"] = actions
        __props__.__dict__["content_selector"] = content_selector
        __props__.__dict__["description"] = description
        __props__.__dict__["domain"] = domain
        __props__.__dict__["format"] = format
        __props__.__dict__["name"] = name
        __props__.__dict__["pattern"] = pattern
        __props__.__dict__["repository"] = repository
        __props__.__dict__["script_name"] = script_name
        __props__.__dict__["type"] = type
        return NexusPrivilege(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Sequence[str]]:
        """
        Actions for the privilege (browse, read, edit, add, delete, all and run)
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="contentSelector")
    def content_selector(self) -> pulumi.Output[Optional[str]]:
        """
        The content selector for the privilege
        """
        return pulumi.get(self, "content_selector")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the privilege
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[Optional[str]]:
        """
        The domain of the privilege
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[Optional[str]]:
        """
        The format of the privilege. Possible values: `apt`, `bower`, `conan`, `docker`, `gitlfs`, `go`, `helm`, `maven2`,
        `npm`, `nuget`, `p2`, `pypi`, `raw`, `rubygems`, `yum`
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the privilege
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[Optional[str]]:
        """
        The wildcard privilege pattern
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter
    def repository(self) -> pulumi.Output[Optional[str]]:
        """
        The repository of the privilege
        """
        return pulumi.get(self, "repository")

    @property
    @pulumi.getter(name="scriptName")
    def script_name(self) -> pulumi.Output[Optional[str]]:
        """
        The script name related to the privilege
        """
        return pulumi.get(self, "script_name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the privilege. Possible values: `application`, `repository-admin`, `repository-content-selector`,
        `repository-view`, `script`, `wildcard`
        """
        return pulumi.get(self, "type")

