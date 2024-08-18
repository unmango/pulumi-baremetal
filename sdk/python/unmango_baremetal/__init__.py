# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .command import *
from .provider import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import unmango_baremetal.config as __config
    config = __config
    import unmango_baremetal.coreutils as __coreutils
    coreutils = __coreutils
    import unmango_baremetal.kubeadm as __kubeadm
    kubeadm = __kubeadm
else:
    config = _utilities.lazy_import('unmango_baremetal.config')
    coreutils = _utilities.lazy_import('unmango_baremetal.coreutils')
    kubeadm = _utilities.lazy_import('unmango_baremetal.kubeadm')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "baremetal",
  "mod": "coreutils",
  "fqn": "unmango_baremetal.coreutils",
  "classes": {
   "baremetal:coreutils:Chmod": "Chmod",
   "baremetal:coreutils:Mkdir": "Mkdir",
   "baremetal:coreutils:Mktemp": "Mktemp",
   "baremetal:coreutils:Mv": "Mv",
   "baremetal:coreutils:Rm": "Rm",
   "baremetal:coreutils:Tar": "Tar",
   "baremetal:coreutils:Tee": "Tee",
   "baremetal:coreutils:Wget": "Wget"
  }
 },
 {
  "pkg": "baremetal",
  "mod": "index",
  "fqn": "unmango_baremetal",
  "classes": {
   "baremetal:index:Command": "Command"
  }
 },
 {
  "pkg": "baremetal",
  "mod": "kubeadm",
  "fqn": "unmango_baremetal.kubeadm",
  "classes": {
   "baremetal:kubeadm:Kubeadm": "Kubeadm"
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
