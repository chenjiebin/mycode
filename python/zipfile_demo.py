#!/usr/bin/env python
#coding=utf-8

import zipfile

#apk包文件路径
zipfilePath = '/home/chenjiebin/Downloads/86dada8ec12c639fc2a814b9d1bcab78.apk'
z = zipfile.ZipFile(zipfilePath, 'r')

#获取文件路径
for f in z.namelist():
	print f

#读取文件信息
for i in z.infolist():
	print i.filename, i.file_size, i.header_offset


#提取单个文件
print z.read(z.namelist()[0])


#创建压缩包
import os 
filename = '/home/chenjiebin/Downloads/test.rar'
z = zipfile.ZipFile(filename, 'w')
testdir = '/home/chenjiebin/mycode/mycode/c'
if os.path.isdir(testdir):
	for d in os.listdir(testdir):
		z.write(testdir+os.sep+d)

z.close()




