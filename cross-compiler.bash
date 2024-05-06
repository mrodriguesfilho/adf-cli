#!/usr/bin/env bash

package="adf-cli"
platforms=("windows/amd64" "darwin/arm64" "darwin/amd64" "linux/amd64")

if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

rm -rf "release"

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package
	output_dir="release/$GOOS-$GOARCH"

	mkdir -p "$output_dir"
	
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o "$output_dir/$output_name" $package
	if [ $? -ne 0 ]; then
   		echo 'Erro ao buildar a versão $GOOS/$GOARCH! Abortando'
		exit 1
	fi
done

echo "Build completado com sucesso!"