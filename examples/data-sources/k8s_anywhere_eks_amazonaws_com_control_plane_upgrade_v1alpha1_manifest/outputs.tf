output "manifests" {
  value = {
    "example" = data.k8s_anywhere_eks_amazonaws_com_control_plane_upgrade_v1alpha1_manifest.example.yaml
  }
}
