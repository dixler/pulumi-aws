# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Pipeline(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The codepipeline ARN.
    """
    artifact_store: pulumi.Output[dict]
    """
    An artifact_store block. Artifact stores are documented below.
    * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
    """
    name: pulumi.Output[str]
    """
    The name of the pipeline.
    """
    role_arn: pulumi.Output[str]
    """
    A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
    """
    stages: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, artifact_store=None, name=None, role_arn=None, stages=None, __name__=None, __opts__=None):
        """
        Provides a CodePipeline.
        
        > **NOTE on `aws_codepipeline`:** - the `GITHUB_TOKEN` environment variable must be set if the GitHub provider is specified.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] artifact_store: An artifact_store block. Artifact stores are documented below.
               * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
        :param pulumi.Input[str] name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        :param pulumi.Input[list] stages
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if artifact_store is None:
            raise TypeError('Missing required property artifact_store')
        __props__['artifact_store'] = artifact_store

        __props__['name'] = name

        if role_arn is None:
            raise TypeError('Missing required property role_arn')
        __props__['role_arn'] = role_arn

        if stages is None:
            raise TypeError('Missing required property stages')
        __props__['stages'] = stages

        __props__['arn'] = None

        super(Pipeline, __self__).__init__(
            'aws:codepipeline/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

