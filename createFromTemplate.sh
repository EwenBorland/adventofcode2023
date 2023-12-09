#! /usr/bin/bash

dayNum=$1;
echo "Creating for day : $dayNum";
 
dir="day_$dayNum"
mkdir -p $dir;

cp day_template/day_template.go $dir/$dir.go
cp day_template/day_template_test.go $dir/$dir\_test.go
>$dir/input.txt
>$dir/mock_input.txt

sed -i -e "s/day_template/$dir/g" $dir/$dir.go
sed -i -e "s/day_template/$dir/g" $dir/$dir\_test.go
sed -i -e "s/Day_template/Day_$dayNum/g" $dir/$dir.go
sed -i -e "s/Day_template/Day_$dayNum/g" $dir/$dir\_test.go