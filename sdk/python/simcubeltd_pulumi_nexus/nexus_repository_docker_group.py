# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['NexusRepositoryDockerGroupArgs', 'NexusRepositoryDockerGroup']

@pulumi.input_type
class NexusRepositoryDockerGroupArgs:
    def __init__(__self__, *,
                 docker: pulumi.Input['NexusRepositoryDockerGroupDockerArgs'],
                 group: pulumi.Input['NexusRepositoryDockerGroupGroupArgs'],
                 storage: pulumi.Input['NexusRepositoryDockerGroupStorageArgs'],
                 name: Optional[pulumi.Input[str]] = None,
                 online: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a NexusRepositoryDockerGroup resource.
        :param pulumi.Input['NexusRepositoryDockerGroupDockerArgs'] docker: docker contains the configuration of the docker repository
        :param pulumi.Input['NexusRepositoryDockerGroupGroupArgs'] group: Configuration for repository group
        :param pulumi.Input['NexusRepositoryDockerGroupStorageArgs'] storage: The storage configuration of the repository
        :param pulumi.Input[str] name: A unique identifier for this repository
        :param pulumi.Input[bool] online: Whether this repository accepts incoming requests
        """
        pulumi.set(__self__, "docker", docker)
        pulumi.set(__self__, "group", group)
        pulumi.set(__self__, "storage", storage)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if online is not None:
            pulumi.set(__self__, "online", online)

    @property
    @pulumi.getter
    def docker(self) -> pulumi.Input['NexusRepositoryDockerGroupDockerArgs']:
        """
        docker contains the configuration of the docker repository
        """
        return pulumi.get(self, "docker")

    @docker.setter
    def docker(self, value: pulumi.Input['NexusRepositoryDockerGroupDockerArgs']):
        pulumi.set(self, "docker", value)

    @property
    @pulumi.getter
    def group(self) -> pulumi.Input['NexusRepositoryDockerGroupGroupArgs']:
        """
        Configuration for repository group
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: pulumi.Input['NexusRepositoryDockerGroupGroupArgs']):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Input['NexusRepositoryDockerGroupStorageArgs']:
        """
        The storage configuration of the repository
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: pulumi.Input['NexusRepositoryDockerGroupStorageArgs']):
        pulumi.set(self, "storage", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for this repository
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def online(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this repository accepts incoming requests
        """
        return pulumi.get(self, "online")

    @online.setter
    def online(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "online", value)


