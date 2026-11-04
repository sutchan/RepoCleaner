// src/app/layout.tsx v1.2.3
import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "RepoCleaner · 项目清洁助手",
  description: "一键让所有前端项目保持干净：无冗余 node_modules、无缓存垃圾。",
};

export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="zh" id="root-html">
      <body id="root-body" className="font-sans antialiased">
        <div id="app-shell" className="min-h-screen bg-background text-foreground">
          {children}
        </div>
      </body>
    </html>
  );
}
