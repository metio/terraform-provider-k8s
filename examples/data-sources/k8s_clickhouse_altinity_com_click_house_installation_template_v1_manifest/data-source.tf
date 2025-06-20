data "k8s_clickhouse_altinity_com_click_house_installation_template_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
