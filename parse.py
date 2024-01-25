import json

import sys

d = open(sys.argv[1]).read()
data = json.loads(d)

for row in data:
    if row['elapsed'].endswith('ms'):
        n = float(row['elapsed'].strip('ms')) / 1E3
        row['elapsed'] = n
        continue
    if row['elapsed'].endswith('s'):
        n = float(row['elapsed'].strip('s'))
        row['elapsed'] = n
        continue

print(json.dumps(data))
