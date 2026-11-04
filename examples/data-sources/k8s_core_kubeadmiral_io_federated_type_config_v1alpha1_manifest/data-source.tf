data "k8s_core_kubeadmiral_io_federated_type_config_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"

  }

  spec = {
    source_type = {
      kind        = "some-value"
      plural_name = "some-value"
      scope       = "some-value"
      version     = "some-value"
    }
  }
}
