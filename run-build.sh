#!/bin/bash

SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_ROOT_DIR="$( cd ${SCRIPT_DIR} && git rev-parse --show-toplevel )"
PACKAGE_NAME=$(basename "$REPO_ROOT_DIR")




## Create bin dir if it not exist
mkdir -p ${REPO_ROOT_DIR}/bin
BIN_DIR=${REPO_ROOT_DIR}/bin

##Debug info
printf "DEBUG: Path of script ${SCRIPT_NAME} location = ${SCRIPT_DIR}\n"
printf "DEBUG: Path of REPO location = ${REPO_ROOT_DIR}\n"
printf "DEBUG: Package name is: ${PACKAGE_NAME}\n"
printf "DEBUG: Path of BIN_DIR location = ${BIN_DIR}\n"




rm ${BIN_DIR}/${PACKAGE_NAME}-*

## Create array of systems that will be supported for compilation
declare -a GOOS=("darwin" "linux" "windows")

for i in "${GOOS[@]}"
do
	printf "DEBUG: Now will be compiled for ${i}\n"
   case ${i} in
   windows)
           go build -o ${BIN_DIR}/${PACKAGE_NAME}-${i}.exe
           ;;
   *)
           go build -o ${BIN_DIR}/${PACKAGE_NAME}-${i}
           ;;
   esac
   	if [ -e ${BIN_DIR}/${PACKAGE_NAME}-${i}* ]
then
    echo "ok"
else
    echo "nok"
fi
   printf "\n\n"
   awk -v line=$(tput cols) 'BEGIN{for (i = 1; i <= line; ++i){printf "-";}}'
	printf "\n\n"

   # or do whatever with individual element of the array
done



git log --pretty=format:"- %cd: %s%n%b" --date=short  > ${REPO_ROOT_DIR}/CHANGELOG

# zip ${PACKAGE_NAME}.zip ${PACKAGE_NAME}-*