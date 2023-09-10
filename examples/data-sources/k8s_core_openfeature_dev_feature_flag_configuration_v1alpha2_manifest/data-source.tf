data "k8s_core_openfeature_dev_feature_flag_configuration_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
