output "manifests" {
  value = {
    "example" = data.k8s_chaos_mesh_org_workflow_node_v1alpha1_manifest.example.yaml
  }
}
