resource "k8s_networking_istio_io_proxy_config_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
