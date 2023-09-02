output "manifests" {
  value = {
    "example" = data.k8s_lb_lbconfig_carlosedp_com_external_load_balancer_v1_manifest.example.yaml
  }
}
