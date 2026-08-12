// e:/Github/RepoCleaner/prototype/src/components/ui/badge.tsx v1.2.0
import * as React from "react";
import { cn } from "@/lib/utils";

type Tone = "primary" | "secondary" | "destructive" | "success" | "warning";

const tones: Record<Tone, string> = {
  primary: "bg-primary/15 text-primary border-primary/30",
  secondary: "bg-muted text-muted-foreground border-border",
  destructive: "bg-destructive/15 text-destructive border-destructive/30",
  success: "bg-success/15 text-success border-success/30",
  warning: "bg-warning/15 text-warning border-warning/30",
};

export function Badge({
  tone = "secondary", className, ...props
}: React.HTMLAttributes<HTMLSpanElement> & { tone?: Tone }) {
  return (
    <span
      className={cn(
        "inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-medium",
        tones[tone], className
      )}
      {...props}
    />
  );
}
