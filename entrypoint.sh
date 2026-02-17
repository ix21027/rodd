#!/bin/sh

while true; do
    echo "run app: $(date)"
    
    xvfb-run --auto-servernum --server-args="-screen 0 1280x1024x24" ./app_binary
    
    echo "DONE. WAIT 13 min..."
    sleep 780
done