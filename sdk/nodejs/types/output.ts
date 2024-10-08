// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ProvisionerConnection {
    address: string;
    caPem?: string;
    certPem?: string;
    keyPem?: string;
    port?: string;
}

export namespace coreutils {
    export interface CatArgs {
        e?: boolean;
        files: string[];
        help?: boolean;
        number?: boolean;
        numberNonblank?: boolean;
        showAll?: boolean;
        showEnds?: boolean;
        showNonprinting?: boolean;
        showTabs?: boolean;
        squeezeBlank?: boolean;
        t?: boolean;
        version?: boolean;
    }

    export interface ChmodArgs {
        changes?: boolean;
        files: string[];
        help?: boolean;
        mode?: string[];
        noPreserveRoot?: boolean;
        octalMode?: string;
        preserveRoot?: boolean;
        quiet?: boolean;
        recursive?: boolean;
        reference?: string;
        verbose?: boolean;
        version?: boolean;
    }

    export interface MkdirArgs {
        directory: string[];
        help?: boolean;
        mode?: string;
        parents?: boolean;
        verbose?: boolean;
        version?: boolean;
    }

    export interface MktempArgs {
        directory?: boolean;
        dryRun?: boolean;
        help?: boolean;
        p?: string;
        quiet?: boolean;
        suffix?: string;
        t?: boolean;
        template?: string;
        tmpdir?: boolean;
        version?: boolean;
    }

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
        anchored?: boolean;
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
        noAnchored?: boolean;
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
        version?: boolean;
        xz?: boolean;
        zstd?: boolean;
    }

    export interface TeeArgs {
        append?: boolean;
        files: string[];
        stdin?: string;
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

export namespace kubeadm {
    export interface KubeadmArgs {
        commands: string[];
    }

}
