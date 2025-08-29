data "k8s_radapp_io_deployment_template_v1alpha3_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
