# Cloud Sync

Die Anwendung kann Daten über eine einfache REST-API sichern und wiederherstellen. 
Im Konfigurationsfile werden dazu folgende Werte gesetzt:

```json
{
  "cloudUploadURL": "https://example.com/upload",
  "cloudDownloadURL": "https://example.com/download",
  "cloudToken": "my-secret-token"
}
```

* **cloudUploadURL** – Endpunkt für das Hochladen der SQLite-Datei.
* **cloudDownloadURL** – Endpunkt für das Herunterladen der Datei.
* **cloudToken** – Bearer-Token für die Authentifizierung.

Beide Endpunkte erwarten/liefern die rohen Dateidaten. Bei Fehlern antwortet
der Server mit `400` oder höher und einem JSON-Objekt `{"error": "nachricht"}`.
Die Anwendung gibt diese Meldung direkt an den Benutzer weiter.

Ohne gesetzte URLs fällt die Implementierung auf eine lokale Kopie im
Verzeichnis `syncdata/` zurück.
