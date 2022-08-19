package main

import (
	"os/exec"
	"strings"
)

var usableNouns = []string{
	"BiosCharacteristics",
	"BIOSVersion",
	"BuildNumber",
	"Caption",
	"CodeSet",
	"CurrentLanguage",
	"Description",
	"EmbeddedControllerMajorVersion",
	"EmbeddedControllerMinorVersion",
	"IdentificationCode",
	"InstallableLanguages",
	"InstallDate",
	"LanguageEdition",
	"ListOfLanguages",
	"Manufacturer",
	"Name",
	"OtherTargetOS",
	"PrimaryBIOS",
	"ReleaseDate",
	"SerialNumber",
	"SMBIOSBIOSVersion",
	"SMBIOSMajorVersion",
	"SMBIOSMinorVersion",
	"SMBIOSPresent",
	"SoftwareElementID",
	"SoftwareElementState",
	"Status",
	"SystemBiosMajorVersion",
	"SystemBiosMinorVersion",
	"TargetOperatingSystem",
	"Version",
}

func SystemInfo(Noun string) (result string, UsableNouns []string) {
	UsableNouns = usableNouns
	c := exec.Command("wmic", "bios", "get", Noun)
	a, _ := c.Output()
	b := string(a)
	res := strings.Split(strings.Replace(b, " ", "", 1), "\n")
	res1 := strings.TrimSpace(res[1])
	return res1, UsableNouns
}
