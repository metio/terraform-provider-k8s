output "manifests" {
  value = {
    "example" = data.k8s_jetstream_nats_io_consumer_v1beta1_manifest.example.yaml
  }
}
