// e:/Github/RepoCleaner/prototype/src/components/ui/alert.tsx v1.2.0
import * as React from "react";
import { cn } from "@/lib/utils";

type Tone = "primary" | "destructive" | "success" | "warning" | "secondary";

const tones: Record<Tone, string> = {
  primary: "border-primary/30 bg-primary/10 text-primary",
  destructive: "border-destructive/40 bg-destructive/10 text-destructive",
  success: "border-success/40 bg-success/10 text-success",
  warning: "border-warning/40 bg-warning/10 text-warning",
  secondary: "border-border bg-muted text-muted-foreground",
};

export function Alert({
  tone = "secondary", className, children, ...props
}: React.HTMLAttributes<HTMLDivElement> & { tone?: Tone }) {
  return (
    <div
      role="alert"
      className={cn("flex items-start gap-3 rounded-[10px] border p-4 text-sm", tones[tone], className)}
      {...props}
    >
      {children}
    </div>
  );
}
