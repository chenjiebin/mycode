#!/usr/bin/env python
#coding=utf-8

import zipfile

#apk包文件路径
zipfilePath = '/home/chenjiebin/Downloads/86dada8ec12c639fc2a814b9d1bcab78.apk'
z = zipfile.ZipFile(zipfilePath, 'r')

#获取文件路径
for f in z.namelist():
	print f.split('/')
        
