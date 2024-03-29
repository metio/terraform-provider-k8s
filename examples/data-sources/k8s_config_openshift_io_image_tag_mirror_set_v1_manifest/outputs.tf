output "manifests" {
  value = {
    "example" = data.k8s_config_openshift_io_image_tag_mirror_set_v1_manifest.example.yaml
  }
}
