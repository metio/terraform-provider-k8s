output "manifests" {
  value = {
    "example" = data.k8s_rocketmq_apache_org_broker_v1alpha1_manifest.example.yaml
  }
}
