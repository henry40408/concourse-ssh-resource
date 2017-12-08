#!/bin/sh
set -e
echo "---------running out tests---------------"

echo "-------------- payload_out_placeholders_value"
/opt/resource/out /tests/assets < /tests/assets/payload_out_placeholders_value.json
echo "payload_out_placeholders_value succeeded"


echo "-------------- payload_out_placeholders_file"
/opt/resource/out /tests/assets < /tests/assets/payload_out_placeholders_file.json | jq .
echo "payload_out_placeholders_file succeeded"
