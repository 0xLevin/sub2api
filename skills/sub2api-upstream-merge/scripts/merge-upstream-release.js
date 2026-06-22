#!/usr/bin/env node

const { spawnSync } = require("node:child_process");
const fs = require("node:fs");
const path = require("node:path");

const LOCAL_ROOT = "/home/ubuntu/dev/sub2api";
const EXPECTED_ROOT = process.env.SUB2API_MERGE_EXPECTED_ROOT
  ? path.resolve(process.env.SUB2API_MERGE_EXPECTED_ROOT)
  : process.env.GITHUB_ACTIONS === "true"
    ? path.resolve(process.env.GITHUB_WORKSPACE || process.cwd())
    : LOCAL_ROOT;
const EXPECTED_BRANCH = process.env.SUB2API_MERGE_EXPECTED_BRANCH || "product-edition";
const VERSION_FILE = "backend/cmd/server/VERSION";
const CRITICAL_VITEST = [
  "src/views/auth/__tests__/LinuxDoCallbackView.spec.ts",
  "src/views/auth/__tests__/WechatCallbackView.spec.ts",
  "src/views/user/__tests__/PaymentView.spec.ts",
  "src/views/user/__tests__/PaymentResultView.spec.ts",
  "src/components/user/profile/__tests__/ProfileInfoCard.spec.ts",
  "src/views/admin/__tests__/SettingsView.spec.ts",
];

function run(command, args, options = {}) {
  const result = spawnSync(command, args, {
    cwd: options.cwd || EXPECTED_ROOT,
    encoding: "utf8",
    stdio: options.capture ? ["ignore", "pipe", "pipe"] : "inherit",
  });
  if (result.status !== 0) {
    const rendered = [command, ...args].join(" ");
    if (options.capture) {
      process.stderr.write(result.stdout || "");
      process.stderr.write(result.stderr || "");
    }
    throw new Error(`Command failed: ${rendered}`);
  }
  return (result.stdout || "").trim();
}

function usage() {
  console.error("Usage: node skills/sub2api-upstream-merge/scripts/merge-upstream-release.js v0.1.137 [--push]");
}

function assertCleanTree() {
  const status = run("git", ["status", "--porcelain"], { capture: true });
  if (status) {
    throw new Error("Working tree is not clean. Commit, stash, or inspect local changes before merging.");
  }
}

function main() {
  const args = process.argv.slice(2);
  const tag = args.find((arg) => !arg.startsWith("--"));
  const push = args.includes("--push");
  if (!tag || !/^v\d+\.\d+\.\d+$/.test(tag)) {
    usage();
    process.exit(2);
  }

  const root = run("git", ["rev-parse", "--show-toplevel"], { capture: true });
  if (path.resolve(root) !== EXPECTED_ROOT) {
    throw new Error(`Refusing to run outside ${EXPECTED_ROOT}; current root is ${root}`);
  }

  const branch = run("git", ["branch", "--show-current"], { capture: true });
  if (branch !== EXPECTED_BRANCH) {
    throw new Error(`Expected branch ${EXPECTED_BRANCH}, got ${branch || "(detached)"}`);
  }

  const upstreamUrl = run("git", ["remote", "get-url", "upstream"], { capture: true });
  if (!/Wei-Shaw\/sub2api(?:\.git)?$/.test(upstreamUrl)) {
    throw new Error(`Expected upstream remote to point at Wei-Shaw/sub2api, got ${upstreamUrl}`);
  }

  assertCleanTree();

  run("git", ["fetch", "upstream", "tag", tag]);
  run("git", ["merge", "--no-ff", tag]);

  const version = tag.slice(1);
  const versionPath = path.join(EXPECTED_ROOT, VERSION_FILE);
  const currentVersion = fs.readFileSync(versionPath, "utf8").trim();
  let versionCommit = "";
  if (currentVersion !== version) {
    fs.writeFileSync(versionPath, `${version}\n`);
    run("git", ["add", VERSION_FILE]);
    run("git", ["commit", "-m", `chore: sync version to ${version}`]);
    versionCommit = run("git", ["rev-parse", "--short", "HEAD"], { capture: true });
  }

  run("go", ["test", "./..."], { cwd: path.join(EXPECTED_ROOT, "backend") });
  run("make", ["build"], { cwd: path.join(EXPECTED_ROOT, "backend") });
  run("pnpm", ["--dir", "frontend", "run", "lint:check"]);
  run("pnpm", ["--dir", "frontend", "run", "typecheck"]);
  run("pnpm", ["--dir", "frontend", "exec", "vitest", "run", ...CRITICAL_VITEST]);
  run("pnpm", ["--dir", "frontend", "run", "build"]);

  assertCleanTree();

  if (push) {
    run("git", ["push", "origin", `HEAD:${EXPECTED_BRANCH}`]);
  }

  const head = run("git", ["rev-parse", "--short", "HEAD"], { capture: true });
  console.log(JSON.stringify({ tag, branch, head, version, versionCommit: versionCommit || null, pushed: push }, null, 2));
}

try {
  main();
} catch (error) {
  console.error(error.message);
  process.exit(1);
}
