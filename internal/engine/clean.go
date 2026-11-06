// internal/engine/clean.go v2.0.0

package engine

import (
	"os"
	"path/filepath"

	"repo-cleaner/internal/platform"
)

// CleanNextCache 功能 3：清理并禁用 .next 缓存（对齐 OPENSPEC §3.3）。
func CleanNextCache(p platform.Platform, repoRoot string, log Logger) error {
	log.Logf("扫描并清理 Next.js 缓存...")
	dirs, err := p.ListProjectDirs(repoRoot)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if !platform.IsNextProject(dir) {
			continue
		}
		log.Logf("Processing: %s", filepath.Base(dir))
		next := filepath.Join(dir, ".next")
		if info, err := os.Stat(next); err == nil {
			if info.IsDir() {
				if err := p.RemoveAll(next); err != nil {
					log.Logf("[错误] 删除目录失败: %v", err)
					continue
				}
			} else if err := p.ClearHidden(next); err != nil {
				log.Logf("[警告] 清除属性失败: %v", err)
			}
		}
		// 写入同名文件阻止 Next.js 重建缓存目录，再设置隐藏属性
		if err := os.WriteFile(next, []byte("\n"), 0o644); err != nil {
			log.Logf("[错误] 写入 .next 文件失败: %v", err)
			continue
		}
		if err := p.SetHidden(next); err != nil {
			log.Logf("[警告] 设置隐藏属性失败: %v", err)
		}
	}
	return nil
}

// RestoreNextCache 功能 4：恢复 Next.js 缓存（对齐 OPENSPEC §3.4）。
// 关键：必须先清除隐藏属性再删除，否则删除静默失败（沿用 v1.2.3 修复约束）。
func RestoreNextCache(p platform.Platform, repoRoot string, log Logger) error {
	log.Logf("恢复 Next.js 缓存功能...")
	dirs, err := p.ListProjectDirs(repoRoot)
	if err != nil {
		return err
	}
	for _, dir := range dirs {
		if !platform.IsNextProject(dir) {
			continue
		}
		log.Logf("Restoring: %s", filepath.Base(dir))
		next := filepath.Join(dir, ".next")
		info, err := os.Stat(next)
		if err != nil || info.IsDir() {
			continue // 不存在或是正常缓存目录，无需处理
		}
		if err := p.ClearHidden(next); err != nil {
			log.Logf("[警告] 清除属性失败: %v", err)
		}
		if err := p.RemoveFile(next); err != nil {
			log.Logf("[错误] 删除失败: %v", err)
		}
	}
	return nil
}

// CleanAllCaches 功能 6：清理所有项目缓存（对齐 OPENSPEC §3.6）。
func CleanAllCaches(p platform.Platform, repoRoot string, log Logger) error {
	return cleanByWhitelist(p, repoRoot, cleanDirsF6, lockFilesF6, log)
}

// CleanViteCache 功能 7：清理 Vite/构建缓存（对齐 OPENSPEC §3.7）。
func CleanViteCache(p platform.Platform, repoRoot string, log Logger) error {
	return cleanByWhitelist(p, repoRoot, cleanDirsF7, nil, log)
}

// cleanByWhitelist 遍历 REPO_ROOT 直接子目录，仅清理白名单内的目录与文件。
func cleanByWhitelist(p platform.Platform, repoRoot string, dirNames, fileNames []string, log Logger) error {
	projects, err := p.ListProjectDirs(repoRoot)
	if err != nil {
		return err
	}
	for _, project := range projects {
		log.Logf("Cleaning: %s", filepath.Base(project))
		for _, name := range dirNames {
			target := filepath.Join(project, name)
			if _, err := os.Stat(target); err == nil {
				if err := p.RemoveAll(target); err != nil {
					log.Logf("[错误] 删除 %s 失败: %v", name, err)
				}
			}
		}
		for _, name := range fileNames {
			if err := p.RemoveFile(filepath.Join(project, name)); err == nil {
				log.Logf("Removed: %s", name)
			}
		}
	}
	return nil
}
