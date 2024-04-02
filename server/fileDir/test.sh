#!/bin/bash

arr=(46 26 14 25 36)
for (( i = 1 ; i < ${#arr[@]};i++));do
    for (( j = 0 ; j < ${#arr[@]}-1;j++));do
        if [ ${arr[j]} -gt ${arr[j+1]} ];then
            temp=${arr[j]}
            arr[j]=${arr[j+1]}
            arr[j+1]=$temp
        fi
    done
done
echo ${arr[@]}