variable "length" {
  type = number
  default = 16
}

resource "random_string" "random" {
  length           = var.length
  special          = true
  override_special = "/@£$"
}

output "result" {
  value = random_string.random.result
}