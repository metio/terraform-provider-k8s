terraform {
  required_providers {
    # define local alias banzaicloud provider
    banzaicloud-k8s = {
      source  = "banzaicloud/k8s"
      version = "0.9.1"
    }
    # define another alias for metio provider
    metio-k8s = {
      source  = "metio/k8s"
      version = "> 2022.10.15"
    }
  }
}

provider "banzaicloud-k8s" {
  # Configuration options for banzaicloud's provider
}
provider "metio-k8s" {
  # we need no configuration
}

# declare any resource from the k8s provider
resource "k8s_monitoring_coreos_com_pod_monitor_v1" "example" {
  # be explicit about the provider to use
  provider = metio-k8s

  metadata = {
    name = "example"
  }
  spec = {
    pod_metrics_endpoints = [
      {
        path = "/metrics"
        port = "metrics"
      }
    ]
    selector = {
      match_labels = {
        "app.kubernetes.io/name" = "some-name"
      }
    }
  }
}

# use the 'yaml' attribute as input for the k8s provider by banzaicloud
resource "k8s_manifest" "example" {
  # be explicit about the provider to use
  provider = banzaicloud-k8s

  content = k8s_monitoring_coreos_com_pod_monitor_v1.example.yaml
}
