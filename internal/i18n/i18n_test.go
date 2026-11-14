// internal/i18n/i18n_test.go v2.0.0

package i18n

import "testing"

func TestDictsComplete(t *testing.T) {
	for _, lang := range []Lang{Zh, Hant, En} {
		d := T(lang)
		if d.AppTitle == "" || d.TargetDir == "" || d.Save == "" || d.OutputTitle == "" {
			t.Fatalf("%s: 基础文案缺失", lang)
		}
		for i, name := range d.FuncNames {
			if name == "" {
				t.Fatalf("%s: 功能 %d 文案为空", lang, i+1)
			}
		}
		if d.WarningF1 == "" || d.WarningAll == "" || d.WarningVite == "" || d.Confirm == "" {
			t.Fatalf("%s: 确认/警告文案缺失", lang)
		}
	}
}

func TestNormalize(t *testing.T) {
	if Normalize("en") != En {
		t.Fatal("en 应归一化为 En")
	}
	if Normalize("fr") != Zh {
		t.Fatal("非法语言应回落 Zh")
	}
	if T("fr").AppTitle != T(Zh).AppTitle {
		t.Fatal("非法语言应取中文文案")
	}
}

func TestHantDictAndAliases(t *testing.T) {
	for _, alias := range []string{"zh-Hant", "tw", "zh-TW", "zh-Hant-TW"} {
		if Normalize(alias) != Hant {
			t.Fatalf("%s 应归一化为 Hant", alias)
		}
	}
	if T(Hant).AppTitle == "" || T(Hant).FuncNames[0] == "" {
		t.Fatal("繁中界面字典缺失")
	}
	if Logs(Hant).F1Step1 == "" || Logs(Hant).GroupNames[0] == "" {
		t.Fatal("繁中执行日志字典缺失")
	}
	if Normalize("fr") != Zh {
		t.Fatal("非法语言应回落简中")
	}
}
