resource "k8s_hazelcast_com_map_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    hazelcast_resource_name = "some-name"
  }
}
