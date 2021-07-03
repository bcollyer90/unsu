#!/usr/bin/python3

import os
import re
import argparse

parser = argparse.ArgumentParser()
parser.add_argument("cmd", nargs='+')

args, unknown = parser.parse_known_args()
command = ' '.join(args.cmd) + ' ' + ' '.join(unknown)

rx = re.compile("[sudo]{4}")
if rx.match(command):
    os.system(command.replace("sudo ", ""))
else:
    os.system(command)
