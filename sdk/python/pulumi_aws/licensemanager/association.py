# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Association(pulumi.CustomResource):
    license_configuration_arn: pulumi.Output[str]
    """
    ARN of the license configuration.
    """
    resource_arn: pulumi.Output[str]
    """
    ARN of the resource associated with the license configuration.
    """
    def __init__(__self__, resource_name, opts=None, license_configuration_arn=None, resource_arn=None, __name__=None, __opts__=None):
        """
        Provides a License Manager association.
        
        > **Note:** License configurations can also be associated with launch templates by specifying the `license_specifications` block for an `aws_launch_template`.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] license_configuration_arn: ARN of the license configuration.
        :param pulumi.Input[str] resource_arn: ARN of the resource associated with the license configuration.
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

        if license_configuration_arn is None:
            raise TypeError('Missing required property license_configuration_arn')
        __props__['license_configuration_arn'] = license_configuration_arn

        if resource_arn is None:
            raise TypeError('Missing required property resource_arn')
        __props__['resource_arn'] = resource_arn

        super(Association, __self__).__init__(
            'aws:licensemanager/association:Association',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

