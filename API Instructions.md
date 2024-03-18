# 接口文档

## I.管理员端

****

### 一、分类的接口

#### 1. 新增分类

**(1) URL**

> **admin/category** 

**(2) 方法: POST**

**(3) 请求参数：JSON数据**

```json
{
  "name": "个人成长" // 必填
}
```

**(4) 返回参数：JSON数据**

```json
{
	"data": {
		"id": 8,
		"name": "个人成长"
	},
	"message": "OK",
	"status": 200
}
```

#### 2. 获取分类信息

**(1) URL**

> **admin/category/:id**  

**(2) 方法: GET**

**(3) 请求参数：无**

**(4) 返回参数：JSON数据**

```json
{
	"data": {
		"id": 2,
		"name": "前端开发"
	},
	"message": "OK",
	"status": 200
}
```

#### 3. 获取所有分类列表

**(1) URL**

> **admin/categories**

**(2) 方法: GET**

**(3) 请求参数：Query查询参数**

| 参数名      | 类型  | 说明  |
| -------- | --- | --- |
| pagesize | int | 选填  |
| pagenum  | int | 选填  |

**(4) 返回参数：JSON数据**

```json
{
	"data": [
		{
			"id": 3,
			"name": "后端开发"
		},
		{
			"id": 4,
			"name": "机器学习"
		}
	],
	"message": "OK",
	"status": 200,
	"total": 6
}
```

#### 4. 更新分类信息

**(1) URL**

> **admin/category/:id**

**(2) 方法: PUT**

**(3) 请求参数：JSON数据**

```json
{
    "name": "操作系统"
}
```

**(4) 返回参数：JSON数据**

```json
{
    "data": {
        "id": 8,
        "name": "操作系统"
    },
    "message": "OK",
    "status": 200
}
```

#### 5. 删除分类

**(1) URL**

> admin/category/:id  

**(2) 方法: DELETE**

**(3) 请求参数：无**

**(4) 返回参数：JSON数据**

```json
{
	"message": "OK",
	"status": 200
}
```

****

### 二、文章的接口

#### 1. 新增文章

(1) URL: admin/article
(2) 方法: POST

(3) 请求参数：JSON数据

```json
{
  "title": "如何管理时间",
  "desc": "有效的时间管理技巧和压力缓解方法。",
  "content": "时间和压力是现代生活中的两大挑战，学会管理它们可以带来更高的生活质量...",
  "category_id": 5,
  "user_id": 3
}
```

(4) 返回参数：JSON数据

```json
{
	"data": {
		"id": 19,
		"created_at": "2023-09-25T20:47:15.03+08:00",
		"updated_at": "2023-09-25T20:47:15.03+08:00",
		"title": "如何管理时间",
		"desc": "有效的时间管理技巧和压力缓解方法。",
		"content": "时间和压力是现代生活中的两大挑战，学会管理它们可以带来更高的生活质量...",
		"img": "",
		"comment_count": 0,
		"read_count": 0,
		"category": {
			"category_id": 5,
			"category_name": "个人成长"
		},
		"author": {
			"user_id": 3,
			"full_name": "Mike Smith"
		}
	},
	"message": "OK",
	"status": 200
}
```

#### 2. 根据用户获取文章列表

(1) URL: admin/articles/user/:id  
(2) 方法: GET

(3) 请求参数：Query查询参数

| 参数名      | 类型  | 说明  |
| -------- | --- | --- |
| pagesize | int | 选填  |
| pagenum  | int | 选填  |

(4) 返回参数：JSON数据

```json
{
	"data": [
		{
			"id": 1,
			"created_at": "2023-09-24T15:08:54.844+08:00",
			"updated_at": "2023-09-24T15:08:54.844+08:00",
			"title": "深入浅出Python编程",
			"desc": "一篇介绍Python的核心概念和应用的文章。",
			"content": "Python是一种广泛使用的高级编程语言，著名的简洁明了...",
			"img": "",
			"comment_count": 0,
			"read_count": 0,
			"category": {
				"category_id": 1,
				"category_name": "编程语言"
			},
			"author": {
				"user_id": 1,
				"full_name": ""
			}
		},
		{
			"id": 2,
			"created_at": "2023-09-24T15:09:05.334+08:00",
			"updated_at": "2023-09-24T15:09:05.334+08:00",
			"title": "Java与C++的对比分析",
			"desc": "分析Java和C++两种流行语言的特点和适用场景。",
			"content": "Java和C++都是强大的编程语言，但在某些方面它们有所不同...",
			"img": "",
			"comment_count": 0,
			"read_count": 5,
			"category": {
				"category_id": 1,
				"category_name": "编程语言"
			},
			"author": {
				"user_id": 1,
				"full_name": ""
			}
		}
	],
	"message": "OK",
	"status": 200,
	"total": 3
}
```

#### 3. 根据分类获取文章列表

(1) URL: admin/articles/category/:id  
(2) 方法: GET

(3) 请求参数：Query查询参数

