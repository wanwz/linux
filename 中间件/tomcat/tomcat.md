## Tomcat

### 安装

1. jdk安装

   ```shell
   yum install java-1.8.0-openjdk -y
   ```

2. 下载tomcat

   ```shell
   wget http://mirrors.tuna.tsinghua.edu.cn/apache/tomcat/tomcat-8/v8.5.47/bin/apache-tomcat-8.5.47.tar.gz
   ```

3. 测试

   ```shell
   tar xvf apache-tomcat-8.5.47.tar.gz -C /opt #解压到/opt目录下
   mv apache-tomcat-8.5.47 tomcat8 
   /opt/tomcat8/bin/startup.sh #启动tomcat
   service iptables stop #关闭防火墙
   iptables -I INPUT -p tcp -m tcp --dport 8080 -j ACCEPT #或者开通8080端口访问
   http://ip:8080 #浏览器访问
   ```

4. tomcat优化

   1. tomcat后台配置用户查看

      ```shell
      ls /opt/tomcat8/webapps/ #查看manager项目是否存在，存在后执行下一步
      vim /opt/tomcat8/conf/tomcat-users.xml
      <role rolename="tomcat"/>
      <role rolename="manager-gui"/>
      <role rolename="manager-status"/>
      
      <user username="tomcat" password="123456" roles="tomcat,manager-gui,manager-status"/>
      vim /opt/tomcat8/webapps/manager/META-INF/context.xml #正常情况下manager-gui界面只允许127.0.0.1访问，需要关闭这个限制
      <Context antiResourceLocking="false" privileged="true" >
        <!--
        <Valve className="org.apache.catalina.valves.RemoteAddrValve" allow="127\.\d+\.\d+\.\d+|::1|0:0:0:0:0:0:0:1" />
        -->
        <Manager sessionAttributeValueClassNameFilter="java\.lang\.(?:Boolean|Integer|Long|Number|String)|org\.apache\.catalina\.filters\.CsrfPreventionFilter\$LruCache(?:\$1)?|jav
      </Context>
      
      http://ip:8080/manager/status #重启tomcat后访问,可以看到jvm堆栈信息
      Max threads: 200 #配置文件里面限制的最大线程数
      Current thread count: 10 #当前线程数
      Current thread busy: 1 #当前繁忙的线程数，如果当前繁忙线程已经是接近最大线程数，那基本可以表示负载到了
      Keep alive sockets count: 1 #保持连接数
      ```

   2. 连接池配置

      ```
      vim /opt/tomcat8/conf/server.xml
      默认值：
      <!--
      <Executor name="tomcatThreadPool" namePrefix="catalina-exec-"
          maxThreads="150" minSpareThreads="4"/>
      -->
      修改为：
      <Executor
              name="tomcatThreadPool"
              namePrefix="catalina-exec-"
              maxThreads="500"
              minSpareThreads="30"
              maxIdleTime="60000"
              prestartminSpareThreads = "true"
              maxQueueSize = "100"
      />
      重点参数解释：
      maxThreads，最大并发数，默认设置 200，一般建议在 500 ~ 800，根据硬件设施和业务来判断
      minSpareThreads，Tomcat 初始化时创建的线程数，默认设置 25
      prestartminSpareThreads，在 Tomcat 初始化的时候就初始化 minSpareThreads 的参数值，如果不等于 true，minSpareThreads 的值就没啥效果了
      maxQueueSize，最大的等待队列数，超过则拒绝请求
      maxIdleTime，如果当前线程大于初始化线程，那空闲线程存活的时间，单位毫秒，默认60000=60秒=1分钟。
      ```

   3. 链接参数配置

      ```shell
      默认值：
      <Connector 
          port="8080" 
          protocol="HTTP/1.1" 
          connectionTimeout="20000" 
          redirectPort="8443" 
      />
      修改为：
      <Connector 
         executor="tomcatThreadPool"
         port="8080" 
         protocol="org.apache.coyote.http11.Http11Nio2Protocol" 
         connectionTimeout="20000" 
         maxConnections="10000" 
         redirectPort="8443" 
         enableLookups="false" 
         acceptCount="100" 
         maxPostSize="10485760" 
         maxHttpHeaderSize="8192" 
         disableUploadTimeout="true" 
         URIEncoding="utf-8"
      />
      重点参数解释：
      protocol，Tomcat 8 设置 nio2 更好：org.apache.coyote.http11.Http11Nio2Protocol（如果这个用不了，就用下面那个）
      protocol，Tomcat 6、7 设置 nio 更好：org.apache.coyote.http11.Http11NioProtocol
      enableLookups，禁用DNS查询，tomcat 8 默认已经是禁用了。
      maxConnections，最大连接数，tomcat 8 默认设置 10000
      acceptCount，指定当所有可以使用的处理请求的线程数都被使用时，可以放到处理队列中的请求数，超过这个数的请求将不予处理，默认设置 100
      maxPostSize，以 FORM URL 参数方式的 POST 提交方式，限制提交最大的大小，默认是 2097152(2兆)，它使用的单位是字节。10485760 为 10M。如果要禁用限制，则可以设置为 -1。
      maxHttpHeaderSize，http请求头信息的最大程度，超过此长度的部分不予处理。一般8K。
      ```

   4. 禁用AJP（如果服务器没有使用Apache）

      ```shell
      注释该行，默认是开启的
      <!-- <Connector port="8009" protocol="AJP/1.3" redirectPort="8443" /> -->
      关闭自动部署功能
      旧值：
      <Host name="localhost"  appBase="webapps" unpackWARs="true" autoDeploy="true">
      新值：
      <Host name="localhost"  appBase="webapps" unpackWARs="true" autoDeploy="false">
      ```

   5. jvm优化

