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

## 版本号规则

当前 fork 使用“上游版本 + 自定义版本”的版本号：

- 上游版本来自 `backend/cmd/server/VERSION`
- 自定义版本来自 `backend/cmd/server/FORK_VERSION`
- 最终应用版本形如 `0.1.126-product.1`

同步上游后，如果上游修改了 `backend/cmd/server/VERSION`，最终版本会自动变成新的上游版本加当前自定义版本。例如上游更新到 `0.1.127` 后，当前 fork 构建版本会变成 `0.1.127-product.1`。

当自己的私有功能发生需要区分生产镜像的改动时，递增 `backend/cmd/server/FORK_VERSION`：

```bash
# 例如 product.1 -> product.2
vim backend/cmd/server/FORK_VERSION
git add backend/cmd/server/FORK_VERSION
git commit -m "chore: bump fork version"
git push
```

`Fork Image` workflow 会使用这个版本号：

- 应用内显示版本：`0.1.126-product.1`
- Docker tag：`<DOCKERHUB_USERNAME>/sub2api:0.1.126-product.1`
- 同时仍保留 `product-edition` 和 `product-edition-<short-sha>` tag

## 更新检测

原版会在后台检查 `Wei-Shaw/sub2api` 的最新 release，并支持直接下载上游二进制进行更新。当前 fork 包含私有改动，不能直接套用这个方式，否则可能把自定义功能替换成上游原版。

当前策略：

- 后台仍会检查上游最新版本，方便知道是否需要同步上游
- 自定义 fork 构建不会显示“一键更新”按钮
- 自定义 fork 构建调用自动更新接口会被拒绝
- 正确更新方式是先同步上游到 `main`，再把 `main` 合入 `product-edition`，最后重新构建/拉取 fork 镜像

生产环境推荐更新流程：

```bash
git checkout main
git fetch upstream
git merge upstream/main
git push origin main

git checkout product-edition
git merge main
git push

# 等待 Fork Image workflow 构建完成后，在生产环境拉取新镜像
docker compose pull sub2api
docker compose up -d sub2api
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
