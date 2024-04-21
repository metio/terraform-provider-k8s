data "k8s_servicemesh_cisco_com_peer_istio_control_plane_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
