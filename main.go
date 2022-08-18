package main

import (
	"os/exec"
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
	return string(a), UsableNouns
}
