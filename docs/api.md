# 博客系统 API 文档

Base URL（开发环境）：`http://localhost:8080/api`

## 通用约定

- **请求/响应格式**：JSON
- **认证方式**：JWT
  - 登录成功后返回 `token`
  - 需要认证的接口必须在请求头中带上：

    ```http
    Authorization: Bearer <token>
    ```

- **错误返回格式（约定）**：

  ```json
  { "error": "错误信息" }
  ```

---

## 1. 用户与认证

### 1.1 注册

- **URL**：`POST /user/register`
- **认证**：不需要
- **请求体**：

  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

- **成功响应（200）**：

  ```json
  { "message": "register success" }
  ```

- **失败响应**：
  - 400：参数错误
  - 500：用户名已存在或其他错误（`error` 字段说明原因）

---

### 1.2 登录

- **URL**：`POST /user/login`
- **认证**：不需要
- **请求体**：

  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

- **成功响应（200）**：

  ```json
  {
    "token": "jwt-token-string",
    "user_id": 1,
    "username": "yourname"
  }
  ```

- **失败响应**：
  - 400：参数错误
  - 401：用户名或密码错误
  - 500：生成 token 失败

> 当前代码中 token 有效期约为 **48 小时**。

---

## 2. 文章（Post）

`model.Post` 结构（响应中字段）：

```json
{
  "id": 1,
  "title": "标题",
  "content": "正文",
  "author_id": 1,
  "category_id": 1,
  "status": 1,
  "cover_image": "",
  "views": 0,
  "likes": 0,
  "is_top": 0,
  "created_at": "2025-11-15T20:05:48Z",
  "updated_at": "2025-11-15T20:05:48Z"
}
```

### 2.1 创建文章

- **URL**：`POST /posts`
- **认证**：需要（任意登录用户）
- **请求头**：

  ```http
  Authorization: Bearer <token>
  ```

- **请求体**：

  ```json
  {
    "title": "string",
    "content": "string",
    "category_id": 1,
    "status": 1
  }
  ```

  > `author_id` 不需要传，后端根据 token 中的 `user_id` 设置作者。

- **成功响应（200）**：

  ```json
  {
    "message": "create post success",
    "id": 1
  }
  ```

- **失败响应**：
  - 400：参数错误
  - 401：未登录或 token 无效
  - 500：数据库错误

---

### 2.2 获取文章详情

- **URL**：`GET /posts/:id`
- **认证**：不需要
- **成功响应（200）**：
  - 返回单个 `Post` 对象
- **失败响应**：
  - 400：`id` 非法
  - 404：文章不存在
  - 500：其他错误

---

### 2.3 文章列表（分页）

- **URL**：`GET /posts`
- **认证**：不需要
- **查询参数**：
  - `page`：页码，默认 `1`
  - `page_size`：每页数量，默认 `10`

- **成功响应（200）**：

  ```json
  {
    "list": [ /* Post 数组 */ ],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
  ```

- **失败响应**：
  - 500：数据库错误

---

### 2.4 更新文章

- **URL**：`PUT /posts/:id`
- **认证**：需要
  - 只允许 **文章作者本人** 或 **管理员 (role == 1)**

- **请求体**：

  ```json
  {
    "title": "new title",
    "content": "new content",
    "category_id": 1,
    "status": 1
  }
  ```

- **成功响应（200）**：

  ```json
  { "message": "update post success" }
  ```

- **失败响应**：
  - 400：`id` 或 body 非法
  - 401：未登录
  - 403：无权限（不是作者且不是管理员）
  - 404：文章不存在
  - 500：数据库错误

---

### 2.5 删除文章

- **URL**：`DELETE /posts/:id`
- **认证**：需要
  - 只允许 **文章作者** 或 **管理员**

- **成功响应（200）**：

  ```json
  { "message": "delete post success" }
  ```

