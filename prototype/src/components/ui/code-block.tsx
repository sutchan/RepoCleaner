// e:/Github/RepoCleaner/prototype/src/components/ui/code-block.tsx v1.2.0
"use client";
import * as React from "react";
import { cn } from "@/lib/utils";

export function CodeBlock({ code, className }: { code: string; className?: string }) {
  const [copied, setCopied] = React.useState(false);
  const copy = async () => {
    try { await navigator.clipboard.writeText(code); setCopied(true); setTimeout(() => setCopied(false), 1500); }
    catch { /* clipboard 不可用时静默 */ }
  };
  return (
    <div className={cn("relative rounded-[10px] border border-border bg-[#0d0d0f]", className)}>
      <button
        onClick={copy}
        aria-label="复制代码"
        className="absolute right-2 top-2 rounded-md border border-border bg-card px-2 py-1 text-xs text-muted-foreground hover:text-foreground"
      >
        {copied ? "已复制" : "复制"}
      </button>
      <pre className="overflow-x-auto p-4 pr-16 text-[13px] leading-5 text-muted-foreground">
        <code className="font-mono">{code}</code>
      </pre>
    </div>
  );
}
