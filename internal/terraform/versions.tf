terraform {
  required_version = ">=1"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=5"
    }
    azure = {
      source  = "hashicorp/azurerm"
      version = ">=3"
    }
    google = {
      source  = "hashicorp/google"
      version = ">=4"
    }
  }
}
