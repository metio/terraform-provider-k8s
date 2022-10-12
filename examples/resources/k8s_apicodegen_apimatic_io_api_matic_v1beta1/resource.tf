resource "k8s_apicodegen_apimatic_io_api_matic_v1beta1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}

resource "k8s_apicodegen_apimatic_io_api_matic_v1beta1" "example" {
  metadata = {
    name = "apimatic-sample"
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
