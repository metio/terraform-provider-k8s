terraform {
  required_providers {
    k8s = {
      source  = "localhost/metio/k8s"
      version = "9999.99.99"
    }
  }
}

provider "k8s" {}
