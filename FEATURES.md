# Features & Roadmap : Portman

## 🚀 Fonctionnalités Développées
- [x] Initialisation de l'application TUI (Go + Bubble Tea).
- [x] Parsing asynchrone basique des ports TCP/UDP avec `lsof`.
- [x] Définition des compétences des agents (Kill, Firewall, NetScanner).
- [x] Catégorisation des processus (`System`, `Apps`, `Network`, `Dev`, `Other`).
- [x] UI avancée : Onglets (Tabs) de filtrage par catégorie.
- [x] UI avancée : Barre de recherche interactive.
- [x] UI avancée : Ajout de la colonne d'adresse IP (`Address`) et date de démarrage (`Started`).
- [x] Sauvegarde persistante des processus favoris dans `~/.portman.json`.
- [x] Terminaison de processus (kill) simple ou multiple depuis l'UI.
- [x] Détection automatique des conflits de ports (Address already in use).
- [x] Raccourci pour ouvrir directement un port dans le navigateur local.

## 🚧 En Cours de Développement
- [ ] Intégration de `pfctl` pour bloquer/débloquer un port directement depuis l'UI.

## 🗺️ Roadmap (À Venir)
- [ ] Support d'IPv6 (Affichage différencié).
- [ ] Exportation de la liste des ports en JSON ou CSV.
