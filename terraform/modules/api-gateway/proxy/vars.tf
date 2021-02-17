variable "name" {
  description = "Name of the API gateway"
  type = string
}

variable "lambda" {
  description = "The lambda to do stuff"
  type = object({
    invoke_arn = string
    function_name = string
  })
}

variable tags {
  description = "Any tags we want to associate"
  type        = map(string)
  default     = {}
}
