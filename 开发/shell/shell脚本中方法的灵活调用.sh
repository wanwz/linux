#!/usr/bin/env bash
# 直接通过传参来调用脚本中的方法

dog() {
  echo "wang~~"
}

cat() {
  echo "miao~~"
}

$1 2>/dev/null || echo "没有此方法"

]# sh test.sh dog
wang~~
]# sh test.sh cat
miao~~
]# sh test.sh ksdf
没有此方法
