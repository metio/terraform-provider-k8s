resource "k8s_kyverno_io_update_request_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
