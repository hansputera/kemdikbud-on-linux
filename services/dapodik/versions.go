package dapodik

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"github.com/hansputera/kemdikbud-on-linux/constants"
	"github.com/hansputera/kemdikbud-on-linux/utils"
)

func GetCurrentVersion() *DapodikVersion {
	version := &DapodikVersion{}

	collector := colly.NewCollector()

	// Tracking version
	collector.OnHTML("p.lead", func(h *colly.HTMLElement) {
		if strings.Contains(h.Text, "Versi") {
			nums := utils.ExtractNumbersOnly(h.Text, 1)
			version.Version = fmt.Sprintf("Dapodik_%d", nums[0])
		}
	})

	// Tracking URL
	collector.OnHTML("a.btn[href]", func(h *colly.HTMLElement) {
		url := h.Attr("href")
		if strings.HasPrefix(url, constants.DAPO_CDN_RELEASE) && h.DOM.HasClass("btn btn-raised btn-success btn-lg") {
			if len(version.Url) > 0 {
				version.VokasiUrl = url
			} else {
				version.Url = url
			}
		}
	})

	// Tracking patches
	collector.OnHTML("li a[href]", func(h *colly.HTMLElement) {
		patchurl := h.Attr("href")
		if strings.Contains(patchurl, constants.DAPO_CDN_RELEASE) {
			v := utils.ExtractVersion(h.Text)
			if len(v) > 0 {
				version.Patches = append(version.Patches, &DapodikPatch{
					PatchName:  v,
					PatchUrl:   patchurl,
					IsVokasi:   strings.Contains(patchurl, "SMK"),
					Categories: strings.Split(strings.ReplaceAll(strings.ReplaceAll(h.Text, "(", ""), ")", ""), ","),
				})
			}
		}
	})

	collector.Visit(constants.DAPO_UNDUHAN)

	return version
}
