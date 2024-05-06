data "k8s_config_karmada_io_resource_interpreter_customization_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    customizations = {}
    target = {
      api_version = "v1"
      kind        = "Service"
      name        = "some-service"
    }
  }
}
