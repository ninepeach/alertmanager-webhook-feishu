#!/usr/bin/env bash
alerts_message='[
  {
    "labels": {
       "alertname": "磁盘已满",
       "dev": "sda1",
       "instance": "实例1",
       "msgtype": "testing",
       "dc": "caohejing"
     },
     "annotations": {
        "info": "程序员小王提示您：这个磁盘sda1已经满了，快处理！",
        "summary": "请检查实例示例1"
      }
  },
  {
    "labels": {
       "alertname": "磁盘已满",
       "dev": "sda2",
       "instance": "实例1",
       "msgtype": "testing",
       "dc": "caohejing"
     },
     "annotations": {
        "info": "程序员小王提示您：这个磁盘sda2已经满了，快处理！",
        "summary": "请检查实例示例1",
        "runbook": "以下链接http://test-url应该是可点击的"
      }
  }
]'

curl -XPOST -d"$alerts_message" http://127.0.0.1:9093/api/v1/alerts

