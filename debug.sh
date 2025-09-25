#!/bin/bash

# Beendet das Skript sofort, wenn ein Befehl fehlschlägt
set -e

# Name deines systemd-Services für das Go-Backend
SYSTEMD_SERVICE_NAME="pliqutech.service"

#!/bin/bash

# --- UFW Firewall-Regeln für die Entwicklung ---
echo "Konfiguriere UFW für Entwicklungsports..."

# Prüfe, ob UFW aktiv ist
UFW_STATUS=$(sudo ufw status | grep Status | awk '{print $2}')

if [ "$UFW_STATUS" = "active" ]; then
    # Erlaube eingehende Verbindungen auf den benötigten Ports
    sudo ufw allow 5173/tcp comment 'Vue Frontend Dev Server'
    sudo ufw allow 8080/tcp comment 'Go Backend API Dev' # Passe den Port an, falls dein Go-Backend einen anderen nutzt
    sudo ufw allow 2345/tcp comment 'Delve Debugger'

    echo "UFW-Regeln für 5173, 8080 und 2345 hinzugefügt."
else
    echo "UFW ist inaktiv. Keine Firewall-Regeln angewendet."
    echo "WARNUNG: Dein Server ist möglicherweise ungeschützt, wenn UFW nicht aktiv ist."
fi

# --- Bestehende PM2-Dienste stoppen ---
echo "Stoppe bestehende PM2-Dienste..."
sudo systemctl stop pm2-karion.service # Stoppt den systemd-Dienst, der PM2 verwaltet

echo "Produktionsdienste gestoppt. Bereit für Debugging."

# HINWEIS: Dein Go-Backend und Vue-Frontend werden dann separat gestartet
# (z.B. Go-Backend mit Delve im Terminal, Vue-Frontend mit npm run dev in einem anderen Terminal)
# Oder du fügst die Startbefehle hier direkt hinzu, wenn du sie nicht manuell starten möchtest.

echo "▶️  Stoppe den Produktionsdienst '$SYSTEMD_SERVICE_NAME'..."
# Dieser Befehl erfordert sudo und wird dich nach dem Passwort fragen
sudo systemctl stop "$SYSTEMD_SERVICE_NAME" --no-block || echo "Info: Service war nicht aktiv oder konnte nicht gestoppt werden."

echo "✅ Produktionsdienst gestoppt. Du kannst jetzt den Debugger in Code-Server starten."
echo "Denk daran, nach dem Debuggen dein Deployment-Skript auszuführen, um zur Produktion zurückzukehren!"