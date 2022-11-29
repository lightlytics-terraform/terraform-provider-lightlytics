## Lightlytics Terraform Provider



The Lightlytics Terraform provider is used to connect your cloud account to Lightlytics and integrate with various available features.

In order to use this provider, you must have an active account with [Lightlytics](https://www.lightlytics.com).

You can [start free](https://www.lightlytics.com/treemium) or check out our [plans](https://www.lightlytics.com/plans) and [contact us](https://www.lightlytics.com/contact-us) or [book a demo](https://www.lightlytics.com/book-demo).


## Building the provider
1. Clone [this](terraform-provider-lightlytics) repository
2. Navigate to the provider directory
3. Install the provider by running the following command:
```
make install
```

## Inputs
| Variable Name                     | Description                                               | Notes                              | Type           | Required? | Default |
| :-------------------------------- | :-------------------------------------------------------  | :----------------------------------|:---------------|:--------- |:--------|
| host                              | Your environment URL including https://                   | e.g `https://org.lightlytics.com`  | `string`       | Yes       | n/a     |
| username                          | Your Lightlytics user Email                               |                                    | `string`       | Yes       | n/a     |
| password                          | Your Lightlytics user password                            |                                    | `string`       | Yes       | n/a     |                                                                              |
| aws_account_id                    | Your AWS account ID                                       |                       			 | `string`       | Yes       | n/a     |                                              
| display_name                      | Your integration display name within Lightlytics platform |                                    | `string`       | Yes       | n/a     |
| aws_regions                       | Desired regions to be scanned                             |                                    | `list(string)` | Yes       | n/a     |
| stack_region                      | The main region for the IAM Role to be deployed           |                                    | `string`       | Yes       | n/a     |



## Usage

Configure provider credentials and host

```
terraform {
  required_providers {
    lightlytics = {
      version = "0.2"
      source  = "lightlytics.com/api/lightlytics"
    }
  }
}

provider "lightlytics" {
  host = "https://<env_name>.lightlytics.com"
  username = ""
  password = ""
}
```

Configure AWS account


```
resource "lightlytics_account" "aws" {
  account_type = "AWS"
  aws_account_id = "123234818678"
  display_name = "test-user"
  aws_regions = ["us-east-1", "us-east-2"]
  stack_region = "us-east-1"
}
```

Find more examples in `/examples` 
