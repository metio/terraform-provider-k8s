output "manifests" {
  value = {
    "example" = data.k8s_appprotectdos_f5_com_ap_dos_policy_v1beta1_manifest.example.yaml
  }
}
