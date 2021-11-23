echo -e 'Doing this will overwrite data on \033[1mTHIS PC\033[0m with data from \033[1mGITHUB\033[0m are you sure you want to continue '
read -p " (y/n) - " choice
case "$choice" in 
  y|Y ) git pull; cd Backend/; go install; cd ..;;
  n|N ) echo "Canceled";;
  * ) echo "Please enter (y/n)";;
esac