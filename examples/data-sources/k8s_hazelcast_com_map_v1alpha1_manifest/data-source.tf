data "k8s_hazelcast_com_map_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    hazelcast_resource_name = "some-name"
  }
}
