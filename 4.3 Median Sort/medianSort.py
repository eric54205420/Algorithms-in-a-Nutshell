#!/usr/bin/env python
# -*- coding: utf-8 -*-

import random
import itertools

def medianSort(l):
    m = len(l)/2

    left = []
    right = []
    for i in xrange(0, len(l)):
        if i == m:
            continue
        if l[i]>l[m]:
            right.append(l[i])
        else:
            left.append(l[i])
    if len(left)>1:
        left = medianSort(left)
    if len(right)>1:
        right = medianSort(right)

    l = list(itertools.chain(left, [l[m]], right))
    return l



if __name__ == "__main__":
    l = []
    for x in xrange(10):
        l.append(int(random.random()*100))
    print l
    print medianSort(l)
