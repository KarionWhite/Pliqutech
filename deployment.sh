#!/bin/bash

set -e

# --- Hier die Pfade eintragen, die du mit 'which' gefunden hast ---
NPM_PATH="/home/karion/.nvm/versions/node/v22.15.1/bin/npm"
GO_PATH="/usr/local/go/bin/go"
# -----------------------------------------------------------

DEPLOY_PATH="/var/www/pliqutech"
PROJECT_DIR="/home/karion/myPliqutech" # Sicherstellen, dass wir absolute Pfade haben

echo "🚀 Starte lokales Deployment für Pliqutech..."

# 1. Baue das Vue.js Frontend
echo "▶️  1/3: Baue Frontend..."
cd ${PROJECT_DIR}/vue-frontend && ${NPM_PATH} run build

# 2. Baue das Go Backend
echo "▶️  2/3: Baue Backend..."
cd ${PROJECT_DIR}/go-backend && ${GO_PATH} build -o ${DEPLOY_PATH}/pliqutech-server

# 3. Kopiere das Frontend
echo "▶️  3/3: Kopiere Frontend-Dateien..."
rm -rf ${DEPLOY_PATH}/dist
cp -r ${PROJECT_DIR}/vue-frontend/dist ${DEPLOY_PATH}/
sudo chown -R karion:karion /var/www/pliqutech

echo "▶️  Fast fertig! Starte Dienste neu..."
# Diese Befehle brauchen sudo und bleiben unverändert
sudo systemctl daemon-reload
sudo systemctl restart pliqutech.service
sudo systemctl reload nginx
sudo systemctl restart pliqutech.service
sudo systemctl status pliqutech.service

# --- UFW Firewall-Regeln für die Entwicklung wieder schließen ---
echo "🔒 Schließe UFW-Ports für die Entwicklung..."

# Prüfe, ob UFW aktiv ist, bevor Regeln gelöscht werden
UFW_STATUS=$(sudo ufw status | grep Status | awk '{print $2}')

if [ "$UFW_STATUS" = "active" ]; then
    # Versuche, die spezifischen Entwicklungsports zu löschen
    # Diese Befehle werden fehlschlagen, wenn die Regeln nicht existieren,
    # aber das ist in Ordnung, da 'set -e' nur bei einem echten Fehler abbricht.
    sudo ufw delete allow 5173/tcp || true # Vue Frontend Dev Server
    sudo ufw delete allow 8080/tcp || true # Go Backend API Dev (Beispielport)
    sudo ufw delete allow 2345/tcp || true # Delve Debugger

    echo "UFW-Regeln für 5173, 8080 und 2345 (falls vorhanden) entfernt."
else
    echo "UFW ist inaktiv. Keine Firewall-Regeln zu entfernen."
fi

echo "✅ Deployment erfolgreich abgeschlossen!"