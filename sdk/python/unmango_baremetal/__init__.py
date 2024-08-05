# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import unmango_baremetal.cmd as __cmd
    cmd = __cmd
    import unmango_baremetal.config as __config
    config = __config
else:
    cmd = _utilities.lazy_import('unmango_baremetal.cmd')
    config = _utilities.lazy_import('unmango_baremetal.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "baremetal",
  "mod": "cmd",
  "fqn": "unmango_baremetal.cmd",
  "classes": {
   "baremetal:cmd:Mkdir": "Mkdir",
   "baremetal:cmd:Mktemp": "Mktemp",
   "baremetal:cmd:Mv": "Mv",
   "baremetal:cmd:Rm": "Rm",
   "baremetal:cmd:Tar": "Tar",
   "baremetal:cmd:Tee": "Tee",
   "baremetal:cmd:Wget": "Wget"
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
