data "k8s_infinispan_org_infinispan_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
