output "manifests" {
  value = {
    "example" = data.k8s_k8s_nginx_org_policy_v1alpha1_manifest.example.yaml
  }
}
