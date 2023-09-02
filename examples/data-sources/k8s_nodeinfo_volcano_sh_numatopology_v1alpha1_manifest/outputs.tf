output "manifests" {
  value = {
    "example" = data.k8s_nodeinfo_volcano_sh_numatopology_v1alpha1_manifest.example.yaml
  }
}
