resource "k8s_infinispan_org_batch_v2alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
