// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export namespace cmd {
    export interface MvArgs {
        backup?: string;
        destination?: string;
        directory?: string;
        force?: boolean;
        help?: boolean;
        noClobber?: boolean;
        noTargetDirectory?: boolean;
        source: string[];
        stripTrailingSlashes?: boolean;
        suffix?: string;
        targetDirectory?: string;
        update?: boolean;
        verbose?: boolean;
        version?: boolean;
    }

    export interface RmArgs {
        dir?: boolean;
        files: string[];
        force?: boolean;
        help?: boolean;
        oneFileSystem?: boolean;
        recursive?: boolean;
        verbose?: boolean;
    }

    export interface TarArgs {
        append?: boolean;
        args?: string[];
        bzip2?: boolean;
        create?: boolean;
        delete?: boolean;
        diff?: boolean;
        directory?: string;
        exclude?: string;
        excludeFrom?: string;
        excludeVcs?: boolean;
        excludeVcsIgnores?: boolean;
        extract?: boolean;
        file?: string;
        gzip?: boolean;
        ignoreCommandError?: boolean;
        keepDirectorySymlink?: boolean;
        keepNewerFiles?: boolean;
        keepOldfiles?: boolean;
        list?: boolean;
        lzip?: boolean;
        lzma?: boolean;
        lzop?: boolean;
        noOverwriteDir?: boolean;
        noSeek?: boolean;
        overwrite?: boolean;
        overwriteDir?: boolean;
        removeFiles?: boolean;
        skipOldFiles?: boolean;
        sparse?: boolean;
        stripComponents?: number;
        suffix?: string;
        toStdout?: boolean;
        transform?: string;
        unlinkFirst?: boolean;
        update?: boolean;
        verbose?: boolean;
        verify?: boolean;
        xz?: boolean;
        zstd?: boolean;
    }

    export interface TeeArgs {
        append?: boolean;
        content: string;
        files: string[];
    }

    export interface WgetArgs {
        appendOutput?: string;
        background?: boolean;
        base?: string;
        caCertificateFile?: string;
        caDirectory?: string;
        certificate?: string;
        certificateType?: string;
        config?: string;
        continue?: boolean;
        crlFile?: string;
        cutDirs?: number;
        debug?: boolean;
        directoryPrefix?: string;
        execute?: string[];
        forceDirectories?: boolean;
        forceHtml?: boolean;
        help?: boolean;
        httpsOnly?: boolean;
        inet4Only?: boolean;
        inputFile?: string;
        keepSessionCookies?: boolean;
        noClobber?: boolean;
        noDirectories?: boolean;
        noDnsCache?: boolean;
        noVerbose?: boolean;
        outputDocument?: string;
        outputFile?: string;
        password?: string;
        privateKey?: string;
        privateKeyType?: string;
        progress?: string;
        quiet?: boolean;
        randomWait?: boolean;
        reportSpeed?: string;
        saveCookies?: string;
        showProgress?: boolean;
        startPos?: string;
        timeout?: string;
        timestamping?: boolean;
        tries?: number;
        urls: string[];
        user?: string;
        userAgent?: string;
        verbose?: boolean;
        version?: string;
        wait?: string;
    }

}
