resource "k8s_kyverno_io_generate_request_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    context  = {}
    policy   = "some-policy"
    resource = {}
  }
}
