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
    'GetNexusRepositoryYumProxyResult',
    'AwaitableGetNexusRepositoryYumProxyResult',
    'get_nexus_repository_yum_proxy',
    'get_nexus_repository_yum_proxy_output',
]

@pulumi.output_type
class GetNexusRepositoryYumProxyResult:
    """
    A collection of values returned by GetNexusRepositoryYumProxy.
    """
    def __init__(__self__, cleanups=None, http_clients=None, id=None, name=None, negative_caches=None, online=None, proxies=None, routing_rule=None, storages=None, yum_signings=None):
        if cleanups and not isinstance(cleanups, list):
            raise TypeError("Expected argument 'cleanups' to be a list")
        pulumi.set(__self__, "cleanups", cleanups)
        if http_clients and not isinstance(http_clients, list):
            raise TypeError("Expected argument 'http_clients' to be a list")
        pulumi.set(__self__, "http_clients", http_clients)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if negative_caches and not isinstance(negative_caches, list):
            raise TypeError("Expected argument 'negative_caches' to be a list")
        pulumi.set(__self__, "negative_caches", negative_caches)
        if online and not isinstance(online, bool):
            raise TypeError("Expected argument 'online' to be a bool")
        pulumi.set(__self__, "online", online)
        if proxies and not isinstance(proxies, list):
            raise TypeError("Expected argument 'proxies' to be a list")
        pulumi.set(__self__, "proxies", proxies)
        if routing_rule and not isinstance(routing_rule, str):
            raise TypeError("Expected argument 'routing_rule' to be a str")
        pulumi.set(__self__, "routing_rule", routing_rule)
        if storages and not isinstance(storages, list):
            raise TypeError("Expected argument 'storages' to be a list")
        pulumi.set(__self__, "storages", storages)
        if yum_signings and not isinstance(yum_signings, list):
            raise TypeError("Expected argument 'yum_signings' to be a list")
        pulumi.set(__self__, "yum_signings", yum_signings)

    @property
    @pulumi.getter
    def cleanups(self) -> Sequence['outputs.GetNexusRepositoryYumProxyCleanupResult']:
        return pulumi.get(self, "cleanups")

    @property
    @pulumi.getter(name="httpClients")
    def http_clients(self) -> Sequence['outputs.GetNexusRepositoryYumProxyHttpClientResult']:
        return pulumi.get(self, "http_clients")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="negativeCaches")
    def negative_caches(self) -> Sequence['outputs.GetNexusRepositoryYumProxyNegativeCachResult']:
        return pulumi.get(self, "negative_caches")

    @property
    @pulumi.getter
    def online(self) -> bool:
        return pulumi.get(self, "online")

    @property
    @pulumi.getter
    def proxies(self) -> Sequence['outputs.GetNexusRepositoryYumProxyProxyResult']:
        return pulumi.get(self, "proxies")

    @property
    @pulumi.getter(name="routingRule")
    def routing_rule(self) -> str:
        return pulumi.get(self, "routing_rule")

    @property
    @pulumi.getter
    def storages(self) -> Sequence['outputs.GetNexusRepositoryYumProxyStorageResult']:
        return pulumi.get(self, "storages")

    @property
    @pulumi.getter(name="yumSignings")
    def yum_signings(self) -> Sequence['outputs.GetNexusRepositoryYumProxyYumSigningResult']:
        return pulumi.get(self, "yum_signings")


class AwaitableGetNexusRepositoryYumProxyResult(GetNexusRepositoryYumProxyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNexusRepositoryYumProxyResult(
            cleanups=self.cleanups,
            http_clients=self.http_clients,
            id=self.id,
            name=self.name,
            negative_caches=self.negative_caches,
            online=self.online,
            proxies=self.proxies,
            routing_rule=self.routing_rule,
            storages=self.storages,
            yum_signings=self.yum_signings)


def get_nexus_repository_yum_proxy(name: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNexusRepositoryYumProxyResult:
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
    __ret__ = pulumi.runtime.invoke('nexus:index/getNexusRepositoryYumProxy:GetNexusRepositoryYumProxy', __args__, opts=opts, typ=GetNexusRepositoryYumProxyResult).value

    return AwaitableGetNexusRepositoryYumProxyResult(
        cleanups=__ret__.cleanups,
        http_clients=__ret__.http_clients,
        id=__ret__.id,
        name=__ret__.name,
        negative_caches=__ret__.negative_caches,
        online=__ret__.online,
        proxies=__ret__.proxies,
        routing_rule=__ret__.routing_rule,
        storages=__ret__.storages,
        yum_signings=__ret__.yum_signings)


@_utilities.lift_output_func(get_nexus_repository_yum_proxy)
def get_nexus_repository_yum_proxy_output(name: Optional[pulumi.Input[str]] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNexusRepositoryYumProxyResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
