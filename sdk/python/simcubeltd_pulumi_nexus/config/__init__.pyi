# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

insecure: Optional[bool]
"""
Boolean to specify wether insecure SSL connections are allowed or not. Reading environment variable
NEXUS_INSECURE_SKIP_VERIFY. Default:`true`
"""

password: Optional[str]
"""
Password of user to connect to API. Reading environment variable NEXUS_PASSWORD. Default:`admin123`
"""

url: Optional[str]
"""
URL of Nexus to reach API. Reading environment variable NEXUS_URL. Default:`http://127.0.0.1:8080`
"""

username: Optional[str]
"""
Username used to connect to API. Reading environment variable NEXUS_USERNAME. Default:`admin`
"""

