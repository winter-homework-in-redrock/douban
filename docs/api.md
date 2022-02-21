# api

## ❗❗❗所有api基本返回参数

## 一、用户登录

- **返回参数，所有api返回`Content-Type`:`application/json`**，当请求资源失败会返回某项资源的空字符串
- **不用携带token的情况会有注明，没有注明则需要携带token**

| 返回参数               | 说明                                                  |
| ---------------------- | ----------------------------------------------------- |
| status                 | 状态码                                                |
| error                  | status为1时，error为空字符串；0时，返回服务端报错信息 |
| 其他多级或单级json参数 | 会自定义标识某项资源，详情如以下api                   |

| status | 含义                                                 |
| :----- | ---------------------------------------------------- |
| 0      | 失败，会有其他多级或单级json参数选项                 |
| 1      | 成功，会有其他多级或单级json参数选项                 |
| -1     | Authorization为空，只有status、error两项单级json返回 |
| -2     | Bearer不存在，只有status、error两项单级json返回      |
| -3     | 无效的token，只有status、error两项单级json返回       |

### 1. 获取网站验证码api

- 访问方法

```http
GET /user/randCode
无需携带token
```
- 请求参数

| 请求参数 | 类型                          | 说明       |
| -------- | ----------------------------- | ---------- |
| phone    | `application/form-data`，必选 | 注册手机号 |

- 其他返回参数

  | 返回参数  | 说明                        |
  | --------- | --------------------------- |
  | rand_code | 有效时间为5分钟的五位随机码 |

- 返回实例

  ```json
  {
      "error": "",
      "rand_code": 92508,
      "status": "1"
  }
  ```

### 2. 注册api

- 访问方法

```http
POST /user/registerOrLoginByPhone
无需携带token
```

- 请求参数

| 请求参数  | 类型`Content-Type`          | 说明                                    |
| --------- | --------------------------- | --------------------------------------- |
| user_name | `multipart/form-data`，可选 | 昵称，格式要求英文大小写字母数字8到16位 |
| password  | `multipart/form-data`，必选 | 密码，格式要求英文大小写字母数字8到16位 |
| phone     | `multipart/form-data`，必选 | 手机号作为登录账号，格式要求11位数字    |
| answer    | `multipart/form-data`，可选 | 忘记密码时，需要的问题                  |
| question  | `multipart/form-data`，可选 | 忘记密码时，问题对应的答案              |

- 其他返回参数

| 其他返回参数 | 说明                                                      |
| ------------ | --------------------------------------------------------- |
| token        | 请求成功返回token字符串，12小时过期；请求失败返回空字符串 |

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IjEyMzQ1Njc4OTAxIiwiZXhwIjoxNjQ1MDYxNzY4LCJpYXQiOjE2NDUwMTg1NjgsImlzcyI6ImNvbGQgYmluIFx1MDAyNiB0YW8gcnVpIn0.tz6R3mjK9pwOM_4_WmRX51JrUSOkObBvy_rHmfDA3_k"
  }
  ```

### 3. 密码登录api

- 访问方法

```http
GET /user/loginByPwd
不需要token
```

- 请求参数

| 请求参数 | 类型                       | 说明                                    |
| -------- | -------------------------- | --------------------------------------- |
| phone    | `multipart/form-data`,必选 | 注册时的手机号，11位数字即可            |
| password | `multipart/form-data`,必选 | 密码，格式要求英文大小写字母数字8到16位 |

- 其他返回参数

| 返回参数 | 说明                                                      |
| -------- | --------------------------------------------------------- |
| token    | 请求成功返回token字符串，12小时过期；请求失败返回空字符串 |

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IjEyMzQ1Njc4OTAxIiwiZXhwIjoxNjQ1MDYxNzY4LCJpYXQiOjE2NDUwMTg1NjgsImlzcyI6ImNvbGQgYmluIFx1MDAyNiB0YW8gcnVpIn0.tz6R3mjK9pwOM_4_WmRX51JrUSOkObBvy_rHmfDA3_k"
  }
  ```

### 4. 修改密码api

- 访问方法

```http
PATCH  /user/alterPwd
```

- 请求参数

