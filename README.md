# eosdev-go

## 目前的问题

- 如何发币？

## 已经解决的问题

- 创建账户需要一定量的EOS币，那么这个EOS币最小的量是多少？
  > 目前每个新的账户，最少需要2996 Bytes，但是，考虑内存还有部分需要手续费什么的，目前代码里面设置最小需要3020 Bytes.
- 创建账户是否消耗大账户的内存？
  > 是的
- 从大账户给小账户转账是否消耗大账户的内存？
  > 是的
- 创建代笔的过程中需要指定最大的量，那么这个量最大能指定多少？
  > 目前没有找到最大值的算法，但是可以给个估计值：“300000000000000.0000 CYC” （3 * 10^14）
- 如何部署合约？
  > 参考“部署合约”文件夹的内容