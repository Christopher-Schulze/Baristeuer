# Cloud Sync (Preview)

Dieses Projekt wird zukünftig eine optionale Synchronisierung der SQLite-Datenbank über einen Cloud-Service unterstützen. 
Aktuell existiert lediglich eine lokale Platzhalter-Implementierung, die Dateien in ein Unterverzeichnis `syncdata/` kopiert.

Zum Testen kann einfach `SyncUpload` bzw. `SyncDownload` über die Anwendung aufgerufen werden.
Die eigentliche Integration einer REST-API oder eines anderen Dienstes folgt später.