| 请求参数     | 类型                       | 说明                                      |
| ------------ | -------------------------- | ----------------------------------------- |
| phone        | `multipart/form-data`,必选 | 11位数字的手机号                          |
| old_password | `multipart/form-data`,必选 | 旧密码，格式要求英文大小写字母数字8到16位 |
| new_password | `multipart/form-data`,必选 | 新密码，格式要求英文大小写字母数字8到16位 |

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1"
  }
  ```

### 5. 获取用户设置的问题api

- 访问方法

```http
GET /user/question
```

- 请求参数

| 请求参数 | 类型 | 说明                                                         |
| -------- | ---- | ------------------------------------------------------------ |
| 无       | 无   | 该api通过中间件校验token,自动获取用户手机号的字段，从而查找对应用户设置的问题 |

- 其他返回参数

  | 返回参数 | 说明                         |
  | -------- | ---------------------------- |
  | question | 问题内容。未设置时，默认为无 |

- 返回实例

  ```json
  {
      "error": "",
      "question": "mylove",
      "status": "1"
  }
  ```

### 6. **获取用户问题对应答案api**

- 访问方法

```http
GET /user/answer
```

- 请求参数

| 请求参数 | 类型 | 说明                                                         |
| -------- | ---- | ------------------------------------------------------------ |
| 无       | 无   | 该api也是通过中间件自动获取用户手机号字段，从而查找对应用户设置的问题 |

- 其他返回参数

  | 返回参数 | 说明                         |
  | -------- | ---------------------------- |
  | answer   | 答案内容。未设置时，默认为无 |

- 返回实例

  ```json
  {
      "answer": "lxf",
      "error": "",
      "status": "1"
  }
  ```

### 7. **获取错误token的api**

- 访问方法

```http
GET /wrongToken
```

- 请求参数

| 请求参数 | 类型 | 说明 |
| -------- | ---- | ---- |
| 无       | 无   | 无   |

- 其他返回参数

| 返回参数 | 说明                                     |
| -------- | ---------------------------------------- |
| token    | 返回空token，退出登录时更换本地为此token |

- 返回实例

  ```json
  {
  		"status": "1",
  		"error":  "",
  		"token":  "this is a wrong token",
  }
  ```

### 8. 获取正确token的api

- 访问方法

```http
GET /token
无需携带token
```

- 请求参数

| 请求参数 | 类型                          | 说明                                    |
| -------- | ----------------------------- | --------------------------------------- |
| phone    | `application/form-data`，必选 | 注册手机号                              |
| password | `application/form-data`，必选 | 密码，格式要求英文大小写字母数字8到16位 |

- 其他返回参数

| 返回参数 | 说明                                                      |
| -------- | --------------------------------------------------------- |
| token    | 请求成功返回token字符串，12小时过期；请求失败返回空字符串 |

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwaG9uZSI6IjEyMzQ1Njc4OTAxIiwiZXhwIjoxNjQ1MDYxNzY4LCJpYXQiOjE2NDUwMTg1NjgsImlzcyI6ImNvbGQgYmluIFx1MDAyNiB0YW8gcnVpIn0.tz6R3mjK9pwOM_4_WmRX51JrUSOkObBvy_rHmfDA3_k"
  }
  ```

## 二、主页

### 1. 搜索api

- 访问方法

```http
GET    /search
```

- 请求参数

| 请求参数 | 类型                          | 说明                                                         |
| -------- | ----------------------------- | ------------------------------------------------------------ |
| search   | `application/form-data`，必选 | 正则字符串："[a-zA-Z0-9_\u4e00-\u9fa5]+"，只输入中英文数字或下划线 |

- 其他返回参数

| 返回参数 | 说明                                                         |
| -------- | ------------------------------------------------------------ |
| movies   | 请求成功返回参数如下多级json，会有多个电影数据，失败返回`""` |
| staff    | 请求成功返回,如下多级json，只有一个影人数据，失败返回`""`    |

- 返回实例

  ```json
  {
      "error": "",
      "movies": [
          {
              "id": 1,  //电影唯一标识
              "name": "奇迹·笨小孩 (2022)",//电影名字
              "picture": ".././douban/static/picture/mvpicture/奇迹·笨小孩 (2022).webp",//电影头像资源地址
              "director": " 文牧野",//导演
              "lead_role": "易烊千玺 / 田雨 / 陈哈琳 / 齐溪 / 公磊 / 许君聪 / 王宁 / 黄尧 / 巩金国 / 田壮壮 / 王传君 / 章宇 / 张志坚 / 咏梅 / 杨新鸣 / 徐峥 / 岳小军 / 朱俊麟 / 王丽涵 / 贾弘逍 / 韩笑 / 孙征宇 / 黄艺馨 / 修梦迪 / 苏子航 / 郑伊倩 / 丁文博 / 陈翊曈",//主演
              "produce_where": "中国大陆",//制片地区
              "release_time": "2022-02-01",//上映日期
              "duration": 106,//片长，默认单位位分钟
              "score": 0//总评分
          }
      ],
      "staff":
      {   "id": "3",//影人唯一标识
          "name": "田雨 Yu Tian",//影人姓名
          "avatar": ".././douban/static/picture/perfpicture/田雨 Yu Tian.jpg",//影人头像资源地址
          "jobs": "演员"//职业
      },
      "status": "1"
  }
  ```

### 2. 获取正在热映电影api**（上线30天内的电影）**

- 访问方法

```http
GET    /movie/hot
```

- 请求参数无

- 其他返回参数

| 返回参数 | 说明                                           |
| -------- | ---------------------------------------------- |
| movies   | 请求成功返回参数如下多级json，会有多个电影数据 |

