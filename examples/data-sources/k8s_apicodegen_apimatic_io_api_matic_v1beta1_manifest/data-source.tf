data "k8s_apicodegen_apimatic_io_api_matic_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    licensespec = {
      license_source_name = "apimaticlicense"
      license_source_type = "ConfigMap"
    }
    podspec = {
      apimatic_container_spec = {
        image             = "apimaticio/apimatic-codegen"
        image_pull_secret = "apimaticimagesecret"
      }
    }
    replicas = 3
    servicespec = {
      apimaticserviceport = {
        node_port = 32000
        port      = 8070
      }
      servicetype = "NodePort"
    }
  }
}
