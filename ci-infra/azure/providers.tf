terraform {
  required_version = ">=0.12"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>3.0"
    }
  }

  backend "azurerm" {
    resource_group_name  = "test-ci-setup"
    storage_account_name = "caatestterraformstate"
    container_name       = "terraform-state"
    key                  = "ci.terraform.tfstate"
  }
}

provider "azurerm" {
  features {}
}