- 返回实例

  ```json
  {
      "error": "",
      "movies": [
          {
              "id": 1,
              "name": "奇迹·笨小孩 (2022)",
              "picture": ".././douban/static/picture/mvpicture/奇迹·笨小孩 (2022).webp",
              "director": " 文牧野",
              "lead_role": "易烊千玺 / 田雨 / 陈哈琳 / 齐溪 / 公磊 / 许君聪 / 王宁 / 黄尧 / 巩金国 / 田壮壮 / 王传君 / 章宇 / 张志坚 / 咏梅 / 杨新鸣 / 徐峥 / 岳小军 / 朱俊麟 / 王丽涵 / 贾弘逍 / 韩笑 / 孙征宇 / 黄艺馨 / 修梦迪 / 苏子航 / 郑伊倩 / 丁文博 / 陈翊曈",
              "produce_where": "中国大陆",
              "release_time": "2022-02-01",
              "duration": 106,
              "score": 0
          },
          {
              "id": 2,
              "name": "长津湖之水门桥 (2022)",
              "picture": ".././douban/static/picture/mvpicture/长津湖之水门桥 (2022).webp",
              "director": "徐克",
              "lead_role": "吴京 / 易烊千玺 / 朱亚文 / 李晨 / 韩东君 / 张涵予 / 耿乐 / 杜淳 / 段奕宏 / 胡军 / 王丽坤 / 杨一威 / 李卓扬 / 何跃飞 / 唐志强 / 刘治威 / 庄小龙 / 辛玉波 / 张跃 / 许明虎 / 王宁 / 王振威 / 陈泽轩 / 李小锋 / 詹姆斯·菲尔伯德 / 约翰·克鲁兹",
              "produce_where": "中国大陆 / 中国香港",
              "release_time": "2022-02-01",
              "duration": 149,
              "score": 0
          }
      ],
      "status": "1"
  }
  ```

### 3. 获取即将上映电影api

- 访问方法

```http
GET    /movie/future
```

- 请求参数无

- 其他返回参数

| 返回参数 | 说明                                           |
| -------- | ---------------------------------------------- |
| movies   | 请求成功返回参数如下多级json，会有多个电影数据 |

- 返回实例

  ```json
  {
      "error": "",
      "movies": [
          {
              "id": 1,
              "name": "奇迹·笨小孩 (2022)",
              "picture": ".././douban/static/picture/mvpicture/奇迹·笨小孩 (2022).webp",
              "director": " 文牧野",
              "lead_role": "易烊千玺 / 田雨 / 陈哈琳 / 齐溪 / 公磊 / 许君聪 / 王宁 / 黄尧 / 巩金国 / 田壮壮 / 王传君 / 章宇 / 张志坚 / 咏梅 / 杨新鸣 / 徐峥 / 岳小军 / 朱俊麟 / 王丽涵 / 贾弘逍 / 韩笑 / 孙征宇 / 黄艺馨 / 修梦迪 / 苏子航 / 郑伊倩 / 丁文博 / 陈翊曈",
              "produce_where": "中国大陆",
              "release_time": "2022-02-01",
              "duration": 106,
              "score": 0
          },
          {
              "id": 2,
              "name": "长津湖之水门桥 (2022)",
              "picture": ".././douban/static/picture/mvpicture/长津湖之水门桥 (2022).webp",
              "director": "徐克",
              "lead_role": "吴京 / 易烊千玺 / 朱亚文 / 李晨 / 韩东君 / 张涵予 / 耿乐 / 杜淳 / 段奕宏 / 胡军 / 王丽坤 / 杨一威 / 李卓扬 / 何跃飞 / 唐志强 / 刘治威 / 庄小龙 / 辛玉波 / 张跃 / 许明虎 / 王宁 / 王振威 / 陈泽轩 / 李小锋 / 詹姆斯·菲尔伯德 / 约翰·克鲁兹",
              "produce_where": "中国大陆 / 中国香港",
              "release_time": "2022-02-01",
              "duration": 149,
              "score": 0
          }
      ],
      "status": "1"
  }
  ```

### 三、电影详情

### 1. 请求单个电影api

- 访问方法

```http
GET    /movie/:mv_id/info
mv_id指的是请求电影返回数据的id选项
```

- 请求参数无

- 其他返回参数

| 返回参数 | 说明                                           |
| -------- | ---------------------------------------------- |
| movie    | 请求成功返回参数如下多级json，会有多个电影数据 |

- 返回实例

  ```json
  {
      "error": "",
      "movie": {
          "id": 1,
          "form": "电影",//影视的形式
          "name": "奇迹·笨小孩 (2022)",
          "picture": ".././douban/static/picture/mvpicture/奇迹·笨小孩 (2022).webp",
          "director": " 文牧野",//导演
          "writer": "周楚岑 / 修梦迪 / 文牧野 / 韩晓邯 / 钟伟",//编剧
          "lead_role": "易烊千玺 / 田雨 / 陈哈琳 / 齐溪 / 公磊 / 许君聪 / 王宁 / 黄尧 / 巩金国 / 田壮壮 / 王传君 / 章宇 / 张志坚 / 咏梅 / 杨新鸣 / 徐峥 / 岳小军 / 朱俊麟 / 王丽涵 / 贾弘逍 / 韩笑 / 孙征宇 / 黄艺馨 / 修梦迪 / 苏子航 / 郑伊倩 / 丁文博 / 陈翊曈",
          "type": "剧情",//电影类型
          "special": "",//特色
          "produce_where": "中国大陆",
          "language": "汉语普通话",//语言
          "release_time": "2022-02-01",
          "duration": 106,
          "a_name": "奇迹 / 奇迹年代 / Nice View",//又名
          "imdb": "tt15783462",//imdb编号
          "plot_introduction": "二十岁的景浩（易烊千玺 饰）独自带着年幼的妹妹来到深圳生活，兄妹俩生活温馨却拮据。为了妹妹高昂的手术费，机缘巧合之下，景浩得到一个机会，本以为美好生活即将来临，却不料遭遇重创。在时间和金钱的双重压力下，毫无退路的景浩决定孤注一掷，而他陷入困境的平凡人生，又能否燃起希望的火花？\r\n　　电影《奇迹》是中宣部国家电影局2021年重点电影项目，也是2021年重点建党百年献礼片，描述十八大以后新时代年轻人在深圳创业的影片。",//剧情简介
          "one_star_num": 0,//表示评价1星人数
          "two_star_num": 0,//表示评价2星人数
          "three_star_num": 0,//表示评价3星人数
          "four_star_num": 0,//表示评价4星人数
          "five_star_num": 0//表示评价5星人数
      },
      "status": "1"
  }
  ```

