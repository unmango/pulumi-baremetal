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
                 stdin: pulumi.Input[str],
                 create: Optional[pulumi.Input['TeeOptsArgs']] = None):
        """
        The set of arguments for constructing a Tee resource.
        """
        pulumi.set(__self__, "stdin", stdin)
        if create is not None:
            pulumi.set(__self__, "create", create)

    @property
    @pulumi.getter
    def stdin(self) -> pulumi.Input[str]:
        return pulumi.get(self, "stdin")

    @stdin.setter
    def stdin(self, value: pulumi.Input[str]):
        pulumi.set(self, "stdin", value)

    @property
    @pulumi.getter
    def create(self) -> Optional[pulumi.Input['TeeOptsArgs']]:
        return pulumi.get(self, "create")

    @create.setter
    def create(self, value: Optional[pulumi.Input['TeeOptsArgs']]):
        pulumi.set(self, "create", value)


class Tee(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 create: Optional[pulumi.Input[Union['TeeOptsArgs', 'TeeOptsArgsDict']]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
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
                 create: Optional[pulumi.Input[Union['TeeOptsArgs', 'TeeOptsArgsDict']]] = None,
                 stdin: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeeArgs.__new__(TeeArgs)

            __props__.__dict__["create"] = create
            if stdin is None and not opts.urn:
                raise TypeError("Missing required property 'stdin'")
            __props__.__dict__["stdin"] = stdin
            __props__.__dict__["created_files"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Tee, __self__).__init__(
            'baremetal:cmd:Tee',
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

        __props__.__dict__["create"] = None
        __props__.__dict__["created_files"] = None
        __props__.__dict__["stderr"] = None
        __props__.__dict__["stdin"] = None
        __props__.__dict__["stdout"] = None
        return Tee(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def create(self) -> pulumi.Output[Optional['outputs.TeeOpts']]:
        return pulumi.get(self, "create")

    @property
    @pulumi.getter
    def created_files(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "created_files")

    @property
    @pulumi.getter
    def stderr(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stderr")

    @property
    @pulumi.getter
    def stdin(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stdin")

    @property
    @pulumi.getter
    def stdout(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stdout")

