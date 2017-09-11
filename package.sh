#!/bin/bash -e

GLOB='summon-file-*-amd64'

echo "==> Packaging..."

rm -rf output/dist && mkdir -p output/dist

pushd output

for binary_name in $GLOB; do
  pushd dist

  cp ../$binary_name summon-file && \
  tar -cvzf $binary_name.tar.gz summon-file && \
  rm -f summon-file

  popd
done

popd

# # Make the checksums
echo "==> Checksumming..."
pushd output/dist
shasum -a256 * > SHA256SUMS.txt
popd
