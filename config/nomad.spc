connection "nomad" {
  plugin = "nomad"

  # Address is required for requests. Required.
  # This can also be set via the NOMAD_ADDR environment variable.
  # address = "http://18.118.164.168:4646"

  # The secret ID of ACL token is required for ACL-enabled Nomad servers. Optional.
  # For more information on the ACL Token, please see https://developer.hashicorp.com/nomad/tutorials/access-control/access-control-tokens.
  # This can also be set via the NOMAD_TOKEN environment variable.
  # secret_id = "c178b810-8b18-6f38-016f-725ddec5d58"

  # Namespace is required for Nomad Enterprise access. Optional.
  # For more information on the Namespace, please see https://developer.hashicorp.com/nomad/tutorials/manage-clusters/namespaces.
  # This can also be set via the NOMAD_NAMESPACE environment variable.
  # "*" indicates all the namespaces available.
  # namespace = "*"
}

