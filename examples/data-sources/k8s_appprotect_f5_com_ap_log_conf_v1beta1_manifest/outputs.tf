output "manifests" {
  value = {
    "example" = data.k8s_appprotect_f5_com_ap_log_conf_v1beta1_manifest.example.yaml
  }
}
