cd ..
echo -e 'Doing this will \033[1m OVERWRITE ALL DATA ON THE GITHUB\033[0m are you sure you want to continue '
read -p " (y/n) - " choice
case "$choice" in 
  y|Y ) read -p "Please enter a commit message and dont make it dumb - " commitmsg; git commit -m "$commitmsg"; git push -u origin main;;
  n|N ) echo "Canceled";;
  * ) echo "Please enter (y/n)";;
esac
