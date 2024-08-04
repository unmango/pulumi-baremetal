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

__all__ = ['WgetArgs', 'Wget']

@pulumi.input_type
class WgetArgs:
    def __init__(__self__, *,
                 urls: pulumi.Input[Sequence[pulumi.Input[str]]],
                 append_output: Optional[pulumi.Input[str]] = None,
                 background: Optional[pulumi.Input[bool]] = None,
                 base: Optional[pulumi.Input[str]] = None,
                 ca_certificate_file: Optional[pulumi.Input[str]] = None,
                 ca_directory: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_type: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 continue_: Optional[pulumi.Input[bool]] = None,
                 crl_file: Optional[pulumi.Input[str]] = None,
                 cut_dirs: Optional[pulumi.Input[int]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 execute: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force_directories: Optional[pulumi.Input[bool]] = None,
                 force_html: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 inet4_only: Optional[pulumi.Input[bool]] = None,
                 input_file: Optional[pulumi.Input[str]] = None,
                 keep_session_cookies: Optional[pulumi.Input[bool]] = None,
                 no_clobber: Optional[pulumi.Input[bool]] = None,
                 no_directories: Optional[pulumi.Input[bool]] = None,
                 no_dns_cache: Optional[pulumi.Input[bool]] = None,
                 no_verbose: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 output_file: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_type: Optional[pulumi.Input[str]] = None,
                 progress: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 random_wait: Optional[pulumi.Input[bool]] = None,
                 report_speed: Optional[pulumi.Input[str]] = None,
                 save_cookies: Optional[pulumi.Input[str]] = None,
                 show_progress: Optional[pulumi.Input[bool]] = None,
                 start_pos: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None,
                 tries: Optional[pulumi.Input[int]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 wait: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Wget resource.
        """
        pulumi.set(__self__, "urls", urls)
        if append_output is not None:
            pulumi.set(__self__, "append_output", append_output)
        if background is not None:
            pulumi.set(__self__, "background", background)
        if base is not None:
            pulumi.set(__self__, "base", base)
        if ca_certificate_file is not None:
            pulumi.set(__self__, "ca_certificate_file", ca_certificate_file)
        if ca_directory is not None:
            pulumi.set(__self__, "ca_directory", ca_directory)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if certificate_type is not None:
            pulumi.set(__self__, "certificate_type", certificate_type)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if continue_ is not None:
            pulumi.set(__self__, "continue_", continue_)
        if crl_file is not None:
            pulumi.set(__self__, "crl_file", crl_file)
        if cut_dirs is not None:
            pulumi.set(__self__, "cut_dirs", cut_dirs)
        if debug is not None:
            pulumi.set(__self__, "debug", debug)
        if directory_prefix is not None:
            pulumi.set(__self__, "directory_prefix", directory_prefix)
        if execute is not None:
            pulumi.set(__self__, "execute", execute)
        if force_directories is not None:
            pulumi.set(__self__, "force_directories", force_directories)
        if force_html is not None:
            pulumi.set(__self__, "force_html", force_html)
        if help is not None:
            pulumi.set(__self__, "help", help)
        if https_only is not None:
            pulumi.set(__self__, "https_only", https_only)
        if inet4_only is not None:
            pulumi.set(__self__, "inet4_only", inet4_only)
        if input_file is not None:
            pulumi.set(__self__, "input_file", input_file)
        if keep_session_cookies is not None:
            pulumi.set(__self__, "keep_session_cookies", keep_session_cookies)
        if no_clobber is not None:
            pulumi.set(__self__, "no_clobber", no_clobber)
        if no_directories is not None:
            pulumi.set(__self__, "no_directories", no_directories)
        if no_dns_cache is not None:
            pulumi.set(__self__, "no_dns_cache", no_dns_cache)
        if no_verbose is not None:
            pulumi.set(__self__, "no_verbose", no_verbose)
        if output_document is not None:
            pulumi.set(__self__, "output_document", output_document)
        if output_file is not None:
            pulumi.set(__self__, "output_file", output_file)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if private_key_type is not None:
            pulumi.set(__self__, "private_key_type", private_key_type)
        if progress is not None:
            pulumi.set(__self__, "progress", progress)
        if quiet is not None:
            pulumi.set(__self__, "quiet", quiet)
        if random_wait is not None:
            pulumi.set(__self__, "random_wait", random_wait)
        if report_speed is not None:
            pulumi.set(__self__, "report_speed", report_speed)
        if save_cookies is not None:
            pulumi.set(__self__, "save_cookies", save_cookies)
        if show_progress is not None:
            pulumi.set(__self__, "show_progress", show_progress)
        if start_pos is not None:
            pulumi.set(__self__, "start_pos", start_pos)
        if timeout is not None:
            pulumi.set(__self__, "timeout", timeout)
        if timestamping is not None:
            pulumi.set(__self__, "timestamping", timestamping)
        if tries is not None:
            pulumi.set(__self__, "tries", tries)
        if user is not None:
            pulumi.set(__self__, "user", user)
        if user_agent is not None:
            pulumi.set(__self__, "user_agent", user_agent)
        if verbose is not None:
            pulumi.set(__self__, "verbose", verbose)
        if version is not None:
            pulumi.set(__self__, "version", version)
        if wait is not None:
            pulumi.set(__self__, "wait", wait)

    @property
    @pulumi.getter
    def urls(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "urls")

    @urls.setter
    def urls(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "urls", value)

    @property
    @pulumi.getter(name="appendOutput")
    def append_output(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "append_output")

    @append_output.setter
    def append_output(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "append_output", value)

    @property
    @pulumi.getter
    def background(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "background")

    @background.setter
    def background(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "background", value)

    @property
    @pulumi.getter
    def base(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "base")

    @base.setter
    def base(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base", value)

    @property
    @pulumi.getter(name="caCertificateFile")
    def ca_certificate_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ca_certificate_file")

    @ca_certificate_file.setter
    def ca_certificate_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_certificate_file", value)

    @property
    @pulumi.getter(name="caDirectory")
    def ca_directory(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "ca_directory")

    @ca_directory.setter
    def ca_directory(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_directory", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="certificateType")
    def certificate_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "certificate_type")

    @certificate_type.setter
    def certificate_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_type", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="continue")
    def continue_(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "continue_")

    @continue_.setter
    def continue_(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "continue_", value)

    @property
    @pulumi.getter(name="crlFile")
    def crl_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "crl_file")

    @crl_file.setter
    def crl_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crl_file", value)

    @property
    @pulumi.getter(name="cutDirs")
    def cut_dirs(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "cut_dirs")

    @cut_dirs.setter
    def cut_dirs(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cut_dirs", value)

    @property
    @pulumi.getter
    def debug(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "debug")

    @debug.setter
    def debug(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "debug", value)

    @property
    @pulumi.getter(name="directoryPrefix")
    def directory_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "directory_prefix")

    @directory_prefix.setter
    def directory_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "directory_prefix", value)

    @property
    @pulumi.getter
    def execute(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "execute")

    @execute.setter
    def execute(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "execute", value)

    @property
    @pulumi.getter(name="forceDirectories")
    def force_directories(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "force_directories")

    @force_directories.setter
    def force_directories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_directories", value)

    @property
    @pulumi.getter(name="forceHtml")
    def force_html(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "force_html")

    @force_html.setter
    def force_html(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "force_html", value)

    @property
    @pulumi.getter
    def help(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "help")

    @help.setter
    def help(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "help", value)

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "https_only")

    @https_only.setter
    def https_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https_only", value)

    @property
    @pulumi.getter(name="inet4Only")
    def inet4_only(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "inet4_only")

    @inet4_only.setter
    def inet4_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "inet4_only", value)

    @property
    @pulumi.getter(name="inputFile")
    def input_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "input_file")

    @input_file.setter
    def input_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "input_file", value)

    @property
    @pulumi.getter(name="keepSessionCookies")
    def keep_session_cookies(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "keep_session_cookies")

    @keep_session_cookies.setter
    def keep_session_cookies(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_session_cookies", value)

    @property
    @pulumi.getter(name="noClobber")
    def no_clobber(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_clobber")

    @no_clobber.setter
    def no_clobber(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_clobber", value)

    @property
    @pulumi.getter(name="noDirectories")
    def no_directories(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_directories")

    @no_directories.setter
    def no_directories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_directories", value)

    @property
    @pulumi.getter(name="noDnsCache")
    def no_dns_cache(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_dns_cache")

    @no_dns_cache.setter
    def no_dns_cache(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_dns_cache", value)

    @property
    @pulumi.getter(name="noVerbose")
    def no_verbose(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "no_verbose")

    @no_verbose.setter
    def no_verbose(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_verbose", value)

    @property
    @pulumi.getter(name="outputDocument")
    def output_document(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "output_document")

    @output_document.setter
    def output_document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_document", value)

    @property
    @pulumi.getter(name="outputFile")
    def output_file(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "output_file")

    @output_file.setter
    def output_file(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "output_file", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="privateKeyType")
    def private_key_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "private_key_type")

    @private_key_type.setter
    def private_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key_type", value)

    @property
    @pulumi.getter
    def progress(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "progress")

    @progress.setter
    def progress(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "progress", value)

    @property
    @pulumi.getter
    def quiet(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "quiet")

    @quiet.setter
    def quiet(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "quiet", value)

    @property
    @pulumi.getter(name="randomWait")
    def random_wait(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "random_wait")

    @random_wait.setter
    def random_wait(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "random_wait", value)

    @property
    @pulumi.getter(name="reportSpeed")
    def report_speed(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "report_speed")

    @report_speed.setter
    def report_speed(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "report_speed", value)

    @property
    @pulumi.getter(name="saveCookies")
    def save_cookies(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "save_cookies")

    @save_cookies.setter
    def save_cookies(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "save_cookies", value)

    @property
    @pulumi.getter(name="showProgress")
    def show_progress(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "show_progress")

    @show_progress.setter
    def show_progress(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "show_progress", value)

    @property
    @pulumi.getter(name="startPos")
    def start_pos(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "start_pos")

    @start_pos.setter
    def start_pos(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_pos", value)

    @property
    @pulumi.getter
    def timeout(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "timeout")

    @timeout.setter
    def timeout(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timeout", value)

    @property
    @pulumi.getter
    def timestamping(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "timestamping")

    @timestamping.setter
    def timestamping(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "timestamping", value)

    @property
    @pulumi.getter
    def tries(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "tries")

    @tries.setter
    def tries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tries", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter(name="userAgent")
    def user_agent(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_agent")

    @user_agent.setter
    def user_agent(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_agent", value)

    @property
    @pulumi.getter
    def verbose(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "verbose")

    @verbose.setter
    def verbose(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verbose", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def wait(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "wait")

    @wait.setter
    def wait(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wait", value)


class Wget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 append_output: Optional[pulumi.Input[str]] = None,
                 background: Optional[pulumi.Input[bool]] = None,
                 base: Optional[pulumi.Input[str]] = None,
                 ca_certificate_file: Optional[pulumi.Input[str]] = None,
                 ca_directory: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_type: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 continue_: Optional[pulumi.Input[bool]] = None,
                 crl_file: Optional[pulumi.Input[str]] = None,
                 cut_dirs: Optional[pulumi.Input[int]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 execute: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force_directories: Optional[pulumi.Input[bool]] = None,
                 force_html: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 inet4_only: Optional[pulumi.Input[bool]] = None,
                 input_file: Optional[pulumi.Input[str]] = None,
                 keep_session_cookies: Optional[pulumi.Input[bool]] = None,
                 no_clobber: Optional[pulumi.Input[bool]] = None,
                 no_directories: Optional[pulumi.Input[bool]] = None,
                 no_dns_cache: Optional[pulumi.Input[bool]] = None,
                 no_verbose: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 output_file: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_type: Optional[pulumi.Input[str]] = None,
                 progress: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 random_wait: Optional[pulumi.Input[bool]] = None,
                 report_speed: Optional[pulumi.Input[str]] = None,
                 save_cookies: Optional[pulumi.Input[str]] = None,
                 show_progress: Optional[pulumi.Input[bool]] = None,
                 start_pos: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None,
                 tries: Optional[pulumi.Input[int]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 wait: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Wget resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WgetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Wget resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param WgetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WgetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 append_output: Optional[pulumi.Input[str]] = None,
                 background: Optional[pulumi.Input[bool]] = None,
                 base: Optional[pulumi.Input[str]] = None,
                 ca_certificate_file: Optional[pulumi.Input[str]] = None,
                 ca_directory: Optional[pulumi.Input[str]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_type: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 continue_: Optional[pulumi.Input[bool]] = None,
                 crl_file: Optional[pulumi.Input[str]] = None,
                 cut_dirs: Optional[pulumi.Input[int]] = None,
                 debug: Optional[pulumi.Input[bool]] = None,
                 directory_prefix: Optional[pulumi.Input[str]] = None,
                 execute: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 force_directories: Optional[pulumi.Input[bool]] = None,
                 force_html: Optional[pulumi.Input[bool]] = None,
                 help: Optional[pulumi.Input[bool]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 inet4_only: Optional[pulumi.Input[bool]] = None,
                 input_file: Optional[pulumi.Input[str]] = None,
                 keep_session_cookies: Optional[pulumi.Input[bool]] = None,
                 no_clobber: Optional[pulumi.Input[bool]] = None,
                 no_directories: Optional[pulumi.Input[bool]] = None,
                 no_dns_cache: Optional[pulumi.Input[bool]] = None,
                 no_verbose: Optional[pulumi.Input[bool]] = None,
                 output_document: Optional[pulumi.Input[str]] = None,
                 output_file: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 private_key_type: Optional[pulumi.Input[str]] = None,
                 progress: Optional[pulumi.Input[str]] = None,
                 quiet: Optional[pulumi.Input[bool]] = None,
                 random_wait: Optional[pulumi.Input[bool]] = None,
                 report_speed: Optional[pulumi.Input[str]] = None,
                 save_cookies: Optional[pulumi.Input[str]] = None,
                 show_progress: Optional[pulumi.Input[bool]] = None,
                 start_pos: Optional[pulumi.Input[str]] = None,
                 timeout: Optional[pulumi.Input[str]] = None,
                 timestamping: Optional[pulumi.Input[bool]] = None,
                 tries: Optional[pulumi.Input[int]] = None,
                 urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 user_agent: Optional[pulumi.Input[str]] = None,
                 verbose: Optional[pulumi.Input[bool]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 wait: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WgetArgs.__new__(WgetArgs)

            __props__.__dict__["append_output"] = append_output
            __props__.__dict__["background"] = background
            __props__.__dict__["base"] = base
            __props__.__dict__["ca_certificate_file"] = ca_certificate_file
            __props__.__dict__["ca_directory"] = ca_directory
            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["certificate_type"] = certificate_type
            __props__.__dict__["config"] = config
            __props__.__dict__["continue_"] = continue_
            __props__.__dict__["crl_file"] = crl_file
            __props__.__dict__["cut_dirs"] = cut_dirs
            __props__.__dict__["debug"] = debug
            __props__.__dict__["directory_prefix"] = directory_prefix
            __props__.__dict__["execute"] = execute
            __props__.__dict__["force_directories"] = force_directories
            __props__.__dict__["force_html"] = force_html
            __props__.__dict__["help"] = help
            __props__.__dict__["https_only"] = https_only
            __props__.__dict__["inet4_only"] = inet4_only
            __props__.__dict__["input_file"] = input_file
            __props__.__dict__["keep_session_cookies"] = keep_session_cookies
            __props__.__dict__["no_clobber"] = no_clobber
            __props__.__dict__["no_directories"] = no_directories
            __props__.__dict__["no_dns_cache"] = no_dns_cache
            __props__.__dict__["no_verbose"] = no_verbose
            __props__.__dict__["output_document"] = output_document
            __props__.__dict__["output_file"] = output_file
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["private_key"] = None if private_key is None else pulumi.Output.secret(private_key)
            __props__.__dict__["private_key_type"] = None if private_key_type is None else pulumi.Output.secret(private_key_type)
            __props__.__dict__["progress"] = progress
            __props__.__dict__["quiet"] = quiet
            __props__.__dict__["random_wait"] = random_wait
            __props__.__dict__["report_speed"] = report_speed
            __props__.__dict__["save_cookies"] = save_cookies
            __props__.__dict__["show_progress"] = show_progress
            __props__.__dict__["start_pos"] = start_pos
            __props__.__dict__["timeout"] = timeout
            __props__.__dict__["timestamping"] = timestamping
            __props__.__dict__["tries"] = tries
            if urls is None and not opts.urn:
                raise TypeError("Missing required property 'urls'")
            __props__.__dict__["urls"] = urls
            __props__.__dict__["user"] = user
            __props__.__dict__["user_agent"] = user_agent
            __props__.__dict__["verbose"] = verbose
            __props__.__dict__["version"] = version
            __props__.__dict__["wait"] = wait
            __props__.__dict__["args"] = None
            __props__.__dict__["created_files"] = None
            __props__.__dict__["exit_code"] = None
            __props__.__dict__["stderr"] = None
            __props__.__dict__["stdout"] = None
        super(Wget, __self__).__init__(
            'baremetal:cmd:Wget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Wget':
        """
        Get an existing Wget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WgetArgs.__new__(WgetArgs)

        __props__.__dict__["args"] = None
        __props__.__dict__["created_files"] = None
        __props__.__dict__["exit_code"] = None
        __props__.__dict__["stderr"] = None
        __props__.__dict__["stdout"] = None
        return Wget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def args(self) -> pulumi.Output['outputs.WgetArgs']:
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
