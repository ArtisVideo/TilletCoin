echo -e 'Doing this will \033[1m OVERWRITE ALL DATA ON THE GITHUB\033[0m are you sure you want to continue '
read -p " (y/n) - " choice
case "$choice" in 
  y|Y ) git push -u origin main;;
  n|N ) echo "Canceled";;
  * ) echo "Please enter (y/n)";;
esac