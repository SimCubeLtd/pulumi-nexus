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
    'GetNexusBlobstoreFileResult',
    'AwaitableGetNexusBlobstoreFileResult',
    'get_nexus_blobstore_file',
    'get_nexus_blobstore_file_output',
]

@pulumi.output_type
class GetNexusBlobstoreFileResult:
    """
    A collection of values returned by GetNexusBlobstoreFile.
    """
    def __init__(__self__, available_space_in_bytes=None, blob_count=None, id=None, name=None, path=None, soft_quotas=None, total_size_in_bytes=None):
        if available_space_in_bytes and not isinstance(available_space_in_bytes, int):
            raise TypeError("Expected argument 'available_space_in_bytes' to be a int")
        pulumi.set(__self__, "available_space_in_bytes", available_space_in_bytes)
        if blob_count and not isinstance(blob_count, int):
            raise TypeError("Expected argument 'blob_count' to be a int")
        pulumi.set(__self__, "blob_count", blob_count)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if soft_quotas and not isinstance(soft_quotas, list):
            raise TypeError("Expected argument 'soft_quotas' to be a list")
        pulumi.set(__self__, "soft_quotas", soft_quotas)
        if total_size_in_bytes and not isinstance(total_size_in_bytes, int):
            raise TypeError("Expected argument 'total_size_in_bytes' to be a int")
        pulumi.set(__self__, "total_size_in_bytes", total_size_in_bytes)

    @property
    @pulumi.getter(name="availableSpaceInBytes")
    def available_space_in_bytes(self) -> int:
        return pulumi.get(self, "available_space_in_bytes")

    @property
    @pulumi.getter(name="blobCount")
    def blob_count(self) -> int:
        return pulumi.get(self, "blob_count")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="softQuotas")
    def soft_quotas(self) -> Sequence['outputs.GetNexusBlobstoreFileSoftQuotaResult']:
        return pulumi.get(self, "soft_quotas")

    @property
    @pulumi.getter(name="totalSizeInBytes")
    def total_size_in_bytes(self) -> int:
        return pulumi.get(self, "total_size_in_bytes")


class AwaitableGetNexusBlobstoreFileResult(GetNexusBlobstoreFileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNexusBlobstoreFileResult(
            available_space_in_bytes=self.available_space_in_bytes,
            blob_count=self.blob_count,
            id=self.id,
            name=self.name,
            path=self.path,
            soft_quotas=self.soft_quotas,
            total_size_in_bytes=self.total_size_in_bytes)


def get_nexus_blobstore_file(name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNexusBlobstoreFileResult:
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
    __ret__ = pulumi.runtime.invoke('nexus:index/getNexusBlobstoreFile:GetNexusBlobstoreFile', __args__, opts=opts, typ=GetNexusBlobstoreFileResult).value

    return AwaitableGetNexusBlobstoreFileResult(
        available_space_in_bytes=__ret__.available_space_in_bytes,
        blob_count=__ret__.blob_count,
        id=__ret__.id,
        name=__ret__.name,
        path=__ret__.path,
        soft_quotas=__ret__.soft_quotas,
        total_size_in_bytes=__ret__.total_size_in_bytes)


@_utilities.lift_output_func(get_nexus_blobstore_file)
def get_nexus_blobstore_file_output(name: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNexusBlobstoreFileResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
