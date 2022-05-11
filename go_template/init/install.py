import os
from pathlib import Path

LOG_PATH="./bin/logs/"

print("start install...")

if not os.path.isdir(LOG_PATH): 
    os.mkdir(LOG_PATH)
    print(">>> log path created...")
else: 
    print(">>> log path exist")

print("DONE!")

