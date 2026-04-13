# 电子元件仓库管理系统

一个用于管理电子元件的仓库管理系统，支持元件的入库、出库、导入导出等功能。

## 技术栈

### 后端
- Go 1.21+
- Gin (Web 框架)
- GORM (ORM 框架)
- MySQL (数据库)
- Viper (配置管理)
- Zap (日志库)
- Excelize (Excel 处理)

### 前端
- Vue 3
- Vite
- Element Plus (UI 组件库)
- Axios (HTTP 客户端)
- Vue Router (路由)

## 功能特性

- ✅ 电子元件信息管理（增删改查）
- ✅ 入库操作
- ✅ 出库操作
- ✅ Excel 文件导入
- ✅ Excel 文件导出
- ✅ 库存实时监控
- ✅ 出入库记录追踪
- ✅ 搜索和分页功能
- ✅ 简洁高效的用户界面

## 快速开始

### 环境要求

- Go 1.21+
- Node.js 18+
- MySQL 5.7+

### 1. 数据库配置

创建数据库并执行初始化脚本：

```bash
mysql -u root -p < backend/schema.sql
```

或者手动创建数据库 `cms`，然后执行 `backend/schema.sql` 中的 SQL 语句。

### 2. 后端启动

```bash
cd backend

# 安装依赖
go mod tidy

# 修改配置文件 config.yaml 中的数据库连接信息

# 运行
go run cmd/main.go
```

后端服务将在 `http://localhost:8080` 启动

### 3. 前端启动

```bash
cd frontend

# 安装依赖
npm install

# 运行开发服务器
npm run dev
```

前端服务将在 `http://localhost:3000` 启动

## 项目结构

```
cms/
├── backend/
│   ├── cmd/
│   │   └── main.go              # 应用入口
│   ├── internal/
│   │   ├── config/              # 配置管理
│   │   ├── database/            # 数据库连接
│   │   ├── handlers/            # HTTP 处理器
│   │   ├── middleware/          # 中间件
│   │   ├── models/              # 数据模型
│   │   └── service/             # 业务逻辑
│   ├── pkg/
│   │   └── utils/               # 工具函数
│   ├── storage/                 # 文件存储
│   ├── uploads/                 # 上传文件
│   ├── config.yaml              # 配置文件
│   ├── go.mod                   # Go 模块定义
│   └── schema.sql               # 数据库初始化脚本
└── frontend/
    ├── src/
    │   ├── api/                 # API 接口
    │   ├── components/          # 组件
    │   ├── router/              # 路由配置
    │   ├── utils/               # 工具函数
    │   ├── views/               # 页面视图
    │   ├── App.vue              # 根组件
    │   └── main.js              # 应用入口
    ├── index.html
    ├── package.json
    └── vite.config.js
```

## API 接口

### 元件管理

- `GET /api/components` - 获取元件列表
- `GET /api/components/:id` - 获取单个元件
- `POST /api/components` - 创建元件
- `PUT /api/components/:id` - 更新元件
- `DELETE /api/components/:id` - 删除元件

### 出入库管理

- `POST /api/components/:id/stock-in` - 入库
- `POST /api/components/:id/stock-out` - 出库
- `GET /api/components/:id/records` - 获取出入库记录

### Excel 管理

- `GET /api/excel/export` - 导出 Excel
- `POST /api/excel/import` - 导入 Excel

## Excel 导入格式

导入的 Excel 文件应包含以下列：

| 序号 | 商品编号 | 品牌 | 厂家型号 | 封装 | 商品名称 | 订购数量 | 商品金额 |
|------|----------|------|----------|------|----------|----------|----------|

## 配置说明

编辑 `backend/config.yaml` 文件：

```yaml
server:
  port: 8080          # 服务端口
  mode: debug         # 运行模式：debug/release

database:
  host: 127.0.0.1     # 数据库地址
  port: 3306          # 数据库端口
  user: root          # 数据库用户名
  password: root      # 数据库密码
  dbname: cms_db      # 数据库名称

log:
  level: debug        # 日志级别
  filename: backend/logs/app.log
```

## 开发说明

### 后端开发

```bash
# 安装依赖
go mod tidy

# 运行
go run cmd/main.go

# 构建
go build -o cms-backend cmd/main.go
```

### 前端开发

```bash
# 安装依赖
npm install

# 开发模式
npm run dev

# 构建生产版本
npm run build
```

## 注意事项

1. 确保 MySQL 数据库已启动并可访问
2. 首次运行前需要执行数据库初始化脚本
3. 修改配置文件中的数据库连接信息
4. 入库和出库操作会自动更新库存数量
5. 出库时如果库存不足会提示错误

## 许可证

MIT License
