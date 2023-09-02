resource "k8s_kyverno_io_policy_exception_v2alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
