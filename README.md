# 改动

## 1. 签名接口 Signature
```$xslt
入参: 移除 IDlen 和 message_len 两处长度描述
出参: 添加一个 v 参数
改动: 签名结果满足所有特殊签名标准
```

## 2. 验签接口 Verify
```$xslt
入参: 移除 IDlen 和 message_len 两处长度描述
```

## 3. 协商相关接口
```$xslt
入参: 移除 IDinitiator_len 和 IDresponder_len 两处长度描述
```

## 4. 不再支持接口
```$xslt
PreprocessRandomNum
GetCurveOrder
```
