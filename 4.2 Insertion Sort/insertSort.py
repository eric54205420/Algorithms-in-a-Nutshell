#!/usr/bin/env python
# -*- coding: utf-8 -*-

import random

def insertSort(l):
    for i in xrange(len(l)):
        for j in xrange(i):
            if l[i] < l[j]:
                l.insert(j, l.pop(i))

if __name__ == "__main__":
    l = []
    for x in xrange(10):
        l.append(int(random.random()*100))
    print l
    insertSort(l)
    print l
