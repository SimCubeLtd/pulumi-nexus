# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetNexusBlobstoreS3Result',
    'AwaitableGetNexusBlobstoreS3Result',
    'get_nexus_blobstore_s3',
    'get_nexus_blobstore_s3_output',
]

@pulumi.output_type
class GetNexusBlobstoreS3Result:
    """
    A collection of values returned by GetNexusBlobstoreS3.
    """
    def __init__(__self__, blob_count=None, bucket_configurations=None, id=None, name=None, soft_quotas=None, total_size_in_bytes=None):
        if blob_count and not isinstance(blob_count, int):
            raise TypeError("Expected argument 'blob_count' to be a int")
        pulumi.set(__self__, "blob_count", blob_count)
        if bucket_configurations and not isinstance(bucket_configurations, list):
            raise TypeError("Expected argument 'bucket_configurations' to be a list")
        pulumi.set(__self__, "bucket_configurations", bucket_configurations)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if soft_quotas and not isinstance(soft_quotas, list):
            raise TypeError("Expected argument 'soft_quotas' to be a list")
        pulumi.set(__self__, "soft_quotas", soft_quotas)
        if total_size_in_bytes and not isinstance(total_size_in_bytes, int):
            raise TypeError("Expected argument 'total_size_in_bytes' to be a int")
        pulumi.set(__self__, "total_size_in_bytes", total_size_in_bytes)

    @property
    @pulumi.getter(name="blobCount")
    def blob_count(self) -> int:
        return pulumi.get(self, "blob_count")

    @property
    @pulumi.getter(name="bucketConfigurations")
    def bucket_configurations(self) -> Sequence['outputs.GetNexusBlobstoreS3BucketConfigurationResult']:
        return pulumi.get(self, "bucket_configurations")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="softQuotas")
    def soft_quotas(self) -> Sequence['outputs.GetNexusBlobstoreS3SoftQuotaResult']:
        return pulumi.get(self, "soft_quotas")

    @property
    @pulumi.getter(name="totalSizeInBytes")
    def total_size_in_bytes(self) -> int:
        return pulumi.get(self, "total_size_in_bytes")


class AwaitableGetNexusBlobstoreS3Result(GetNexusBlobstoreS3Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNexusBlobstoreS3Result(
            blob_count=self.blob_count,
            bucket_configurations=self.bucket_configurations,
            id=self.id,
            name=self.name,
            soft_quotas=self.soft_quotas,
            total_size_in_bytes=self.total_size_in_bytes)


def get_nexus_blobstore_s3(name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNexusBlobstoreS3Result:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('nexus:index/getNexusBlobstoreS3:GetNexusBlobstoreS3', __args__, opts=opts, typ=GetNexusBlobstoreS3Result).value

    return AwaitableGetNexusBlobstoreS3Result(
        blob_count=__ret__.blob_count,
        bucket_configurations=__ret__.bucket_configurations,
        id=__ret__.id,
        name=__ret__.name,
        soft_quotas=__ret__.soft_quotas,
        total_size_in_bytes=__ret__.total_size_in_bytes)


@_utilities.lift_output_func(get_nexus_blobstore_s3)
def get_nexus_blobstore_s3_output(name: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNexusBlobstoreS3Result]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
