# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .tee import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import unmango_baremetal.config as __config
    config = __config
    import unmango_baremetal.v1alpha1 as __v1alpha1
    v1alpha1 = __v1alpha1
else:
    config = _utilities.lazy_import('unmango_baremetal.config')
    v1alpha1 = _utilities.lazy_import('unmango_baremetal.v1alpha1')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "baremetal",
  "mod": "index",
  "fqn": "unmango_baremetal",
  "classes": {
   "baremetal:index:Tee": "Tee"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "baremetal",
  "token": "pulumi:providers:baremetal",
  "fqn": "unmango_baremetal",
  "class": "Provider"
 }
]
"""
)
