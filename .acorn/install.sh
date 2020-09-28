if command -v apt &> /dev/null; then
	make aptinstall
elif command -v aptitude &> /dev/null; then
        make aptitude
elif command -v brew &> /dev/null; then
	make brewinstall
elif command -v zypper &> /dev/null; then
        make zyppinstall
elif command -v yay &> /dev/null; then
        make yayinstall
elif command -v pacman &> /dev/null; then
        make pacinstall
fi
