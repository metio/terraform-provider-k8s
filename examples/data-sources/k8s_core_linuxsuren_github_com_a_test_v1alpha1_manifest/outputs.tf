output "manifests" {
  value = {
    "example" = data.k8s_core_linuxsuren_github_com_a_test_v1alpha1_manifest.example.yaml
  }
}
