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

__all__ = ['MvArgs', 'Mv']

@pulumi.input_type
class MvArgs:
    def __init__(__self__, *,
                 source: pulumi.Input[Sequence[pulumi.Input[str]]],
                 backup: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 no_clobber: Optional[pulumi.Input[bool]] = None,
                 no_target_directory: Optional[pulumi.Input[bool]] = None,
                 strip_trailing_slashes: Optional[pulumi.Input[bool]] = None,
                 suffix: Optional[pulumi.Input[str]] = None,
                 target_directory: Optional[pulumi.Input[str]] = None,
                 update: Optional[pulumi.Input[bool]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Mv resource.
        """
        pulumi.set(__self__, "source", source)
        if backup is not None:
            pulumi.set(__self__, "backup", backup)
        if destination is not None:
            pulumi.set(__self__, "destination", destination)
        if directory is not None:
            pulumi.set(__self__, "directory", directory)
        if force is not None:
            pulumi.set(__self__, "force", force)
        if help is not None:
            pulumi.set(__self__, "help", help)
        if no_clobber is not None:
            pulumi.set(__self__, "no_clobber", no_clobber)
        if no_target_directory is not None:
            pulumi.set(__self__, "no_target_directory", no_target_directory)
        if strip_trailing_slashes is not None:
            pulumi.set(__self__, "strip_trailing_slashes", strip_trailing_slashes)
        if suffix is not None:
            pulumi.set(__self__, "suffix", suffix)
        if target_directory is not None:
            pulumi.set(__self__, "target_directory", target_directory)
        if update is not None:
            pulumi.set(__self__, "update", update)
        if verbose is not None:
            pulumi.set(__self__, "verbose", verbose)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def backup(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "backup")

    @backup.setter
    def backup(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup", value)

    @property
    @pulumi.getter
    def destination(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter
    def directory(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "directory")

    @directory.setter
    def directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory", value)

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
    @pulumi.getter(name="noClobber")
    def no_clobber(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_clobber")

    @no_clobber.setter
    def no_clobber(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_clobber", value)

    @property
    @pulumi.getter(name="noTargetDirectory")
    def no_target_directory(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_target_directory")

    @no_target_directory.setter
    def no_target_directory(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_target_directory", value)

    @property
    @pulumi.getter(name="stripTrailingSlashes")
    def strip_trailing_slashes(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "strip_trailing_slashes")

    @strip_trailing_slashes.setter
    def strip_trailing_slashes(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "strip_trailing_slashes", value)

    @property
    @pulumi.getter
    def suffix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "suffix")

    @suffix.setter
    def suffix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "suffix", value)

    @property
    @pulumi.getter(name="targetDirectory")
    def target_directory(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "target_directory")

    @target_directory.setter
    def target_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_directory", value)

    @property
    @pulumi.getter
    def update(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "update")

    @update.setter
    def update(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "update", value)

    @property
    @pulumi.getter
    def verbose(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "verbose")

    @verbose.setter
    def verbose(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verbose", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "version", value)


class Mv(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 no_clobber: Optional[pulumi.Input[bool]] = None,
                 no_target_directory: Optional[pulumi.Input[bool]] = None,
                 source: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 strip_trailing_slashes: Optional[pulumi.Input[bool]] = None,
                 suffix: Optional[pulumi.Input[str]] = None,
                 target_directory: Optional[pulumi.Input[str]] = None,
                 update: Optional[pulumi.Input[bool]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a Mv resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MvArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Mv resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param MvArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MvArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup: Optional[pulumi.Input[str]] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 directory: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 no_clobber: Optional[pulumi.Input[bool]] = None,
                 no_target_directory: Optional[pulumi.Input[bool]] = None,
                 source: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 strip_trailing_slashes: Optional[pulumi.Input[bool]] = None,
                 suffix: Optional[pulumi.Input[str]] = None,
                 target_directory: Optional[pulumi.Input[str]] = None,
                 update: Optional[pulumi.Input[bool]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MvArgs.__new__(MvArgs)

            __props__.__dict__["backup"] = backup
            __props__.__dict__["destination"] = destination
            __props__.__dict__["directory"] = directory
            __props__.__dict__["force"] = force
            __props__.__dict__["help"] = help
            __props__.__dict__["no_clobber"] = no_clobber
            __props__.__dict__["no_target_directory"] = no_target_directory
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["strip_trailing_slashes"] = strip_trailing_slashes
            __props__.__dict__["suffix"] = suffix
            __props__.__dict__["target_directory"] = target_directory
            __props__.__dict__["update"] = update
            __props__.__dict__["verbose"] = verbose
            __props__.__dict__["version"] = version
            __props__.__dict__["args"] = None
            __props__.__dict__["created_files"] = None
            __props__.__dict__["exit_code"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Mv, __self__).__init__(
            'baremetal:cmd:Mv',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Mv':
        """
        Get an existing Mv resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MvArgs.__new__(MvArgs)

        __props__.__dict__["args"] = None
        __props__.__dict__["created_files"] = None
        __props__.__dict__["exit_code"] = None
        __props__.__dict__["stderr"] = None
        __props__.__dict__["stdout"] = None
        return Mv(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output['outputs.MvArgs']:
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
    @pulumi.getter
    def stderr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stderr")

    @property
    @pulumi.getter
    def stdout(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stdout")

