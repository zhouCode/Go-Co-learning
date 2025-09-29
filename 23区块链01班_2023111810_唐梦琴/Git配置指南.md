# Git用户名和邮箱配置指南

## 为什么需要配置Git用户名和邮箱？

配置Git用户名和邮箱是使用Git进行代码提交的基本要求。每次提交代码时，Git会记录这些信息，以便追踪谁做了什么修改。

## 配置步骤

请按照以下步骤在您的电脑上配置Git用户名和邮箱：

### 步骤1：打开终端

在Windows系统中，您可以：
1. 按下 `Win + R` 键
2. 输入 `cmd` 并按回车，打开命令提示符
3. 或者输入 `powershell` 并按回车，打开PowerShell

### 步骤2：执行配置命令

在终端中输入以下命令（注意替换为您自己的姓名和邮箱）：

```bash
# 配置全局Git用户名
git config --global user.name "唐梦琴"

# 配置全局Git邮箱
git config --global user.email "您的邮箱地址"
```

例如，如果您的邮箱是 `tangmq@example.com`，则输入：

```bash
git config --global user.email "tangmq@example.com"
```

### 步骤3：验证配置是否成功

配置完成后，可以使用以下命令验证配置是否成功：

```bash
git config --global --list
```

执行后，您应该能看到类似这样的输出：
```
user.name=唐梦琴
user.email=tangmq@example.com
```

## 注意事项

1. 全局配置（--global）会应用于您电脑上的所有Git仓库
2. 如果需要为特定仓库设置不同的用户名和邮箱，可以在该仓库目录下执行相同命令，但去掉 `--global` 参数
3. 邮箱地址最好使用您常用的邮箱，以便于识别
4. 用户名可以是您的真实姓名或昵称

## 其他有用的Git配置

您还可以配置Git的默认编辑器、颜色显示等：

```bash
# 配置默认编辑器为VS Code
git config --global core.editor "code --wait"

# 启用Git彩色输出
git config --global color.ui auto
```