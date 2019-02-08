# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetBucketResult(object):
    """
    A collection of values returned by getBucket.
    """
    def __init__(__self__, arn=None, bucket_domain_name=None, hosted_zone_id=None, region=None, website_domain=None, website_endpoint=None, id=None):
        if arn and not isinstance(arn, str):
            raise TypeError('Expected argument arn to be a str')
        __self__.arn = arn
        """
        The ARN of the bucket. Will be of format `arn:aws:s3:::bucketname`.
        """
        if bucket_domain_name and not isinstance(bucket_domain_name, str):
            raise TypeError('Expected argument bucket_domain_name to be a str')
        __self__.bucket_domain_name = bucket_domain_name
        """
        The bucket domain name. Will be of format `bucketname.s3.amazonaws.com`.
        """
        if hosted_zone_id and not isinstance(hosted_zone_id, str):
            raise TypeError('Expected argument hosted_zone_id to be a str')
        __self__.hosted_zone_id = hosted_zone_id
        """
        The [Route 53 Hosted Zone ID](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_website_region_endpoints) for this bucket's region.
        """
        if region and not isinstance(region, str):
            raise TypeError('Expected argument region to be a str')
        __self__.region = region
        """
        The AWS region this bucket resides in.
        """
        if website_domain and not isinstance(website_domain, str):
            raise TypeError('Expected argument website_domain to be a str')
        __self__.website_domain = website_domain
        """
        The domain of the website endpoint, if the bucket is configured with a website. If not, this will be an empty string. This is used to create Route 53 alias records.
        """
        if website_endpoint and not isinstance(website_endpoint, str):
            raise TypeError('Expected argument website_endpoint to be a str')
        __self__.website_endpoint = website_endpoint
        """
        The website endpoint, if the bucket is configured with a website. If not, this will be an empty string.
        """
        if id and not isinstance(id, str):
            raise TypeError('Expected argument id to be a str')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

async def get_bucket(bucket=None):
    """
    Provides details about a specific S3 bucket.
    
    This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
    Distribution.
    """
    __args__ = dict()

    __args__['bucket'] = bucket
    __ret__ = await pulumi.runtime.invoke('aws:s3/getBucket:getBucket', __args__)

    return GetBucketResult(
        arn=__ret__.get('arn'),
        bucket_domain_name=__ret__.get('bucketDomainName'),
        hosted_zone_id=__ret__.get('hostedZoneId'),
        region=__ret__.get('region'),
        website_domain=__ret__.get('websiteDomain'),
        website_endpoint=__ret__.get('websiteEndpoint'),
        id=__ret__.get('id'))
