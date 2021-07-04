# 前言
session: 主要儲存在server端，用以驗證client端來訪的cookie是否為合法，且具有時效性的
JWT: 內建一個expire time來做時效性監控，並且會夾帶使用者資訊。

# 動機:
JWT本身可夾帶一個期限，以及對於登入使用者的描述，但是無法做後台管控。因此需要藉由redis來做註冊以及必要時銷毀。

# refer:
- https://learnku.com/articles/22616
- https://medium.com/@sherryhsu/session-vs-token-based-authentication-11a6c5ac45e4
- https://medium.com/@yuliaoletskaya/can-jwt-be-used-for-sessions-4164d124fe23