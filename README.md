## Lightlytics Terraform Provider



The Lightlytics Terraform provider is used to connect your cloud account to Lightlytics and integrate with various available features.

In order to use this provider, you must have an active account with [Lightlytics](https://www.lightlytics.com).

You can [start free](https://www.lightlytics.com/treemium) or check out our [plans](https://www.lightlytics.com/plans) and [contact us](https://www.lightlytics.com/contact-us) or [book a demo](https://www.lightlytics.com/book-demo).


## Requirements
- A Lightlytics account
- Credentials to Lightlytics platform (Email & Password)


## Building the provider
1. Clone the [terraform-provider-lightlytics](https://github.com/lightlytics-terraform/terraform-provider-lightlytics) repository
2. Navigate to the provider directory
3. Install the provider by running the following command:
```
make install
```


## Usage
```hcl
# Configure Lightlytics provider host and credentials

terraform {
  required_providers {
    lightlytics  = {
      version    = "0.3"
      source     = "lightlytics.com/api/lightlytics"
    }
  }
}

provider "lightlytics" {
  host           = "<https://<env_name>.lightlytics.com>"
  username       = "<Your_Lightlytics_Login_Email>"
  password       = "<Your_Lightlytics_Login_Password>"
  workspace_id   = "<Your_Lightlytics_Workspace_ID>"  ## Can be obtained from Lightlytics platform, if not specified, it will use the first WS
}

# Configure AWS account

resource "lightlytics_account" "aws" {
  account_type   = "AWS"
  aws_account_id = "<Your_AWS_Account_ID>"
  display_name   = "<Your_Desired_Lightlytics_Integration_Display_Name>"
  stack_region   = "us-east-1"
  aws_regions    = ["us-east-1", "us-east-2"]
}

resource "lightlytics_kubernetes_account" "k8s" {
  display_name   = "<Your_Desired_Lightlytics_Kubernetes_Integration_Display_Name>"
  eks_arn   = "<Your_EKS_ARN>"
}
```


## Inputs
| Variable Name                     | Description                                                                | Notes                                               | Type           | Required? | Default |
| :-------------------------------- | :------------------------------------------------------------------------- | :-------------------------------------------------- |:---------------|:--------- |:--------|
| host                              | Your environment URL including https://                                    | e.g `https://org.lightlytics.com`                   | `string`       | Yes       | n/a     |
| username                          | Your Lightlytics user Email                                                |                                                     | `string`       | Yes       | n/a     |
| password                          | Your Lightlytics user password                                             |                                                     | `string`       | Yes       | n/a     |
| workspace_id                      | Can be obtained from Lightlytics platform                                  | Will use default workspace in case not specified    | `string`       | No        | n/a     |
| aws_account_id                    | Your AWS account ID                                                        |                       			                   | `string`       | Yes       | n/a     |
| display_name                      | Your integration display name within Lightlytics platform                  |                                                     | `string`       | Yes       | n/a     |
| stack_region                      | The primary region where Lightlytics read access resources will be created |                                                     | `string`       | Yes       | n/a     |
| aws_regions                       | List of desired regions to be scanned                                      | us-east-1 region is mandatory for the integration   | `list(string)` | Yes       | n/a     | 



Documentation
-------------
If you're new to Lightlytics and want to get started, feel free to [contact us](https://www.lightlytics.com/contact-us) or checkout our [documentation](https://docs.lightlytics.com/) website.


Community
---------
- Join Lightlytics community on [Slack](https://join.slack.com/t/lightlyticscommunity/shared_invite/zt-1f7dk2yo7-xBTOU_o4tOnAjoFxfHVF8Q)


Getting Help
------------
Please use these resources for getting help:
- [Slack](https://join.slack.com/t/lightlyticscommunity/shared_invite/zt-1f7dk2yo7-xBTOU_o4tOnAjoFxfHVF8Q)
- Email: support@lightlytics.com
