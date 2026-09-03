#!/bin/bash
set -euxo pipefail
DEPLOY_SHA="__GITHUB_SHA__"
echo "copying from S3"
aws s3 cp "s3://test-s3-bucket-753674274065-ap-southeast-1-an/artifaces/${DEPLOY_SHA}/" /opt/dist
echo "Setting port number as 8080
export PORT_NUMBER= :8080

chmod +x /opt/dist
echo "serving on port 8080"
/opt/dist

