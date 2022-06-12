locals {
  result = yamldecode(data.atmos_spacelift_stack_config.example.output)
}

data "atmos_spacelift_stack_config" "example" {
  process_component_deps     = true
  process_stack_deps         = true
  process_imports            = true
  stack_config_path_template = "stacks/%s.yaml"
}
