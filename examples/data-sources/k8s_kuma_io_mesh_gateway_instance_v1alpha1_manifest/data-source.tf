data "k8s_kuma_io_mesh_gateway_instance_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
