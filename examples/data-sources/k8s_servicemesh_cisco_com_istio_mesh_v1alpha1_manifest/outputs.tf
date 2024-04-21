output "manifests" {
  value = {
    "example" = data.k8s_servicemesh_cisco_com_istio_mesh_v1alpha1_manifest.example.yaml
  }
}