| 参数名      | 类型  | 说明  |
| -------- | --- | --- |
| pagesize | int | 选填  |
| pagenum  | int | 选填  |

(4) 返回参数：JSON数据

```json
{
	"data": [
		{
			"id": 1,
			"created_at": "2023-09-24T15:08:54.844+08:00",
			"updated_at": "2023-09-24T15:08:54.844+08:00",
			"title": "深入浅出Python编程",
			"desc": "一篇介绍Python的核心概念和应用的文章。",
			"content": "Python是一种广泛使用的高级编程语言，著名的简洁明了...",
			"img": "",
			"comment_count": 0,
			"read_count": 0,
			"category": {
				"category_id": 1,
				"category_name": "编程语言"
			},
			"author": {
				"user_id": 1,
				"full_name": "John Doe"
			}
		}
	],
	"message": "OK",
	"status": 200,
	"total": 2
}
```

#### 4. 获取文章信息

(1) URL: admin/article/:id  
(2) 方法: GET

(3) 请求参数：无

(4) 返回参数：JSON数据

```json
{
	"data": {
		"id": 2,
		"created_at": "2023-09-24T15:09:05.334+08:00",
		"updated_at": "2023-09-24T15:09:05.334+08:00",
		"title": "Java与C++的对比分析",
		"desc": "分析Java和C++两种流行语言的特点和适用场景。",
		"content": "Java和C++都是强大的编程语言，但在某些方面它们有所不同...",
		"img": "",
		"comment_count": 0,
		"read_count": 5,
		"category": {
			"category_id": 1,
			"category_name": "编程语言"
		},
		"author": {
			"user_id": 1,
			"full_name": "John Doe"
		}
	},
	"message": "OK",
	"status": 200
}
```

#### 5. 获取所有文章列表

(1) URL: admin/articles  
(2) 方法: GET

(3) 请求参数：Query查询参数

| 参数名      | 类型     | 说明         |
| -------- | ------ | ---------- |
| pagesize | int    | 选填         |
| pagenum  | int    | 选填         |
| title    | string | 选填(模糊查询标题) |

(4) 返回参数：JSON数据

`localhost:9000/admin/articles?pagesize=5&pagenum=1&title=学习`

```json
{
	"data": [
		{
			"id": 7,
			"created_at": "2023-09-24T15:09:46.503+08:00",
			"updated_at": "2023-09-24T15:09:46.503+08:00",
			"title": "机器学习入门指南",
			"desc": "从零开始的机器学习学习路径。",
			"content": "机器学习是AI领域的一个重要分支，它的应用已经渗透到我们生活的各个方面...",
			"img": "",
			"comment_count": 0,
			"read_count": 0,
			"category": {
				"category_id": 4,
				"category_name": "机器学习"
			},
			"author": {
				"user_id": 4,
				"full_name": "Emily Jones"
			}
		}
	],
	"meesage": "OK",
	"status": 200,
	"total": 1
}
```

#### 6. 更新文章信息

(1) URL: admin/article/:id  
(2) 方法: PUT

(3) 请求参数：JSON数据

```json
{
  "title": "深入理解数据库优化3",
  "desc": "从基础到高级，探索数据库优化的各个方面。",
  "content": "数据库是后端开发的核心，其性能直接影响到整个应用的响应速度...",
  "category_id": 3
}
```

(4) 返回参数：JSON数据

```json
{
	"data": {
		"id": 14,
		"created_at": "2023-09-25T19:44:08.262+08:00",
		"updated_at": "2023-09-25T20:11:30.67+08:00",
		"title": "深入理解数据库优化2",
		"desc": "从基础到高级，探索数据库优化的各个方面。",
		"content": "数据库是后端开发的核心，其性能直接影响到整个应用的响应速度...",
		"img": "",
		"comment_count": 0,
		"read_count": 0,
		"category": {
			"category_id": 3,
			"category_name": "后端开发"
		},
		"author": {
			"user_id": 3,
			"full_name": "Mike Smith"
		}
	},
	"message": "OK",
	"status": 200
}
```

#### 7. 删除文章

(1) URL: admin/article/:id  
(2) 方法: DELETE

(3) 请求参数：无

(4) 返回参数：JSON数据

`localhost:9000/admin/article/14`

```json
{
	"message": "OK",
	"status": 200
}
```

### 三、用户的接口

#### 1. 新增用户

(1) URL: admin/user  
(2) 方法: POST



#### 2. 获取所有用户列表

(1) URL: admin/users  
(2) 方法: GET

#### 3. 更新用户基本信息

(1) URL: admin/user/:id  
(2) 方法: PUT

#### 4. 删除用户

(1) URL: admin/user/:id  
(2) 方法: DELETE

### 四、评论的接口

#### 1. 新增评论

(1) URL: admin/comment  
(2) 方法: POST

#### 2. 根据文章ID获取评论

(1) URL: admin/comment/article/:id  
(2) 方法: GET

#### 3. 删除评论

(1) URL: admin/comment/:id  
(2) 方法: DELETE
