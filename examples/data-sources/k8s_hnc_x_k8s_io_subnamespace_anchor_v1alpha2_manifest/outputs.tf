output "manifests" {
  value = {
    "example" = data.k8s_hnc_x_k8s_io_subnamespace_anchor_v1alpha2_manifest.example.yaml
  }
}