### 2. 写顶级讨论api

- 访问方法

```http
POST   /movie/:mv_id/discuss
```

- 请求参数

| 请求参数 | 类型                          | 说明      |
| -------- | ----------------------------- | --------- |
| title    | `application/form-data`，必选 |           |
| content  | `application/form-data`，必选 | 字符<1000 |

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
  }
  ```

### 3. 写顶级讨论下的回复api

- 访问方法

```http
POST   /movie/:mv_id/:discuss_id/uDiscuss
```

- 请求参数

| 请求参数 | 类型                          | 说明           |
| -------- | ----------------------------- | -------------- |
| to_phone | `application/form-data`，必选 | 回复对象手机号 |
| content  | `application/form-data`，必选 | 字符<1000      |

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
  }
  ```

### 4. 获取顶级讨论的回复api    `BUG TODO`

- 访问方法

```http
GET    /movie/:mv_id/:discuss_id/uDiscuss
```

- 请求参数无

- 其他返回参数

| 返回参数      | 说明                                       |
| ------------- | ------------------------------------------ |
| under_discuss | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
      "under_discuss": {
          "0": {
              "id": 1,
              "from_phone": "15712344329",
              "to_phone": "15736469310",
              "from_user_name": "lxflxflxf",
              "from_avatar": ".././douban/static/picture/useravatar/默认头像",
              "from_mv_id": 1,
              "content": "回复1楼",
              "used_num": 0,
              "date_time": "2022-02-16 13:51:37",
              "label": "1"
          },
          "1": {
              "id": 1,
              "from_phone": "12345678901",
              "to_phone": "15736469310",
              "from_user_name": "lxblxblxb",
              "from_avatar": ".././douban/static/picture/useravatar/默认头像",
              "from_mv_id": 1,
              "content": "lxb",
              "used_num": 0,
              "date_time": "2022-02-16 17:49:11",
              "label": "1"
          }
      }
  }
  ```

### 5. 获取顶级讨论api

- 访问方法

```http
GET    /movie/:mv_id/discuss
```

- 请求参数无

- 其他返回参数无

- 返回实例

  ```json
  {
      "discusses": [
          {
              "id": 3,//唯一标识
              "from_phone": "15736469310",
              "from_user_name": "lhblhblhb",
              "from_avatar": ".././douban/static/picture/useravatar/屏幕截图 2021-06-14 142411.png",
              "from_mv_id": 1,
              "title": "1",
              "content": "回复1xf",
              "discuss_num": 1,//讨论数量
              "date_time": "2022-02-16 11:21:57"
          }
      ],
      "error": "",
      "status": "1"
  }
  ```

### 6. 删除顶级讨论及其回复讨论api

- 访问方法

```http
DELETE /movie/:mv_id/:discuss_id/discuss
```

- 请求参数无

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
  }
  ```

### 7. 删除讨论回复api

- 访问方法

```http
PATCH  /movie/:mv_id/:discuss_id/uDiscuss
```

- 请求参数无

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
  }
  ```

### 8. 写短评api

- 访问方法

```http
POST /comment/:mv_id/sComment
mv_id指的是请求电影返回数据的id选项
```

- 请求参数

| 请求参数 | 类型                          | 说明                                       |
| -------- | ----------------------------- | ------------------------------------------ |
| tag      | `application/form-data`，可选 | 标签，<50字符                              |
| content  | `application/form-data`，可选 | 内容，<350个字符                           |
| label    | `application/form-data`，可选 | 只能是0或1，0表示想看，1表示看过           |
| score    | `application/form-data`，可选 | 只能是0，1，2，3，4，5这几个整数，表示评分 |

- 其他返回参数无

- 返回实例

  ```json
  {	
      "status": "1",
  	"error":  "",
  }
  ```

### 9. 删除短评api

- 访问方法

```http
DELETE /comment/:mv_id/:comment_id/sComment
mv_id指的是请求电影返回数据的id选项，
comment_id指的是短评唯一标识id
```

- 请求参数无

- 其他返回参数无

- 返回实例

  ```json
  {	
      "status": "1",
  	"error":  "",
  }
  ```

### 10. 获取短评api

- 访问方法

```http
GET    /comment/:mv_id/sComment
mv_id指的是请求电影返回数据的id选项，
```

- 请求参数无

- 其他返回参数

| 返回参数       | 说明                                       |
| -------------- | ------------------------------------------ |
| short_comments | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "short_comments": [
          {
              "id": 2,//短评唯一标识
              "from_phone": "15736469310",//发表用户的唯一标识
              "from_user_name": "lhblhblhb",//发表用户的用户名
              "from_avatar": ".././douban/static/picture/useravatar/屏幕截图 2021-06-14 142411.png",
              "from_mv_id": 1,//评论电影的id
              "want_or_watched": "1",//想看或看过，1表示看过，0表示想看
              "mv_star": 5,//评分
              "tag": "tag",//标签
              "content": "短评",//内容
              "date_time": "2022-02-17 10:40:44",//发表日期
              "used_num": 0//赞成数
          }
      ],
      "status": "1"
  }
  ```

