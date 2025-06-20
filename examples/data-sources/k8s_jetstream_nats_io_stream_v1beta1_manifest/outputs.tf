output "manifests" {
  value = {
    "example" = data.k8s_jetstream_nats_io_stream_v1beta1_manifest.example.yaml
  }
}
