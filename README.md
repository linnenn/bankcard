# BankCard 

[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/linnenn/backcard/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/linnenn/backcard/?branch=main)
[![Build Status](https://scrutinizer-ci.com/g/linnenn/backcard/badges/build.png?b=main)](https://scrutinizer-ci.com/g/linnenn/backcard/build-status/main)

Introduction

------------
根据银行卡号获取银行信息（银行名称, 信用卡/借记卡/准贷记卡/预付费卡, 银行LOGO 等）支持go1.14+.


## Install
`go version:go1.14+`

`go get -u github.com/linnenn/bankcard`

## Usage

```
bankcard.BankNameMapList() //所有 银行缩写=> 银行名
bankcard.BankNameList() //所有银行

bankcard.BankCardInfo('6214850106608722')
//返回
{
 Validated: true         // bank验证成功与否
 ValidatedMsg: cardNO error //银行卡号错误
 BankENShortName: CMB      // 银行缩写
 BankENFullName: CMB      // 银行英文全称
 BankCNShortName: 招商银行     // 银行中文缩写
 BankCNFullName: 招商银行     // 银行中文全名
 BankImg: https://apimg.alipay.com/combo.png?d=cashier&t=CMB  // 银行LOGO
 CardType: DC  // 卡类型：CC：信用卡 DC：储蓄卡 SCC：准贷记卡 PC：预付费卡
 CardTypeName: 储蓄卡 // 卡类型名称：信用卡 | 储蓄卡 | 准贷记卡 | 预付费卡
 }
```

## License

MIT

