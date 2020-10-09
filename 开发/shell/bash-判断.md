## bash-判断

### 文件判断

- [ -a file ]：如果file存在，则为true
- [ -b file ]：如果file存在并且是一个块设备文件，则为true
- [ -c file ]：如果file存在并且是一个字符设备文件，则为true
- [ -d file ]：如果file存在并且是一个目录，则为true
- [ -e file ]：如果file存在，则为true
- [ -f file ]：如果file存在并且是一个普通文件，则为true
- [ -g file ]：如果file存在并且设置了组ID，则为true
- [ -G file ]：如果file存在并且属于有效的组ID，则为true
- [ -h file ]：如果file存在并且是符号链接，则为true
- [ -k file ]：如果file存在并且设置了"stick bit"，则为true
- [ -L file ]：如果file存在并且是一个符号链接，则为true
- [ -N file ]：如果file存在并且自上次读取后已被修改，则为true
- [ -O file ]：如果file存在并且属于有效的用户ID，则为true
- [ -p file ]：如果file存在并且是一个命名管道，则为true
- [ -r file ]：如果file存在并且可读（当前用户有可读权限），则为true
- [ -s file ]：如果file存在并且其长度大于零，则为true
- [ -S file ]：如果file存在并且是一个网络socket，则为true
- [ -t fd ]：如果fd是一个文件描述符，并且重定向到终端，则为true
- [ -u file ]：如果file存在并且设置了setuid位，则为true
- [ -w file ]：如果file存在并且可写（当前用户拥有可写权限），则为true
- [ -x file ]：如果file存在并且可执行，则为true
- [ file1 -nt file2 ]：如果file1比file2的更新时间最近，或者file1存在而file2不存在，则为true
- [ file1 -ot file2 ]：如果file1比file2的更新时间更旧，或者file2存在而file1不存在，则为true
- [ file1 -ef file2 ]：如果file1和file2引用相同的设备和iNode编号，则为true

### 字符串判断

- [ string ]：如果string不为空（长度大于0），则为true
- [ -n string ]：如果string长度大于0，则为true
- [ -z string ]：如果string长度为0，则为true
- [ string1 = string2 ]：如果string1和string2相同，则为true
- [ string1 == string2 ]：？？？
- [ string1 != string2 ]：如果string1和string2不相同，则为true
- [ string1 '>' string2 ]：如果按照字典顺序string1排列在string2之后，则为true
- [ string1 '<' string2 ]：如果按照字典顺序string1排列在string2之前，则为true

> test和[]中>和<，必须用引号引起来或用反斜杠转义，否则会被shell解释为重定向操作符

### 整数判断

- [ int1 -eq int2 ]：如果int1 == int2，则为true
- [ int1 -ne int2 ]：如果int1 != int2，则为true
- [ int1 -le int2 ]：如果int1 <= int2，则为true
- [ int1 -lt int2 ]：如果int1 < int2，则为true
- [ int1 -ge int2 ]：如果int1 >= int2，则为true
- [ int1 -gt int2 ]：如果int1 > int2，则为true

> 数值判断不能用>,<,=等进行判断

### 正则判断

1. [[ expression ]]支持正则表达式

   1. 格式：[[ string1 =~ regex ]]

      ```shell
      #!/bin/bash
      INT=-3
      if [[ "$INT" =~ ^-?[0-9]+$ ]]; then
      	echo "$INT is an integer."
      	exit 0
      else
      	echo "$INT is not an integer." &>/dev/null
      	exit 1
      fi
      先判断变量INT的字符串形式，是否满足^-?[0-9]+$的正则模式，如果满足就表明它是一个整数
      ```

      

### 算数判断

- ((1))表示判断成立为真，((0))表示判断不成立为假
- ((3 > 2)) = ((3>2)) = (( 3 > 2 )) = ((   3    >     2))

> 注意点
>
> 1. test expression = [ expression ] = [[ expression ]]
> 2. test expression 和 [ expression ]不支持正则表达式，[[ expression ]]支持
> 3. 在 test 和 []中的变量要放在双引号之中，可以防止\$variable为空(例[ -e "$variable" ])，导致[ -e ]判断为值，而放在双引号当中，返回的就总是一个空字符串，[ -e "" ]会判断为假。 但是在[[ -e \$variable ]]中可以不加双引号

### 逻辑运算

- AND = && = -a
- OR = || = -o
- NOT = !

1. 对判断进行逻辑运算

   1. [[  \$INT -ge 1 && \$INT -le 50 ]]

   2. [ !\\( \$INT -ge 1 -a \$INT -le 50 \\)]

      > test命令内部使用的圆括号，必须使用引号或者转义，否则会被Bash解释

2. 对普通命令进行逻辑运算
   1. command1 && command2，先执行command1，成功后才执行command2
   2. command1 || command2，先执行command1，失败后才执行command2

### if结构

1. 格式：

   ```shell
   if commands; then
   	commands
   	[elif commands; then
   		commands...]
   	[else
   		commands]
   fi
   ```

2. commands可以是

   ```shell
   1.
   if test $USER == "root"; then
   	echo "Hi,root."
   else
   	echo "You are not root."
   fi
   2.
   if true
   then
   	echo "hi,true."
   fi
   3.
   if echo "hi"; then echo "Hi,man!";fi
   4.
   if i=1;i++;i<10; then
   	echo "$i"
   else
   	echo "End"
   fi
   这个表示if后面可以跟任意数量的命令，这时所有的命令都会执行，但是判断真伪只看最后一个命令，即使前面所有的命令都失败，只有最后一个命令返回0，就会执行then的部分，可以理解为 if false;true; then echo "Hi";fi
   ```

   

### case结构

1. 格式

   ```shell
   case expression in
   	pattern)
   		commands ;;
   	pattern)
   		commands ;;
   	...
   esac
   ```

2. expression是一个表达式

   ```shell
   #!/bin/bash
   OS=$(uname -s)
   
   case "$OS" in
   	FreeBSD) echo "FreeBSD";;
   	Darwin) echo "Mac OSX";;
   	AIX) echo "AIX";;
   	Minix) echo "Minix";;
   	Linux) echo "Linux";;
   	*) echo "Failed to identify this OS";;
   esac
   ```

   

3. pattern是表达式的值或一个模式（可以使用通配符），可以有多条，用来匹配多个值，每条以两个分号结尾

   ```shell
   #!/bin/bash
   
   echo -n "Enter a letter or number: "
   read character
   case $character in
   	[[:lower:]] | [[:upper:]])
   		echo "Entered letter $character";;
   	[0-9])
   		echo "Entered number $character";;
   	*)
   		echo "Input does not meet requirements."
   esac
   ```

4. Bash4.0之前，case结果只能匹配一个条件，然后就会推出case结构；Bash4.0之后，允许匹配多个条件，这时可以用;;&终止每个条件块

   ```shell
   #!/bin/bash
   
   read -n 1 -p "Enter a character: "
   case $REPLY in
   	[[:upper:]]) echo "$REPLY is a capital letter." ;;&
   	[[:lower:]]) echo "$REPLY is a lowercase letter." ;;&
   	[[:alpha:]]) echo "$REPLY is a letter." ;;&
   	[[:digit:]]) echo "$REPLY is a digit." ;;&
   esac
   ```