### 11. 写影评api

- 访问方法

```http
POST   /comment/:mv_id/lComment
```

- 请求参数

| 请求参数 | 类型                          | 说明 |
| -------- | ----------------------------- | ---- |
| title    | `application/form-data`，必选 |      |
| content  | `application/form-data`，必选 |      |

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1"
  }
  ```

### 12. 写回复影评api

- 访问方法

```http
POST   /comment/:mv_id/:comment_id/uLComment
```

- 请求参数

| 请求参数 | 类型                          | 说明           |
| -------- | ----------------------------- | -------------- |
| to_phone | `application/form-data`，必选 | 回复对象手机号 |
| content  | `application/form-data`，必选 | 内容           |

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1"
  }
  ```

### 13. 删除影评及其回复api

- 访问方法

```http
DELETE /comment/:mv_id/:comment_id/lComment
```

- 请求参数无

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1"
  }
  ```

### 14. 删除影评回复api

- 访问方法

```http
PATCH  /comment/:mv_id/:comment_id/uLComment
```

- 请求参数无

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1"
  }
  ```

### 15. 获取影评api

- 访问方法

```http
GET    /comment/:mv_id/lComment
```

- 请求参数无

- 其他返回参数

| 返回参数      | 说明                                       |
| ------------- | ------------------------------------------ |
| long_comments | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "long_comments": [
          {
              "id": 1,
              "from_phone": "15736469310",
              "from_user_name": "lhblhblhb",
              "from_avatar": ".././douban/static/picture/useravatar/屏幕截图 2021-06-14 142411.png",
              "from_mv_id": 1,
              "mv_star": 4,
              "title": "title",
              "content": "影评",
              "date_time": "2022-02-17 16:04:36",
              "used_num": 0,
              "unused_num": 0
          }
      ],
      "status": "1"
  }
  ```

### 16. 获取影评下的回复api  ` BUG TODO`

- 访问方法

```http
GET    /comment/:mv_id/:comment_id/uLComment
```

- 请求参数无

- 其他返回参数

| 返回参数            | 说明                                       |
| ------------------- | ------------------------------------------ |
| under_long_comments | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
      "under_long_comments": {
          //“0”键表示回复级数为顶级影评下的回复
          "0": {
              "id": 1,
              "from_phone": "12345678901",
              "to_phone": "15736469310",
              "from_user_name": "lxblxblxb",
              "from_avatar": ".././douban/static/picture/useravatar/默认头像",
              "from_mv_id": 1,
              "content": "回复顶楼",
              "date_time": "2022-02-17 18:03:32"
          },
          //“1”键表示回复级数为顶级影评下的回复的回复
          "1": {
              "id": 1,
              "from_phone": "15712344329",
              "to_phone": "12345678901",
              "from_user_name": "lxflxflxf",
              "from_avatar": ".././douban/static/picture/useravatar/默认头像",
              "from_mv_id": 1,
              "content": "回复2楼",
              "date_time": "2022-02-17 18:03:34"
          }
      }
  }
  ```

### 17. 获取影视的参演影人数据api

- 访问方法

```http
GET    /movie/:mv_id/staff
mv_id指的是请求电影返回数据的id选项
```

- 请求参数无

- 其他返回参数

