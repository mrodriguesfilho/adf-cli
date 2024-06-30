#!/usr/bin/env bash

package="adf-cli"
platforms=("windows/amd64" "darwin/arm64" "darwin/amd64" "linux/amd64")

rm -rf "release"

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name="$package-$GOOS-$GOARCH"
	
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	if [ $GOOS = "darwin" ]; then
		output_name=$(echo "$output_name" | sed 's/darwin/mac/g')
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -o "release/$output_name" $package
	if [ $? -ne 0 ]; then
   		echo 'Erro ao buildar a versão $GOOS/$GOARCH! Abortando'
		exit 1
	fi

	echo "$output_name criado"
done

echo "Build completado com sucesso!"