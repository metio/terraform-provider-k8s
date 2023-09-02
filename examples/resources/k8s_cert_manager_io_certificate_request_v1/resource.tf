resource "k8s_cert_manager_io_certificate_request_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
