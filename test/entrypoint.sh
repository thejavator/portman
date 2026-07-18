#!/bin/bash

# Démarrer nginx en arrière-plan
nginx

# Démarrer un serveur HTTP Python en arrière-plan sur le port 8000
python3 -m http.server 8000 &

# Démarrer un listener netcat sur le port 9000 en UDP
nc -u -l -p 9000 &

# Démarrer un listener netcat sur le port 9001 en TCP
nc -l -p 9001 &

# Attendre un peu que les services démarrent
sleep 1

# Lancer portman
exec ./portman
