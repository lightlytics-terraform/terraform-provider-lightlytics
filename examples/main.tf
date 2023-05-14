terraform {
  required_providers {
    lightlytics = {
      version = "0.2"
      source  = "lightlytics.com/api/lightlytics"
    }
  }
}

provider "lightlytics" {
  alias = "test-integration"
  host = "https://<env_name>.lightlytics.com"
  username = "<username>"
  password = "<password>"
  workspace_id = ""
}

resource "lightlytics_account" "aws" {
  provider = lightlytics.test-integration
  account_type = "AWS"
  cloud_account_id = "123234818678"
  display_name = "test-user"
  cloud_regions = ["us-east-1", "us-west-1", "us-east-2"]
  stack_region = "us-east-1"
}