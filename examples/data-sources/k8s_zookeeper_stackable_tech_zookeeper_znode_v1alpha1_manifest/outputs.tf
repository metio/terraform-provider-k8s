output "manifests" {
  value = {
    "example" = data.k8s_zookeeper_stackable_tech_zookeeper_znode_v1alpha1_manifest.example.yaml
  }
}
