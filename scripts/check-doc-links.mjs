// scripts/check-doc-links.mjs v2.3.1
// 校验仓库内所有 Markdown 相对链接的目标是否存在，且文件名大小写严格匹配。
// 用法：npm run check:docs（或在仓库根执行 node scripts/check-doc-links.mjs）
// 退出码：0 = 全部通过；1 = 存在失效链接。
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

// 兼容 Node 18：不使用 import.meta.dirname
const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const SKIP = new Set(["node_modules", ".next", ".git", "out", "dist", "build", "bin", ".tmp", "reports", "coverage", "test-results"]);

function collectMarkdown(dir, out = []) {
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    if (SKIP.has(entry.name)) continue;
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) collectMarkdown(full, out);
    else if (/\.md$/i.test(entry.name)) out.push(full);
  }
  return out;
}

// Windows 文件系统不区分大小写，需逐段比对真实文件名，防止 GitHub / Linux 上 404
function existsWithExactCase(target) {
  const relative = path.relative(ROOT, target);
  if (relative.startsWith("..") || path.isAbsolute(relative)) return false;
  let cursor = ROOT;
  for (const segment of relative.split(path.sep)) {
    if (!segment) continue;
    let names;
    try {
      names = fs.readdirSync(cursor);
    } catch {
      return false;
    }
    if (!names.includes(segment)) return false;
    cursor = path.join(cursor, segment);
  }
  return true;
}

const linkPattern = /\[[^\]]*\]\(([^)]+)\)/g;
const broken = [];
let checked = 0;
const files = collectMarkdown(ROOT);

for (const file of files) {
  const text = fs.readFileSync(file, "utf8");
  let match;
  while ((match = linkPattern.exec(text)) !== null) {
    let raw = match[1].trim();
    // 跳过外部链接、页内锚点、data URI
    if (/^(https?:|mailto:|#|data:)/i.test(raw)) continue;
    raw = raw.replace(/^`|`$/g, "");
    const anchor = raw.indexOf("#");
    if (anchor >= 0) raw = raw.slice(0, anchor);
    if (!raw) continue;
    checked++;
    let target;
    try {
      target = path.resolve(path.dirname(file), decodeURIComponent(raw));
    } catch {
      broken.push(`${path.relative(ROOT, file)} -> ${raw} (无法解析)`);
      continue;
    }
    if (!existsWithExactCase(target)) {
      const line = text.slice(0, match.index).split("\n").length;
      broken.push(`${path.relative(ROOT, file)}:${line} -> ${raw}`);
    }
  }
}

console.log(`[check:docs] Markdown 文件 ${files.length} 个，相对链接 ${checked} 个，失效 ${broken.length} 个`);
if (broken.length) {
  for (const item of broken) console.error(`  BROKEN ${item}`);
  process.exitCode = 1;
} else {
  console.log("[check:docs] 全部链接可达且大小写正确");
}