- **失败响应**：
  - 400：`id` 非法
  - 401：未登录
  - 403：无权限
  - 404：文章不存在
  - 500：数据库错误

---

## 3. 分类（Category）

`model.Category` 结构：

```json
{
  "id": 1,
  "name": "Golang",
  "slug": "golang",
  "description": "描述",
  "created_at": "2025-11-15T...",
  "updated_at": "2025-11-15T..."
}
```

### 3.1 创建分类（仅管理员）

- **URL**：`POST /categories`
- **认证**：需要
  - 仅 **管理员 (role == 1)**

- **请求体**：

  ```json
  {
    "name": "string",
    "slug": "string",
    "description": "string"
  }
  ```

- **成功响应（200）**：

  ```json
  {
    "message": "create category success",
    "id": 1
  }
  ```

- **失败响应**：
  - 400：参数错误
  - 401：未登录
  - 403：非管理员
  - 500：数据库错误

---

### 3.2 更新分类（仅管理员）

- **URL**：`PUT /categories/:id`
- **认证**：需要管理员

- **请求体**：同创建

- **成功响应（200）**：

  ```json
  { "message": "update category success" }
  ```

- **失败响应**：
  - 400：参数错误
  - 401：未登录
  - 403：非管理员
  - 404：分类不存在
  - 500：数据库错误

---

### 3.3 删除分类（仅管理员）

- **URL**：`DELETE /categories/:id`
- **认证**：需要管理员

- **成功响应（200）**：

  ```json
  { "message": "delete category success" }
  ```

- **失败响应**：
  - 400：`id` 非法
  - 401：未登录
  - 403：非管理员
  - 404：`{ "error": "category not found" }`
  - 500：其他错误

---

### 3.4 获取单个分类

- **URL**：`GET /categories/:id`
- **认证**：不需要
- **成功响应（200）**：返回 Category 对象
- **失败响应**：
  - 400：`id` 非法
  - 404：`{ "error": "category not found" }`
  - 500：其他错误

---

### 3.5 分类列表

- **URL**：`GET /categories`
- **认证**：不需要

- **成功响应（200）**：

  ```json
  {
    "list": [ /* Category 数组 */ ]
  }
  ```

- **失败响应**：
  - 500：数据库错误

---

## 4. 评论（Comment）

`model.Comment` 结构：

```json
{
  "id": 1,
  "user_id": 2,
  "post_id": 1,
  "content": "评论内容",
  "parent_id": 0,
  "created_at": "2025-11-15T...",
  "updated_at": "2025-11-15T..."
}
```

### 4.1 发表评论

- **URL**：`POST /posts/:id/comments`
- **认证**：需要（任意登录用户）
- **含义**：对文章 `:id` 发表评论

- **请求体**：

  ```json
  {
    "content": "评论内容",
    "parent_id": 0
  }
  ```

  > 评论者 ID `user_id` 从 token 中获取，无需前端传。

- **成功响应（200）**：

  ```json
  {
    "message": "create comment success",
    "id": 1
  }
  ```

- **失败响应**：
  - 400：文章 ID 或 body 非法
  - 401：未登录
  - 500：数据库错误

---

### 4.2 获取文章下评论列表

- **URL**：`GET /posts/:id/comments`
- **认证**：不需要

- **成功响应（200）**：

  ```json
  {
    "list": [ /* Comment 数组 */ ]
  }
  ```

- **失败响应**：
  - 400：文章 ID 非法
  - 500：数据库错误

---

### 4.3 删除评论

- **URL**：`DELETE /comments/:id`
- **认证**：需要
  - 只允许 **评论作者** 或 **管理员 (role == 1)**

- **成功响应（200）**：

  ```json
  { "message": "delete comment success" }
  ```

- **失败响应**：
  - 400：评论 ID 非法
  - 401：未登录
  - 403：无权限
  - 404：`{ "error": "comment not found" }`
  - 500：其他错误
