package models

const (
	ReferenceFileName = "references.json"
)

type References struct {
	InstalledBundles []BundleInstalled
}

type BundleInstalled struct {
	Version       string
	DirectoryPath string
}

func NewReference(version, directoryPath string) References {
	reference := References{
		InstalledBundles: []BundleInstalled{
			{
				Version:       version,
				DirectoryPath: directoryPath,
			},
		},
	}

	return reference
}
