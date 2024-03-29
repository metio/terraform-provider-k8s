output "manifests" {
  value = {
    "example" = data.k8s_m4e_krestomat_io_routine_v1alpha1_manifest.example.yaml
  }
}