| 返回参数 | 说明                                       |
| -------- | ------------------------------------------ |
| staffs   | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "staffs": [
          {
              "id": "1",
              "name": "文牧野 Muye Wen",
              "avatar": ".././douban/static/picture/perfpicture/文牧野 Muye Wen.jpg",
              "jobs": "导演/编剧/演员/剪辑"
          },
          {
              "id": "2",
              "name": "易烊千玺 Jackson Yee",
              "avatar": ".././douban/static/picture/perfpicture/易烊千玺 Jackson Yee.jpg",
              "jobs": "演员/配音/音乐"
          }
      ],
      "status": "1"
  }
  ```

## 四、影人

### 1. 获取单个影人属性api

- 访问方法

```http
GET    /staff/:staff_id/info
staff_id指的是影人数据的id选项
```

- 请求参数无

- 其他返回参数

| 返回参数 | 说明                                       |
| -------- | ------------------------------------------ |
| staff    | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "staff": {
          "id": "2",//影人身份标识
          "name": "易烊千玺 Jackson Yee",
          "sex": "男",
          "avatar": ".././douban/static/picture/perfpicture/易烊千玺 Jackson Yee.jpg",
          "constellation": "射手座",
          "birthday": "2000-11-28",
          "birthplace": "中国,湖南,怀化",
          "jobs": "演员/配音/音乐",
          "ac_name": "无",//更多中文名
          "ae_name": "Yi Yangqianxi / Jackson Yi / Jackson Yee",//更多英文名
          "family": "无",//影视家庭成员
          "imdb": "nm3725393",
          "introduction": "易烊千玺，2000年11月28日出生于湖南怀化，中国内地新生代男歌手、演员、舞者、TFBOYS成员。\r\n\r\n2005年，易烊千玺首登电视荧屏，开始参演各类综艺节目。\r\n\r\n2009年，加入“飞炫少年”组合，先后亮相CCTV和全国各大卫视，并接拍多部影视剧、广告、MV等。2011年底，退出“飞炫少年”组合。\r\n\r\n2013年，1月发行首支个人单曲《梦想摩天楼》。6月，因舞蹈出众、多才多艺获邀加入TF家族。8月6日，以TFBOYS组合成员形式出道，是组合中年纪最小的成员。\r\n\r\n2015年，为法国动画电影《小王子》男主角“小王子”配音。\r\n\r\n2016年，1月获得湖南卫视《全员加速中》第一季总冠军。2月，登上中央电视台春节联欢晚会。同年电视剧作品《超少年密码》、《青云志》、《小别离》播出。\r\n\r\n2017年，1月登上中央电视台春节联欢晚会。3月，首档常驻综艺节目《放开我北鼻》第二季播出。7月，电视剧作品《我们的少年时代》播出。\r\n\r\n2018年，2月《这！就是街舞》第一季播出，易烊千玺担任“易燃装置”战队队长，最终带领队伍卫冕冠军。4月，成为天猫首位代言人。7月确认出演金马班底校园暴力题材电影《少年的你》男主角小北。9月，以双料第一的优异成绩进入中央戏剧学院2018级话剧影视表演系本科班就读。\r\n\r\n2019年，年初主演的首部电影《少年的你》入围柏林电影节新生代单元。1月，由易烊千玺担任“金牌经理人”，携手国家体育总局推出的冰雪成长秀综艺节目《大冰小将》播出。3月，确认以队长身份继续加盟《这！就是街舞》第二季。 6月27号，易烊千玺主演的《长安十二时辰》热播。10月25号《少年的你》上映反响轰动，最终取得15.58亿票房的成绩，口碑票房双赢。\r\n\r\n2020年1月，特别出演的《热血同行》播出。并于1月凭《少年的你》提名香港电影评论学会大奖最佳男演员。2月12号，提名第39届香港金像奖最佳男演员和最佳新人。"
      },
      "status": "0"
  }
  ```

### 2. 获取影人参演电影的数据api

**注意：staffs为json键，获取对应staff_id对应演员参演的影视作品数据，本来该改成movies的，都部署上去了，就算了，将就用吧**

- 访问方法

```http
GET    /staff/:staff_id/movie
```

- 请求参数无
- 其他返回参数

| 返回参数 | 说明                                       |
| -------- | ------------------------------------------ |
| staffs   | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "staffs": [
          {
              "id": 1,
              "name": "奇迹·笨小孩 (2022)",
              "picture": ".././douban/static/picture/mvpicture/奇迹·笨小孩 (2022).webp",
              "director": " 文牧野",
              "lead_role": "易烊千玺 / 田雨 / 陈哈琳 / 齐溪 / 公磊 / 许君聪 / 王宁 / 黄尧 / 巩金国 / 田壮壮 / 王传君 / 章宇 / 张志坚 / 咏梅 / 杨新鸣 / 徐峥 / 岳小军 / 朱俊麟 / 王丽涵 / 贾弘逍 / 韩笑 / 孙征宇 / 黄艺馨 / 修梦迪 / 苏子航 / 郑伊倩 / 丁文博 / 陈翊曈",
              "produce_where": "中国大陆",
              "release_time": "2022-02-01",
              "duration": 106,
              "score": 0
          },
          {
              "id": 2,
              "name": "长津湖之水门桥 (2022)",
              "picture": ".././douban/static/picture/mvpicture/长津湖之水门桥 (2022).webp",
              "director": "徐克",
              "lead_role": "吴京 / 易烊千玺 / 朱亚文 / 李晨 / 韩东君 / 张涵予 / 耿乐 / 杜淳 / 段奕宏 / 胡军 / 王丽坤 / 杨一威 / 李卓扬 / 何跃飞 / 唐志强 / 刘治威 / 庄小龙 / 辛玉波 / 张跃 / 许明虎 / 王宁 / 王振威 / 陈泽轩 / 李小锋 / 詹姆斯·菲尔伯德 / 约翰·克鲁兹",
              "produce_where": "中国大陆 / 中国香港",
              "release_time": "2022-02-01",
              "duration": 149,
              "score": 0
          }
      ],
      "status": "1"
  }
  ```

## 五、获取排行榜api

- 访问方法

```http
GET    /rank?type_name=

