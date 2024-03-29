package main

var stargate_text string = `
██████╗ ███████╗██████╗  ██████╗     ███████╗████████╗ █████╗ ██████╗  ██████╗  █████╗ ████████╗███████╗
██╔══██╗██╔════╝██╔══██╗██╔═══██╗    ██╔════╝╚══██╔══╝██╔══██╗██╔══██╗██╔════╝ ██╔══██╗╚══██╔══╝██╔════╝
██║  ██║█████╗  ██████╔╝██║   ██║    ███████╗   ██║   ███████║██████╔╝██║  ███╗███████║   ██║   █████╗  
██║  ██║██╔══╝  ██╔══██╗██║   ██║    ╚════██║   ██║   ██╔══██║██╔══██╗██║   ██║██╔══██║   ██║   ██╔══╝  
██████╔╝███████╗██║  ██║╚██████╔╝    ███████║   ██║   ██║  ██║██║  ██║╚██████╔╝██║  ██║   ██║   ███████╗
╚═════╝ ╚══════╝╚═╝  ╚═╝ ╚═════╝     ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝
`

var stargate_3d string = `
_______  ________ _______   ______        ______    __                                           __              
|       \|        \       \ /      \      /      \  |  \                                         |  \             
| ▓▓▓▓▓▓▓\ ▓▓▓▓▓▓▓▓ ▓▓▓▓▓▓▓\  ▓▓▓▓▓▓\    |  ▓▓▓▓▓▓\_| ▓▓_    ______   ______   ______   ______  _| ▓▓_    ______  
| ▓▓  | ▓▓ ▓▓__   | ▓▓__| ▓▓ ▓▓  | ▓▓    | ▓▓___\▓▓   ▓▓ \  |      \ /      \ /      \ |      \|   ▓▓ \  /      \ 
| ▓▓  | ▓▓ ▓▓  \  | ▓▓    ▓▓ ▓▓  | ▓▓     \▓▓    \ \▓▓▓▓▓▓   \▓▓▓▓▓▓\  ▓▓▓▓▓▓\  ▓▓▓▓▓▓\ \▓▓▓▓▓▓\\▓▓▓▓▓▓ |  ▓▓▓▓▓▓\
| ▓▓  | ▓▓ ▓▓▓▓▓  | ▓▓▓▓▓▓▓\ ▓▓  | ▓▓     _\▓▓▓▓▓▓\ | ▓▓ __ /      ▓▓ ▓▓   \▓▓ ▓▓  | ▓▓/      ▓▓ | ▓▓ __| ▓▓    ▓▓
| ▓▓__/ ▓▓ ▓▓_____| ▓▓  | ▓▓ ▓▓__/ ▓▓    |  \__| ▓▓ | ▓▓|  \  ▓▓▓▓▓▓▓ ▓▓     | ▓▓__| ▓▓  ▓▓▓▓▓▓▓ | ▓▓|  \ ▓▓▓▓▓▓▓▓
| ▓▓    ▓▓ ▓▓     \ ▓▓  | ▓▓\▓▓    ▓▓     \▓▓    ▓▓  \▓▓  ▓▓\▓▓    ▓▓ ▓▓      \▓▓    ▓▓\▓▓    ▓▓  \▓▓  ▓▓\▓▓     \
 \▓▓▓▓▓▓▓ \▓▓▓▓▓▓▓▓\▓▓   \▓▓ \▓▓▓▓▓▓       \▓▓▓▓▓▓    \▓▓▓▓  \▓▓▓▓▓▓▓\▓▓      _\▓▓▓▓▓▓▓ \▓▓▓▓▓▓▓   \▓▓▓▓  \▓▓▓▓▓▓▓
                                                                             |  \__| ▓▓                           
                                                                              \▓▓    ▓▓                           
                                                                               \▓▓▓▓▓▓                            

`

var main_stargate string = `
▄▄▄▄▄▄  ▄▄▄▄▄▄▄ ▄▄▄▄▄▄   ▄▄▄▄▄▄▄    ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄ ▄▄▄▄▄▄   ▄▄▄▄▄▄▄ ▄▄▄▄▄▄ ▄▄▄▄▄▄▄ ▄▄▄▄▄▄▄ 
█      ██       █   ▄  █ █       █  █       █       █      █   ▄  █ █       █      █       █       █
█  ▄    █    ▄▄▄█  █ █ █ █   ▄   █  █  ▄▄▄▄▄█▄     ▄█  ▄   █  █ █ █ █   ▄▄▄▄█  ▄   █▄     ▄█    ▄▄▄█
█ █ █   █   █▄▄▄█   █▄▄█▄█  █ █  █  █ █▄▄▄▄▄  █   █ █ █▄█  █   █▄▄█▄█  █  ▄▄█ █▄█  █ █   █ █   █▄▄▄ 
█ █▄█   █    ▄▄▄█    ▄▄  █  █▄█  █  █▄▄▄▄▄  █ █   █ █      █    ▄▄  █  █ █  █      █ █   █ █    ▄▄▄█
█       █   █▄▄▄█   █  █ █       █   ▄▄▄▄▄█ █ █   █ █  ▄   █   █  █ █  █▄▄█ █  ▄   █ █   █ █   █▄▄▄ 
█▄▄▄▄▄▄██▄▄▄▄▄▄▄█▄▄▄█  █▄█▄▄▄▄▄▄▄█  █▄▄▄▄▄▄▄█ █▄▄▄█ █▄█ █▄▄█▄▄▄█  █▄█▄▄▄▄▄▄▄█▄█ █▄▄█ █▄▄▄█ █▄▄▄▄▄▄▄█

`

