#encoding=utf-8
import time
from test2 import printl
print 'hello jenkins heheh'
with open('text.log', 'w') as fw:
    fw.write('hello jenkins')
for i in range(100):
    printl(i)
    time.sleep(1)
