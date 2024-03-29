output "manifests" {
  value = {
    "example" = data.k8s_onepassword_com_one_password_item_v1_manifest.example.yaml
  }
}
