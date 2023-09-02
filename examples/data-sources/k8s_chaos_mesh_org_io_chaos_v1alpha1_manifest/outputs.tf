output "manifests" {
  value = {
    "example" = data.k8s_chaos_mesh_org_io_chaos_v1alpha1_manifest.example.yaml
  }
}
