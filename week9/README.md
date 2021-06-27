# Week9 Homework

## socket 粘包的解包方式
- fix length: 包都是固定長度, 不滿固定長度用空格補齊, 可能會浪費帶寬, 擴展性差
- delimiter based:  每個包結尾補上特定的分隔符, 像是 \r\n, \n etc., 同時要確保分隔符不會用在包的內容之中\
  eg. redis RECP
- length field based frame decoder:  包的 header 包含了長度訊息, 擴展性佳\
  eg. http, protocol buffer