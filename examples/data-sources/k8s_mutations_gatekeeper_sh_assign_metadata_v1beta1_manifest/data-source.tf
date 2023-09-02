data "k8s_mutations_gatekeeper_sh_assign_metadata_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"
    
  }
}
