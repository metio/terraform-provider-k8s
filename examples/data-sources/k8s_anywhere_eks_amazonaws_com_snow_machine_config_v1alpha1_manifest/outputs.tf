output "manifests" {
  value = {
    "example" = data.k8s_anywhere_eks_amazonaws_com_snow_machine_config_v1alpha1_manifest.example.yaml
  }
}
