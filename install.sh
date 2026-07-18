#!/usr/bin/env bash
set -e

# Configuration
REPO="thejavator/portman"
BIN_DIR="/usr/local/bin"
BIN_NAME="portman"

echo "🚀 Installation de Portman..."

# Détection de l'OS
OS="$(uname -s)"
if [ "$OS" = "Darwin" ]; then
    OS="Darwin"
elif [ "$OS" = "Linux" ]; then
    OS="Linux"
else
    echo "❌ OS non supporté: $OS"
    exit 1
fi

# Détection de l'architecture
ARCH="$(uname -m)"
if [ "$ARCH" = "x86_64" ]; then
    ARCH="x86_64"
elif [ "$ARCH" = "arm64" ] || [ "$ARCH" = "aarch64" ]; then
    ARCH="arm64"
else
    echo "❌ Architecture non supportée: $ARCH"
    exit 1
fi

echo "🔍 Système détecté : $OS ($ARCH)"

# Récupérer la dernière version
echo "📦 Recherche de la dernière version..."
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_RELEASE" ]; then
    echo "❌ Impossible de trouver la dernière version de Portman."
    exit 1
fi

echo "✅ Dernière version trouvée : $LATEST_RELEASE"

# Construction de l'URL de téléchargement
TAR_FILE="portman_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/$TAR_FILE"

# Téléchargement et extraction dans un dossier temporaire
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

echo "⬇️  Téléchargement de $DOWNLOAD_URL..."
if ! curl -sL "$DOWNLOAD_URL" -o "$TAR_FILE"; then
    echo "❌ Échec du téléchargement."
    exit 1
fi

echo "📦 Extraction de l'archive..."
tar -xzf "$TAR_FILE"

# Installation du binaire
echo "🛠  Installation dans $BIN_DIR (peut nécessiter le mot de passe sudo)..."
sudo mv "$BIN_NAME" "$BIN_DIR/"
sudo chmod +x "$BIN_DIR/$BIN_NAME"

# Nettoyage
cd - > /dev/null
rm -rf "$TMP_DIR"

echo "🎉 Portman a été installé avec succès !"
echo "👉 Lancez la commande 'sudo portman' pour commencer."
