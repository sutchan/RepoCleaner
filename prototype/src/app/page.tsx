// src/app/page.tsx v1.2.1
"use client";

import { useState } from "react";
import {
  Github,
  Menu,
  Package,
  RotateCcw,
  Trash2,
  RotateCw,
  Download,
  Brush,
  Zap,
  ArrowRight,
  Check,
} from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent } from "@/components/ui/card";
import { features, projectTypes, steps, i18n, type Lang } from "@/lib/data";

const ICONS = {
  package: Package,
  "rotate-ccw": RotateCcw,
  trash: Trash2,
  "rotate-cw": RotateCw,
  download: Download,
  broom: Brush,
  zap: Zap,
} as const;

export default function Home() {
  const [lang, setLang] = useState<Lang>("zh");
  const t = i18n[lang];

  return (
    <div id="page-home" className="flex flex-col">
      {/* 顶部导航 */}
      <header id="site-header" className="sticky top-0 z-50 border-b border-border bg-background/80 backdrop-blur">
        <div id="site-header-inner" className="mx-auto flex h-16 max-w-6xl items-center justify-between px-6">
          <div id="brand" className="flex items-center gap-2">
            <div id="brand-logo" className="flex h-8 w-8 items-center justify-center rounded-lg bg-primary text-primary-foreground">
              <Brush className="h-4 w-4" />
            </div>
            <span id="brand-name" className="font-semibold">{t.appName}</span>
          </div>
          <nav id="site-nav" className="hidden items-center gap-6 md:flex">
            <a id="nav-features" href="#features" className="text-sm text-muted-foreground hover:text-foreground">{t.navFeatures}</a>
            <a id="nav-projects" href="#projects" className="text-sm text-muted-foreground hover:text-foreground">{t.navProjects}</a>
            <a id="nav-steps" href="#steps" className="text-sm text-muted-foreground hover:text-foreground">{t.navSteps}</a>
          </nav>
          <div id="lang-switch" className="flex items-center gap-2">
            <span id="lang-label" className="text-xs text-muted-foreground">{t.langLabel}</span>
            <button
              id="lang-toggle"
              type="button"
              onClick={() => setLang((l) => (l === "zh" ? "en" : "zh"))}
              className="rounded-md border border-border px-2 py-1 text-xs hover:bg-accent"
            >
              {lang === "zh" ? "EN" : "中"}
            </button>
          </div>
        </div>
      </header>

      {/* Hero */}
      <section id="hero" className="relative overflow-hidden">
        <div id="hero-bg" className="pointer-events-none absolute inset-0 bg-gradient-to-b from-primary/5 to-transparent" />
        <div id="hero-inner" className="relative mx-auto max-w-6xl px-6 py-20 text-center md:py-28">
          <Badge id="hero-badge" tone="secondary" className="mb-4">
            {t.heroBadge}
          </Badge>
          <h1 id="hero-title" className="mx-auto max-w-3xl text-4xl font-bold tracking-tight md:text-5xl">
            {t.heroTitle}
          </h1>
          <p id="hero-subtitle" className="mx-auto mt-4 max-w-2xl text-lg text-muted-foreground">
            {t.heroSubtitle}
          </p>
          <div id="hero-actions" className="mt-8 flex items-center justify-center gap-3">
            <Button id="cta-primary" size="lg">
              {t.ctaPrimary}
              <ArrowRight className="ml-2 h-4 w-4" />
            </Button>
            <Button id="cta-secondary" variant="outline" size="lg">
              {t.ctaSecondary}
            </Button>
          </div>
          <div id="hero-stats" className="mx-auto mt-12 grid max-w-2xl grid-cols-3 gap-6">
            <div id="stat-saved" className="text-center">
              <div className="text-2xl font-bold text-primary">10GB+</div>
              <div className="text-xs text-muted-foreground">{t.statSaved}</div>
            </div>
            <div id="stat-projects" className="text-center">
              <div className="text-2xl font-bold text-primary">4</div>
              <div className="text-xs text-muted-foreground">{t.statProjects}</div>
            </div>
            <div id="stat-languages" className="text-center">
              <div className="text-2xl font-bold text-primary">2</div>
              <div className="text-xs text-muted-foreground">{t.statLanguages}</div>
            </div>
          </div>
        </div>
      </section>

      {/* 功能 */}
      <section id="features" className="border-t border-border">
        <div id="features-inner" className="mx-auto max-w-6xl px-6 py-16">
          <div id="features-head" className="mb-10 text-center">
            <h2 id="features-title" className="text-3xl font-bold">{t.featuresTitle}</h2>
            <p id="features-subtitle" className="mt-2 text-muted-foreground">{t.featuresSubtitle}</p>
          </div>
          <div id="features-grid" className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
            {features.map((f) => {
              const Icon = ICONS[f.icon];
              return (
                <Card id={`feature-card-${f.id}`} key={f.id}>
                  <CardContent id={`feature-card-content-${f.id}`} className="pt-6">
                    <div id={`feature-icon-${f.id}`} className="mb-4 flex h-10 w-10 items-center justify-center rounded-lg bg-primary/10 text-primary">
                      <Icon className="h-5 w-5" />
                    </div>
                    <h3 id={`feature-title-${f.id}`} className="font-semibold">
                      [{f.id}] {f.title[lang]}
                    </h3>
                    <p id={`feature-desc-${f.id}`} className="mt-2 text-sm text-muted-foreground">
                      {f.desc[lang]}
                    </p>
                  </CardContent>
                </Card>
              );
            })}
          </div>
        </div>
      </section>

      {/* 项目类型 */}
      <section id="projects" className="border-t border-border bg-muted/30">
        <div id="projects-inner" className="mx-auto max-w-6xl px-6 py-16">
          <div id="projects-head" className="mb-10 text-center">
            <h2 id="projects-title" className="text-3xl font-bold">{t.projectsTitle}</h2>
            <p id="projects-subtitle" className="mt-2 text-muted-foreground">{t.projectsSubtitle}</p>
          </div>
          <div id="projects-grid" className="grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
            {projectTypes.map((p) => (
              <Card id={`project-card-${p.type}`} key={p.type}>
                <CardContent id={`project-card-content-${p.type}`} className="pt-6">
                  <div id={`project-icon-${p.type}`} className="mb-3 text-3xl">{p.icon}</div>
                  <h3 id={`project-name-${p.type}`} className="font-semibold">{p.type}</h3>
                  <p id={`project-size-${p.type}`} className="mt-1 text-sm font-medium text-primary">{p.size[lang]}</p>
                  <p id={`project-desc-${p.type}`} className="mt-2 text-xs text-muted-foreground">{p.desc[lang]}</p>
                </CardContent>
              </Card>
            ))}
          </div>
          <p id="projects-avg" className="mt-6 text-center text-sm text-muted-foreground">
            {t.projectsAvg}：≈ 290 MB / 项目
          </p>
        </div>
      </section>

      {/* 使用步骤 */}
      <section id="steps" className="border-t border-border">
        <div id="steps-inner" className="mx-auto max-w-6xl px-6 py-16">
          <div id="steps-head" className="mb-10 text-center">
            <h2 id="steps-title" className="text-3xl font-bold">{t.stepsTitle}</h2>
            <p id="steps-subtitle" className="mt-2 text-muted-foreground">{t.stepsSubtitle}</p>
          </div>
          <div id="steps-grid" className="grid gap-4 md:grid-cols-3">
            {steps.map((s, i) => (
              <Card id={`step-card-${i + 1}`} key={i}>
                <CardContent id={`step-card-content-${i + 1}`} className="pt-6">
                  <div id={`step-index-${i + 1}`} className="mb-3 flex h-8 w-8 items-center justify-center rounded-full bg-primary text-sm font-bold text-primary-foreground">
                    {i + 1}
                  </div>
                  <h3 id={`step-title-${i + 1}`} className="font-semibold">{s.title[lang]}</h3>
                  <p id={`step-desc-${i + 1}`} className="mt-2 text-sm text-muted-foreground">{s.desc[lang]}</p>
                </CardContent>
              </Card>
            ))}
          </div>
          <div id="steps-commands" className="mx-auto mt-8 max-w-xl">
            <code
              id="steps-copy-command"
              className="block rounded-lg border border-border bg-muted px-4 py-3 text-center text-sm"
            >
              RepoCleaner.bat "D:\Projects"
            </code>
          </div>
        </div>
      </section>

      {/* 页脚 */}
      <footer id="site-footer" className="border-t border-border">
        <div id="footer-inner" className="mx-auto flex max-w-6xl flex-col items-center gap-2 px-6 py-10 text-center">
          <div id="footer-brand" className="flex items-center gap-2">
            <Github className="h-4 w-4" />
            <span id="footer-name" className="text-sm font-semibold">{t.appTagline}</span>
          </div>
          <p id="footer-desc" className="text-sm text-muted-foreground">{t.footerDesc}</p>
          <p id="footer-built" className="text-xs text-muted-foreground/70">{t.footerBuilt}</p>
          <p id="footer-check" className="mt-1 flex items-center gap-1 text-xs text-primary">
            <Check className="h-3 w-3" /> v1.2.1
          </p>
        </div>
      </footer>
    </div>
  );
}
