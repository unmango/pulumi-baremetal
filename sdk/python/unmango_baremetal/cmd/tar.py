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
from ._inputs import *

__all__ = ['TarArgs', 'Tar']

@pulumi.input_type
class TarArgs:
    def __init__(__self__, *,
                 args: pulumi.Input['TarArgsArgs'],
                 custom_delete: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_update: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None):
        """
        The set of arguments for constructing a Tar resource.
        """
        pulumi.set(__self__, "args", args)
        if custom_delete is not None:
            pulumi.set(__self__, "custom_delete", custom_delete)
        if custom_update is not None:
            pulumi.set(__self__, "custom_update", custom_update)
        if triggers is not None:
            pulumi.set(__self__, "triggers", triggers)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Input['TarArgsArgs']:
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: pulumi.Input['TarArgsArgs']):
        pulumi.set(self, "args", value)

    @property
    @pulumi.getter(name="customDelete")
    def custom_delete(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_delete")

    @custom_delete.setter
    def custom_delete(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_delete", value)

    @property
    @pulumi.getter(name="customUpdate")
    def custom_update(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "custom_update")

    @custom_update.setter
    def custom_update(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "custom_update", value)

    @property
    @pulumi.getter
    def triggers(self) -> Optional[pulumi.Input[Sequence[Any]]]:
        return pulumi.get(self, "triggers")

    @triggers.setter
    def triggers(self, value: Optional[pulumi.Input[Sequence[Any]]]):
        pulumi.set(self, "triggers", value)


class Tar(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[Union['TarArgsArgs', 'TarArgsArgsDict']]] = None,
                 custom_delete: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_update: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 __props__=None):
        """
        Create a Tar resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TarArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Tar resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param TarArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TarArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[Union['TarArgsArgs', 'TarArgsArgsDict']]] = None,
                 custom_delete: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_update: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TarArgs.__new__(TarArgs)

            if args is None and not opts.urn:
                raise TypeError("Missing required property 'args'")
            __props__.__dict__["args"] = args
            __props__.__dict__["custom_delete"] = custom_delete
            __props__.__dict__["custom_update"] = custom_update
            __props__.__dict__["triggers"] = triggers
            __props__.__dict__["created_files"] = None
            __props__.__dict__["exit_code"] = None
            __props__.__dict__["moved_files"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Tar, __self__).__init__(
            'baremetal:cmd:Tar',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Tar':
        """
        Get an existing Tar resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TarArgs.__new__(TarArgs)

        __props__.__dict__["args"] = None
        __props__.__dict__["created_files"] = None
        __props__.__dict__["custom_delete"] = None
        __props__.__dict__["custom_update"] = None
        __props__.__dict__["exit_code"] = None
        __props__.__dict__["moved_files"] = None
        __props__.__dict__["stderr"] = None
        __props__.__dict__["stdout"] = None
        __props__.__dict__["triggers"] = None
        return Tar(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output['outputs.TarArgs']:
        return pulumi.get(self, "args")

    @property
    @pulumi.getter(name="createdFiles")
    def created_files(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "created_files")

    @property
    @pulumi.getter(name="customDelete")
    def custom_delete(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "custom_delete")

    @property
    @pulumi.getter(name="customUpdate")
    def custom_update(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "custom_update")

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

    @property
    @pulumi.getter
    def triggers(self) -> pulumi.Output[Optional[Sequence[Any]]]:
        return pulumi.get(self, "triggers")

