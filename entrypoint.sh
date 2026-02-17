#!/bin/sh
xvfb-run --auto-servernum --server-args="-screen 0 1280x1024x24" ./voeru

if [ -f "./download/linux-1520176/chrome-linux/chrome" ]; then
    echo "chrome found!"
    chmod +x ./download/linux-1520176/chrome-linux/chrome
else
    echo "chrome not found"
    exit 1
fi

while true; do
    echo "run app: $(date)"
    
    xvfb-run --auto-servernum --server-args="-screen 0 1280x1024x24" ./rodd
    
    echo "DONE. WAIT 13 min..."
    sleep 780
done