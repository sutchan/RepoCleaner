// e:/Github/RepoCleaner/prototype/src/components/ui/tabs.tsx v1.2.1
"use client";
import * as React from "react";
import { cn } from "@/lib/utils";

interface TabsCtx { value: string; setValue: (v: string) => void; }
const Ctx = React.createContext<TabsCtx | null>(null);

export function Tabs({ value, onValueChange, children }: {
  value: string; onValueChange: (v: string) => void; children: React.ReactNode;
}) {
  return <Ctx.Provider value={{ value, setValue: onValueChange }}>{children}</Ctx.Provider>;
}

export function TabsList({ className, children }: { className?: string; children: React.ReactNode }) {
  return <div className={cn("inline-flex rounded-[8px] bg-muted p-1", className)}>{children}</div>;
}

export function TabsTrigger({ value, children }: { value: string; children: React.ReactNode }) {
  const ctx = React.useContext(Ctx)!;
  const active = ctx.value === value;
  return (
    <button
      onClick={() => ctx.setValue(value)}
      className={cn(
        "rounded-[6px] px-3 py-1.5 text-sm font-medium transition-colors",
        active ? "bg-card text-foreground shadow-sm" : "text-muted-foreground hover:text-foreground"
      )}
    >
      {children}
    </button>
  );
}

export function TabsContent({ value, children, className }: { value: string; children: React.ReactNode; className?: string }) {
  const ctx = React.useContext(Ctx)!;
  if (ctx.value !== value) return null;
  return <div className={cn("mt-4", className)}>{children}</div>;
}
