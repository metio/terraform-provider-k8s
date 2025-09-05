output "manifests" {
  value = {
    "example" = data.k8s_topology_volcano_sh_hyper_node_v1alpha1_manifest.example.yaml
  }
}
