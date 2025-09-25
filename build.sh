#!/bin/bash

# Beendet das Skript sofort, wenn ein Befehl fehlschlägt
set -e

echo "▶️  1/4: Baue das Vue.js Frontend..."
# Gehe in den Frontend-Ordner und führe den Build aus
cd vue-frontend && npm run build && cd ..

echo "▶️  2/4: Baue das Go Backend..."
# Gehe in den Backend-Ordner und baue die ausführbare Datei
cd go-backend && go build -o ../pliqutech-server && cd ..
# -o ../pliqutech-server sorgt dafür, dass die fertige Datei im Hauptordner landet

echo "▶️  3/4: Bereite den Deployment-Ordner vor..."
# Erstelle einen finalen Ordner und lösche alte Inhalte, falls vorhanden
rm -rf deployment
mkdir -p deployment/dist

echo "▶️  4/4: Kopiere fertige Dateien..."
# Kopiere die Go-Anwendung und den dist-Ordner in den deployment-Ordner
mv pliqutech-server deployment/
mv vue-frontend/dist deployment/

echo "✅ Fertig! Der 'deployment'-Ordner ist bereit zum Hochladen auf den Server."