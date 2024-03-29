data "k8s_memorydb_services_k8s_aws_acl_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
