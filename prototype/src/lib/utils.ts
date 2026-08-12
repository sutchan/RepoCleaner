// e:/Github/RepoCleaner/prototype/src/lib/utils.ts v1.2.1
export type ClassValue = string | number | null | false | undefined | ClassValue[];

export function cn(...inputs: ClassValue[]): string {
  const out: string[] = [];
  for (const i of inputs) {
    if (!i) continue;
    if (Array.isArray(i)) out.push(cn(...i));
    else out.push(String(i));
  }
  return out.join(" ").replace(/\s+/g, " ").trim();
}
