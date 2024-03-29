output "manifests" {
  value = {
    "example" = data.k8s_core_kubeadmiral_io_override_policy_v1alpha1_manifest.example.yaml
  }
}
