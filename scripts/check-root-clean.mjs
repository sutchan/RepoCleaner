// scripts/check-root-clean.mjs v2.3.3
// 校验仓库根目录只保留白名单条目：临时脚本 / 日志 / 产物必须落 .tmp/ 或 reports/。
// 用法：npm run check:root-clean（或在仓库根执行 node scripts/check-root-clean.mjs）
// 退出码：0 = 根目录清洁；1 = 存在散落条目或 git 不可用。
import { execFileSync } from "node:child_process";
import path from "node:path";
import { fileURLToPath } from "node:url";

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");

// 根目录白名单：工程清单、入口配置与约定目录，其余一律视为散落产物
const ALLOWED_FILES = new Set([
  ".gitignore",
  ".gitattributes",
  ".nvmdrc",
  "next-env.d.ts",
  "CHANGELOG.MD",
  "LICENSE",
  "README.MD",
  "RepoCleaner.bat",
  "lang.ini",
  "eslint.config.mjs",
  "go.mod",
  "go.sum",
  "next.config.ts",
  "package-lock.json",
  "package.json",
  "postcss.config.mjs",
  "tsconfig.json",
]);
const ALLOWED_DIRS = new Set([
  ".github",
  ".tmp",
  "bin",
  "cmd",
  "coverage",
  "docs",
  "internal",
  "landing",
  "prototype",
  "public",
  "reports",
  "scripts",
  "src",
  "test-results",
]);

// 只校验 git 可见条目：已被 .gitignore 排除的运行时配置 / 缓存不视为仓库污染
function visibleRootEntries() {
  const out = execFileSync(
    "git",
    ["ls-files", "--cached", "--others", "--exclude-standard"],
    { cwd: ROOT, encoding: "utf8" },
  );
  return new Set(
    out
      .split(/\r?\n/)
      .filter(Boolean)
      .map((line) => line.split(/[\\/]/)[0]),
  );
}

let entries;
try {
  entries = visibleRootEntries();
} catch (error) {
  console.error("[check:root-clean] 无法读取 git 文件清单：", error.message);
  process.exit(1);
}

// .code-workspace 是工作区配置（禁止删除），按后缀放行
const violations = [...entries].filter(
  (name) =>
    !ALLOWED_FILES.has(name) &&
    !ALLOWED_DIRS.has(name) &&
    !name.endsWith(".code-workspace"),
);

if (violations.length) {
  console.error("[check:root-clean] 根目录出现未登记的条目：");
  for (const name of violations.sort()) console.error(`  - ${name}`);
  console.error(
    "一次性脚本 / 日志 / 中间产物请移入 .tmp/，工具产物请移入 reports/；确属工程清单的请登记白名单。",
  );
  process.exitCode = 1;
} else {
  console.log(
    `[check:root-clean] 根目录 ${entries.size} 个条目全部合规（临时文件统一在 .tmp/ 与 reports/）`,
  );
}
