#!/bin/bash
DEPLOY_SHA="__GITHUB_SHA__"

aws s3 cp "s3://test-s3-bucket-753674274065-ap-southeast-1-an/artifaces/${DEPLOY_SHA}/" /opt/dist

export PORT= :8080

chmod +x /opt/dist

/opt/dist

