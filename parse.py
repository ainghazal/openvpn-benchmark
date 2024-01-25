import json
import sys

from timelength import TimeLength


d = open(sys.argv[1]).read()
data = json.loads(d)

d = open(sys.argv[2]).read()
data = data + json.loads(d)


def fix(data):
    for row in data:
        t = row['elapsed']
        length = TimeLength(t)
        row['elapsed'] = length.to_seconds()

fix(data)
print(json.dumps(data))
