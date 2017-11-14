#encoding=utf-8
import datetime


def printl(txt):
    print '%s %s' % (datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S'),
                     txt)
