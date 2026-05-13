# Fork 工作流说明

本文档说明当前 fork 仓库的推荐开发、同步上游和镜像打包流程。

当前 remote 约定：

- `origin`: 自己的 fork，`git@github.com:0xLevin/sub2api.git`
- `upstream`: 原始上游仓库，`git@github.com:Wei-Shaw/sub2api.git`

## 分支职责

推荐长期保持两个主要分支：

- `main`: 尽量保持干净，用来同步上游 `upstream/main`
- `product-edition`: 放自己的长期自定义功能

自定义功能大概率不会合并回上游时，不建议长期直接改 `main`。这样可以降低以后同步上游代码时的冲突和维护成本。

## 初次确认 remote

查看 remote：

```bash
git remote -v
```

如果还没有 `upstream`，添加：

```bash
git remote add upstream git@github.com:Wei-Shaw/sub2api.git
git fetch upstream
```

如果 `upstream` 已存在但地址不对，修正：

```bash
git remote set-url upstream git@github.com:Wei-Shaw/sub2api.git
git fetch upstream
```

## 创建自定义分支

如果还没有自定义分支，基于当前 `main` 创建：

```bash
git checkout main
git checkout -b product-edition
git push -u origin product-edition
```

之后自己的功能改动都提交到 `product-edition`。

```bash
git checkout product-edition

# 修改代码后
git add .
git commit -m "custom: describe your change"
git push
```

## 同步上游最新代码

先更新干净的 `main`：

```bash
git checkout main
git fetch upstream
git merge upstream/main
git push origin main
```

然后把最新 `main` 合入自定义分支：

```bash
git checkout product-edition
git merge main
git push
```

如果发生冲突：

```bash
# 手动解决冲突后
git add .
git commit
git push
```

## 是否使用 rebase

也可以用 `rebase` 让历史更线性：

```bash
git checkout product-edition
git rebase main
```

如果有冲突：

```bash
# 手动解决冲突后
git add .
git rebase --continue
```

rebase 完成后推送：

```bash
git push --force-with-lease
```

如果这个 fork 只有自己维护，可以优先用 `rebase`。如果多人一起基于 `product-edition` 开发，建议用 `merge`，避免改写共享分支历史。

## 镜像打包

镜像应该从自定义分支打包，而不是从 `main` 打包。

推荐流程：

```bash
git checkout main
git fetch upstream
git merge upstream/main
git push origin main

git checkout product-edition
git merge main
git push

docker build -t your-image-name .
```

原因：

- `main` 只负责跟随上游
- `product-edition` 包含上游最新代码和自己的私有功能
- 打包镜像需要包含自己的私有功能，所以应从 `product-edition` 构建

## GitHub Actions

当前 fork 保留 CI 和安全扫描，用来发现同步上游或自定义修改带来的问题。

上游专用的 CLA 和 Release workflow 已加仓库限制，只会在 `Wei-Shaw/sub2api` 执行。这样 fork 中不会因为缺少上游密钥、CLA 分支或发布权限而误跑失败。

fork 自己的镜像由 `Fork Image` workflow 构建：

- 推送 `product-edition` 分支时自动构建并推送 DockerHub 镜像
- 也可以在 GitHub Actions 页面手动运行，填写 `image_tag` 生成额外 tag
- 需要在 fork 仓库配置 `DOCKERHUB_USERNAME` 和 `DOCKERHUB_TOKEN` 两个 secrets
- 默认镜像地址为 `<DOCKERHUB_USERNAME>/sub2api`
- 默认 tag 包含 `product-edition` 和 `product-edition-<short-sha>`
- 手动运行时可勾选 `push_latest`，额外推送 `latest`

使用镜像示例：

```bash
docker pull <DOCKERHUB_USERNAME>/sub2api:product-edition
```

## 不推荐的做法

不推荐把 `product-edition` 合并回 `main`：

```bash
git checkout main
git merge product-edition
```

这样会让 `main` 混入私有功能，后续同步上游时更容易产生冲突，也会失去 `main` 作为干净上游同步分支的作用。

## 常用检查命令

查看当前分支和工作区状态：

```bash
git status --short --branch
```

查看分支：

```bash
git branch -vv
```

查看本地和远程分支：

```bash
git branch -a
```

查看提交关系：

```bash
git log --oneline --graph --decorate --all -n 30
```
