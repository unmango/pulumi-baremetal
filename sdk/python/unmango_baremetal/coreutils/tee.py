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

__all__ = ['TeeArgs', 'Tee']

@pulumi.input_type
class TeeArgs:
    def __init__(__self__, *,
                 args: pulumi.Input['TeeArgsArgs'],
                 custom_delete: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_update: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None):
        """
        The set of arguments for constructing a Tee resource.
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
    def args(self) -> pulumi.Input['TeeArgsArgs']:
        return pulumi.get(self, "args")

    @args.setter
    def args(self, value: pulumi.Input['TeeArgsArgs']):
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


class Tee(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[Union['TeeArgsArgs', 'TeeArgsArgsDict']]] = None,
                 custom_delete: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 custom_update: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 triggers: Optional[pulumi.Input[Sequence[Any]]] = None,
                 __props__=None):
        """
        TEE(1)                           User Commands                          TEE(1)

        NAME
               tee - read from standard input and write to standard output and files

        SYNOPSIS
               tee [OPTION]... [FILE]...

        DESCRIPTION
               Copy standard input to each FILE, and also to standard output.

               -a, --append
                      append to the given FILEs, do not overwrite
            
               -i, --ignore-interrupts
                      ignore interrupt signals
            
               -p     operate in a more appropriate MODE with pipes.
            
               --output-error[=MODE]
                      set behavior on write error.  See MODE below
            
               --help display this help and exit
            
               --version
                      output version information and exit

           MODE determines behavior with write errors on the outputs:
               warn   diagnose errors writing to any output

               warn-nopipe
                      diagnose errors writing to any output not a pipe
            
               exit   exit on error writing to any output
            
               exit-nopipe
                      exit on error writing to any output not a pipe
            
               The  default  MODE  for  the -p option is 'warn-nopipe'.  With "nopipe"
               MODEs, exit immediately if all outputs become broken  pipes.   The  de‐
               fault  operation when --output-error is not specified, is to exit imme‐
               diately on error writing to a pipe, and diagnose errors writing to  non
               pipe outputs.

        AUTHOR
               Written by Mike Parker, Richard M. Stallman, and David MacKenzie.

        REPORTING BUGS
               GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
               Report any translation bugs to <https://translationproject.org/team/>

        COPYRIGHT
               Copyright  ©  2024  Free Software Foundation, Inc.  License GPLv3+: GNU
               GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
               This is free software: you are free  to  change  and  redistribute  it.
               There is NO WARRANTY, to the extent permitted by law.

        SEE ALSO
               Full documentation <https://www.gnu.org/software/coreutils/tee>
               or available locally via: info '(coreutils) tee invocation'

        GNU coreutils 9.5                 March 2024                            TEE(1)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TeeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        TEE(1)                           User Commands                          TEE(1)

        NAME
               tee - read from standard input and write to standard output and files

        SYNOPSIS
               tee [OPTION]... [FILE]...

        DESCRIPTION
               Copy standard input to each FILE, and also to standard output.

               -a, --append
                      append to the given FILEs, do not overwrite
            
               -i, --ignore-interrupts
                      ignore interrupt signals
            
               -p     operate in a more appropriate MODE with pipes.
            
               --output-error[=MODE]
                      set behavior on write error.  See MODE below
            
               --help display this help and exit
            
               --version
                      output version information and exit

           MODE determines behavior with write errors on the outputs:
               warn   diagnose errors writing to any output

               warn-nopipe
                      diagnose errors writing to any output not a pipe
            
               exit   exit on error writing to any output
            
               exit-nopipe
                      exit on error writing to any output not a pipe
            
               The  default  MODE  for  the -p option is 'warn-nopipe'.  With "nopipe"
               MODEs, exit immediately if all outputs become broken  pipes.   The  de‐
               fault  operation when --output-error is not specified, is to exit imme‐
               diately on error writing to a pipe, and diagnose errors writing to  non
               pipe outputs.

        AUTHOR
               Written by Mike Parker, Richard M. Stallman, and David MacKenzie.

        REPORTING BUGS
               GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
               Report any translation bugs to <https://translationproject.org/team/>

        COPYRIGHT
               Copyright  ©  2024  Free Software Foundation, Inc.  License GPLv3+: GNU
               GPL version 3 or later <https://gnu.org/licenses/gpl.html>.
               This is free software: you are free  to  change  and  redistribute  it.
               There is NO WARRANTY, to the extent permitted by law.

        SEE ALSO
               Full documentation <https://www.gnu.org/software/coreutils/tee>
               or available locally via: info '(coreutils) tee invocation'

        GNU coreutils 9.5                 March 2024                            TEE(1)

        :param str resource_name: The name of the resource.
        :param TeeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 args: Optional[pulumi.Input[Union['TeeArgsArgs', 'TeeArgsArgsDict']]] = None,
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
            __props__ = TeeArgs.__new__(TeeArgs)

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
        super(Tee, __self__).__init__(
            'baremetal:coreutils:Tee',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Tee':
        """
        Get an existing Tee resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TeeArgs.__new__(TeeArgs)

        __props__.__dict__["args"] = None
        __props__.__dict__["created_files"] = None
        __props__.__dict__["custom_delete"] = None
        __props__.__dict__["custom_update"] = None
        __props__.__dict__["exit_code"] = None
        __props__.__dict__["moved_files"] = None
        __props__.__dict__["stderr"] = None
        __props__.__dict__["stdout"] = None
        __props__.__dict__["triggers"] = None
        return Tee(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output['outputs.TeeArgs']:
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
