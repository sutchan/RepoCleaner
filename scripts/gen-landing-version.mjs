// scripts/gen-landing-version.mjs v2.3.3
// 从版本权威源 internal/buildinfo/version.go 读取版本号，按本地时间生成 10 位 build 编号，
// 写入 landing/assets/version.js（构建期注入，页面运行时不生成、不硬编码版本号）。
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const ROOT = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const source = fs.readFileSync(path.join(ROOT, 'internal/buildinfo/version.go'), 'utf8');
const match = source.match(/const Version = "([^"]+)"/);
if (!match) {
  throw new Error('版本权威源读取失败：internal/buildinfo/version.go');
}
const version = match[1];
const now = new Date();
const pad = (n) => String(n).padStart(2, '0');
const buildId = [
  String(now.getFullYear()).slice(-2),
  pad(now.getMonth() + 1),
  pad(now.getDate()),
  pad(now.getHours()),
  pad(now.getMinutes()),
].join('');
const label = `V${version}· build ${buildId}`;
const target = path.join(ROOT, 'landing/assets/version.js');
fs.mkdirSync(path.dirname(target), { recursive: true });
fs.writeFileSync(
  target,
  `// 由 scripts/gen-landing-version.mjs 生成，请勿手改\nwindow.LANDING_VERSION = ${JSON.stringify({ version, buildId, label }, null, 2)};\n`,
  'utf8',
);
console.log(label);
