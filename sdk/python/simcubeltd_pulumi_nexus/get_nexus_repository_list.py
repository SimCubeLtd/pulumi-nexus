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
    'GetNexusRepositoryListResult',
    'AwaitableGetNexusRepositoryListResult',
    'get_nexus_repository_list',
]

@pulumi.output_type
class GetNexusRepositoryListResult:
    """
    A collection of values returned by GetNexusRepositoryList.
    """
    def __init__(__self__, id=None, items=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if items and not isinstance(items, list):
            raise TypeError("Expected argument 'items' to be a list")
        pulumi.set(__self__, "items", items)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def items(self) -> Sequence['outputs.GetNexusRepositoryListItemResult']:
        return pulumi.get(self, "items")


class AwaitableGetNexusRepositoryListResult(GetNexusRepositoryListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNexusRepositoryListResult(
            id=self.id,
            items=self.items)


def get_nexus_repository_list(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNexusRepositoryListResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('nexus:index/getNexusRepositoryList:GetNexusRepositoryList', __args__, opts=opts, typ=GetNexusRepositoryListResult).value

    return AwaitableGetNexusRepositoryListResult(
        id=__ret__.id,
        items=__ret__.items)