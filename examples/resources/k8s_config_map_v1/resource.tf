resource "k8s_config_map_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_config_map_v1" "example" {
  metadata = {
    name = "test"
  }
  data = {
    api_host             = "myhost:443"
    db_host              = "dbhost:5432"
    "my_config_file.yml" = file("${path.module}/main.tf")
  }

  binary_data = {
    "my_payload.bin" = filebase64("${path.module}/outputs.tf")
  }
}
