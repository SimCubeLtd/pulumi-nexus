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

__all__ = ['NexusBlobstoreArgs', 'NexusBlobstore']

@pulumi.input_type
class NexusBlobstoreArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 bucket_configuration: Optional[pulumi.Input['NexusBlobstoreBucketConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 soft_quota: Optional[pulumi.Input['NexusBlobstoreSoftQuotaArgs']] = None):
        """
        The set of arguments for constructing a NexusBlobstore resource.
        :param pulumi.Input[str] type: The type of the blobstore. Possible values: `S3` or `File`
        :param pulumi.Input['NexusBlobstoreBucketConfigurationArgs'] bucket_configuration: The S3 bucket configuration. Needed for blobstore type 'S3'
        :param pulumi.Input[str] name: Blobstore name
        :param pulumi.Input[str] path: The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can
               be a path relative to the sonatype-work directory
        :param pulumi.Input['NexusBlobstoreSoftQuotaArgs'] soft_quota: Soft quota of the blobstore
        """
        pulumi.set(__self__, "type", type)
        if bucket_configuration is not None:
            pulumi.set(__self__, "bucket_configuration", bucket_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if soft_quota is not None:
            pulumi.set(__self__, "soft_quota", soft_quota)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of the blobstore. Possible values: `S3` or `File`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="bucketConfiguration")
    def bucket_configuration(self) -> Optional[pulumi.Input['NexusBlobstoreBucketConfigurationArgs']]:
        """
        The S3 bucket configuration. Needed for blobstore type 'S3'
        """
        return pulumi.get(self, "bucket_configuration")

    @bucket_configuration.setter
    def bucket_configuration(self, value: Optional[pulumi.Input['NexusBlobstoreBucketConfigurationArgs']]):
        pulumi.set(self, "bucket_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Blobstore name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can
        be a path relative to the sonatype-work directory
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="softQuota")
    def soft_quota(self) -> Optional[pulumi.Input['NexusBlobstoreSoftQuotaArgs']]:
        """
        Soft quota of the blobstore
        """
        return pulumi.get(self, "soft_quota")

    @soft_quota.setter
    def soft_quota(self, value: Optional[pulumi.Input['NexusBlobstoreSoftQuotaArgs']]):
        pulumi.set(self, "soft_quota", value)


@pulumi.input_type
class _NexusBlobstoreState:
    def __init__(__self__, *,
                 available_space_in_bytes: Optional[pulumi.Input[str]] = None,
                 blob_count: Optional[pulumi.Input[int]] = None,
                 bucket_configuration: Optional[pulumi.Input['NexusBlobstoreBucketConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 soft_quota: Optional[pulumi.Input['NexusBlobstoreSoftQuotaArgs']] = None,
                 total_size_in_bytes: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NexusBlobstore resources.
        :param pulumi.Input[str] available_space_in_bytes: Available space in Bytes
        :param pulumi.Input[int] blob_count: Count of blobs
        :param pulumi.Input['NexusBlobstoreBucketConfigurationArgs'] bucket_configuration: The S3 bucket configuration. Needed for blobstore type 'S3'
        :param pulumi.Input[str] name: Blobstore name
        :param pulumi.Input[str] path: The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can
               be a path relative to the sonatype-work directory
        :param pulumi.Input['NexusBlobstoreSoftQuotaArgs'] soft_quota: Soft quota of the blobstore
        :param pulumi.Input[int] total_size_in_bytes: The total size of the blobstore in Bytes
        :param pulumi.Input[str] type: The type of the blobstore. Possible values: `S3` or `File`
        """
        if available_space_in_bytes is not None:
            pulumi.set(__self__, "available_space_in_bytes", available_space_in_bytes)
        if blob_count is not None:
            pulumi.set(__self__, "blob_count", blob_count)
        if bucket_configuration is not None:
            pulumi.set(__self__, "bucket_configuration", bucket_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if soft_quota is not None:
            pulumi.set(__self__, "soft_quota", soft_quota)
        if total_size_in_bytes is not None:
            pulumi.set(__self__, "total_size_in_bytes", total_size_in_bytes)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="availableSpaceInBytes")
    def available_space_in_bytes(self) -> Optional[pulumi.Input[str]]:
        """
        Available space in Bytes
        """
        return pulumi.get(self, "available_space_in_bytes")

    @available_space_in_bytes.setter
    def available_space_in_bytes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "available_space_in_bytes", value)

    @property
    @pulumi.getter(name="blobCount")
    def blob_count(self) -> Optional[pulumi.Input[int]]:
        """
        Count of blobs
        """
        return pulumi.get(self, "blob_count")

    @blob_count.setter
    def blob_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "blob_count", value)

    @property
    @pulumi.getter(name="bucketConfiguration")
    def bucket_configuration(self) -> Optional[pulumi.Input['NexusBlobstoreBucketConfigurationArgs']]:
        """
        The S3 bucket configuration. Needed for blobstore type 'S3'
        """
        return pulumi.get(self, "bucket_configuration")

    @bucket_configuration.setter
    def bucket_configuration(self, value: Optional[pulumi.Input['NexusBlobstoreBucketConfigurationArgs']]):
        pulumi.set(self, "bucket_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Blobstore name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can
        be a path relative to the sonatype-work directory
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="softQuota")
    def soft_quota(self) -> Optional[pulumi.Input['NexusBlobstoreSoftQuotaArgs']]:
        """
        Soft quota of the blobstore
        """
        return pulumi.get(self, "soft_quota")

    @soft_quota.setter
    def soft_quota(self, value: Optional[pulumi.Input['NexusBlobstoreSoftQuotaArgs']]):
        pulumi.set(self, "soft_quota", value)

    @property
    @pulumi.getter(name="totalSizeInBytes")
    def total_size_in_bytes(self) -> Optional[pulumi.Input[int]]:
        """
        The total size of the blobstore in Bytes
        """
        return pulumi.get(self, "total_size_in_bytes")

    @total_size_in_bytes.setter
    def total_size_in_bytes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "total_size_in_bytes", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the blobstore. Possible values: `S3` or `File`
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class NexusBlobstore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_configuration: Optional[pulumi.Input[pulumi.InputType['NexusBlobstoreBucketConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 soft_quota: Optional[pulumi.Input[pulumi.InputType['NexusBlobstoreSoftQuotaArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NexusBlobstore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NexusBlobstoreBucketConfigurationArgs']] bucket_configuration: The S3 bucket configuration. Needed for blobstore type 'S3'
        :param pulumi.Input[str] name: Blobstore name
        :param pulumi.Input[str] path: The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can
               be a path relative to the sonatype-work directory
        :param pulumi.Input[pulumi.InputType['NexusBlobstoreSoftQuotaArgs']] soft_quota: Soft quota of the blobstore
        :param pulumi.Input[str] type: The type of the blobstore. Possible values: `S3` or `File`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NexusBlobstoreArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NexusBlobstore resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NexusBlobstoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NexusBlobstoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_configuration: Optional[pulumi.Input[pulumi.InputType['NexusBlobstoreBucketConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 soft_quota: Optional[pulumi.Input[pulumi.InputType['NexusBlobstoreSoftQuotaArgs']]] = None,
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
            __props__ = NexusBlobstoreArgs.__new__(NexusBlobstoreArgs)

            __props__.__dict__["bucket_configuration"] = bucket_configuration
            __props__.__dict__["name"] = name
            __props__.__dict__["path"] = path
            __props__.__dict__["soft_quota"] = soft_quota
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["available_space_in_bytes"] = None
            __props__.__dict__["blob_count"] = None
            __props__.__dict__["total_size_in_bytes"] = None
        super(NexusBlobstore, __self__).__init__(
            'nexus:index/nexusBlobstore:NexusBlobstore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            available_space_in_bytes: Optional[pulumi.Input[str]] = None,
            blob_count: Optional[pulumi.Input[int]] = None,
            bucket_configuration: Optional[pulumi.Input[pulumi.InputType['NexusBlobstoreBucketConfigurationArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            soft_quota: Optional[pulumi.Input[pulumi.InputType['NexusBlobstoreSoftQuotaArgs']]] = None,
            total_size_in_bytes: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'NexusBlobstore':
        """
        Get an existing NexusBlobstore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] available_space_in_bytes: Available space in Bytes
        :param pulumi.Input[int] blob_count: Count of blobs
        :param pulumi.Input[pulumi.InputType['NexusBlobstoreBucketConfigurationArgs']] bucket_configuration: The S3 bucket configuration. Needed for blobstore type 'S3'
        :param pulumi.Input[str] name: Blobstore name
        :param pulumi.Input[str] path: The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can
               be a path relative to the sonatype-work directory
        :param pulumi.Input[pulumi.InputType['NexusBlobstoreSoftQuotaArgs']] soft_quota: Soft quota of the blobstore
        :param pulumi.Input[int] total_size_in_bytes: The total size of the blobstore in Bytes
        :param pulumi.Input[str] type: The type of the blobstore. Possible values: `S3` or `File`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NexusBlobstoreState.__new__(_NexusBlobstoreState)

        __props__.__dict__["available_space_in_bytes"] = available_space_in_bytes
        __props__.__dict__["blob_count"] = blob_count
        __props__.__dict__["bucket_configuration"] = bucket_configuration
        __props__.__dict__["name"] = name
        __props__.__dict__["path"] = path
        __props__.__dict__["soft_quota"] = soft_quota
        __props__.__dict__["total_size_in_bytes"] = total_size_in_bytes
        __props__.__dict__["type"] = type
        return NexusBlobstore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availableSpaceInBytes")
    def available_space_in_bytes(self) -> pulumi.Output[str]:
        """
        Available space in Bytes
        """
        return pulumi.get(self, "available_space_in_bytes")

    @property
    @pulumi.getter(name="blobCount")
    def blob_count(self) -> pulumi.Output[int]:
        """
        Count of blobs
        """
        return pulumi.get(self, "blob_count")

    @property
    @pulumi.getter(name="bucketConfiguration")
    def bucket_configuration(self) -> pulumi.Output[Optional['outputs.NexusBlobstoreBucketConfiguration']]:
        """
        The S3 bucket configuration. Needed for blobstore type 'S3'
        """
        return pulumi.get(self, "bucket_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Blobstore name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        The path to the blobstore contents. This can be an absolute path to anywhere on the system nxrm has access to or it can
        be a path relative to the sonatype-work directory
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="softQuota")
    def soft_quota(self) -> pulumi.Output[Optional['outputs.NexusBlobstoreSoftQuota']]:
        """
        Soft quota of the blobstore
        """
        return pulumi.get(self, "soft_quota")

    @property
    @pulumi.getter(name="totalSizeInBytes")
    def total_size_in_bytes(self) -> pulumi.Output[int]:
        """
        The total size of the blobstore in Bytes
        """
        return pulumi.get(self, "total_size_in_bytes")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the blobstore. Possible values: `S3` or `File`
        """
        return pulumi.get(self, "type")

