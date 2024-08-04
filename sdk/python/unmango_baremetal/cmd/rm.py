# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = ['RmArgs', 'Rm']

@pulumi.input_type
class RmArgs:
    def __init__(__self__, *,
                 files: pulumi.Input[Sequence[pulumi.Input[str]]],
                 dir: Optional[pulumi.Input[bool]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 one_file_system: Optional[pulumi.Input[bool]] = None,
                 recursive: Optional[pulumi.Input[bool]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Rm resource.
        """
        pulumi.set(__self__, "files", files)
        if dir is not None:
            pulumi.set(__self__, "dir", dir)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if help is not None:
            pulumi.set(__self__, "help", help)
        if one_file_system is not None:
            pulumi.set(__self__, "one_file_system", one_file_system)
        if recursive is not None:
            pulumi.set(__self__, "recursive", recursive)
        if verbose is not None:
            pulumi.set(__self__, "verbose", verbose)

    @property
    @pulumi.getter
    def files(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "files")

    @files.setter
    def files(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "files", value)

    @property
    @pulumi.getter
    def dir(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "dir")

    @dir.setter
    def dir(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dir", value)

    @property
    @pulumi.getter
    def force(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "force")

    @force.setter
    def force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force", value)

    @property
    @pulumi.getter
    def help(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "help")

    @help.setter
    def help(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "help", value)

    @property
    @pulumi.getter(name="oneFileSystem")
    def one_file_system(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "one_file_system")

    @one_file_system.setter
    def one_file_system(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "one_file_system", value)

    @property
    @pulumi.getter
    def recursive(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "recursive")

    @recursive.setter
    def recursive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "recursive", value)

    @property
    @pulumi.getter
    def verbose(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "verbose")

    @verbose.setter
    def verbose(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verbose", value)


class Rm(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dir: Optional[pulumi.Input[bool]] = None,
                 files: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 one_file_system: Optional[pulumi.Input[bool]] = None,
                 recursive: Optional[pulumi.Input[bool]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a Rm resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RmArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Rm resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RmArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RmArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dir: Optional[pulumi.Input[bool]] = None,
                 files: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 one_file_system: Optional[pulumi.Input[bool]] = None,
                 recursive: Optional[pulumi.Input[bool]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RmArgs.__new__(RmArgs)

            __props__.__dict__["dir"] = dir
            if files is None and not opts.urn:
                raise TypeError("Missing required property 'files'")
            __props__.__dict__["files"] = files
            __props__.__dict__["force"] = force
            __props__.__dict__["help"] = help
            __props__.__dict__["one_file_system"] = one_file_system
            __props__.__dict__["recursive"] = recursive
            __props__.__dict__["verbose"] = verbose
            __props__.__dict__["args"] = None
            __props__.__dict__["created_files"] = None
            __props__.__dict__["exit_code"] = None
            __props__.__dict__["moved_files"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Rm, __self__).__init__(
            'baremetal:cmd:Rm',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Rm':
        """
        Get an existing Rm resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RmArgs.__new__(RmArgs)

        __props__.__dict__["args"] = None
        __props__.__dict__["created_files"] = None
        __props__.__dict__["exit_code"] = None
        __props__.__dict__["moved_files"] = None
        __props__.__dict__["stderr"] = None
        __props__.__dict__["stdout"] = None
        return Rm(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output['outputs.RmArgs']:
        return pulumi.get(self, "args")

    @property
    @pulumi.getter(name="createdFiles")
    def created_files(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "created_files")

    @property
    @pulumi.getter(name="exitCode")
    def exit_code(self) -> pulumi.Output[int]:
        return pulumi.get(self, "exit_code")

    @property
    @pulumi.getter(name="movedFiles")
    def moved_files(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "moved_files")

    @property
    @pulumi.getter
    def stderr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stderr")

    @property
    @pulumi.getter
    def stdout(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stdout")

