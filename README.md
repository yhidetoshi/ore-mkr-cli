# 使い方(コマンドオプション)

### host
```
  ■ 一覧
    > $ ore-mkr -org=<ORG> -type=host

  ■ statusを変更する
    > $ ore-mkr -org=<ORG> -type=host -<STATUS> target=<HOSTID>
      ※) STATUS: working standby maintenance poweroff retire
```

### 監視設定
```
  ■ 一覧 
    > $ ore-mkr -org=<ORG> -type=monitor
```

### アラート一覧
```
  ■ 一覧
    > $ ore-mkr -org=<ORG> -type=alert
```