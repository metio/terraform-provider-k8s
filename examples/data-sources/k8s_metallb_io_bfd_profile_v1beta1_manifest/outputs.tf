output "manifests" {
  value = {
    "example" = data.k8s_metallb_io_bfd_profile_v1beta1_manifest.example.yaml
  }
}
