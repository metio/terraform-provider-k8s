output "manifests" {
  value = {
    "example" = data.k8s_chaos_mesh_org_physical_machine_v1alpha1_manifest.example.yaml
  }
}
