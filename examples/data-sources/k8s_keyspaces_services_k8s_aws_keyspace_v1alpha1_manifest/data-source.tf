data "k8s_keyspaces_services_k8s_aws_keyspace_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