5. Tomcat日志分割

   ```shell
   # 安装epel源
   yum install cronolog -y
   which cronolog #记录目录位置
   vim /opt/tomcat8/bin/catalina.sh # 修改配置文件
   # 448行左右
   448   shift
   449   touch "$CATALINA_OUT"
   450   if [ "$1" = "-security" ] ; then
   451     if [ $have_tty -eq 1 ]; then
   452       echo "Using Security Manager"
   453     fi
   454     shift
   455     eval $_NOHUP "\"$_RUNJAVA\"" "\"$LOGGING_CONFIG\"" $LOGGING_MANAGER $JAVA_OPTS $CATALINA_OPTS \
   456       -D$ENDORSED_PROP="\"$JAVA_ENDORSED_DIRS\"" \
   457       -classpath "\"$CLASSPATH\"" \
   458       -Djava.security.manager \
   459       -Djava.security.policy=="\"$CATALINA_BASE/conf/catalina.policy\"" \
   460       -Dcatalina.base="\"$CATALINA_BASE\"" \
   461       -Dcatalina.home="\"$CATALINA_HOME\"" \
   462       -Djava.io.tmpdir="\"$CATALINA_TMPDIR\"" \
   463       org.apache.catalina.startup.Bootstrap "$@" start \
   464       >> "$CATALINA_OUT" 2>&1 "&"
   465 
   466   else
   467     eval $_NOHUP "\"$_RUNJAVA\"" "\"$LOGGING_CONFIG\"" $LOGGING_MANAGER $JAVA_OPTS $CATALINA_OPTS \
   468       -D$ENDORSED_PROP="\"$JAVA_ENDORSED_DIRS\"" \
   469       -classpath "\"$CLASSPATH\"" \
   470       -Dcatalina.base="\"$CATALINA_BASE\"" \
   471       -Dcatalina.home="\"$CATALINA_HOME\"" \
   472       -Djava.io.tmpdir="\"$CATALINA_TMPDIR\"" \
   473       org.apache.catalina.startup.Bootstrap "$@" start \
   474       >> "$CATALINA_OUT" 2>&1 "&"
   475 
   476   fi
   修改为
   448   shift
   449   # touch "$CATALINA_OUT"
   450   if [ "$1" = "-security" ] ; then
   451     if [ $have_tty -eq 1 ]; then
   452       echo "Using Security Manager"
   453     fi
   454     shift
   455     eval $_NOHUP "\"$_RUNJAVA\"" "\"$LOGGING_CONFIG\"" $LOGGING_MANAGER $JAVA_OPTS $CATALINA_OPTS \
   456       -D$ENDORSED_PROP="\"$JAVA_ENDORSED_DIRS\"" \
   457       -classpath "\"$CLASSPATH\"" \
   458       -Djava.security.manager \
   459       -Djava.security.policy=="\"$CATALINA_BASE/conf/catalina.policy\"" \
   460       -Dcatalina.base="\"$CATALINA_BASE\"" \
   461       -Dcatalina.home="\"$CATALINA_HOME\"" \
   462       -Djava.io.tmpdir="\"$CATALINA_TMPDIR\"" \
   463       org.apache.catalina.startup.Bootstrap "$@" start \
   464       >> "$CATALINA_OUT" 2>&1 | /usr/sbin/cronolog "$CATALINA_BASE"/logs/catalina.%Y-%m-%d.out >> /dev/null &
   465 
   466   else
   467     eval $_NOHUP "\"$_RUNJAVA\"" "\"$LOGGING_CONFIG\"" $LOGGING_MANAGER $JAVA_OPTS $CATALINA_OPTS \
   468       -D$ENDORSED_PROP="\"$JAVA_ENDORSED_DIRS\"" \
   469       -classpath "\"$CLASSPATH\"" \
   470       -Dcatalina.base="\"$CATALINA_BASE\"" \
   471       -Dcatalina.home="\"$CATALINA_HOME\"" \
   472       -Djava.io.tmpdir="\"$CATALINA_TMPDIR\"" \
   473       org.apache.catalina.startup.Bootstrap "$@" start \
   474       >> "$CATALINA_OUT" 2>&1 | /usr/sbin/cronolog "$CATALINA_BASE"/logs/catalina.%Y-%m-%d.out >> /dev/null &
   475 
   476   fi
   ```

   

6. 禁止外网通过8080端口访问tomcat

   ```shell
   iptables -t filter -A INPUT -p tcp -m tcp --dport 8080 -s localhost -j ACCEPt
   iptables -t filter -A INPUT -p tcp -m tcp --dport 8080 -j REJECT
   service iptables save
   service iptables restart
   ```
