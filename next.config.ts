// next.config.ts v2.0.1
import type { NextConfig } from "next";
import { readFileSync } from "node:fs";
import { join } from "node:path";

// 版本号唯一来源：package.json 的 version，禁止在页面/组件中硬编码
function readAppVersion(): string {
  try {
    const raw = readFileSync(join(process.cwd(), "package.json"), "utf8");
    const pkg = JSON.parse(raw) as { version?: string };
    return pkg.version ?? "0.0.0";
  } catch {
    return "0.0.0";
  }
}

// build 编号 yymmddhhmm（10 位：2 位年 + 月日时分），dev / build 均自动注入
function makeBuildId(): string {
  const pad = (n: number) => String(n).padStart(2, "0");
  const d = new Date();
  return [
    String(d.getFullYear()).slice(-2),
    pad(d.getMonth() + 1),
    pad(d.getDate()),
    pad(d.getHours()),
    pad(d.getMinutes()),
  ].join("");
}

const nextConfig: NextConfig = {
  env: {
    NEXT_PUBLIC_APP_VERSION: readAppVersion(),
    NEXT_PUBLIC_BUILD_ID: makeBuildId(),
  },
};

export default nextConfig;
