# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetNexusUserResult',
    'AwaitableGetNexusUserResult',
    'get_nexus_user',
    'get_nexus_user_output',
]

@pulumi.output_type
class GetNexusUserResult:
    """
    A collection of values returned by GetNexusUser.
    """
    def __init__(__self__, email=None, firstname=None, id=None, lastname=None, roles=None, status=None, userid=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if firstname and not isinstance(firstname, str):
            raise TypeError("Expected argument 'firstname' to be a str")
        pulumi.set(__self__, "firstname", firstname)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lastname and not isinstance(lastname, str):
            raise TypeError("Expected argument 'lastname' to be a str")
        pulumi.set(__self__, "lastname", lastname)
        if roles and not isinstance(roles, list):
            raise TypeError("Expected argument 'roles' to be a list")
        pulumi.set(__self__, "roles", roles)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if userid and not isinstance(userid, str):
            raise TypeError("Expected argument 'userid' to be a str")
        pulumi.set(__self__, "userid", userid)

    @property
    @pulumi.getter
    def email(self) -> str:
        return pulumi.get(self, "email")

    @property
    @pulumi.getter
    def firstname(self) -> str:
        return pulumi.get(self, "firstname")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def lastname(self) -> str:
        return pulumi.get(self, "lastname")

    @property
    @pulumi.getter
    def roles(self) -> Sequence[str]:
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def userid(self) -> str:
        return pulumi.get(self, "userid")


class AwaitableGetNexusUserResult(GetNexusUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNexusUserResult(
            email=self.email,
            firstname=self.firstname,
            id=self.id,
            lastname=self.lastname,
            roles=self.roles,
            status=self.status,
            userid=self.userid)


def get_nexus_user(userid: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNexusUserResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['userid'] = userid
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('nexus:index/getNexusUser:GetNexusUser', __args__, opts=opts, typ=GetNexusUserResult).value

    return AwaitableGetNexusUserResult(
        email=__ret__.email,
        firstname=__ret__.firstname,
        id=__ret__.id,
        lastname=__ret__.lastname,
        roles=__ret__.roles,
        status=__ret__.status,
        userid=__ret__.userid)


@_utilities.lift_output_func(get_nexus_user)
def get_nexus_user_output(userid: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNexusUserResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...