data "k8s_core_openfeature_dev_feature_flag_configuration_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
