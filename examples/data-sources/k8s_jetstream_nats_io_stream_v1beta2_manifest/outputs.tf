output "manifests" {
  value = {
    "example" = data.k8s_jetstream_nats_io_stream_v1beta2_manifest.example.yaml
  }
}
