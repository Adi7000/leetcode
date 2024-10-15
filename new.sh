#!/bin/bash

if [ $# -eq 2 ]; then
    cd src

    #Create the program file
    file="$1.go"
    touch "$1.go"
    echo "package leetcode" >> $file
    echo "" >> $file

    #Create the program function
    echo "func $2() {" >> $file
    echo "" >> $file
    echo "}" >> $file

    #Create the test file
    functionName=$(echo "$functionName" | awk '{print toupper(substr($0,1,1)) tolower(substr($0,2))}')
    testFile="${1}_test.go"
    touch $testFile
    echo "package leetcode" >> $testFile
    echo "" >> $testFile
    echo "import (" >> $testFile
	echo "    \"testing\"" >> $testFile
    echo ")" >> $testFile
    echo "func Test${functionName}(t *testing.T) {" >> $testFile
    echo "" >> $testFile
    echo "}" >> $testFile

else
    echo "Usage: ./new \${fileName} \${functionName}"
fi