// eslint.config.mjs v2.3.3
import { defineConfig, globalIgnores } from "eslint/config";
import nextVitals from "eslint-config-next/core-web-vitals";
import nextTs from "eslint-config-next/typescript";

const eslintConfig = defineConfig([
  ...nextVitals,
  ...nextTs,
  // Override default ignores of eslint-config-next.
  globalIgnores([
    // Default ignores of eslint-config-next:
    ".next/**",
    "out/**",
    "build/**",
    "next-env.d.ts",
    // 临时与产物目录不参与门禁，避免一次性脚本造成假红
    ".tmp/**",
    "reports/**",
    "coverage/**",
    "test-results/**",
  ]),
]);

export default eslintConfig;
