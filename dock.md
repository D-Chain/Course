环境是 ubuntu16.04 LTS

安装docker(https://www.cnblogs.com/lighten/p/6034984.html)

1.已经尽可能的把依赖的类库放到了vendor里面，这样减少打包出现的错误。

2.现在Dockerfile文件到一个指定的文件夹(假设是demo) 

3.cd demo

4.手动下载 git clone https://github.com/D-Chain/Course.git 后

  cd Course,执行 "sudo docker build -t d-chain ." 

5.安装完成后可用sudo docker images，看是否有d-chain存在。

6.然后执行sudo docker run d-chain即可（后面可以正常的加参数）
