# Portman - Agent Team Configuration

Ce projet utilise une équipe d'agents Antigravity pour la maintenance et l'évolution de l'application TUI `portman`.

## ArchitectAgent
- **Rôle :** Responsable du cycle de vie global, des structures de données (modèles de domaine), et de l'orchestration des commandes macOS (`lsof`, `pfctl`).
- **Focus :** Maintenir une architecture modulaire et s'assurer que les appels système n'impactent pas les performances de l'interface.
- **Skills :** Doit maîtriser le formatage des données système pour l'UI.

## TuiAgent
- **Rôle :** Spécialiste des interfaces textuelles, responsable de l'expérience utilisateur et du framework Bubble Tea.
- **Focus :** Boucle de rendu (Model-Update-View), gestion des états de l'UI (listes, sélection multiple, popups, inputs clavier), et réactivité de l'application.
- **Skills :** Doit s'assurer que les actions bloquantes sont envoyées sous forme de `tea.Cmd`.

## SecurityAgent
- **Rôle :** Garant de la sécurité des opérations.
- **Focus :** Gestion des privilèges `sudo` (détection, alertes), validation des pids avant un `kill`, filtrage des favoris pour éviter l'arrêt de processus critiques (ex: Docker, Postgres).
- **Skills :** Validation des entrées et prévention d'exécution de commandes non autorisées.
