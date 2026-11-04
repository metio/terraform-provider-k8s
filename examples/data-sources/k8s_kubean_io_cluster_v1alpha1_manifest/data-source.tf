data "k8s_kubean_io_cluster_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }

  spec = {
    hosts_conf_ref = {
      name      = "some-value"
      namespace = "some-value"
    }
    vars_conf_ref = {
      name      = "some-value"
      namespace = "some-value"
    }
  }
}
