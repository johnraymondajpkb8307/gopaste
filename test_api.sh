# test_api.sh
#!/bin/bash

# 创建paste
echo "Testing CREATE paste..."
CREATE_RESPONSE=$(curl -s -X POST -H "Content-Type: application/json" \
    -d '{"content":"Test content","expire_hours":24}' \
    http://localhost:8080/api/paste)
PASTE_ID=$(echo $CREATE_RESPONSE | jq -r '.id')
echo "Created paste ID: $PASTE_ID"

# 获取paste
echo "Testing GET paste..."
curl -s -X GET http://localhost:8080/api/paste/$PASTE_ID | jq '.'