根据url可以拼接的参数如下：
剧情 喜剧 动作 爱情 科幻 动画 悬疑 惊悚 恐怖 纪录片 短片 情色 同性 音乐 歌舞 家庭 儿童 传记 历史 战争 犯罪 西部 奇幻 冒险 灾难 武侠 古装 运动 黑色电影
```

- 请求参数无

- 其他返回参数

| 返回参数 | 说明                                       |
| -------- | ------------------------------------------ |
| movies   | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "movies": [
          {
              "id": 1,
              "name": "奇迹·笨小孩 (2022)",
              "picture": ".././douban/static/picture/mvpicture/奇迹·笨小孩 (2022).webp",
              "director": " 文牧野",
              "lead_role": "易烊千玺 / 田雨 / 陈哈琳 / 齐溪 / 公磊 / 许君聪 / 王宁 / 黄尧 / 巩金国 / 田壮壮 / 王传君 / 章宇 / 张志坚 / 咏梅 / 杨新鸣 / 徐峥 / 岳小军 / 朱俊麟 / 王丽涵 / 贾弘逍 / 韩笑 / 孙征宇 / 黄艺馨 / 修梦迪 / 苏子航 / 郑伊倩 / 丁文博 / 陈翊曈",
              "produce_where": "中国大陆",
              "release_time": "2022-02-01",
              "duration": 106,
              "score": 5
          },
          {
              "id": 2,
              "name": "长津湖之水门桥 (2022)",
              "picture": ".././douban/static/picture/mvpicture/长津湖之水门桥 (2022).webp",
              "director": "徐克",
              "lead_role": "吴京 / 易烊千玺 / 朱亚文 / 李晨 / 韩东君 / 张涵予 / 耿乐 / 杜淳 / 段奕宏 / 胡军 / 王丽坤 / 杨一威 / 李卓扬 / 何跃飞 / 唐志强 / 刘治威 / 庄小龙 / 辛玉波 / 张跃 / 许明虎 / 王宁 / 王振威 / 陈泽轩 / 李小锋 / 詹姆斯·菲尔伯德 / 约翰·克鲁兹",
              "produce_where": "中国大陆 / 中国香港",
              "release_time": "2022-02-01",
              "duration": 149,
              "score": 0
          }
      ],
      "status": "1"
  }
  ```

## 六、分类找电影api

- 访问方法

```http
/type?kind=&form=&place=&age=&special=

根据url可以拼接的参数如下：
kind:
剧情 喜剧 动作 爱情 科幻 动画 悬疑 惊悚 恐怖 犯罪 同性 音乐 歌舞 传记 历史 战争 西部 奇幻 冒险 灾难 武侠 情色
form:电影 电视剧 综艺 动漫 纪录片 短片
place:中国大陆 欧美 美国 中国香港 中国台湾 日本 韩国 英国 法国 德国 意大利 西班牙 印度 泰国 俄罗斯 伊朗 加拿大 澳大利 亚爱尔兰 瑞典 巴西 丹麦
age:2022 2021 2020 2019 2010 2000 1990 1980 1970 1960注：一定要整年数（四位数字） 
special:经典 青春 文艺 搞笑 励志 魔幻 感人 
```

- 请求参数无

- 其他返回参数

| 返回参数 | 说明                                       |
| -------- | ------------------------------------------ |
| movies   | 请求成功时返回如下多级json数据，键、值如下 |

- 返回实例

  ```json
  {
      "error": "",
      "movies": [
          {
              "id": 1,
              "name": "奇迹·笨小孩 (2022)",
              "picture": ".././douban/static/picture/mvpicture/奇迹·笨小孩 (2022).webp",
              "director": " 文牧野",
              "lead_role": "易烊千玺 / 田雨 / 陈哈琳 / 齐溪 / 公磊 / 许君聪 / 王宁 / 黄尧 / 巩金国 / 田壮壮 / 王传君 / 章宇 / 张志坚 / 咏梅 / 杨新鸣 / 徐峥 / 岳小军 / 朱俊麟 / 王丽涵 / 贾弘逍 / 韩笑 / 孙征宇 / 黄艺馨 / 修梦迪 / 苏子航 / 郑伊倩 / 丁文博 / 陈翊曈",
              "produce_where": "中国大陆",
              "release_time": "2022-02-01",
              "duration": 106,
              "score": 5
          },
          {
              "id": 2,
              "name": "长津湖之水门桥 (2022)",
              "picture": ".././douban/static/picture/mvpicture/长津湖之水门桥 (2022).webp",
              "director": "徐克",
              "lead_role": "吴京 / 易烊千玺 / 朱亚文 / 李晨 / 韩东君 / 张涵予 / 耿乐 / 杜淳 / 段奕宏 / 胡军 / 王丽坤 / 杨一威 / 李卓扬 / 何跃飞 / 唐志强 / 刘治威 / 庄小龙 / 辛玉波 / 张跃 / 许明虎 / 王宁 / 王振威 / 陈泽轩 / 李小锋 / 詹姆斯·菲尔伯德 / 约翰·克鲁兹",
              "produce_where": "中国大陆 / 中国香港",
              "release_time": "2022-02-01",
              "duration": 149,
              "score": 0
          }
      ],
      "status": "1"
  }
  ```

