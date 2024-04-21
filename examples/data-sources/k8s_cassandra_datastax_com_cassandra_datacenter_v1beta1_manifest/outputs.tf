output "manifests" {
  value = {
    "example" = data.k8s_cassandra_datastax_com_cassandra_datacenter_v1beta1_manifest.example.yaml
  }
}
