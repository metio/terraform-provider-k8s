data "k8s_registry_apicur_io_apicurio_registry_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