## 七、个人页面

### 1. 上传头像api

- 访问方法

```http
POST   /user/avatar
```

- 请求参数

| 请求参数    | 类型                          | 说明     |
| ----------- | ----------------------------- | -------- |
| user_avatar | `application/form-data`，必选 | 图片文件 |

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1"
  }
  ```

### 2. 编写自我介绍api

- 访问方法

```http
POST   /user/introduction
```

- 请求参数

| 请求参数          | 类型`Content-Type`          | 说明                    |
| ----------------- | --------------------------- | ----------------------- |
| user_introduction | `multipart/form-data`，必选 | 用户自我介绍，字符<1000 |

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
  }
  ```

### 3. 编写用户签名api

- 访问方法

```http
POST   /user/introduction
```

- 请求参数

| 请求参数  | 类型`Content-Type`          | 说明               |
| --------- | --------------------------- | ------------------ |
| user_sign | `multipart/form-data`，必选 | 用户签名，字符<200 |

- 其他返回参数无

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
  }
  ```

### 4. 获取用户想看和看过的api

- 访问方法

```http
GET    /user/movie/:label
label：0->想看，1->看过，label只能为0或1
```

- 请求参数无

- 其他返回参数

| 返回参数 | 说明                                                         |
| -------- | ------------------------------------------------------------ |
| movies   | 请求成功返回参数如下多级json，会有多个电影数据，失败返回`""` |

- 返回实例

  ```json
  {
      "error": "",
      "movies": [
          {
              "id": 1,
              "name": "奇迹·笨小孩 (2022)",
              "picture": ".././douban/static/picture/mvpicture/奇迹·笨小孩 (2022).webp",
              "director": " 文牧野",
              "lead_role": "易烊千玺 / 田雨 / 陈哈琳 / 齐溪 / 公磊 / 许君聪 / 王宁 / 黄尧 / 巩金国 / 田壮壮 / 王传君 / 章宇 / 张志坚 / 咏梅 / 杨新鸣 / 徐峥 / 岳小军 / 朱俊麟 / 王丽涵 / 贾弘逍 / 韩笑 / 孙征宇 / 黄艺馨 / 修梦迪 / 苏子航 / 郑伊倩 / 丁文博 / 陈翊曈",
              "produce_where": "中国大陆",
              "release_time": "2022-02-01",
              "duration": 106,
              "score": 5
          }
      ],
      "status": "1"
  }
  ```

### 5. 创建用户想看和看过的api 

即写短评api

### 6. 获取用户的影评api

- 访问方法

```http
GET /user/lComment
```

- 请求参数无

- 其他返回参数

| 返回参数      | 说明                                                         |
| ------------- | ------------------------------------------------------------ |
| long_comments | 请求成功返回参数如下多级json，会有多个电影数据，失败返回`""` |

- 返回实例

  ```json
  {
      "error": "",
      "long_comments": [
          {
              "id": 1,
              "from_phone": "15736469310",
              "from_user_name": "lhblhblhb",
              "from_avatar": ".././douban/static/picture/mvpicture/屏幕截图 2021-06-14 142411.png",
              "from_mv_id": 1,
              "mv_star": 4,
              "title": "title",
              "content": "影评",
              "date_time": "2022-02-17 16:04:36",
              "used_num": 0,
              "unused_num": 0
          }
      ],
      "status": "1"
  }
  ```

### 7. 获取自己的个人页面信息api

- 访问方法

```http
GET    /user/info
```

- 请求参数无

- 其他返回参数

| 返回参数  | 说明                                                         |
| --------- | ------------------------------------------------------------ |
| user_info | 请求成功返回参数如下多级json，会有多个电影数据，失败返回`""` |

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
      "user_info": {
          "user_id": "1",
          "phone": "15736469310",
          "avatar": ".././douban/static/picture/useravatar/屏幕截图 2021-06-14 142411.png",
          "user_name": "lhblhblhb",
          "user_introduction": "我是lhb，我来自火星",
          "user_sign": "人生如茶，久浓醇香",
          "register_time": "2022-02-14"
      }
  }
  ```

### 8. 获取其他人页面信息api

- 访问方法

```http
GET    /user/info
```

- 请求参数

| 请求参数 | 类型                          | 说明                     |
| -------- | ----------------------------- | ------------------------ |
| phone    | `application/form-data`，必选 | 查看用户信息对应的手机号 |

- 其他返回参数

| 返回参数  | 说明                                                         |
| --------- | ------------------------------------------------------------ |
| user_info | 请求成功返回参数如下多级json，会有多个电影数据，失败返回`""` |

- 返回实例

  ```json
  {
      "error": "",
      "status": "1",
      "user_info": {
          "user_id": "1",
          "phone": "15736469310",
          "avatar": ".././douban/static/picture/useravatar/屏幕截图 2021-06-14 142411.png",
          "user_name": "lhblhblhb",
          "user_introduction": "我是lhb，我来自火星",
          "user_sign": "人生如茶，久浓醇香",
          "register_time": "2022-02-14"
      }
  }
  ```
