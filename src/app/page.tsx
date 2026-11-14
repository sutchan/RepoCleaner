// src/app/page.tsx v2.3.3
"use client";
import { useEffect, useState } from "react";

import { scanResults, tweaks } from "@/lib/data";
import { HTML_LANG, SUPPORTED_LANGS, i18n, type Lang } from "@/lib/i18n";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Badge } from "@/components/ui/badge";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { AlertCircle, CheckCircle2, FolderTree, ScanLine, ShieldCheck, Trash2 } from "@/components/icons";

// 页脚是版本号与 build 编号的唯一展示位置，格式 vX.Y.Z · build yymmddhhmm（10 位）
// 两者由 next.config.ts 在构建时注入，注入缺失时回退为仅显示版本号
const APP_VERSION = process.env.NEXT_PUBLIC_APP_VERSION ?? "0.0.0";
const BUILD_ID = process.env.NEXT_PUBLIC_BUILD_ID ?? "";

const LANG_STORAGE_KEY = "repo-manager-lang";

export default function Home() {
  const [lang, setLang] = useState<Lang>("zh");
  const t = i18n[lang];

  // 挂载后恢复语言偏好（避免 SSR hydration 不一致），并做枚举校验
  useEffect(() => {
    const stored = window.localStorage.getItem(LANG_STORAGE_KEY);
    if (stored && SUPPORTED_LANGS.includes(stored as Lang)) {
      // eslint-disable-next-line react-hooks/set-state-in-effect
      setLang(stored as Lang);
    }
  }, []);

  // 同步 html lang 与持久化偏好；三语词条齐全，繁中缺失由字典保证回退
  useEffect(() => {
    document.documentElement.lang = HTML_LANG[lang];
    window.localStorage.setItem(LANG_STORAGE_KEY, lang);
  }, [lang]);

  const totalScan = scanResults.length;
  const safeCount = scanResults.filter((r) => r.risk === "safe").length;

  return (
    <main id="repo-manager-app" className="min-h-screen bg-background py-8 text-foreground">
      <div id="app-container" className="mx-auto max-w-5xl px-4">
        {/* 顶部栏：标题 + 语言切换 */}
        <header id="app-header" className="mb-6 flex items-center justify-between">
          <h1 id="app-title" className="text-2xl font-bold">
            {t.appTitle}
          </h1>
          <div id="lang-switch" className="flex items-center gap-2">
            <span className="text-sm text-muted-foreground">{t.language}</span>
            <Button
              id="lang-btn-zh"
              variant={lang === "zh" ? "default" : "outline"}
              size="sm"
              onClick={() => setLang("zh")}
            >
              简体中文
            </Button>
            <Button
              id="lang-btn-zh-hant"
              variant={lang === "zh-Hant" ? "default" : "outline"}
              size="sm"
              onClick={() => setLang("zh-Hant")}
            >
              繁體中文
            </Button>
            <Button
              id="lang-btn-en"
              variant={lang === "en" ? "default" : "outline"}
              size="sm"
              onClick={() => setLang("en")}
            >
              English
            </Button>
          </div>
        </header>

        {/* 概览卡片 */}
        <section id="overview" className="mb-6 grid grid-cols-1 gap-4 sm:grid-cols-3">
          <Card id="stat-total">
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-sm text-muted-foreground">
                <ScanLine className="h-4 w-4" /> {t.statTotal}
              </CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-3xl font-bold">{totalScan}</p>
            </CardContent>
          </Card>
          <Card id="stat-safe">
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-sm text-muted-foreground">
                <ShieldCheck className="h-4 w-4" /> {t.statSafe}
              </CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-3xl font-bold">{safeCount}</p>
            </CardContent>
          </Card>
          <Card id="stat-tweaks">
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-sm text-muted-foreground">
                <FolderTree className="h-4 w-4" /> {t.statTweaks}
              </CardTitle>
            </CardHeader>
            <CardContent>
              <p className="text-3xl font-bold">{tweaks.length}</p>
            </CardContent>
          </Card>
        </section>

        {/* 安全提示 */}
        <Alert id="security-note" tone="warning" className="mb-6">
          <AlertCircle className="h-4 w-4" />
          <AlertTitle>{t.securityTitle}</AlertTitle>
          <AlertDescription>{t.securityDesc}</AlertDescription>
        </Alert>

        {/* 主内容：扫描结果 / 优化项 */}
        <Tabs id="main-tabs" defaultValue="scan">
          <TabsList id="main-tabs-list">
            <TabsTrigger id="tab-scan" value="scan">
              {t.tabScan}
            </TabsTrigger>
            <TabsTrigger id="tab-tweaks" value="tweaks">
              {t.tabTweaks}
            </TabsTrigger>
          </TabsList>

          <TabsContent id="tab-content-scan" value="scan">
            <section id="scan-results" className="space-y-3">
              {scanResults.map((r) => (
                <Card key={r.repo} id={`repo-card-${r.repo}`}>
                  <CardContent className="flex items-center justify-between py-4">
                    <div id={`repo-info-${r.repo}`}>
                      <p className="font-semibold">{r.repo}</p>
                      <p className="text-sm text-muted-foreground">{r.path}</p>
                    </div>
                    <div id={`repo-status-${r.repo}`} className="flex items-center gap-2">
                      {r.risk === "safe" ? (
                        <Badge tone="success">
                          <CheckCircle2 className="mr-1 h-3 w-3" /> {t.safe}
                        </Badge>
                      ) : (
                        <Badge tone="destructive">
                          <AlertCircle className="mr-1 h-3 w-3" /> {t.warning}
                        </Badge>
                      )}
                      <Button id={`btn-clean-${r.repo}`} variant="outline" size="sm">
                        <Trash2 className="mr-1 h-3 w-3" /> {t.clean}
                      </Button>
                    </div>
                  </CardContent>
                </Card>
              ))}
            </section>
          </TabsContent>

          <TabsContent id="tab-content-tweaks" value="tweaks">
            <section id="tweak-list" className="space-y-3">
              {tweaks.map((tw) => (
                <Card key={tw.id} id={`tweak-card-${tw.id}`}>
                  <CardContent className="py-4">
                    <div id={`tweak-head-${tw.id}`} className="mb-1 flex items-center gap-2">
                      <Badge tone="primary">{tw.category}</Badge>
                      <span className="font-semibold">{tw.title}</span>
                    </div>
                    <p className="text-sm text-muted-foreground">{tw.desc}</p>
                  </CardContent>
                </Card>
              ))}
            </section>
          </TabsContent>
        </Tabs>

        <footer id="app-footer" className="mt-8 text-center text-xs text-muted-foreground">
          {t.footer} · v{APP_VERSION}{BUILD_ID ? ` · build ${BUILD_ID}` : ""}
        </footer>
      </div>
    </main>
  );
}
