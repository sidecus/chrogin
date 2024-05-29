#!/bin/bash
PDF=${2:-output.pdf}
URL=${1:-http://localhost:8080/report}
echo Processing $URL to $PDF
chromium --no-sandbox --headless -disable-gpu --no-pdf-header-footer --timeout=5000 --print-to-pdf=${PDF} ${URL}