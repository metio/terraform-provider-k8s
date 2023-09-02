output "manifests" {
  value = {
    "example" = data.k8s_cilium_io_cilium_cidr_group_v2alpha1_manifest.example.yaml
  }
}
