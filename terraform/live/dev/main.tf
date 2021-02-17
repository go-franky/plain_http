resource "aws_iam_role" "lambda_role" {
  name = "${var.name}-lambda-role"
  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Action" : "sts:AssumeRole",
        "Principal" : {
          "Service" : "lambda.amazonaws.com"
        },
        "Effect" : "Allow",
      }
    ]
  })
  path = "/service-role/"
  tags = {
    terraform = true
  }
}

resource "aws_lambda_function" "main" {
  function_name = var.name
  handler       = "lambda"
  role          = aws_iam_role.lambda_role.arn
  memory_size = 128

  runtime = "go1.x"
  filename = "lambda.zip"
  source_code_hash = filebase64sha256("lambda.zip")
  tags = {
    terraform = true
  }
}

module "api-gateway" {
  source = "../../modules/api-gateway/proxy"
  name = var.name
  lambda = aws_lambda_function.main
}
