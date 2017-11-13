#encoding=utf-8
import time
print 'hello jenkins heheh'
with open('text.log','w') as fw:
  fw.write('hello jenkins')
for i in range(10):
  print i
  time.sleep(1)
raise Exception("test")