@pulumi.input_type
class _NexusRepositoryDockerGroupState:
    def __init__(__self__, *,
                 docker: Optional[pulumi.Input['NexusRepositoryDockerGroupDockerArgs']] = None,
                 group: Optional[pulumi.Input['NexusRepositoryDockerGroupGroupArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 online: Optional[pulumi.Input[bool]] = None,
                 storage: Optional[pulumi.Input['NexusRepositoryDockerGroupStorageArgs']] = None):
        """
        Input properties used for looking up and filtering NexusRepositoryDockerGroup resources.
        :param pulumi.Input['NexusRepositoryDockerGroupDockerArgs'] docker: docker contains the configuration of the docker repository
        :param pulumi.Input['NexusRepositoryDockerGroupGroupArgs'] group: Configuration for repository group
        :param pulumi.Input[str] name: A unique identifier for this repository
        :param pulumi.Input[bool] online: Whether this repository accepts incoming requests
        :param pulumi.Input['NexusRepositoryDockerGroupStorageArgs'] storage: The storage configuration of the repository
        """
        if docker is not None:
            pulumi.set(__self__, "docker", docker)
        if group is not None:
            pulumi.set(__self__, "group", group)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if online is not None:
            pulumi.set(__self__, "online", online)
        if storage is not None:
            pulumi.set(__self__, "storage", storage)

    @property
    @pulumi.getter
    def docker(self) -> Optional[pulumi.Input['NexusRepositoryDockerGroupDockerArgs']]:
        """
        docker contains the configuration of the docker repository
        """
        return pulumi.get(self, "docker")

    @docker.setter
    def docker(self, value: Optional[pulumi.Input['NexusRepositoryDockerGroupDockerArgs']]):
        pulumi.set(self, "docker", value)

    @property
    @pulumi.getter
    def group(self) -> Optional[pulumi.Input['NexusRepositoryDockerGroupGroupArgs']]:
        """
        Configuration for repository group
        """
        return pulumi.get(self, "group")

    @group.setter
    def group(self, value: Optional[pulumi.Input['NexusRepositoryDockerGroupGroupArgs']]):
        pulumi.set(self, "group", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique identifier for this repository
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def online(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this repository accepts incoming requests
        """
        return pulumi.get(self, "online")

    @online.setter
    def online(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "online", value)

    @property
    @pulumi.getter
    def storage(self) -> Optional[pulumi.Input['NexusRepositoryDockerGroupStorageArgs']]:
        """
        The storage configuration of the repository
        """
        return pulumi.get(self, "storage")

    @storage.setter
    def storage(self, value: Optional[pulumi.Input['NexusRepositoryDockerGroupStorageArgs']]):
        pulumi.set(self, "storage", value)


class NexusRepositoryDockerGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 docker: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupDockerArgs']]] = None,
                 group: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupGroupArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 online: Optional[pulumi.Input[bool]] = None,
                 storage: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupStorageArgs']]] = None,
                 __props__=None):
        """
        Create a NexusRepositoryDockerGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupDockerArgs']] docker: docker contains the configuration of the docker repository
        :param pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupGroupArgs']] group: Configuration for repository group
        :param pulumi.Input[str] name: A unique identifier for this repository
        :param pulumi.Input[bool] online: Whether this repository accepts incoming requests
        :param pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupStorageArgs']] storage: The storage configuration of the repository
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NexusRepositoryDockerGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NexusRepositoryDockerGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NexusRepositoryDockerGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NexusRepositoryDockerGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 docker: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupDockerArgs']]] = None,
                 group: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupGroupArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 online: Optional[pulumi.Input[bool]] = None,
                 storage: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupStorageArgs']]] = None,
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
            __props__ = NexusRepositoryDockerGroupArgs.__new__(NexusRepositoryDockerGroupArgs)

            if docker is None and not opts.urn:
                raise TypeError("Missing required property 'docker'")
            __props__.__dict__["docker"] = docker
            if group is None and not opts.urn:
                raise TypeError("Missing required property 'group'")
            __props__.__dict__["group"] = group
            __props__.__dict__["name"] = name
            __props__.__dict__["online"] = online
            if storage is None and not opts.urn:
                raise TypeError("Missing required property 'storage'")
            __props__.__dict__["storage"] = storage
        super(NexusRepositoryDockerGroup, __self__).__init__(
            'nexus:index/nexusRepositoryDockerGroup:NexusRepositoryDockerGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            docker: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupDockerArgs']]] = None,
            group: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupGroupArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            online: Optional[pulumi.Input[bool]] = None,
            storage: Optional[pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupStorageArgs']]] = None) -> 'NexusRepositoryDockerGroup':
        """
        Get an existing NexusRepositoryDockerGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupDockerArgs']] docker: docker contains the configuration of the docker repository
        :param pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupGroupArgs']] group: Configuration for repository group
        :param pulumi.Input[str] name: A unique identifier for this repository
        :param pulumi.Input[bool] online: Whether this repository accepts incoming requests
        :param pulumi.Input[pulumi.InputType['NexusRepositoryDockerGroupStorageArgs']] storage: The storage configuration of the repository
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NexusRepositoryDockerGroupState.__new__(_NexusRepositoryDockerGroupState)

        __props__.__dict__["docker"] = docker
        __props__.__dict__["group"] = group
        __props__.__dict__["name"] = name
        __props__.__dict__["online"] = online
        __props__.__dict__["storage"] = storage
        return NexusRepositoryDockerGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def docker(self) -> pulumi.Output['outputs.NexusRepositoryDockerGroupDocker']:
        """
        docker contains the configuration of the docker repository
        """
        return pulumi.get(self, "docker")

    @property
    @pulumi.getter
    def group(self) -> pulumi.Output['outputs.NexusRepositoryDockerGroupGroup']:
        """
        Configuration for repository group
        """
        return pulumi.get(self, "group")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique identifier for this repository
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def online(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether this repository accepts incoming requests
        """
        return pulumi.get(self, "online")

    @property
    @pulumi.getter
    def storage(self) -> pulumi.Output['outputs.NexusRepositoryDockerGroupStorage']:
        """
        The storage configuration of the repository
        """
        return pulumi.get(self, "storage")

