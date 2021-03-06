# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NexusSecurityUserTokenArgs', 'NexusSecurityUserToken']

@pulumi.input_type
class NexusSecurityUserTokenArgs:
    def __init__(__self__, *,
                 enabled: pulumi.Input[bool],
                 protect_content: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a NexusSecurityUserToken resource.
        :param pulumi.Input[bool] enabled: Activate the feature of user tokens.
        :param pulumi.Input[bool] protect_content: Require user tokens for repository authentication. This does not effect UI access.
        """
        pulumi.set(__self__, "enabled", enabled)
        if protect_content is not None:
            pulumi.set(__self__, "protect_content", protect_content)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Input[bool]:
        """
        Activate the feature of user tokens.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: pulumi.Input[bool]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="protectContent")
    def protect_content(self) -> Optional[pulumi.Input[bool]]:
        """
        Require user tokens for repository authentication. This does not effect UI access.
        """
        return pulumi.get(self, "protect_content")

    @protect_content.setter
    def protect_content(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "protect_content", value)


@pulumi.input_type
class _NexusSecurityUserTokenState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 protect_content: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering NexusSecurityUserToken resources.
        :param pulumi.Input[bool] enabled: Activate the feature of user tokens.
        :param pulumi.Input[bool] protect_content: Require user tokens for repository authentication. This does not effect UI access.
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if protect_content is not None:
            pulumi.set(__self__, "protect_content", protect_content)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Activate the feature of user tokens.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="protectContent")
    def protect_content(self) -> Optional[pulumi.Input[bool]]:
        """
        Require user tokens for repository authentication. This does not effect UI access.
        """
        return pulumi.get(self, "protect_content")

    @protect_content.setter
    def protect_content(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "protect_content", value)


class NexusSecurityUserToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 protect_content: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a NexusSecurityUserToken resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Activate the feature of user tokens.
        :param pulumi.Input[bool] protect_content: Require user tokens for repository authentication. This does not effect UI access.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NexusSecurityUserTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NexusSecurityUserToken resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NexusSecurityUserTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NexusSecurityUserTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 protect_content: Optional[pulumi.Input[bool]] = None,
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
            __props__ = NexusSecurityUserTokenArgs.__new__(NexusSecurityUserTokenArgs)

            if enabled is None and not opts.urn:
                raise TypeError("Missing required property 'enabled'")
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["protect_content"] = protect_content
        super(NexusSecurityUserToken, __self__).__init__(
            'nexus:index/nexusSecurityUserToken:NexusSecurityUserToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            protect_content: Optional[pulumi.Input[bool]] = None) -> 'NexusSecurityUserToken':
        """
        Get an existing NexusSecurityUserToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Activate the feature of user tokens.
        :param pulumi.Input[bool] protect_content: Require user tokens for repository authentication. This does not effect UI access.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NexusSecurityUserTokenState.__new__(_NexusSecurityUserTokenState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["protect_content"] = protect_content
        return NexusSecurityUserToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        """
        Activate the feature of user tokens.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="protectContent")
    def protect_content(self) -> pulumi.Output[Optional[bool]]:
        """
        Require user tokens for repository authentication. This does not effect UI access.
        """
        return pulumi.get(self, "protect_content")

