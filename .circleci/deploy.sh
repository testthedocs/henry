#!/usr/bin/env bash
set -e # halt script on error

zip -r website.zip ~/sphinx/html/_build/html

curl -H "Content-Type: application/zip" \
     -H "Authorization: Bearer $NETLIFYKEY" \
     --data-binary "@website.zip" \
     https://api.netlify.com/api/v1/sites/$API/deploys
