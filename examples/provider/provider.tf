provider "k8s" {
  # uses the 'current-context' from your KUBECONFIG
}

# terraform configuration overrides environment variables
provider "k8s" {
  kubeconfig    = "path/to/kube/config"
  context       = "some-context"
  field_manager = "bill-lumbergh"
  timeout       = 300
}

# do not connect to a kubernetes cluster
provider "k8s" {
  offline = true
}
