# url shorter
一个简单的url缩短服务端, 使用redis   
A simple url shorter, used redis   

## Usage
### Visit a shortened url
`GET` /{key}   
### Shorten a url
`POST` /new   
expire不填或为0则为永不失效   
if expire is not exists or 0, it will never expire   
key_len不填或为0则为使用配置文件中的值   
if key_len is not exists or 0, the value in the config.toml is used   
```json
{
  "token": "000000",
  "expire": 0,
  "url": "http://www.github.com",
  "key_len": 0
}
```
响应内容(如果成功则状态码为200，出现错误则为非200):   
Result(200 if success status code, non-200 if error):   
```json
{
  "key": "abc"
}
```
接下来访问http://localhost:8081/abc将会跳转到http://www.github.com   
Next visit http://localhost:8081/abc will jump to http://www.github.com   
