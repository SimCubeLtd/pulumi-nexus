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
    'GetNexusSecurityLdapResult',
    'AwaitableGetNexusSecurityLdapResult',
    'get_nexus_security_ldap',
]

@pulumi.output_type
class GetNexusSecurityLdapResult:
    """
    A collection of values returned by GetNexusSecurityLdap.
    """
    def __init__(__self__, id=None, ldaps=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ldaps and not isinstance(ldaps, list):
            raise TypeError("Expected argument 'ldaps' to be a list")
        pulumi.set(__self__, "ldaps", ldaps)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ldaps(self) -> Sequence['outputs.GetNexusSecurityLdapLdapResult']:
        return pulumi.get(self, "ldaps")


class AwaitableGetNexusSecurityLdapResult(GetNexusSecurityLdapResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNexusSecurityLdapResult(
            id=self.id,
            ldaps=self.ldaps)


def get_nexus_security_ldap(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNexusSecurityLdapResult:
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
    __ret__ = pulumi.runtime.invoke('nexus:index/getNexusSecurityLdap:GetNexusSecurityLdap', __args__, opts=opts, typ=GetNexusSecurityLdapResult).value

    return AwaitableGetNexusSecurityLdapResult(
        id=__ret__.id,
        ldaps=__ret__.ldaps)
