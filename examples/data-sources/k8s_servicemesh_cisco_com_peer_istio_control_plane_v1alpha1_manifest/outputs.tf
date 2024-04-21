output "manifests" {
  value = {
    "example" = data.k8s_servicemesh_cisco_com_peer_istio_control_plane_v1alpha1_manifest.example.yaml
  }
}
