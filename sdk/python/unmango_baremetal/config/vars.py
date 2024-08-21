# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('baremetal')


class _ExportableConfig(types.ModuleType):
    @property
    def address(self) -> Optional[str]:
        return __config__.get('address')

    @property
    def ca_pem(self) -> Optional[str]:
        return __config__.get('caPem')

    @property
    def cert_pem(self) -> Optional[str]:
        return __config__.get('certPem')

    @property
    def key_pem(self) -> Optional[str]:
        return __config__.get('keyPem')

    @property
    def port(self) -> Optional[str]:
        return __config__.get('port')
