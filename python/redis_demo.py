#!/usr/bin/env python
#coding=utf-8

'''
redis的并发连接测试
'''
import time
import threading
import redis


def loop(i):
	r = redis.Redis(host='192.168.66.66', port=6379, db=0)
	key = 'thread_count_%d' % i
	r.set (key, i)
	for j in range(30):
		r.incr(key)
		time.sleep(1)

	
	print r.get(key)


threads = []	
for i in range(1000):	
	threads.append(threading.Thread(target=loop, args=(i,)))

for t in threads:
	t.start()

for t in threads:
	t.join()
	

exit(0)
#pool = redis.ConnectionPool(host='192.168.66.66', port=6379, db=0)
#r = redis.Redis(connection_pool=pool)
r = redis.Redis(host='192.168.66.66', port=6379, db=0)


r.info()

time.sleep(30)



#查看redis信息
info = r.info()
for key in info:
	print "%s: %s" % (key, info[key])

#查看数据库大小
print "\ndbsize:%s" % r.dbsize()

#看连接
print "ping %s" % r.ping()