var graphic_dero_startgate string = `
▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
██░▄▄▀██░▄▄▄██░▄▄▀██░▄▄▄░████░▄▄▄░█▄░▄█░▄▄▀█░▄▄▀█░▄▄▄█░▄▄▀█▄░▄█░▄▄
██░██░██░▄▄▄██░▀▀▄██░███░████▄▄▄▀▀██░██░▀▀░█░▀▀▄█░█▄▀█░▀▀░██░██░▄▄
██░▀▀░██░▀▀▀██░██░██░▀▀▀░████░▀▀▀░██▄██▄██▄█▄█▄▄█▄▄▄▄█▄██▄██▄██▄▄▄
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
`

var mod_graphic_hansenmod3 string = `
▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
██ ██ █ ▄▄▀█ ▄▄▀█ ▄▄█ ▄▄█ ▄▄▀█ ▄▄ █ ▄▄ ██ ▄▀▄ █▀▄▄▀█ ▄▀██
██ ▄▄ █ ▀▀ █ ██ █▄▄▀█ ▄▄█ ██ ███▄▀███▄▀██ █ █ █ ██ █ █ ██
██ ██ █▄██▄█▄██▄█▄▄▄█▄▄▄█▄██▄█ ▀▀ █ ▀▀ ██ ███ ██▄▄██▄▄███
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
`

var real_hansen33_mod string = `
░▒█░▒█░█▀▀▄░█▀▀▄░█▀▀░█▀▀░█▀▀▄░█▀▀█░█▀▀█░▒█▀▄▀█░▄▀▀▄░█▀▄░░
░▒█▀▀█░█▄▄█░█░▒█░▀▀▄░█▀▀░█░▒█░░▒▀▄░░▒▀▄░▒█▒█▒█░█░░█░█░█░░
░▒█░▒█░▀░░▀░▀░░▀░▀▀▀░▀▀▀░▀░░▀░█▄▄█░█▄▄█░▒█░░▒█░░▀▀░░▀▀░░░
`

var mod_graphic_hansenmod string = `
▄▄  
▀████▀  ▀████▀▀                                                              ▀████▄     ▄███▀              ▀███  
  ██      ██                                                                   ████    ████                  ██  
  ██      ██   ▄█▀██▄ ▀████████▄  ▄██▀███ ▄▄█▀██▀████████▄   ██▀▀█▄   ██▀▀█▄   █ ██   ▄█ ██   ▄██▀██▄   ▄█▀▀███  
  ██████████  ██   ██   ██    ██  ██   ▀▀▄█▀   ██ ██    ██  ███  ▀██ ███  ▀██  █  █▓  █▀ ██  ██▀   ▀██▄██    ██  
  ▓█      █▓   ▄███▓█   █▓    ██  ▀█████▄▓█▀▀▀▀▀▀ █▓    ██       ▄██      ▄██  ▓  █▓▄█▀  ██  ██     ███▓█    █▓  
  ▓█      █▓  █▓   ▓█   █▓    ▓█       ██▓█▄    ▄ █▓    ▓█     ▀▀██▄    ▀▀██▄  ▓  ▀▓█▀   ██  ██     ▓█▀▓█    █▓  
  ▒▓      ▓▓   ▓▓▓▓▒▓   ▓▓    ▓▓  ▀▓   █▓▓▓▀▀▀▀▀▀ ▓▓    ▓▓       ▓█▓      ▓█▓  ▓  ▓▓▓▓▀  ▓▓  ▓█     ▓▓▓▓▓    ▓▓  
  ▒▓      ▒▓  ▓▓   ▒▓   ▓▓    ▓▓  ▓▓   ▓▓▒▓▓      ▓▓    ▓▓     ▀▀▓▓▓    ▀▀▓▓▓  ▒  ▀▓▓▀   ▓▓  ▓▓▓   ▓▓▓▀▒▓    ▓▒  
▒▒▒ ▒   ▒ ▒▓▒▒▒▓▒ ▒ ▓▒▒ ▒▒▒  ▒▓▒ ▒▒ ▒▓▒   ▒ ▒ ▒▒▒ ▒▒▒  ▒▓▒ ▒      ▒        ▒ ▒ ▒▒▒ ▒   ▒ ▒▒▒  ▒ ▒ ▒ ▒  ▒ ▒ ▒ ▓ ▒ 
                                                           ▒▒▒  ▒▒▒ ▒▒▒  ▒▒▒                                     
                                                            ▒▒▒▒▒▒   ▒▒▒▒▒▒                                      

`
