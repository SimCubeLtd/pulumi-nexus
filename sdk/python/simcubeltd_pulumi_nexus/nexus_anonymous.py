# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NexusAnonymousArgs', 'NexusAnonymous']

@pulumi.input_type
class NexusAnonymousArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 realm_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NexusAnonymous resource.
        :param pulumi.Input[bool] enabled: Activate the anonymous access to the repository manager. Default: false
        :param pulumi.Input[str] realm_name: The name of the used realm. Default: "NexusAuthorizingRealm"
        :param pulumi.Input[str] user_id: The user id used by anonymous access. Default: "anonymous"
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if realm_name is not None:
            pulumi.set(__self__, "realm_name", realm_name)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Activate the anonymous access to the repository manager. Default: false
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="realmName")
    def realm_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the used realm. Default: "NexusAuthorizingRealm"
        """
        return pulumi.get(self, "realm_name")

    @realm_name.setter
    def realm_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_name", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The user id used by anonymous access. Default: "anonymous"
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


@pulumi.input_type
class _NexusAnonymousState:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 realm_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NexusAnonymous resources.
        :param pulumi.Input[bool] enabled: Activate the anonymous access to the repository manager. Default: false
        :param pulumi.Input[str] realm_name: The name of the used realm. Default: "NexusAuthorizingRealm"
        :param pulumi.Input[str] user_id: The user id used by anonymous access. Default: "anonymous"
        """
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if realm_name is not None:
            pulumi.set(__self__, "realm_name", realm_name)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Activate the anonymous access to the repository manager. Default: false
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="realmName")
    def realm_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the used realm. Default: "NexusAuthorizingRealm"
        """
        return pulumi.get(self, "realm_name")

    @realm_name.setter
    def realm_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_name", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The user id used by anonymous access. Default: "anonymous"
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class NexusAnonymous(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 realm_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NexusAnonymous resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Activate the anonymous access to the repository manager. Default: false
        :param pulumi.Input[str] realm_name: The name of the used realm. Default: "NexusAuthorizingRealm"
        :param pulumi.Input[str] user_id: The user id used by anonymous access. Default: "anonymous"
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NexusAnonymousArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NexusAnonymous resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NexusAnonymousArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NexusAnonymousArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 realm_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = NexusAnonymousArgs.__new__(NexusAnonymousArgs)

            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["realm_name"] = realm_name
            __props__.__dict__["user_id"] = user_id
        super(NexusAnonymous, __self__).__init__(
            'nexus:index/nexusAnonymous:NexusAnonymous',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            realm_name: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'NexusAnonymous':
        """
        Get an existing NexusAnonymous resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Activate the anonymous access to the repository manager. Default: false
        :param pulumi.Input[str] realm_name: The name of the used realm. Default: "NexusAuthorizingRealm"
        :param pulumi.Input[str] user_id: The user id used by anonymous access. Default: "anonymous"
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NexusAnonymousState.__new__(_NexusAnonymousState)

        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["realm_name"] = realm_name
        __props__.__dict__["user_id"] = user_id
        return NexusAnonymous(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Activate the anonymous access to the repository manager. Default: false
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="realmName")
    def realm_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the used realm. Default: "NexusAuthorizingRealm"
        """
        return pulumi.get(self, "realm_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[Optional[str]]:
        """
        The user id used by anonymous access. Default: "anonymous"
        """
        return pulumi.get(self, "user_id")
